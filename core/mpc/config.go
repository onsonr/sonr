package mpc

import (
	"time"
)

// Option is a function that can be applied to ExchangeProtocol config
type Option func(*options)

// options for ExchangeProtocol config
type options struct {
	// location        *types.Location
	interval        time.Duration
	olcCode         string
	autoPushEnabled bool
}

// defaultOptions for ExchangeProtocol config
func defaultOptions() *options {
	return &options{
		//location:        api.DefaultLocation(),
		interval:        time.Second * 5,
		autoPushEnabled: true,
	}
}
