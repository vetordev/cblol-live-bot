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

func (b *Bot) Reply(chatId int64, messageId int, text string) {

	msg := tgbot.NewMessage(chatId, text)
	msg.ReplyToMessageID = messageId

	if _, err := b.bot.Send(msg); err != nil {
		fmt.Printf("could not send message: %s", err.Error())
	}
}

func (b *Bot) handleUpdate(update tgbot.Update) {
	if update.Message == nil {
		return
	}

	command := newCommand(update.Message)
	response := command.exec()

	b.Reply(update.Message.Chat.ID, update.Message.MessageID, response)
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
