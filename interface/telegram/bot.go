package telegrambot

import (
	"fmt"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Bot struct {
	bot *tgbot.BotAPI
}

func (b *Bot) Reply(chatId int64, messageId int, text string) {
	msg := b.newMessage(chatId, text)
	msg.ReplyToMessageID = messageId

	if _, err := b.bot.Send(msg); err != nil {
		fmt.Printf("could not send message: %s", err.Error())
	}
}

func (b *Bot) Send(chatId int64, text string) {
	msg := b.newMessage(chatId, text)

	if _, err := b.bot.Send(msg); err != nil {
		fmt.Printf("could not send message: %s", err.Error())
	}
}

func (b *Bot) newMessage(chatId int64, text string) tgbot.MessageConfig {
	msg := tgbot.NewMessage(chatId, text)
	msg.ParseMode = "HTML"

	return msg
}

func (b *Bot) handleUpdate(update tgbot.Update, commandHandler *CommandHandler) {
	if update.Message == nil {
		return
	}

	response := commandHandler.Exec(update.Message)

	b.Reply(update.Message.Chat.ID, update.Message.MessageID, response)
}

func (b *Bot) Run(commandHandler *CommandHandler) {
	updateConfig := tgbot.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := b.bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		b.handleUpdate(update, commandHandler)
	}
}

func New(telegramToken string, debugMode bool) *Bot {

	bot, err := tgbot.NewBotAPI(telegramToken)
	if err != nil {
		log.Fatalf("Cannot initialize telegram bot: %s", err.Error())
	}

	bot.Debug = debugMode

	return &Bot{bot}
}
