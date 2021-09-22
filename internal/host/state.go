package host

import "github.com/sonr-io/core/tools/state"

const (
	// ### - Connect Method Types -
	// 1. States
	// ConnectionFailed is state for when Connect Fails
	ConnectionFailed state.StateType = "ConnectionFailed"

	// BootstrapFailed is state for when Bootstrap Fails
	BootstrapFailed state.StateType = "BootstrapFailed"

	// JoinNetworkFailed is state for when Bootstrap Fails
	JoinNetworkFailed state.StateType = "JoinNetworkFailed"

	// Connected is state for when Connect Succeeds
	Connected state.StateType = "Connected"

	// 2. Events
	// Connect Method Events
	FailConnect    state.EventType = "FailConnect"
	SucceedConnect state.EventType = "SucceedConnect"

	// Bootstrap Method Events
	FailBootstrap    state.EventType = "FailBootstrap"
	SucceedBootstrap state.EventType = "SucceedBootstrap"
)
