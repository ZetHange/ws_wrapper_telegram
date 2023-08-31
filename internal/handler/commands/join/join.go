package join

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	botInternal "websocket_to_telegram/internal/bot"
	"websocket_to_telegram/internal/handler/commands/chat"
	"websocket_to_telegram/internal/storage"
	"websocket_to_telegram/internal/ws"
	"websocket_to_telegram/pkg/utils"
)

func HandleJoin(update tgbotapi.Update, user storage.User) {
	text := strings.Split(update.Message.Text, " ")

	if len(text) < 2 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Не указан тип чата\nПодробнее в /chat")
		msg.ReplyToMessageID = update.Message.MessageID
		msg.ParseMode = "html"
		botInternal.SendMessage(msg)
		return
	}

	if user.InChat != "" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы уже находитесь в чате")
		msg.ReplyToMessageID = update.Message.MessageID
		msg.ParseMode = "html"
		botInternal.SendMessage(msg)
		return
	}

	chats := chat.GetChats(user.Header)

	if !utils.IncludeString(chats, text[1]) {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Тип чата указан неверно\nДоступные чаты можно посмотреть /chat")
		msg.ReplyToMessageID = update.Message.MessageID
		msg.ParseMode = "html"
		botInternal.SendMessage(msg)
		return
	}

	fmt.Println(user.Server)
	storage.SetInChat(user.TelegramId, text[1])
	go ws.Subscribe(text[1], user.Header, user.Server, update)
}
