package match

import (
	"cblol-bot/domain/model/match"
	"cblol-bot/domain/model/team"
	matchsvc "cblol-bot/domain/service/match"
	"cblol-bot/domain/service/week"
	"cblol-bot/util/date"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const NoGames = "Não há jogos hoje!"

const (
	ScheduleEndpoint = "https://esports-api.lolesports.com/persisted/gw/getSchedule"
	LeagueId         = "98767991332355509"
)

type Application struct {
	apiKey string
	lang   string
}

func (a *Application) ListMatchesFromAPI() ([]*match.Match, error) {

	req, err := http.NewRequest(http.MethodGet, ScheduleEndpoint, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("x-api-key", a.apiKey)

	query := req.URL.Query()
	query.Add("hl", a.lang)
	query.Add("leagueId", LeagueId)

	req.URL.RawQuery = query.Encode()

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var scheduleData DataDto
	err = json.Unmarshal(body, &scheduleData)

	if err != nil {
		return nil, err
	}

	var matches []*match.Match

	events := scheduleData.Data.Schedule.Events

	location, _ := time.LoadLocation("America/Sao_Paulo")

	for _, event := range events {
		if event.Type != Match {
			continue
		}

		startTime, err := time.Parse(time.RFC3339, event.StartTime)
		startTime = startTime.In(location)

		if err != nil {
			fmt.Println(err)
			return nil, err
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

		matches = append(matches, match.New(startTime, event.BlockName, matchState, teams[0], teams[1], winner))
	}

	return matches, nil
}

func (a *Application) ListTodayMatchesFromAPI() ([]*match.Match, error) {
	matches, err := a.ListMatchesFromAPI()

	if err != nil {
		return nil, err
	}

	var todayMatches []*match.Match

	location, _ := time.LoadLocation("America/Sao_Paulo")
	today := date.ResetHours(time.Now().In(location))

	for _, m := range matches {
		if date.ResetHours(m.Schedule).Equal(today) {
			todayMatches = append(todayMatches, m)
		}
	}

	return todayMatches, nil
}

func (a *Application) GetWeekMatches() string {

	matches, err := a.ListMatchesFromAPI()

	if err != nil {
		fmt.Println(err)
		return CouldNotGetWeekMatches
	}

	weekService := week.New()

	return weekService.FormatWeekMatches(matches)
}

func (a *Application) GetTodayMatches() string {

	matches, err := a.ListTodayMatchesFromAPI()

	if err != nil {
		fmt.Println(err)
		return CouldNotGetTodayMatches
	}

	if len(matches) == 0 {
		return NoGames
	}

	return matchsvc.FormatMatchesPerDay(time.Now(), matches)
}

func New(apiKey string, lang string) *Application {
	return &Application{apiKey, lang}
}
