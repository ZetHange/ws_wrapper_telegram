package middleware

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	botInternal "websocket_to_telegram/internal/bot"
	"websocket_to_telegram/internal/storage"
)

func AuthMiddleware(update tgbotapi.Update, next func(update tgbotapi.Update, user *storage.User)) {
	isAuth, user := storage.ContainsUser(int(update.Message.From.ID))
	if !isAuth {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы не авторизованы, используйте /login для авторизации")
		msg.ReplyToMessageID = update.Message.MessageID
		msg.ParseMode = "html"
		botInternal.SendMessage(msg)
		return
	}

	next(update, user)
}
