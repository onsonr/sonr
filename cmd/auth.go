package bind

import (
	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// @ Method Checks if SName is Valid
func (mn *Node) CheckSName(buf []byte) []byte {
	// Unmarshal Request
	req := &md.AuthenticationRequest{}
	err := proto.Unmarshal(buf, req)
	if err != nil {
		sentry.CaptureException(errors.Wrap(err, "Unmarshalling Authentication Request"))
	}

	// Handle SName Check
	resp := mn.auth.CheckSName(req)
	if resp != nil {
		// Callback Result
		bytes, err := proto.Marshal(resp)
		if err != nil {
			return bytes
		}
	}
	return nil
}

// @ Method Saves Given SName to Records
func (mn *Node) SaveSName(buf []byte) []byte {
	// Unmarshal Request
	req := &md.AuthenticationRequest{}
	err := proto.Unmarshal(buf, req)
	if err != nil {
		sentry.CaptureException(errors.Wrap(err, "Unmarshalling Authentication Request"))
	}

	// Handle SName Save
	resp := mn.auth.SaveSName(req)
	if resp != nil {
		// Callback Result
		bytes, err := proto.Marshal(resp)
		if err != nil {
			return bytes
		}
	}
	return nil
}
