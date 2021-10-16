// Forked from: https://github.com/duanckham/hands
package state

import (
	"context"
	"errors"
	"log"
	"sort"
	"sync/atomic"
)

var (
	errTaskPanic = errors.New("task panic")
)

type taskOptions struct {
	name     string
	priority int32
}

type handOptions struct {
	all         bool
	fastest     bool
	betweenFrom int32
	betweenTo   int32
	in          []int32
	percentage  float32
	ctx         context.Context
}

type task struct {
	options *taskOptions
	f       func(ctx context.Context) error
}

type taskResult struct {
	err      error
	name     string
	priority int32
}

// TaskOption sets some metadata for the task.
type TaskOption func(options *taskOptions)

// Name is the task's name.
func Name(name string) TaskOption {
	return func(o *taskOptions) {
		o.name = name
	}
}

// Priority is the task's priority, high priority tasks will be
// executed first.
func Priority(priority int32) TaskOption {
	return func(o *taskOptions) {
		o.priority = priority
	}
}

// P is an alias for `.Priority()`.
func P(priority int32) TaskOption {
	return Priority(priority)
}

// HandOption configures decide on the execution strategy of the task.
type HandOption func(options *handOptions)

// Fastest option return result as fast as possible.
func Fastest() HandOption {
	return func(o *handOptions) {
		o.fastest = true
	}
}

// Percentage option return resutls when a certain percentage of tasks
// are completed.
func Percentage(percentage float32) HandOption {
	return func(o *handOptions) {
		o.percentage = percentage
	}
}

// Between option wait for the task in the priority interval to end.
func Between(l, r int32) HandOption {
	return func(o *handOptions) {
		o.betweenFrom = l
		o.betweenTo = r
	}
}

// In option wait for the task in the specified priority list.
func In(in []int32) HandOption {
	return func(o *handOptions) {
		o.in = in
	}
}

// WithContext method will make the tasks use the specified context.
func WithContext(ctx context.Context) HandOption {
	return func(o *handOptions) {
		o.ctx = ctx
	}
}

// HandController defines hand controller interface.
type HandController interface {
	Do(f func(ctx context.Context) error, options ...TaskOption) HandController
	Done(callback func())
	Run(options ...HandOption) error
	RunAll(options ...HandOption) error
}

type handImpl struct {
	tasks    []task
	size     int32
	waiting  int32
	running  bool
	options  *handOptions
	callback func()
}

// NewHands method create a new hands controller.
func NewHands() HandController {
	return &handImpl{
		options: &handOptions{
			betweenFrom: -1,
			betweenTo:   -1,
			ctx:         context.Background(),
		},
	}
}

// Do method create a new task into hands.
func (h *handImpl) Do(f func(ctx context.Context) error, options ...TaskOption) HandController {
	t := task{f: f, options: &taskOptions{}}

	for _, setter := range options {
		setter(t.options)
	}

	// If `priority` < 0, return directly.
	if t.options.priority < 0 {
		return h
	}

	h.addTask(t)
	return h
}

// Done method will be called when all tasks completed.
func (h *handImpl) Done(callback func()) {
	if h.running {
		log.Print("can not setting the `Done()` method after the `Run()` or `RunAll()`")
		return
	}

	h.callback = callback
}

