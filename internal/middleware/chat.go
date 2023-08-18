package middleware

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	botInternal "websocket_to_telegram/internal/bot"
	"websocket_to_telegram/internal/storage"
)

func ChatMiddleware(update tgbotapi.Update, next func(update tgbotapi.Update)) {
	if storage.GetInChat(int(update.Message.From.ID)) == "" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы не находитесь в чате")
		msg.ReplyToMessageID = update.Message.MessageID
		msg.ParseMode = "html"
		botInternal.SendMessage(msg)
		return
	}

	next(update)
}
