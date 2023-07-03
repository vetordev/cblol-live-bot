package telegrambot

import (
	"cblol-bot/application/match"
	"cblol-bot/application/ranking"
	"log"
	"os"
)

const InvalidCommand = "Oops! Comando inválido."

type Command struct {
	command   string
	arguments string

	rankingApplication *ranking.Application
	matchApplication   *match.Application
}

func (c *Command) exec() string {

	var response string

	switch c.command {
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

func NewCommand(command string, arguments string) *Command {
	lolApiKey := os.Getenv("LOL_API_KEY")

	if lolApiKey == "" {
		log.Fatal("LOL_API_KEY is empty")
	}

	lang := os.Getenv("LOL_API_LANG")

	if lolApiKey == "" {
		log.Fatal("LOL_API_LANG is empty")
	}

	rankingApplication := ranking.New(lolApiKey, lang)
	matchApplication := match.New(lolApiKey, lang)

	return &Command{command, arguments, rankingApplication, matchApplication}
}
