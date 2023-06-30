package week

import (
	"cblol-bot/domain/model/match"
	"cblol-bot/domain/model/team"
	"cblol-bot/domain/service/week"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

const (
	ScheduleEndpoint = "https://esports-api.lolesports.com/persisted/gw/getSchedule"
	LeagueId         = "98767991332355509"
)

type Application struct {
	apiKey string
	lang   string
}

var onlyNumbers = regexp.MustCompile("[^0-9]")

func (a *Application) GetWeekMatches() string {
	req, err := http.NewRequest(http.MethodGet, ScheduleEndpoint, nil)

	if err != nil {
		fmt.Println(err)
		return CouldNotGetWeekMatches
	}

	req.Header.Add("x-api-key", a.apiKey)

	query := req.URL.Query()
	query.Add("hl", a.lang)
	query.Add("leagueId", LeagueId)

	req.URL.RawQuery = query.Encode()

	resp, err := http.DefaultClient.Do(req)

	body, err := io.ReadAll(resp.Body)

	var scheduleData DataDto
	err = json.Unmarshal(body, &scheduleData)

	if err != nil {
		fmt.Println(err)
		return CouldNotGetWeekMatches
	}

	var matches []*match.Match

	events := scheduleData.Data.Schedule.Events

	location, _ := time.LoadLocation("America/Sao_Paulo")

	for _, event := range events {
		startTime, err := time.ParseInLocation(time.RFC3339, event.StartTime, location)

		if err != nil {
			fmt.Println(err)
			return CouldNotGetWeekMatches
		}

		weekNum, err := strconv.Atoi(onlyNumbers.ReplaceAllString(event.BlockName, ""))

		if err != nil {
			fmt.Println(err)
			continue
		}

		var teams []*team.Team
		var winner *team.Team

		for _, t := range event.Match.Teams {
			matchTeam := team.New(t.Name, t.Record.Wins, t.Record.Losses)
			teams = append(teams, matchTeam)
			if t.Result.Outcome == Win {
				winner = matchTeam
			}
		}

		matchState := match.Unstarted
		if event.State == Completed {
			matchState = match.Completed
		}

		matches = append(matches, match.New(&startTime, weekNum, matchState, teams[0], teams[1], winner))
	}

	weekService := week.New(matches)

	return weekService.MatchesPerDay()
}

func New(apiKey string, lang string) *Application {
	return &Application{apiKey, lang}
}
