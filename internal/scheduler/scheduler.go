package scheduler

import (
	"fmt"
	"sync"

	pkg_errors "github.com/jakewan/gcron/pkg/errors"
)

func NewScheduler() Scheduler {
	return Scheduler{
		jobs:       map[string]job{},
		jobsLocker: &sync.RWMutex{},
	}
}

type Scheduler struct {
	jobs       map[string]job
	jobsLocker *sync.RWMutex
}

// AddJob implements gcron.Scheduler.
func (s Scheduler) AddJob(name string, schedule string) error {
	s.jobsLocker.Lock()
	defer s.jobsLocker.Unlock()
	if _, found := s.jobs[name]; found {
		return fmt.Errorf("job %s: %w", name, pkg_errors.ErrAlreadyExists)
	} else {
		s.jobs[name] = job{}
		return nil
	}
}

// StartJob implements gcron.Scheduler.
func (s Scheduler) StartJob(name string) error {
	s.jobsLocker.Lock()
	defer s.jobsLocker.Unlock()
	if data, found := s.jobs[name]; !found {
		return fmt.Errorf("job %s: %w", name, pkg_errors.ErrNotFound)
	} else if data.started {
		return fmt.Errorf("job %s: %w", name, pkg_errors.ErrAlreadyStarted)
	} else {
		data.started = true
		s.jobs[name] = data
		return nil
	}
}

// StopJob implements gcron.Scheduler.
func (s Scheduler) StopJob(name string) error {
	s.jobsLocker.Lock()
	defer s.jobsLocker.Unlock()
	if data, found := s.jobs[name]; !found {
		return fmt.Errorf("job %s: %w", name, pkg_errors.ErrNotFound)
	} else if !data.started {
		return fmt.Errorf("job %s: %w", name, pkg_errors.ErrNotStarted)
	} else {
		data.started = false
		s.jobs[name] = data
		return nil
	}
}