// Run method will start all tasks.
func (h *handImpl) Run(options ...HandOption) error {
	h.running = true

	if len(h.tasks) == 0 {
		return nil
	}

	h.setOptions(options...)
	h.size = int32(len(h.tasks))
	h.waiting = h.size

	// `.In()`
	if len(h.options.in) > 0 {
		sort.Slice(h.options.in, func(i, j int) bool {
			return h.options.in[i] < h.options.in[j]
		})

		i, j, k := 0, 0, 0

		for i < len(h.tasks) && j < len(h.options.in) {
			if h.tasks[i].options.priority > h.options.in[j] {
				j++
				continue
			}

			if h.tasks[i].options.priority == h.options.in[j] {
				if i > k {
					h.tasks[i], h.tasks[k] = h.tasks[k], h.tasks[i]
				}

				k++
			}

			i++
		}

		h.updateWaitingCount(int32(k))

		err := h.runTasks(h.options.ctx, h.tasks[:k])
		if err != nil {
			return err
		}

		if h.options.all {
			go func() {
				err := h.runTasks(context.Background(), h.tasks[k:])
				if err != nil {
					log.Fatalf("an error occurred in a non-essential task running asynchronously: %v", err)
				}
			}()
		}

		return nil
	}

	// `.Between()`
	if h.options.betweenFrom > -1 && h.options.betweenTo > -1 {
		if h.options.betweenFrom > h.options.betweenTo {
			return errors.New("the left border cannot be larger than the right border")
		}

		l := search(h.tasks, h.options.betweenFrom)
		r := search(h.tasks, h.options.betweenTo)

		if l > 0 {
			l = l - 1
		}

		h.updateWaitingCount(int32(r - l))

		err := h.runTasks(h.options.ctx, h.tasks[l:r])
		if err != nil {
			return err
		}

		if h.options.all {
			go func() {
				err := h.runTasks(context.Background(), h.tasks[:l])
				if err != nil {
					log.Fatalf("an error occurred in a non-essential task running asynchronously: %v", err)
				}
			}()

			go func() {
				err := h.runTasks(context.Background(), h.tasks[r:])
				if err != nil {
					log.Fatalf("an error occurred in a non-essential task running asynchronously: %v", err)
				}
			}()
		}

		return nil
	}

	return h.runTasks(h.options.ctx, h.tasks)
}

// RunAll method will cause non-high priority tasks to be executed asynchronously.
func (h *handImpl) RunAll(options ...HandOption) error {
	h.options.all = true
	return h.Run(options...)
}

func (h *handImpl) addTask(t task) {
	i := search(h.tasks, t.options.priority)
	h.tasks = append(h.tasks, t)
	copy(h.tasks[i+1:], h.tasks[i:])
	h.tasks[i] = t
}

func (h *handImpl) runTasks(ctx context.Context, tasks []task) error {
	c := len(tasks)
	if c == 0 {
		return nil
	}

	results := make(chan taskResult, c)

	for i := range tasks {
		t := tasks[c-i-1]

		go func() {
			defer func() {
				if r := recover(); r != nil {
					errFromPanic, ok := r.(error)
					if !ok {
						errFromPanic = errTaskPanic
					}

					results <- taskResult{
						err:      errFromPanic,
						name:     t.options.name,
						priority: t.options.priority,
					}
				}
			}()

			err := t.f(ctx)

			results <- taskResult{
				err:      err,
				name:     t.options.name,
				priority: t.options.priority,
			}

			h.runTasksDone()
		}()
	}

	switch {
	case h.options.fastest:
		return h.processFastest(results)
	case h.options.percentage > 0:
		return h.processPercentage(c, results)
	default:
		return h.process(c, results)
	}
}

func (h *handImpl) runTasksDone() {
	if atomic.AddInt32(&h.waiting, -1) <= 0 {
		if h.callback != nil {
			h.callback()
		}
	}
}

func (h *handImpl) setOptions(options ...HandOption) {
	for _, setter := range options {
		setter(h.options)
	}
}

func (h *handImpl) updateWaitingCount(count int32) {
	if !h.options.all {
		h.waiting = count
	}
}

func (h *handImpl) process(c int, results chan taskResult) error {
	for i := 0; i < c; i++ {
		select {
		case <-h.options.ctx.Done():
			return h.options.ctx.Err()
		case r := <-results:
			if r.err != nil {
				return r.err
			}
		}
	}

	return nil
}

func (h *handImpl) processFastest(results chan taskResult) error {
	select {
	case <-h.options.ctx.Done():
		return h.options.ctx.Err()
	case r := <-results:
		if r.err != nil {
			return r.err
		}
		return nil
	}
}

func (h *handImpl) processPercentage(c int, results chan taskResult) error {
	cycles := int(float32(c) * h.options.percentage)

	for i := 0; i < cycles; i++ {
		select {
		case <-h.options.ctx.Done():
			return h.options.ctx.Err()
		case r := <-results:
			if r.err != nil {
				return r.err
			}
		}
	}

	return nil
}

// Utility method.

func search(tasks []task, priority int32) int {
	return sort.Search(len(tasks), func(i int) bool {
		return tasks[i].options.priority > priority
	})
}
