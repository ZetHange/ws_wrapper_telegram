package ws

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"nhooyr.io/websocket"
	"strings"
	"time"
	botInternal "websocket_to_telegram/internal/bot"
	"websocket_to_telegram/internal/models"
)

var Cancel context.CancelFunc

func Subscribe(channel string, header string, server string, update tgbotapi.Update, withoutMessage bool) {
	ctx, cancel := context.WithCancel(context.Background())
	Cancel = cancel
	defer cancel()

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

	url := "wss://" + string(decodeString) + "@" + server + ".artux.net/pdanetwork/" + typeChannel

	c, _, err := websocket.Dial(ctx, url, nil)
	c.SetReadLimit(10737418240) // 10mb

	conn := Connect{
		conn:      c,
		ctxGlobal: &ctx,
		tgId:      int(update.Message.From.ID),
	}
	Connections = append(Connections, conn)

	if err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка подключения по WebSocket: "+err.Error())
		msg.ParseMode = "html"

		botInternal.SendMessage(msg)
		log.Printf("Ошибка при подключении: %v\n", err)
	}
	defer c.Close(websocket.StatusInternalError, "the sky is falling")

	if !withoutMessage {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы успешно подключены к чату, показаны первые 5 сообщений\n\n/send сообщение - отправка сообщения\n/leave - ливнуть\n\nЧтобы отправить пользователя подумать о грустном нужно ответить на сообщение с ним и прописать команду:\n/ban - забанит на всегда\n/ban 60 причина сообщение - где 60 время в минутах, причина - одно слово, сообщение многа слов")
		msg.ParseMode = "html"
		botInternal.SendMessage(msg)
	}

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
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка при чтении сообщения: "+err.Error()+"\n\nПереподключение к серверу: <code>/join "+channel+"</code>")
			msg.ParseMode = "html"

			botInternal.SendMessage(msg)

			Unsubscribe(update)

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

			if len(messageJSON.Updates) > 5 && !withoutMessage {
				for _, text := range messageJSON.Updates[len(messageJSON.Updates)-5:] {
					messageText := fmt.Sprintf("<b><a href=\"%s\">%s</a></b>: %s [<code>%s</code>]\n%s", "https://admin.artux.net/panel/users/"+text.Author.ID, text.Author.Login, text.Author.Role, text.Author.ID, text.Content)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
					msg.DisableWebPagePreview = true
					msg.ParseMode = "html"

					botInternal.SendMessage(msg)
				}
			} else if !withoutMessage {
				for _, text := range messageJSON.Updates {
					messageText := fmt.Sprintf("<b><a href=\"%s\">%s</a></b>: %s [<code>%s</code>]\n%s", "https://admin.artux.net/panel/users/"+text.Author.ID, text.Author.Login, text.Author.Role, text.Author.ID, text.Content)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
					msg.DisableWebPagePreview = true
					msg.ParseMode = "html"

					botInternal.SendMessage(msg)
				}
			}
			for _, text := range messageJSON.Events {
				messageText := fmt.Sprintf("<b>%s</b>: %s\n%s", "SYSTEM", "ADMIN", text.Content)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
				msg.DisableWebPagePreview = true
				msg.ParseMode = "html"

				botInternal.SendMessage(msg)
			}
		}
	}

}

func Unsubscribe(update tgbotapi.Update) {
	conn := GetConnByTg(int(update.Message.From.ID))
	if conn.tgId != 0 {
		conn.conn.Close(websocket.StatusNormalClosure, "leave")
	}
	RemoveConnById(int(update.Message.From.ID))
	Cancel()
}

func SendMessage(update tgbotapi.Update, message string) {
	conn := GetConnByTg(int(update.Message.From.ID))
	ctxGlobal := *conn.ctxGlobal
	if conn.tgId != 0 {
		err := conn.conn.Write(ctxGlobal, websocket.MessageText, []byte(message))
		if err != nil {
			log.Println(err)
		}
	}
}
