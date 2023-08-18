package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	botInternal "websocket_to_telegram/internal/bot"
)

func HandleUndefined(update tgbotapi.Update) {
	message := fmt.Sprintf("Кажется я не знаю такой команды\nТвой ID: [<code>%v</code>]", update.Message.From.ID)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ReplyToMessageID = update.Message.MessageID
	msg.ParseMode = "html"

	botInternal.SendMessage(msg)
}
