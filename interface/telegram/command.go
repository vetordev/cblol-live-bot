package telegrambot

import (
	"cblol-bot/application/match"
	"cblol-bot/application/ranking"
)

const InvalidCommand = "Oops! Comando inválido."

type CommandHandler struct {
	rankingApplication *ranking.Application
	matchApplication   *match.Application
}

func (c *CommandHandler) Exec(command string, arguments string) string {

	var response string

	switch command {
	case "ranking":
		response = c.rankingApplication.GetRanking()
		break
	case "week":
		response = c.matchApplication.GetWeekMatches()
		break
	case "today":
		response = c.matchApplication.GetTodayMatches()
		break
	default:
		response = InvalidCommand
	}

	return response
}

func NewCommand(lolApiKey string, apiLang string) *CommandHandler {
	rankingApplication := ranking.New(lolApiKey, apiLang)
	matchApplication := match.New(lolApiKey, apiLang)

	return &CommandHandler{rankingApplication, matchApplication}
}
