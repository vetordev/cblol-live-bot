package rankingapp

type StandingsDataDto struct {
	Data struct {
		Standings []StandingsDto `json:"standings"`
	} `json:"data"`
}

type StandingsDto struct {
	Stages []StagesDto `json:"stages"`
}

type StagesDto struct {
	Sections []SectionsDto `json:"sections"`
}

type SectionsDto struct {
	Rankings []RankingDto `json:"rankings"`
}

type RankingDto struct {
	Teams []TeamDto `json:"teams"`
}

type TeamDto struct {
	Name   string `json:"name"`
	Record struct {
		Wins   int `json:"wins"`
		Losses int `json:"losses"`
	} `json:"record"`
}
