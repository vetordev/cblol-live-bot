package telegrambot

import (
	"cblol-bot/application/match"
	"cblol-bot/application/notification"
	"cblol-bot/application/ranking"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const InvalidCommand = "Oops! Comando inválido."

type CommandHandler struct {
	rankingApplication      *ranking.Application
	matchApplication        *match.Application
	notificationApplication *notification.Application
}

func (c *CommandHandler) Exec(message *tgbotapi.Message) string {

	var response string

	switch message.Command() {
	case "ranking":
		response = c.rankingApplication.GetRanking()
		break
	case "week":
		response = c.matchApplication.GetWeekMatches()
		break
	case "today":
		response = c.matchApplication.GetTodayMatches()
		break
	case "notify":
		scheduledTime := message.CommandArguments()
		response = c.notificationApplication.EnableDailyNotificationOfMatches(
			message.Chat.ID,
			message.Chat.UserName,
			scheduledTime,
		)
	default:
		response = InvalidCommand
	}

	return response
}

func NewCommand(
	rankingApplication *ranking.Application,
	matchApplication *match.Application,
	notificationApplication *notification.Application,
) *CommandHandler {
	return &CommandHandler{rankingApplication, matchApplication, notificationApplication}
}
