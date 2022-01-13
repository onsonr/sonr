package did_test

import (
	"crypto/rand"
	"log"
	"testing"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/sonr-io/core/did"

	b "github.com/sonr-io/core/did"
)

func TestCreateBaseDid(t *testing.T) {
	cases := []struct {
		valid   bool
		id      string
		options []b.Option
	}{
		{
			true,
			"eb9a70f4",
			[]b.Option{},
		},
		{
			true,
			"e89a6990",
			[]b.Option{
				b.WithNetwork("testnet"),
			},
		},
		{
			true,
			"ee4718ac",
			[]b.Option{
				b.WithNetwork("mainnet"),
				b.WithPath("/service1"),
			},
		},
		{
			true,
			"07264884",
			[]b.Option{
				b.WithNetwork("testnet"),
				b.WithPath("/service2"),
				b.WithQuery("?test=1"),
			},
		},
		{
			true,
			"18f797d4",
			[]b.Option{
				b.WithNetwork("mainnet"),
				b.WithPath("/service3/module1"),
				b.WithFragment("#channel1"),
			},
		},
		{
			false,
			"1ceafe1a",
			[]b.Option{
				b.WithNetwork("testnet"),
				b.WithPath("/service4/module2/submodule1"),
				b.WithQuery("?test=2"),
				b.WithFragment("#channel2"),
			},
		},
	}

	for i := 0; i < len(cases); i++ {
		_, pubKey, err := crypto.GenerateEd25519Key(rand.Reader)
		if err != nil {
			t.Errorf("failed to generate key: %v", err)
		}

		did, pbstr, err := did.CreateBaseDID(pubKey)
		if err != nil {
			t.Errorf("failed to create did: %v", err)
		}

		log.Printf("did: %s", did.ToString())
		cases[i].id = pbstr
	}

	for i, tc := range cases {
		res, err := b.NewDID(tc.id, tc.options...)
		if err != nil && tc.valid {
			t.Errorf("NewDID(%s) error: %s", tc.id, err)
		} else if err == nil && !tc.valid {
			t.Errorf("NewDID(%s) should be invalid", tc.id)
		} else {
			if tc.valid {
				t.Logf("%s", res)
				log.Println(res.String())
			} else {
				log.Printf("Succesfully identified invalid did at Index: (%d)", i)
			}

		}
	}
}
