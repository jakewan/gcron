package gcron

import (
	"github.com/jakewan/gcron/internal/scheduler"
)

type Scheduler interface {
	AddJob(name string, schedule string) error
	StartJob(name string) error
	StartScheduler() error
	StopJob(name string) error
	StopScheduler() error
}

func NewScheduler() Scheduler {
	return scheduler.NewScheduler()
}
