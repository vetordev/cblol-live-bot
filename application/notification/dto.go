package notification

import (
	"cblol-bot/domain/model/notification"
	"strings"
	"time"
)

func toSchedulingFormat(schedule string) string {

	if schedule == "" {
		return notification.DefaultScheduledTime
	}

	split := strings.Split(schedule, ":")

	if len(split) == 1 {
		schedule += ":00:00"
	}

	if len(split) == 2 {
		schedule += ":00"
	}

	if _, err := time.Parse(time.TimeOnly, schedule); err != nil {
		return ""
	}

	return schedule
}
