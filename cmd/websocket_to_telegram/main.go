package main

import (
	"github.com/joho/godotenv"
	"log"
	"websocket_to_telegram/internal/bot"
	"websocket_to_telegram/internal/handler"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	bot.Init()
	handler.InitHandler()
}
