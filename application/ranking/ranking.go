package rankingapp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const StandingsEndpoint = "https://esports-api.lolesports.com/persisted/gw/getStandings"
const RegularSeasonId = "110413046183015975"

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

	fmt.Println(string(body))

	var standingsData StandingsDataDto
	err = json.Unmarshal(body, &standingsData)

	if err != nil {
		fmt.Println(err)
		return CouldNotGetRankings
	}
	fmt.Println(standingsData)
	return "OK!"
}

func New(apiKey string, lang string) *Application {
	return &Application{apiKey, lang}
}
