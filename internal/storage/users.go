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
	Server     string
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

func SetServer(telegramId int, newServerValue string) {
	for i, u := range Users {
		if u.TelegramId == telegramId {
			Users[i].Server = newServerValue
			break
		}
	}
}

func GetServer(telegramId int) string {
	for _, u := range Users {
		if u.TelegramId == telegramId {
			return u.Server
			break
		}
	}
	return ""
}

func Logout(telegramId int) {
	var result []User

	for _, p := range Users {
		if p.TelegramId != telegramId {
			result = append(result, p)
		}
	}

	Users = result
}
