package telegrambot

import (
	"fmt"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

type Bot struct {
	bot *tgbot.BotAPI
}

func (b *Bot) handleUpdate(update tgbot.Update) {
	if update.Message == nil {
		return
	}

	fmt.Println(update.UpdateID)
	msg := tgbot.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyToMessageID = update.Message.MessageID

	if _, err := b.bot.Send(msg); err != nil {
		// Note that panics are a bad way to handle errors. Telegram can
		// have service outages or network errors, you should retry sending
		// messages or more gracefully handle failures.
		panic(err)
	}
}

func (b *Bot) Run() {
	updateConfig := tgbot.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := b.bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		b.handleUpdate(update)
	}
}

func New(debugMode bool) *Bot {
	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	if telegramToken == "" {
		log.Fatal("TELEGRAM_TOKEN is empty")
	}

	bot, err := tgbot.NewBotAPI(telegramToken)
	if err != nil {
		log.Fatalf("Cannot initialize telegram bot: %s", err.Error())
	}

	bot.Debug = debugMode

	return &Bot{bot}
}
