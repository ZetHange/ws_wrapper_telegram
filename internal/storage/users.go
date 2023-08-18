package storage

import (
	"websocket_to_telegram/internal/models"
)

type User struct {
	User       models.User
	Header     string
	TelegramId int
	ChatId     int
	InChat     string
}

var Users []User

func ContainsUser(telegramId int) (bool, User) {
	for _, user := range Users {
		if user.TelegramId == telegramId {
			return true, user
		}
	}
	return false, User{}
}

func GetInChat(telegramId int) string {
	for _, u := range Users {
		if u.TelegramId == telegramId {
			return u.InChat
			break
		}
	}
	return ""
}

func SetInChat(telegramId int, newInChatValue string) {
	for i, u := range Users {
		if u.TelegramId == telegramId {
			Users[i].InChat = newInChatValue
			break
		}
	}
}
