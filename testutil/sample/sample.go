package sample

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/pkg/did"
)

// AccAddress returns a sample account address
func AccAddress() string {
	pk := ed25519.GenPrivKey().PubKey()
	addr := pk.Address()
	return sdk.AccAddress(addr).String()
}

// Mock Credential object from webauthn test bench https://github.com/psteniusubi/webauthn-tester
func CreateMockCredential() *did.Credential {
	return &did.Credential{
		ID:              []byte("ktIQAlFosR9OMGnyJnGthmKcIodPb323F3UqPVe9kvB-eOYrE-pNchsSuiN4ZE0ICyAaRiCb6vfF-7Y5nrvcoD-D42KQsXzhJd14ciqzibA"),
		AttestationType: "platform",
		PublicKey:       []byte("o2NmbXRkbm9uZWdhdHRTdG10oGhhdXRoRGF0YVjULNeTz6C0GMu_DqhSIoYH2el7Mz1NsKQQF3Zq9ruMdVFFAAAAAK3OAAI1vMYKZIsLJfHwVQMAUJLSEAJRaLEfTjBp8iZxrYZinCKHT299txd1Kj1XvZLwfnjmKxPqTXIbErojeGRNCAsgGkYgm"),
		Authenticator: did.Authenticator{
			AAGUID:    []byte("eyJ0eXBlIjoid2ViYXV0aG4uY3JlYXRlIiwiY2hhbGxlbmdlIjoiOHhBM2t3dUVCM0xtc2UxMkJyT2FrSDlZUWlrIiwib3JpZ2luIjoiaHR0cHM6Ly9wc3Rlbml1c3ViaS5naXRodWIuaW8iLCJjcm9zc09yaWdpbiI6ZmFsc2V9"),
			SignCount: 1,
		},
	}
}
