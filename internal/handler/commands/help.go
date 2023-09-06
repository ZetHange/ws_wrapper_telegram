package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	botInternal "websocket_to_telegram/internal/bot"
)

func HandleHelp(update tgbotapi.Update) {
	message := fmt.Sprintf("/start - Стартовая команда\n/help - Помощь по командам\n/login user:12345678 app - Авторизация\n/user - Получение себя хорошего\n/chat - Все доступные чаты, а также текущие подключения\n/join chat_type - Подключение к чату")
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ReplyToMessageID = update.Message.MessageID
	msg.ParseMode = "html"

	botInternal.SendMessage(msg)
}
