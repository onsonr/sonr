package common

import "errors"

var (
	ErrRequestBody = errors.New("Provided HTTP request body is empty or invalid.")

	// Path Manipulation Errors
	ErrDuplicateFilePathOption    = errors.New("Duplicate file path option")
	ErrPrefixSuffixSetWithReplace = errors.New("Prefix or Suffix set with Replace.")
	ErrSeparatorLength            = errors.New("Separator length must be 1.")
	ErrNoFileNameSet              = errors.New("File name was not set by options.")

	// Device ID Errors
	ErrEmptyDeviceID = errors.New("Device ID cannot be empty")
	ErrMissingEnvVar = errors.New("Cannot set EnvVariable with empty value")

	// Directory errors
	ErrDirectoryInvalid = errors.New("Directory Type is invalid")
	ErrDirectoryUnset   = errors.New("Directory path has not been set")
	ErrDirectoryJoin    = errors.New("Failed to join directory path")

	// Node Errors
	ErrEmptyQueue       = errors.New("No items in Transfer Queue.")
	ErrInvalidQuery     = errors.New("No SName or PeerID provided.")
	ErrMissingParam     = errors.New("Paramater is missing.")
	ErrProtocolsNotSet  = errors.New("Node Protocol has not been initialized.")
	ErrRoutingNotSet    = errors.New("DHT and Host have not been set by Routing Function")
	ErrListenerRequired = errors.New("Listener was not Provided")
	ErrMDNSInvalidConn  = errors.New("Invalid Connection, cannot begin MDNS Service")
	ErrMotorWallet      = errors.New("Motor Wallet is not initialized")
)
