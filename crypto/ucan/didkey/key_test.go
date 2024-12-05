package didkey

import (
	"testing"
)

func TestID(t *testing.T) {
	keyStrED := "did:key:z6MkpTHR8VNsBxYAAWHut2Geadd9jSwuBV8xRoAnwWsdvktH"
	id, err := Parse(keyStrED)
	if err != nil {
		t.Fatal(err)
	}

	if id.String() != keyStrED {
		t.Errorf("string mismatch.\nwant: %q\ngot:  %q", keyStrED, id.String())
	}

	keyStrRSA := "did:key:z2MGw4gk84USotaWf4AkJ83DcnrfgGaceF86KQXRYMfQ7xqnUFp38UZ6Le8JPfkb4uCLGjHBzKpjEXb9hx9n2ftecQWCHXKtKszkke4FmENdTZ7i9sqRmL3pLnEEJ774r3HMuuC7tNRQ6pqzrxatXx2WinCibdhUmvh3FobnA9ygeqkSGtV6WLa7NVFw9cAvnv8Y6oHcaoZK7fNP4ASGs6AHmSC6ydSR676aKYMe95QmEAj4xJptDsSxG7zLAGzAdwCgm56M4fTno8GdWNmU6Pdghnuf6fWyYus9ASwdfwyaf3SDf4uo5T16PRJssHkQh6DJHfK4Rka7RNQLjzfGBPjFLHbUSvmf4EdbHasbVaveAArD68ZfazRCCvjdovQjWr6uyLCwSAQLPUFZBTT8mW"

	id, err = Parse(keyStrRSA)
	if err != nil {
		t.Fatal(err)
	}

	if id.String() != keyStrRSA {
		t.Errorf("string mismatch.\nwant: %q\ngot:  %q", keyStrRSA, id.String())
	}
}
