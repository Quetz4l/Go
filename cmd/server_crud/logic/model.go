package logic

import "time"

type Todo struct {
	Id        int       `json:"id"`
	Task      string    `json:"task"`
	DueDate   time.Time `json:"due_date"`
	Priority  string    `json:"priority"`
	Completed bool      `json:"completed"`
}
