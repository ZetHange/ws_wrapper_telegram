package ws

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"nhooyr.io/websocket"
	"reflect"
	"strings"
	"time"
	botInternal "websocket_to_telegram/internal/bot"
	"websocket_to_telegram/internal/models"
)

var conn *websocket.Conn
var ctxGlobal context.Context

func Subscribe(channel string, header string, update tgbotapi.Update) {
	ctx := context.WithoutCancel(context.Background())
	ctxGlobal = ctx

	var typeChannel string
	switch channel {
	case "GENERAL":
		typeChannel = "chat"
	case "GROUP":
		typeChannel = "groups"
	case "RP":
		typeChannel = "rp"
	}

	decodeString, err := base64.StdEncoding.DecodeString(strings.Split(header, " ")[1])
	if err != nil {
		log.Println(err.Error())
	}

	url := "wss://" + string(decodeString) + "@dev.artux.net/pdanetwork/" + typeChannel

	c, _, err := websocket.Dial(ctx, url, nil)
	c.SetReadLimit(10737418240) // 10mb
	conn = c
	if err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка подключения по WebSocket: "+err.Error())
		msg.ParseMode = "html"

		botInternal.SendMessage(msg)
		log.Println("Ошибка при подключении: %v", err)
	}
	defer c.Close(websocket.StatusInternalError, "the sky is falling")

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы успешно подключены к чату, показаны первые 10 сообщений\n\n/send сообщение - отправка сообщения\n/leave - ливнуть")
	msg.ParseMode = "html"
	botInternal.SendMessage(msg)

	ticker := time.NewTicker(time.Minute)

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				c.Ping(ctx)
			case <-done:
				return
			}
		}
	}()

	for {
		messageType, message, err := c.Read(ctx)
		if err != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка при чтении сообщения: "+err.Error())
			msg.ParseMode = "html"

			botInternal.SendMessage(msg)
			log.Println(message, messageType, reflect.TypeOf(err))

			return
		}

		if messageType == websocket.MessageText {
			var messageJSON models.JSONData
			err = json.Unmarshal(message, &messageJSON)
			if err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка при десериализации данных: "+err.Error())
				msg.ParseMode = "html"

				botInternal.SendMessage(msg)
				log.Println("Ошибка при десериализации данных:", err)
				return
			}

			if len(messageJSON.Updates) > 10 {
				for _, text := range messageJSON.Updates[len(messageJSON.Updates)-10:] {
					messageText := fmt.Sprintf("<b>%s</b>: %s\n%s", text.Author.Login, text.Author.Role, text.Content)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
					msg.ParseMode = "html"

					botInternal.SendMessage(msg)
				}
			} else {
				for _, text := range messageJSON.Updates {
					messageText := fmt.Sprintf("<b>%s</b>: %s\n%s", text.Author.Login, text.Author.Role, text.Content)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
					msg.ParseMode = "html"

					botInternal.SendMessage(msg)
				}
			}
		}
	}

}

func Unsubscribe() {
	conn.Close(websocket.StatusNormalClosure, "leave")
}

func SendMessage(message string) {
	err := conn.Write(ctxGlobal, websocket.MessageText, []byte(message))
	if err != nil {
		log.Println(err)
	}
}
