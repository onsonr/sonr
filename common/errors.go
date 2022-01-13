package common

import "errors"

var (
	ErrFragmentAndQuery = errors.New("Unable to create new DID. Fragment and Query are mutually exclusive")
	ErrParseInvalid     = errors.New("Unable to parse string into DID, invalid format.")
)
