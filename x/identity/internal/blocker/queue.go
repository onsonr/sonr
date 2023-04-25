package blocker

import (
	"context"
	"sync"
)

// Queue holds name, list of jobs and context with cancel.
type Queue struct {
	name   string
	jobs   chan Job
	ctx    context.Context
	cancel context.CancelFunc
}

// Job - holds fmtic to perform some operations during queue execution.
type Job struct {
	Name   string
	Action func() error // A function that should be executed when the job is running.
}

// NewQueue instantiates new queue.
func NewQueue(name string) *Queue {
	ctx, cancel := context.WithCancel(context.Background())

	return &Queue{
		jobs:   make(chan Job),
		name:   name,
		ctx:    ctx,
		cancel: cancel,
	}
}

// PendingJobs returns number of pending jobs.
func (q *Queue) PendingJobs() int {
	return len(q.jobs)
}

// AddJobs adds jobs to the queue and cancels channel.
func (q *Queue) AddJobs(jobs []Job) {
	var wg sync.WaitGroup
	wg.Add(len(jobs))

	for _, job := range jobs {
		// Goroutine which adds job to the queue.
		go func(job Job) {
			q.AddJob(job)
			wg.Done()
		}(job)
	}

	go func() {
		wg.Wait()
		// Cancel queue channel, when all goroutines were done.
		q.cancel()
	}()
}

// AddJob sends job to the channel.
func (q *Queue) AddJob(job Job) {
	q.jobs <- job
}

// Run performs job execution.
func (j Job) Run() error {
	err := j.Action()
	if err != nil {
		return err
	}

	return nil
}
