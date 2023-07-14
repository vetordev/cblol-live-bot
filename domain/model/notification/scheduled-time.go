package notification

import (
	"errors"
	"strings"
	"time"
)

var InvalidScheduledTime = errors.New("invalid scheduled time")

const DefaultScheduledTime = "10:00:00"

func ValidateScheduledTime(schedule string) (string, error) {

	if schedule == "" {
		return DefaultScheduledTime, nil
	}

	split := strings.Split(schedule, ":")

	if len(split) == 1 {
		schedule += ":00:00"
	}

	if len(split) == 2 {
		schedule += ":00"
	}

	if _, err := time.Parse(time.TimeOnly, schedule); err != nil {
		return "", InvalidScheduledTime
	}

	return schedule, nil
}
