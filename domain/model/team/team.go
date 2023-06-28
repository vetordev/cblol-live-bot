package team

type Team struct {
	Name   string
	Wins   int
	Losses int
}

func New(name string, wins int, losses int) *Team {
	return &Team{name, wins, losses}
}
