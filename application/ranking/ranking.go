package ranking

import (
	ranking "cblol-bot/domain/model/ranking"
	team "cblol-bot/domain/model/team"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	StandingsEndpoint = "https://esports-api.lolesports.com/persisted/gw/getStandings"
	RegularSeasonId   = "110413046183015975"
)

type Application struct {
	apiKey string
	lang   string
}

func (a *Application) GetRanking() string {
	req, err := http.NewRequest(http.MethodGet, StandingsEndpoint, nil)

	if err != nil {
		fmt.Println(err)
		return CouldNotGetRankings
	}

	req.Header.Add("x-api-key", a.apiKey)

	query := req.URL.Query()
	query.Add("hl", a.lang)
	query.Add("tournamentId", RegularSeasonId)

	req.URL.RawQuery = query.Encode()

	resp, err := http.DefaultClient.Do(req)

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return CouldNotGetRankings
	}

	var standingsData StandingsDataDto
	err = json.Unmarshal(body, &standingsData)

	if err != nil {
		fmt.Println(err)
		return CouldNotGetRankings
	}

	rankingsDto := standingsData.Data.Standings[0].Stages[0].Sections[0].Rankings

	var teams []*team.Team

	for _, rankingDto := range rankingsDto {
		teamDto := rankingDto.Teams[0]
		teams = append(teams, team.New(teamDto.Name, teamDto.Record.Wins, teamDto.Record.Losses))
	}

	ranking := ranking.New(teams)

	return ranking.Format()
}

func New(apiKey string, lang string) *Application {
	return &Application{apiKey, lang}
}
