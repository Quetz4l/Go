package logic

import (
	"time"
)

func StringToTodoTime(date string) (time.Time, error) {
	format := time.RFC3339
	return time.Parse(format, date)
}
