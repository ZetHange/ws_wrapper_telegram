package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	botInternal "websocket_to_telegram/internal/bot"
	"websocket_to_telegram/internal/storage"
)

func HandleLogout(update tgbotapi.Update, user storage.User) {
	storage.Logout(user.TelegramId)

	message := fmt.Sprintf("Вы успешно вышли")
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ReplyToMessageID = update.Message.MessageID
	msg.ParseMode = "html"

	botInternal.SendMessage(msg)
}
