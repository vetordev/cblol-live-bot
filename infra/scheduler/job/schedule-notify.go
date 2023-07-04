package job

type Alert struct {
}

func (a *Alert) Run() {

}

func NewAlert() *Alert {
	return &Alert{}
}
