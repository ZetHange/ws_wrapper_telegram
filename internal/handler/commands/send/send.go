package send

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	botInternal "websocket_to_telegram/internal/bot"
	"websocket_to_telegram/internal/ws"
)

func HandleSend(update tgbotapi.Update) {
	text := strings.Split(update.Message.Text, "/send ")
	if len(text) < 2 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Не указано само сообщение\nПример: <code>/send pomidor</code>")
		msg.ReplyToMessageID = update.Message.MessageID
		msg.ParseMode = "html"
		botInternal.SendMessage(msg)
		return
	}

	ws.SendMessage(update, text[1])
}
