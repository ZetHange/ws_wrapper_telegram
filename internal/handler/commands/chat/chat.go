package chat

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	botInternal "websocket_to_telegram/internal/bot"
	"websocket_to_telegram/internal/storage"
)

func HandleChat(update tgbotapi.Update, user *storage.User) {
	go func() {
		message := fmt.Sprintf("Доступные чаты: %s, используйте /join для подключения к ним, при этом укажите ID\nПример: <code>/join GENERAL</code>", strings.Join(GetChats(user.Header), ", "))
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
		msg.ReplyToMessageID = update.Message.MessageID
		msg.ParseMode = "html"
		botInternal.SendMessage(msg)
	}()
}
