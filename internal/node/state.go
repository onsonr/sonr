package node

import "github.com/sonr-io/core/tools/state"

const (
	// ### - Supply Method Types -
	// 1. States
	// SupplyInvalid is state for when Supply doesnt contain any paths
	SupplyInvalid state.StateType = "SupplyInvalid"

	// AnalyzeFailed is state for when Analyze Fails
	AnalyzeFailed state.StateType = "AnalyzeFailed"

	// Queued is state for when Analyze Succeeds and no PeerID provided
	Queued state.StateType = "Queued"

	// 2. Events
	// Supply Method Events
	FailSupply    state.EventType = "FailSupply"
	SucceedSupply state.EventType = "SucceedSupply"

	// Analyze Method Events
	FailAnalyze    state.EventType = "FailAnalyze"
	SucceedAnalyze state.EventType = "SucceedAnalyze"
)
