package tasks

import "time"

type Task struct {
	Id          string
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
}

func NewTask(title, description string) Task {
	return Task{
		Id:          0,
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}
}
