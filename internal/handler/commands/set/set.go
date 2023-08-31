package set

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	botInternal "websocket_to_telegram/internal/bot"
	"websocket_to_telegram/internal/storage"
)

func HandleSet(update tgbotapi.Update, user storage.User) {
	text := strings.Split(update.Message.Text, " ")

	if len(text) < 2 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Не указан сервер\n\nПример: <code>/set dev</code>\n\nДоступные сервера: dev, app")
		msg.ReplyToMessageID = update.Message.MessageID
		msg.ParseMode = "html"
		botInternal.SendMessage(msg)
		return
	}

	server := storage.GetServer(user.TelegramId)

	if text[1] != "app" && text[1] != "dev" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Такого сервера не существует")
		msg.ReplyToMessageID = update.Message.MessageID
		msg.ParseMode = "html"
		botInternal.SendMessage(msg)
		return
	}

	if server == text[1] {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы уже находитесь на этом сервере")
		msg.ReplyToMessageID = update.Message.MessageID
		msg.ParseMode = "html"
		botInternal.SendMessage(msg)
		return
	}

	storage.SetServer(user.TelegramId, text[1])
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Сервер "+text[1]+" успешно установлен\n\nИспользуйте /join для подключения")
	msg.ReplyToMessageID = update.Message.MessageID
	msg.ParseMode = "html"
	botInternal.SendMessage(msg)
	return
}
