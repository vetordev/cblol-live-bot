package telegrambot

import tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const InvalidCommand = "Oops! Comando inválido."

type Command struct {
	message *tgbot.Message
}

func (c *Command) exec() string {

	var response string

	switch c.message.Command() {
	case "ranking":
		response = "ranking: "
		break
	default:
		response = InvalidCommand
	}

	return response
}

func newCommand(message *tgbot.Message) *Command {
	return &Command{message}
}
