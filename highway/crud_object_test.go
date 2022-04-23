package highway

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	ot "go.buf.build/grpc/go/sonr-io/blockchain/object"
	"go.buf.build/grpc/go/sonr-io/blockchain/registry"
)

func TestMsgCreateObject_ValidateBasic(t *testing.T) {
	ctx := context.Background()
	highway, err := NewHighway(ctx)
	require.NoError(t, err)

	tests := []struct {
		name     string
		msg      ot.MsgCreateObject
		expected ot.MsgCreateObjectResponse
		err      error
	}{
		{
			name: "invalid object",
			msg:  ot.MsgCreateObject{},
			err:  errors.New("object to register must have fields"),
		}, {
			name: "valid object",
			msg: ot.MsgCreateObject{
				Creator:     "snr1ulu0a0eew3w3sj8nk8lx5z3cmaw46u383kek9t",
				Label:       "a label",
				Description: "a description",
				InitialFields: []*ot.TypeField{
					{
						Name: "another label",
						Kind: 0,
					},
				},
				Session: &registry.Session{
					Whois: &registry.WhoIs{
						Name: "alice",
						Type: 1,
						Credentials: []*registry.Credential{
							{
								PublicKey: []byte("ufJWp8YGlibm1Kd9XQBWN1WAw2jy5In2Xhon9HAqcXE="),
								Authenticator: &registry.Authenticator{
									Aaguid:       []byte("test"),
									SignCount:    2,
									CloneWarning: false,
								},
							},
						},
						Metadata: map[string]string{
							"user": "alice",
						},
					},
					Credential: &registry.Credential{
						Authenticator: &registry.Authenticator{
							Aaguid:       []byte("test"),
							SignCount:    2,
							CloneWarning: false,
						},
					},
				},
			},
			expected: ot.MsgCreateObjectResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			resp, err := highway.CreateObject(ctx, &tt.msg)
			if tt.err != nil {
				require.Error(t, err, tt.err)
				return
			}

			require.ElementsMatch(t, resp, tt.expected)

			require.NoError(t, err)
		})
	}
}
