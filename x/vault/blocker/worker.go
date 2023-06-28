package blocker

import "fmt"

// Worker responsible for queue serving.
type Worker struct {
	Queue *Queue
}

// NewWorker initializes a new Worker.
func NewWorker(queue *Queue) *Worker {
	return &Worker{
		Queue: queue,
	}
}

// DoWork processes jobs from the queue (jobs channel).
func (w *Worker) DoWork(errCh chan error) bool {
	for {
		select {
		// if context was canceled.
		case <-w.Queue.ctx.Done():
			return true
		// if job received.
		case job := <-w.Queue.jobs:
			err := job.Run()
			if err != nil {
				errCh <- fmt.Errorf("job %s failed: %w", job.Name, err)
				continue
			}
		}
	}
}
