package contract

import (
	"time"
)

type Task struct {
	UUID  string `json:"uuid"`
	Value string `json:"value"`
}

type DelayedTask interface {
	AddDelayed(value Task, delay time.Duration)
	GetReadyTasks() []Task
	TemoveTasks(tasks []Task)
}
