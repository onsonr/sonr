package motor

import (
	"encoding/json"

	"github.com/sonr-io/sonr/internal/motor/x/registry"
	prt "go.buf.build/grpc/go/sonr-io/motor/registry/v1"
)

func CreateAccount(requestBytes []byte) (prt.CreateAccountResponse, error) {
	var request prt.CreateAccountRequest
	if err := json.Unmarshal(requestBytes, &request); err != nil {
		return prt.CreateAccountResponse{}, err
	}

	m, err := registry.CreateAccount(request)
	if err != nil {
		return prt.CreateAccountResponse{}, err
	}
	return prt.CreateAccountResponse{
		Address: m.Address,
		AesPsk:  []byte(""),
	}, nil
}
