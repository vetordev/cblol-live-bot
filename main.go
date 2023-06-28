package main

import (
	"fmt"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("cannot load env file: %s", err.Error())
	}
}

func main() {
	loadEnv()

	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	if telegramToken == "" {
		log.Fatal("TELEGRAM_TOKEN is empty")
	}

	bot, err := tgbot.NewBotAPI(telegramToken)
	if err != nil {
		log.Fatalf("Cannot initialize telegram bot: %s", err.Error())
	}

	bot.Debug = true

	updateConfig := tgbot.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		fmt.Println(update.UpdateID)
		msg := tgbot.NewMessage(update.Message.Chat.ID, "Eu amo a vitória!")
		msg.ReplyToMessageID = update.Message.MessageID

		if _, err := bot.Send(msg); err != nil {
			// Note that panics are a bad way to handle errors. Telegram can
			// have service outages or network errors, you should retry sending
			// messages or more gracefully handle failures.
			panic(err)
		}
	}
}
