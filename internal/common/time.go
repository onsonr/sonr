package common

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

// ErrFunc is a function that returns an error
type ErrFunc func() error

// FuncOption is a timed-function option
type FuncOption func(funcOptions)

// WithRequest sets the initialize request.
func WithInterval(i time.Duration) FuncOption {
	return func(o funcOptions) {
		o.interval = i
	}
}

// WithRequest sets the initialize request.
func WithMaxErrors(me int) FuncOption {
	return func(o funcOptions) {
		o.maxErrors = me
	}
}

// WithRequest sets the initialize request.
func WithRetry(times int, interval time.Duration) FuncOption {
	return func(o funcOptions) {
		o.retryCount = times
		o.retryInterval = interval
		o.hasRetry = true
	}
}

// WithRequest sets the initialize request.
func WithTimeout(to time.Duration) FuncOption {
	return func(o funcOptions) {
		o.timeout = to
	}
}

// funcOptions is a struct that holds the options for a timed function
type funcOptions struct {
	hasRetry      bool
	interval      time.Duration
	maxErrors     int
	retryCount    int
	retryInterval time.Duration
	timeout       time.Duration
}

// defaultNodeOptions returns the default node options.
func defaultFuncOptions() funcOptions {
	return funcOptions{
		hasRetry:      false,
		maxErrors:     1,
		retryCount:    0,
		interval:      time.Second * 4,
		retryInterval: time.Second * 8,
		timeout:       time.Minute * 30,
	}
}

// NewRetryFunc creates a new retry function
func NewRetryFunc(f ErrFunc, retries int, interval time.Duration) ErrFunc {
	return func() error {
		var err error
		for i := 0; i < retries; i++ {
			err = f()
			if err == nil {
				return nil
			}
			logger.Warn(fmt.Sprintf("(%v/%v) Retrying...", i, retries), err)
			time.Sleep(interval)
		}
		return errors.Wrap(err, fmt.Sprintf("Failed to call method after %v attempts", retries))
	}
}

// NewPeriodicFunc creates a new ticker function with a given interval
func NewPeriodicFunc(ctx context.Context, f ErrFunc, options ...FuncOption) (func(), chan error) {
	// Set Function Options
	opts := defaultFuncOptions()
	for _, opt := range options {
		opt(opts)
	}

	// Set Caller Function
	var caller ErrFunc
	if opts.hasRetry {
		caller = NewRetryFunc(f, opts.retryCount, opts.retryInterval)
	} else {
		caller = f
	}

	// Create channel
	ctxTimeout, cancel := context.WithTimeout(ctx, opts.timeout)
	ticker := time.NewTicker(opts.interval)
	errChan := make(chan error, 1)

	// Return Channel Function
	return func() {
		for errCount := 0; errCount < opts.maxErrors; {
			select {
			// Tick
			case <-ticker.C:
				// Call function
				err := caller()
				if err != nil {
					errCount++
					errChan <- err
				}

				// Check if we exceeded the max errors
				if errCount >= opts.maxErrors {
					logger.Error(fmt.Sprintf("Exceeded Maximum Errors (%v), closing channel", opts.maxErrors), err)
					cancel()
					errChan <- errors.Wrap(err, "exceeded max errors")
					break
				}

			case <-ctxTimeout.Done():
				ticker.Stop()
				break
			}
		}
	}, errChan
}
