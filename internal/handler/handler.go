package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
	botInternal "websocket_to_telegram/internal/bot"
	"websocket_to_telegram/internal/handler/commands"
	"websocket_to_telegram/internal/handler/commands/ban"
	"websocket_to_telegram/internal/handler/commands/chat"
	"websocket_to_telegram/internal/handler/commands/join"
	"websocket_to_telegram/internal/handler/commands/leave"
	"websocket_to_telegram/internal/handler/commands/login"
	"websocket_to_telegram/internal/handler/commands/send"
	"websocket_to_telegram/internal/middleware"
)

func InitHandler() {
	bot := botInternal.GetBot()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			switch update.Message.Text {
			case "/start":
				commands.HandleStart(update)
			case "/help":
				commands.HandleHelp(update)
			case "/chat":
				middleware.AuthMiddleware(update, chat.HandleChat)
			case "/leave":
				middleware.ChatMiddleware(update, leave.HandleLeave)
			default:
				text := update.Message.Text

				if strings.Contains(text, "/login") {
					login.HandleLogin(update)
					break
				} else if strings.Contains(text, "/join") {
					middleware.AuthMiddleware(update, join.HandleJoin)
					break
				} else if strings.Contains(text, "/send") {
					middleware.ChatMiddleware(update, send.HandleSend)
					break
				} else if strings.Contains(text, "/ban") {
					middleware.AuthMiddleware(update, ban.HandleBan)
					break
				}
				commands.HandleUndefined(update)
			}
		}
	}
}
