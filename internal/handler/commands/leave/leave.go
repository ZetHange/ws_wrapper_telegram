package leave

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	botInternal "websocket_to_telegram/internal/bot"
	"websocket_to_telegram/internal/storage"
	"websocket_to_telegram/internal/ws"
)

func HandleLeave(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ливаешь?")
	msg.ReplyToMessageID = update.Message.MessageID
	msg.ParseMode = "html"
	botInternal.SendMessage(msg)
	ws.Unsubscribe(update)
	storage.SetInChat(int(update.Message.From.ID), "")
}
