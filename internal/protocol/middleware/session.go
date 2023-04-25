package middleware

import "github.com/sonrhq/core/x/service/types"

type WebAuthentication interface {
	// StartRegistration returns the credential creation options as a JSON string

}

type webAuthentication struct {
	serviceRecord *types.ServiceRecord
}
