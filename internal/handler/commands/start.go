package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	botInternal "websocket_to_telegram/internal/bot"
)

func HandleStart(update tgbotapi.Update) {
	message := fmt.Sprintf("Думаю ты знаешь что это за бот)\nЖми /help\nТвой ID: [<code>%v</code>]", update.Message.From.ID)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ReplyToMessageID = update.Message.MessageID
	msg.ParseMode = "html"

	botInternal.SendMessage(msg)
}
