package gcron

import (
	"errors"
	"fmt"
	"sync"
)

var (
	ErrAlreadyExists  = errors.New("already exists")
	ErrAlreadyStarted = errors.New("already started")
	ErrNotFound       = errors.New("not found")
	ErrNotStarted     = errors.New("not started")
)

func NewScheduler() Scheduler {

	return scheduler{
		jobs:       map[string]job{},
		jobsLocker: &sync.RWMutex{},
	}
}

type Scheduler interface {
	AddJob(name string, schedule string) error
	StartJob(name string) error
	StopJob(name string) error
}

type scheduler struct {
	jobs       map[string]job
	jobsLocker *sync.RWMutex
}

// AddJob implements Scheduler.
func (s scheduler) AddJob(name string, schedule string) error {
	s.jobsLocker.Lock()
	defer s.jobsLocker.Unlock()
	if _, found := s.jobs[name]; found {
		return fmt.Errorf("job %s: %w", name, ErrAlreadyExists)
	} else {
		s.jobs[name] = job{}
		return nil
	}
}

// StartJob implements Scheduler.
func (s scheduler) StartJob(name string) error {
	s.jobsLocker.Lock()
	defer s.jobsLocker.Unlock()
	if data, found := s.jobs[name]; !found {
		return fmt.Errorf("job %s: %w", name, ErrNotFound)
	} else if data.started {
		return fmt.Errorf("job %s: %w", name, ErrAlreadyStarted)
	} else {
		data.started = true
		s.jobs[name] = data
		return nil
	}
}

// StopJob implements Scheduler.
func (s scheduler) StopJob(name string) error {
	s.jobsLocker.Lock()
	defer s.jobsLocker.Unlock()
	if data, found := s.jobs[name]; !found {
		return fmt.Errorf("job %s: %w", name, ErrNotFound)
	} else if !data.started {
		return fmt.Errorf("job %s: %w", name, ErrNotStarted)
	} else {
		data.started = false
		s.jobs[name] = data
		return nil
	}
}

type job struct {
	started bool
}
