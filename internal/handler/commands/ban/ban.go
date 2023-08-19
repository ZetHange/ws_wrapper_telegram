package ban

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
	botInternal "websocket_to_telegram/internal/bot"
	"websocket_to_telegram/internal/storage"
)

func HandleBan(update tgbotapi.Update, user storage.User) {
	if storage.GetInChat(int(update.Message.From.ID)) == "" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы не находитесь в чате")
		msg.ReplyToMessageID = update.Message.MessageID
		msg.ParseMode = "html"
		botInternal.SendMessage(msg)
		return
	}

	text := strings.Split(update.Message.Text, " ")
	if update.Message.ReplyToMessage == nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Чтобы забанить пользователя нужно ответить на его сообщение")
		msg.ReplyToMessageID = update.Message.MessageID
		msg.ParseMode = "html"
		botInternal.SendMessage(msg)
		return
	}
	parts := strings.Split(update.Message.ReplyToMessage.Text, "[")
	uuid := strings.Split(parts[1], "]")[0]

	if len(text) == 1 {
		go func() {
			success := BanAlways(user.Header, uuid)
			if success {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Пользователь [<code>%s</code>] успешно забанен НАВСЕГДА ХАХАХАА", uuid))
				msg.ReplyToMessageID = update.Message.MessageID
				msg.ParseMode = "html"
				botInternal.SendMessage(msg)
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Произошла ошибка, F")
				msg.ReplyToMessageID = update.Message.MessageID
				msg.ParseMode = "html"
				botInternal.SendMessage(msg)
			}
		}()
	} else if len(text) >= 4 {
		secs, err := strconv.Atoi(text[1])
		if err != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы не правильно указали время бана")
			msg.ReplyToMessageID = update.Message.MessageID
			msg.ParseMode = "html"
			botInternal.SendMessage(msg)
			return
		}
		reason := text[2]
		message := strings.Join(text[3:], " ")

		go func() {
			success := BanTime(user.Header, uuid, secs*60, reason, message)
			if success {
				messageToUser := fmt.Sprintf("Пользователь [<code>%s</code>] забанен на %v минут\nПричина: %s\nСообщение: %s", uuid, secs, reason, message)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, messageToUser)
				msg.ReplyToMessageID = update.Message.MessageID
				msg.ParseMode = "html"
				botInternal.SendMessage(msg)
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Произошла ошибка, F")
				msg.ReplyToMessageID = update.Message.MessageID
				msg.ParseMode = "html"
				botInternal.SendMessage(msg)
			}
		}()
	} else {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Думаю вы что-то не правильно указали\n\nПримеры:\n<code>/ban</code>\n<code>/ban 60 лох туда его</code>")
		msg.ReplyToMessageID = update.Message.MessageID
		msg.ParseMode = "html"
		botInternal.SendMessage(msg)
	}
}
