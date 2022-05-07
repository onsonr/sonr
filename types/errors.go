package types

import "errors"

var (
	ErrRequestBody = errors.New("Provided HTTP request body is empty or invalid.")
)
