package week

type Outcome string
type MatchStateDto string

const (
	Unstarted MatchStateDto = "unstarted"
	Completed MatchStateDto = "completed"
)

const (
	Loss Outcome = "loss"
	Win  Outcome = "win"
)

type DataDto struct {
	Data struct {
		Schedule ScheduleDto `json:"schedule"`
	} `json:"data"`
}

type ScheduleDto struct {
	Events []EventsDto `json:"events"`
}

type EventsDto struct {
	StartTime string        `json:"startTime"`
	State     MatchStateDto `json:"state"`
	Match     MatchDto      `json:"match"`
	BlockName string        `json:"blockName"`
}

type MatchDto struct {
	Teams []TeamDto `json:"teams"`
}

type TeamDto struct {
	Name   string        `json:"name"`
	Result TeamResultDto `json:"result"`
	Record TeamRecordDto `json:"record"`
}

type TeamResultDto struct {
	Outcome Outcome `json:"outcome"`
}

type TeamRecordDto struct {
	Losses int `json:"losses"`
	Wins   int `json:"wins"`
}
