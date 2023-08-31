package login

import (
	"encoding/base64"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	botInternal "websocket_to_telegram/internal/bot"
	"websocket_to_telegram/internal/storage"
)

func HandleLogin(update tgbotapi.Update) {
	text := strings.Split(update.Message.Text, " ")

	if len(text) < 3 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Не указаны логин и пароль (или сервер)\nПример: <code>/login login:12345678 app</code>\n\nВместо app можно использовать dev")
		msg.ReplyToMessageID = update.Message.MessageID
		msg.ParseMode = "html"
		botInternal.SendMessage(msg)
		return
	}

	if contains := strings.Contains(text[1], ":"); !contains {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Отсутсвует разделитель \":\"\nПример: <code>/login login:12345678</code>")
		msg.ReplyToMessageID = update.Message.MessageID
		msg.ParseMode = "html"
		botInternal.SendMessage(msg)
		return
	}

	if text[2] != "app" && text[2] != "dev" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Такого сервера не существует")
		msg.ReplyToMessageID = update.Message.MessageID
		msg.ParseMode = "html"
		botInternal.SendMessage(msg)
		return
	}

	go func() {
		credentials := "Basic " + base64.StdEncoding.EncodeToString([]byte(text[1]))
		user := Auth(text[2], credentials)
		if user.ID == "" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Логин или пароль неверные")
			msg.ReplyToMessageID = update.Message.MessageID
			msg.ParseMode = "html"
			botInternal.SendMessage(msg)
			return
		}

		if user.Role == "USER" || user.Role == "CREATOR" || user.Role == "TESTER" {
			message := fmt.Sprintf("У вас нет прав, ваша роль: %s\nТреба: ADMIN или MODERATOR", user.Role)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
			msg.ReplyToMessageID = update.Message.MessageID
			msg.ParseMode = "html"
			botInternal.SendMessage(msg)
			return
		}

		message := fmt.Sprintf("Вы успешно вошли как <code>%s</code>\nВаша почта: <code>%s</code>\n\nВы находитесь в "+text[2], user.Login, user.Email)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
		msg.ReplyToMessageID = update.Message.MessageID
		msg.ParseMode = "html"

		storage.Users = append(storage.Users, storage.User{
			User:       user,
			TelegramId: int(update.Message.From.ID),
			ChatId:     update.Message.MessageID,
			Header:     credentials,
			InChat:     "",
			Server:     text[2],
		})

		botInternal.SendMessage(msg)
		return
	}()
}
