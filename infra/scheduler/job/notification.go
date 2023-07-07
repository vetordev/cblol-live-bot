package job

type ScheduleNotification struct {
}

func (a *ScheduleNotification) Run() {

}

func NewScheduleNotification() *ScheduleNotification {
	return &ScheduleNotification{}
}
