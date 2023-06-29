package date

import "time"

func GetWeekDayInPt(weekday time.Weekday) string {
	switch weekday {
	case time.Sunday:
		return "Domingo"
	case time.Monday:
		return "Segunda"
	case time.Tuesday:
		return "Terça"
	case time.Wednesday:
		return "Quarta"
	case time.Thursday:
		return "Quinta"
	case time.Friday:
		return "Sexta"
	case time.Saturday:
		return "Sábado"
	default:
		return ""
	}
}
