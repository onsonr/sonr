package api

import common "github.com/sonr-io/core/common"

// Option is a function that modifies the node options.
type Option func(*Config)

// WithRequest sets the initialize request.
func WithRequest(req *InitializeRequest) Option {
	return func(o *Config) {
		o.location = req.GetLocation()
		o.profile = req.GetProfile()
		o.connection = req.GetConnection()
	}
}

// WithMode starts the Client RPC server as a highway node.
func WithMode(m StubMode) Option {
	return func(o *Config) {
		o.mode = m
	}
}

// Config is a collection of Config for the node.
type Config struct {
	connection common.Connection
	location   *common.Location
	mode       StubMode
	profile    *common.Profile
}

// defaultNodeOptions returns the default node options.
func defaultNodeOptions() *Config {
	return &Config{
		mode:       StubMode_LIB,
		location:   DefaultLocation(),
		connection: common.Connection_WIFI,
		profile:    common.NewDefaultProfile(),
	}
}
