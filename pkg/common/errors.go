package common

import "errors"

var (
	ErrRequestBody = errors.New("provided HTTP request body is empty or invalid")

	// Path Manipulation Errors
	ErrDuplicateFilePathOption    = errors.New("duplicate file path option")
	ErrPrefixSuffixSetWithReplace = errors.New("prefix or Suffix set with Replace")
	ErrSeparatorLength            = errors.New("separator length must be 1")
	ErrNoFileNameSet              = errors.New("file name was not set by options")

	// Device ID Errors
	ErrEmptyDeviceID = errors.New("device ID cannot be empty")
	ErrMissingEnvVar = errors.New("cannot set EnvVariable with empty value")

	// Directory errors
	ErrDirectoryInvalid = errors.New("directory Type is invalid")
	ErrDirectoryUnset   = errors.New("directory path has not been set")
	ErrDirectoryJoin    = errors.New("failed to join directory path")

	// Node Errors
	ErrEmptyQueue                = errors.New("no items in Transfer Queue")
	ErrInvalidQuery              = errors.New("no SName or PeerID provided")
	ErrMissingParam              = errors.New("paramater is missing")
	ErrProtocolsNotSet           = errors.New("node Protocol has not been initialized")
	ErrRoutingNotSet             = errors.New("DHT and Host have not been set by Routing Function")
	ErrListenerRequired          = errors.New("listener was not Provided")
	ErrMDNSInvalidConn           = errors.New("invalid Connection, cannot begin MDNS Service")
	ErrMotorWalletNotInitialized = errors.New("motor Wallet is not initialized")
	ErrDefaultStillImplemented   = errors.New("callback not implemented")
)
