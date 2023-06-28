package team

type Team struct {
	Name  string
	Wins  int
	Loses int
}

func New(name string, wins int, loses int) *Team {
	return &Team{name, wins, loses}
}
