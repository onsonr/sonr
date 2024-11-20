package macaroon_test

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	qt "github.com/frankban/quicktest"

	"github.com/onsonr/sonr/pkg/crypto/macaroon"
)

func never(string) error {
	return fmt.Errorf("condition is never true")
}

func TestNoCaveats(t *testing.T) {
	c := qt.New(t)
	rootKey := []byte("secret")
	m := MustNew(rootKey, []byte("some id"), "a location", macaroon.LatestVersion)
	c.Assert(m.Location(), qt.Equals, "a location")
	c.Assert(m.Id(), qt.DeepEquals, []byte("some id"))

	err := m.Verify(rootKey, never, nil)
	c.Assert(err, qt.IsNil)
}

func TestFirstPartyCaveat(t *testing.T) {
	c := qt.New(t)
	rootKey := []byte("secret")
	m := MustNew(rootKey, []byte("some id"), "a location", macaroon.LatestVersion)

	caveats := map[string]bool{
		"a caveat":       true,
		"another caveat": true,
	}
	tested := make(map[string]bool)

	for cav := range caveats {
		m.AddFirstPartyCaveat([]byte(cav))
	}
	expectErr := fmt.Errorf("condition not met")
	check := func(cav string) error {
		tested[cav] = true
		if caveats[cav] {
			return nil
		}
		return expectErr
	}
	err := m.Verify(rootKey, check, nil)
	c.Assert(err, qt.IsNil)

	c.Assert(tested, qt.DeepEquals, caveats)

	m.AddFirstPartyCaveat([]byte("not met"))
	err = m.Verify(rootKey, check, nil)
	c.Assert(err, qt.Equals, expectErr)

	c.Assert(tested["not met"], qt.Equals, true)
}

func TestThirdPartyCaveat(t *testing.T) {
	c := qt.New(t)
	rootKey := []byte("secret")
	m := MustNew(rootKey, []byte("some id"), "a location", macaroon.LatestVersion)

	dischargeRootKey := []byte("shared root key")
	thirdPartyCaveatId := []byte("3rd party caveat")
	err := m.AddThirdPartyCaveat(dischargeRootKey, thirdPartyCaveatId, "remote.com")
	c.Assert(err, qt.IsNil)

	dm := MustNew(dischargeRootKey, thirdPartyCaveatId, "remote location", macaroon.LatestVersion)
	dm.Bind(m.Signature())
	err = m.Verify(rootKey, never, []*macaroon.Macaroon{dm})
	c.Assert(err, qt.IsNil)
}

func TestThirdPartyCaveatBadRandom(t *testing.T) {
	c := qt.New(t)
	rootKey := []byte("secret")
	m := MustNew(rootKey, []byte("some id"), "a location", macaroon.LatestVersion)
	dischargeRootKey := []byte("shared root key")
	thirdPartyCaveatId := []byte("3rd party caveat")

	err := macaroon.AddThirdPartyCaveatWithRand(m, dischargeRootKey, thirdPartyCaveatId, "remote.com", &macaroon.ErrorReader{})
	c.Assert(err, qt.ErrorMatches, "cannot generate random bytes: fail")
}

func TestSetLocation(t *testing.T) {
	c := qt.New(t)
	rootKey := []byte("secret")
	m := MustNew(rootKey, []byte("some id"), "a location", macaroon.LatestVersion)
	c.Assert(m.Location(), qt.Equals, "a location")
	m.SetLocation("another location")
	c.Assert(m.Location(), qt.Equals, "another location")
}

var equalTests = []struct {
	about  string
	m1, m2 macaroonSpec
	expect bool
}{{
	about: "same_everywhere",
	m1: macaroonSpec{
		rootKey: "root-key",
		id:      "root-id",
		caveats: []caveat{{
			condition: "a",
		}, {
			condition: "d",
		}},
	},
	m2: macaroonSpec{
		rootKey: "root-key",
		id:      "root-id",
		caveats: []caveat{{
			condition: "a",
		}, {
			condition: "d",
		}},
	},
	expect: true,
}, {
	about: "root_key_differs",
	m1: macaroonSpec{
		rootKey: "root-key",
		id:      "root-id",
		caveats: []caveat{{
			condition: "a",
		}, {
			condition: "d",
		}},
	},
	m2: macaroonSpec{
		rootKey: "root-key1",
		id:      "root-id",
		caveats: []caveat{{
			condition: "a",
		}, {
			condition: "d",
		}},
	},
	expect: false,
}, {
	about: "id_differs",
	m1: macaroonSpec{
		rootKey: "root-key",
		id:      "root-id",
	},
	m2: macaroonSpec{
		rootKey: "root-key",
		id:      "root-id1",
	},
	expect: false,
}, {
	about: "extra_caveat_1",
	m1: macaroonSpec{
		rootKey: "root-key",
		id:      "root-id",
		caveats: []caveat{{
			condition: "a",
		}, {
			condition: "d",
		}, {
			condition: "d",
		}},
	},
	m2: macaroonSpec{
		rootKey: "root-key",
		id:      "root-id",
		caveats: []caveat{{
			condition: "a",
		}, {
			condition: "d",
		}},
	},
	expect: false,
}, {
	about: "extra_caveat_2",
	m1: macaroonSpec{
		rootKey: "root-key1",
		id:      "root-id",
		caveats: []caveat{{
			condition: "a",
		}, {
			condition: "d",
		}},
	},
	m2: macaroonSpec{
		rootKey: "root-key",
		id:      "root-id",
		caveats: []caveat{{
			condition: "a",
		}, {
			condition: "d",
		}, {
			condition: "d",
		}},
	},
	expect: false,
}, {
	about: "caveat_condition_differs",
	m1: macaroonSpec{
		rootKey: "root-key",
		id:      "root-id",
		caveats: []caveat{{
			condition: "a",
			location:  "b",
			rootKey:   "c",
		}},
	},
	m2: macaroonSpec{
		rootKey: "root-key",
		id:      "root-id",
		caveats: []caveat{{
			condition: "a1",
			location:  "b",
			rootKey:   "c",
		}},
	},
	expect: false,
}}

func TestEqual(t *testing.T) {
	c := qt.New(t)
	for _, test := range equalTests {
		c.Run(test.about, func(c *qt.C) {
			m1 := makeMacaroon(test.m1)
			m2 := makeMacaroon(test.m2)
			c.Assert(m1.Equal(m2), qt.Equals, test.expect)
		})
	}
}

func TestEqualNil(t *testing.T) {
	c := qt.New(t)
	var nilm *macaroon.Macaroon
	m := MustNew([]byte("k"), []byte("x"), "l", macaroon.LatestVersion)
	c.Assert(nilm.Equal(nilm), qt.Equals, true)
	c.Assert(nilm.Equal(m), qt.Equals, false)
	c.Assert(m.Equal(nilm), qt.Equals, false)
}

type conditionTest struct {
	conditions map[string]bool
	expectErr  string
}

var verifyTests = []struct {
	about      string
	macaroons  []macaroonSpec
	conditions []conditionTest
}{{
	about: "single third party caveat without discharge",
	macaroons: []macaroonSpec{{
		rootKey: "root-key",
		id:      "root-id",
		caveats: []caveat{{
			condition: "wonderful",
		}, {
			condition: "bob-is-great",
			location:  "bob",
			rootKey:   "bob-caveat-root-key",
		}},
	}},
	conditions: []conditionTest{{
		conditions: map[string]bool{
			"wonderful": true,
		},
		expectErr: fmt.Sprintf(`cannot find discharge macaroon for caveat %x`, "bob-is-great"),
	}},
}, {
	about: "single third party caveat with discharge",
	macaroons: []macaroonSpec{{
		rootKey: "root-key",
		id:      "root-id",
		caveats: []caveat{{
			condition: "wonderful",
		}, {
			condition: "bob-is-great",
			location:  "bob",
			rootKey:   "bob-caveat-root-key",
		}},
	}, {
		location: "bob",
		rootKey:  "bob-caveat-root-key",
		id:       "bob-is-great",
	}},
	conditions: []conditionTest{{
		conditions: map[string]bool{
			"wonderful": true,
		},
	}, {
		conditions: map[string]bool{
			"wonderful": false,
		},
		expectErr: `condition "wonderful" not met`,
	}},
}, {
	about: "single third party caveat with discharge with mismatching root key",
	macaroons: []macaroonSpec{{
		rootKey: "root-key",
		id:      "root-id",
		caveats: []caveat{{
			condition: "wonderful",
		}, {
			condition: "bob-is-great",
			location:  "bob",
			rootKey:   "bob-caveat-root-key",
		}},
	}, {
		location: "bob",
		rootKey:  "bob-caveat-root-key-wrong",
		id:       "bob-is-great",
	}},
	conditions: []conditionTest{{
		conditions: map[string]bool{
			"wonderful": true,
		},
		expectErr: `signature mismatch after caveat verification`,
	}},
}, {
	about: "single third party caveat with two discharges",
	macaroons: []macaroonSpec{{
		rootKey: "root-key",
		id:      "root-id",
		caveats: []caveat{{
			condition: "wonderful",
		}, {
			condition: "bob-is-great",
			location:  "bob",
			rootKey:   "bob-caveat-root-key",
		}},
	}, {
		location: "bob",
		rootKey:  "bob-caveat-root-key",
		id:       "bob-is-great",
		caveats: []caveat{{
			condition: "splendid",
		}},
	}, {
		location: "bob",
		rootKey:  "bob-caveat-root-key",
		id:       "bob-is-great",
		caveats: []caveat{{
			condition: "top of the world",
		}},
	}},
	conditions: []conditionTest{{
		conditions: map[string]bool{
			"wonderful": true,
		},
		expectErr: `condition "splendid" not met`,
	}, {
		conditions: map[string]bool{
			"wonderful":        true,
			"splendid":         true,
			"top of the world": true,
		},
		expectErr: `discharge macaroon "bob-is-great" was not used`,
	}, {
		conditions: map[string]bool{
			"wonderful":        true,
			"splendid":         false,
			"top of the world": true,
		},
		expectErr: `condition "splendid" not met`,
	}, {
		conditions: map[string]bool{
			"wonderful":        true,
			"splendid":         true,
			"top of the world": false,
		},
		expectErr: `discharge macaroon "bob-is-great" was not used`,
	}},
}, {
	about: "one discharge used for two macaroons",
	macaroons: []macaroonSpec{{
		rootKey: "root-key",
		id:      "root-id",
		caveats: []caveat{{
			condition: "somewhere else",
			location:  "bob",
			rootKey:   "bob-caveat-root-key",
		}, {
			condition: "bob-is-great",
			location:  "charlie",
			rootKey:   "bob-caveat-root-key",
		}},
	}, {
		location: "bob",
		rootKey:  "bob-caveat-root-key",
		id:       "somewhere else",
		caveats: []caveat{{
			condition: "bob-is-great",
			location:  "charlie",
			rootKey:   "bob-caveat-root-key",
		}},
	}, {
		location: "bob",
		rootKey:  "bob-caveat-root-key",
		id:       "bob-is-great",
	}},
	conditions: []conditionTest{{
		expectErr: `discharge macaroon "bob-is-great" was used more than once`,
	}},
}, {
	about: "recursive third party caveat",
	macaroons: []macaroonSpec{{
		rootKey: "root-key",
		id:      "root-id",
		caveats: []caveat{{
			condition: "bob-is-great",
			location:  "bob",
			rootKey:   "bob-caveat-root-key",
		}},
	}, {
		location: "bob",
		rootKey:  "bob-caveat-root-key",
		id:       "bob-is-great",
		caveats: []caveat{{
			condition: "bob-is-great",
			location:  "charlie",
			rootKey:   "bob-caveat-root-key",
		}},
	}},
	conditions: []conditionTest{{
		expectErr: `discharge macaroon "bob-is-great" was used more than once`,
	}},
}, {
	about: "two third party caveats",
	macaroons: []macaroonSpec{{
		rootKey: "root-key",
		id:      "root-id",
		caveats: []caveat{{
			condition: "wonderful",
		}, {
			condition: "bob-is-great",
			location:  "bob",
			rootKey:   "bob-caveat-root-key",
		}, {
			condition: "charlie-is-great",
			location:  "charlie",
			rootKey:   "charlie-caveat-root-key",
		}},
	}, {
		location: "bob",
		rootKey:  "bob-caveat-root-key",
		id:       "bob-is-great",
		caveats: []caveat{{
			condition: "splendid",
		}},
	}, {
		location: "charlie",
		rootKey:  "charlie-caveat-root-key",
		id:       "charlie-is-great",
		caveats: []caveat{{
			condition: "top of the world",
		}},
	}},
	conditions: []conditionTest{{
		conditions: map[string]bool{
			"wonderful":        true,
			"splendid":         true,
			"top of the world": true,
		},
	}, {
		conditions: map[string]bool{
			"wonderful":        true,
			"splendid":         false,
			"top of the world": true,
		},
		expectErr: `condition "splendid" not met`,
	}, {
		conditions: map[string]bool{
			"wonderful":        true,
			"splendid":         true,
			"top of the world": false,
		},
		expectErr: `condition "top of the world" not met`,
	}},
}, {
	about: "third party caveat with undischarged third party caveat",
	macaroons: []macaroonSpec{{
		rootKey: "root-key",
		id:      "root-id",
		caveats: []caveat{{
			condition: "wonderful",
		}, {
			condition: "bob-is-great",
			location:  "bob",
			rootKey:   "bob-caveat-root-key",
		}},
	}, {
		location: "bob",
		rootKey:  "bob-caveat-root-key",
		id:       "bob-is-great",
		caveats: []caveat{{
			condition: "splendid",
		}, {
			condition: "barbara-is-great",
			location:  "barbara",
			rootKey:   "barbara-caveat-root-key",
		}},
	}},
	conditions: []conditionTest{{
		conditions: map[string]bool{
			"wonderful": true,
			"splendid":  true,
		},
		expectErr: fmt.Sprintf(`cannot find discharge macaroon for caveat %x`, "barbara-is-great"),
	}},
}, {
	about:     "multilevel third party caveats",
	macaroons: multilevelThirdPartyCaveatMacaroons,
	conditions: []conditionTest{{
		conditions: map[string]bool{
			"wonderful":   true,
			"splendid":    true,
			"high-fiving": true,
			"spiffing":    true,
		},
	}, {
		conditions: map[string]bool{
			"wonderful":   true,
			"splendid":    true,
			"high-fiving": false,
			"spiffing":    true,
		},
		expectErr: `condition "high-fiving" not met`,
	}},
}, {
	about: "unused discharge",
	macaroons: []macaroonSpec{{
		rootKey: "root-key",
		id:      "root-id",
	}, {
		rootKey: "other-key",
		id:      "unused",
	}},
	conditions: []conditionTest{{
		expectErr: `discharge macaroon "unused" was not used`,
	}},
}}

var multilevelThirdPartyCaveatMacaroons = []macaroonSpec{{
	rootKey: "root-key",
	id:      "root-id",
	caveats: []caveat{{
		condition: "wonderful",
	}, {
		condition: "bob-is-great",
		location:  "bob",
		rootKey:   "bob-caveat-root-key",
	}, {
		condition: "charlie-is-great",
		location:  "charlie",
		rootKey:   "charlie-caveat-root-key",
	}},
}, {
	location: "bob",
	rootKey:  "bob-caveat-root-key",
	id:       "bob-is-great",
	caveats: []caveat{{
		condition: "splendid",
	}, {
		condition: "barbara-is-great",
		location:  "barbara",
		rootKey:   "barbara-caveat-root-key",
	}},
}, {
	location: "charlie",
	rootKey:  "charlie-caveat-root-key",
	id:       "charlie-is-great",
	caveats: []caveat{{
		condition: "splendid",
	}, {
		condition: "celine-is-great",
		location:  "celine",
		rootKey:   "celine-caveat-root-key",
	}},
}, {
	location: "barbara",
	rootKey:  "barbara-caveat-root-key",
	id:       "barbara-is-great",
	caveats: []caveat{{
		condition: "spiffing",
	}, {
		condition: "ben-is-great",
		location:  "ben",
		rootKey:   "ben-caveat-root-key",
	}},
}, {
	location: "ben",
	rootKey:  "ben-caveat-root-key",
	id:       "ben-is-great",
}, {
	location: "celine",
	rootKey:  "celine-caveat-root-key",
	id:       "celine-is-great",
	caveats: []caveat{{
		condition: "high-fiving",
	}},
}}

func TestVerify(t *testing.T) {
	c := qt.New(t)
	for i, test := range verifyTests {
		c.Logf("test %d: %s", i, test.about)
		rootKey, macaroons := makeMacaroons(test.macaroons)
		for _, cond := range test.conditions {
			c.Logf("conditions %#v", cond.conditions)
			check := func(cav string) error {
				if cond.conditions[cav] {
					return nil
				}
				return fmt.Errorf("condition %q not met", cav)
			}
			err := macaroons[0].Verify(
				rootKey,
				check,
				macaroons[1:],
			)
			if cond.expectErr != "" {
				c.Assert(err, qt.ErrorMatches, cond.expectErr)
			} else {
				c.Assert(err, qt.IsNil)
			}

			// Cloned macaroon should have same verify result.
			cloneErr := macaroons[0].Clone().Verify(rootKey, check, macaroons[1:])
			if err == nil {
				c.Assert(cloneErr, qt.Equals, nil)
			} else {
				c.Assert(cloneErr.Error(), qt.Equals, err.Error())
			}
		}
	}
}

func TestTraceVerify(t *testing.T) {
	c := qt.New(t)
	rootKey, macaroons := makeMacaroons(multilevelThirdPartyCaveatMacaroons)
	traces, err := macaroons[0].TraceVerify(rootKey, macaroons[1:])
	c.Assert(err, qt.Equals, nil)
	c.Assert(traces, qt.HasLen, len(macaroons))
	// Check that we can run through the resulting operations and
	// arrive at the same signature.
	for i, m := range macaroons {
		r := traces[i].Results()
		c.Assert(b64str(r[len(r)-1]), qt.Equals, b64str(m.Signature()), qt.Commentf("macaroon %d", i))
	}
}

func TestTraceVerifyFailure(t *testing.T) {
	c := qt.New(t)
	rootKey, macaroons := makeMacaroons([]macaroonSpec{{
		rootKey: "xxx",
		id:      "hello",
		caveats: []caveat{{
			condition: "cond1",
		}, {
			condition: "cond2",
		}, {
			condition: "cond3",
		}},
	}})
	// Marshal the macaroon, corrupt a condition, then unmarshal
	// it and check we see the expected trace failure.
	data, err := json.Marshal(macaroons[0])
	c.Assert(err, qt.Equals, nil)
	var jm macaroon.MacaroonJSONV2
	err = json.Unmarshal(data, &jm)
	c.Assert(err, qt.Equals, nil)
	jm.Caveats[1].CID = "cond2 corrupted"
	data, err = json.Marshal(jm)
	c.Assert(err, qt.Equals, nil)

	var corruptm *macaroon.Macaroon
	err = json.Unmarshal(data, &corruptm)
	c.Assert(err, qt.Equals, nil)

	traces, err := corruptm.TraceVerify(rootKey, nil)
	c.Assert(err, qt.ErrorMatches, `signature mismatch after caveat verification`)
	c.Assert(traces, qt.HasLen, 1)
	var kinds []macaroon.TraceOpKind
	for _, op := range traces[0].Ops {
		kinds = append(kinds, op.Kind)
	}
	c.Assert(kinds, qt.DeepEquals, []macaroon.TraceOpKind{
		macaroon.TraceMakeKey,
		macaroon.TraceHash, // id
		macaroon.TraceHash, // cond1
		macaroon.TraceHash, // cond2
		macaroon.TraceHash, // cond3
		macaroon.TraceFail, // sig mismatch
	})
}

func b64str(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func TestVerifySignature(t *testing.T) {
	c := qt.New(t)
	rootKey, macaroons := makeMacaroons([]macaroonSpec{{
		rootKey: "xxx",
		id:      "hello",
		caveats: []caveat{{
			rootKey:   "y",
			condition: "something",
			location:  "somewhere",
		}, {
			condition: "cond1",
		}, {
			condition: "cond2",
		}},
	}, {
		rootKey: "y",
		id:      "something",
		caveats: []caveat{{
			condition: "cond3",
		}, {
			condition: "cond4",
		}},
	}})
	conds, err := macaroons[0].VerifySignature(rootKey, macaroons[1:])
	c.Assert(err, qt.IsNil)
	c.Assert(conds, qt.DeepEquals, []string{"cond3", "cond4", "cond1", "cond2"})

	conds, err = macaroons[0].VerifySignature(nil, macaroons[1:])
	c.Assert(err, qt.ErrorMatches, `failed to decrypt caveat 0 signature: decryption failure`)
	c.Assert(conds, qt.IsNil)
}

// TODO(rog) move the following JSON-marshal tests into marshal_test.go.

// jsonTestVersions holds the various possible ways of marshaling a macaroon
// to JSON.
var jsonTestVersions = []macaroon.Version{
	macaroon.V1,
	macaroon.V2,
}

func TestMarshalJSON(t *testing.T) {
	c := qt.New(t)
	for _, vers := range jsonTestVersions {
		c.Run(fmt.Sprintf("version_%v", vers), func(c *qt.C) {
			testMarshalJSONWithVersion(c, vers)
		})
	}
}

func testMarshalJSONWithVersion(c *qt.C, vers macaroon.Version) {
	rootKey := []byte("secret")
	m0 := MustNew(rootKey, []byte("some id"), "a location", vers)
	m0.AddFirstPartyCaveat([]byte("account = 3735928559"))
	m0JSON, err := json.Marshal(m0)
	c.Assert(err, qt.IsNil)
	var m1 macaroon.Macaroon
	err = json.Unmarshal(m0JSON, &m1)
	c.Assert(err, qt.IsNil)
	c.Assert(m0.Location(), qt.Equals, m1.Location())
	c.Assert(string(m0.Id()), qt.Equals, string(m1.Id()))
	c.Assert(
		hex.EncodeToString(m0.Signature()),
		qt.Equals,
		hex.EncodeToString(m1.Signature()))
	c.Assert(m1.Version(), qt.Equals, vers)
}

var jsonRoundTripTests = []struct {
	about string
	// data holds the marshaled data. All the data values hold
	// different encodings of the same macaroon - the same as produced
	// from the second example in libmacaroons
	// example README with the following libmacaroons code:
	//
	// secret = 'this is a different super-secret key; never use the same secret twice'
	// public = 'we used our other secret key'
	// location = 'http://mybank/'
	// M = macaroons.create(location, secret, public)
	// M = M.add_first_party_caveat('account = 3735928559')
	// caveat_key = '4; guaranteed random by a fair toss of the dice'
	// predicate = 'user = Alice'
	// identifier = 'this was how we remind auth of key/pred'
	// M = M.add_third_party_caveat('http://auth.mybank/', caveat_key, identifier)
	// m.serialize_json()
	data                 string
	expectExactRoundTrip bool
	expectVers           macaroon.Version
}{{
	about:                "exact_JSON_as_produced_by_libmacaroons",
	data:                 `{"caveats":[{"cid":"account = 3735928559"},{"cid":"this was how we remind auth of key\/pred","vid":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA027FAuBYhtHwJ58FX6UlVNFtFsGxQHS7uD_w_dedwv4Jjw7UorCREw5rXbRqIKhr","cl":"http:\/\/auth.mybank\/"}],"location":"http:\/\/mybank\/","identifier":"we used our other secret key","signature":"d27db2fd1f22760e4c3dae8137e2d8fc1df6c0741c18aed4b97256bf78d1f55c"}`,
	expectVers:           macaroon.V1,
	expectExactRoundTrip: true,
}, {
	about:      "V2_object_with_std_base-64_binary_values",
	data:       `{"c":[{"i64":"YWNjb3VudCA9IDM3MzU5Mjg1NTk="},{"i64":"dGhpcyB3YXMgaG93IHdlIHJlbWluZCBhdXRoIG9mIGtleS9wcmVk","v64":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA027FAuBYhtHwJ58FX6UlVNFtFsGxQHS7uD/w/dedwv4Jjw7UorCREw5rXbRqIKhr","l":"http://auth.mybank/"}],"l":"http://mybank/","i64":"d2UgdXNlZCBvdXIgb3RoZXIgc2VjcmV0IGtleQ==","s64":"0n2y/R8idg5MPa6BN+LY/B32wHQcGK7UuXJWv3jR9Vw="}`,
	expectVers: macaroon.V2,
}, {
	about:      "V2_object_with_URL_base-64_binary_values",
	data:       `{"c":[{"i64":"YWNjb3VudCA9IDM3MzU5Mjg1NTk"},{"i64":"dGhpcyB3YXMgaG93IHdlIHJlbWluZCBhdXRoIG9mIGtleS9wcmVk","v64":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA027FAuBYhtHwJ58FX6UlVNFtFsGxQHS7uD_w_dedwv4Jjw7UorCREw5rXbRqIKhr","l":"http://auth.mybank/"}],"l":"http://mybank/","i64":"d2UgdXNlZCBvdXIgb3RoZXIgc2VjcmV0IGtleQ","s64":"0n2y_R8idg5MPa6BN-LY_B32wHQcGK7UuXJWv3jR9Vw"}`,
	expectVers: macaroon.V2,
}, {
	about:                "V2_object_with_URL_base-64_binary_values_and_strings_for_ASCII",
	data:                 `{"c":[{"i":"account = 3735928559"},{"i":"this was how we remind auth of key/pred","v64":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA027FAuBYhtHwJ58FX6UlVNFtFsGxQHS7uD_w_dedwv4Jjw7UorCREw5rXbRqIKhr","l":"http://auth.mybank/"}],"l":"http://mybank/","i":"we used our other secret key","s64":"0n2y_R8idg5MPa6BN-LY_B32wHQcGK7UuXJWv3jR9Vw"}`,
	expectVers:           macaroon.V2,
	expectExactRoundTrip: true,
}, {
	about: "V2_base64_encoded_binary",
	data: `"` +
		base64.StdEncoding.EncodeToString([]byte(
			"\x02"+
				"\x01\x0ehttp://mybank/"+
				"\x02\x1cwe used our other secret key"+
				"\x00"+
				"\x02\x14account = 3735928559"+
				"\x00"+
				"\x01\x13http://auth.mybank/"+
				"\x02'this was how we remind auth of key/pred"+
				"\x04\x48\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xd3\x6e\xc5\x02\xe0\x58\x86\xd1\xf0\x27\x9f\x05\x5f\xa5\x25\x54\xd1\x6d\x16\xc1\xb1\x40\x74\xbb\xb8\x3f\xf0\xfd\xd7\x9d\xc2\xfe\x09\x8f\x0e\xd4\xa2\xb0\x91\x13\x0e\x6b\x5d\xb4\x6a\x20\xa8\x6b"+
				"\x00"+
				"\x00"+
				"\x06\x20\xd2\x7d\xb2\xfd\x1f\x22\x76\x0e\x4c\x3d\xae\x81\x37\xe2\xd8\xfc\x1d\xf6\xc0\x74\x1c\x18\xae\xd4\xb9\x72\x56\xbf\x78\xd1\xf5\x5c",
		)) + `"`,
	expectVers: macaroon.V2,
}}

func TestJSONRoundTrip(t *testing.T) {
	c := qt.New(t)
	for _, test := range jsonRoundTripTests {
		c.Run(fmt.Sprintf("v%v_%s", test.expectVers, test.about), func(c *qt.C) {
			testJSONRoundTripWithVersion(c, test.data, test.expectVers, test.expectExactRoundTrip)
		})
	}
}

func testJSONRoundTripWithVersion(c *qt.C, jsonData string, vers macaroon.Version, expectExactRoundTrip bool) {
	var m macaroon.Macaroon
	err := json.Unmarshal([]byte(jsonData), &m)
	c.Assert(err, qt.IsNil)
	assertLibMacaroonsMacaroon(c, &m)
	c.Assert(m.Version(), qt.Equals, vers)

	data, err := m.MarshalJSON()
	c.Assert(err, qt.IsNil)

	if expectExactRoundTrip {
		// The data is in canonical form, so we can check that
		// the round-tripped data is the same as the original
		// data when unmarshalled into an interface{}.
		var got interface{}
		err = json.Unmarshal(data, &got)
		c.Assert(err, qt.IsNil)

		var original interface{}
		err = json.Unmarshal([]byte(jsonData), &original)
		c.Assert(err, qt.IsNil)

		c.Assert(got, qt.DeepEquals, original, qt.Commentf("data: %s", data))
	}
	// Check that we can unmarshal the marshaled data anyway
	// and the macaroon still looks the same.
	var m1 macaroon.Macaroon
	err = m1.UnmarshalJSON(data)
	c.Assert(err, qt.IsNil)
	assertLibMacaroonsMacaroon(c, &m1)
	c.Assert(m.Version(), qt.Equals, vers)
}

// assertLibMacaroonsMacaroon asserts that m looks like the macaroon
// created in the README of the libmacaroons documentation.
// In particular, the signature is the same one reported there.
func assertLibMacaroonsMacaroon(c *qt.C, m *macaroon.Macaroon) {
	c.Assert(fmt.Sprintf("%x", m.Signature()), qt.Equals,
		"d27db2fd1f22760e4c3dae8137e2d8fc1df6c0741c18aed4b97256bf78d1f55c")
	c.Assert(m.Location(), qt.Equals, "http://mybank/")
	c.Assert(string(m.Id()), qt.Equals, "we used our other secret key")
	c.Assert(m.Caveats(), qt.DeepEquals, []macaroon.Caveat{{
		Id: []byte("account = 3735928559"),
	}, {
		Id:             []byte("this was how we remind auth of key/pred"),
		VerificationId: decodeB64("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA027FAuBYhtHwJ58FX6UlVNFtFsGxQHS7uD_w_dedwv4Jjw7UorCREw5rXbRqIKhr"),
		Location:       "http://auth.mybank/",
	}})
}

var jsonDecodeErrorTests = []struct {
	about       string
	data        string
	expectError string
}{{
	about:       "ambiguous id #1",
	data:        `{"i": "hello", "i64": "abcd", "s64": "ZDI3ZGIyZmQxZjIyNzYwZTRjM2RhZTgxMzdlMmQ4ZmMK"}`,
	expectError: "invalid identifier: ambiguous field encoding",
}, {
	about:       "ambiguous signature",
	data:        `{"i": "hello", "s": "345", "s64": "543467"}`,
	expectError: "invalid signature: ambiguous field encoding",
}, {
	about:       "signature too short",
	data:        `{"i": "hello", "s64": "0n2y/R8idg5MPa6BN+LY/B32wHQcGK7UuXJWv3jR9Q"}`,
	expectError: "signature has unexpected length 31",
}, {
	about:       "signature too long",
	data:        `{"i": "hello", "s64": "0n2y/R8idg5MPa6BN+LY/B32wHQcGK7UuXJWv3jR9dP1"}`,
	expectError: "signature has unexpected length 33",
}, {
	about:       "invalid caveat id",
	data:        `{"i": "hello", "s64": "0n2y/R8idg5MPa6BN+LY/B32wHQcGK7UuXJWv3jR9Vw", "c": [{"i": "hello", "i64": "00"}]}`,
	expectError: "invalid cid in caveat: ambiguous field encoding",
}, {
	about:       "invalid caveat vid",
	data:        `{"i": "hello", "s64": "0n2y/R8idg5MPa6BN+LY/B32wHQcGK7UuXJWv3jR9Vw", "c": [{"i": "hello", "v": "hello", "v64": "00"}]}`,
	expectError: "invalid vid in caveat: ambiguous field encoding",
}}

func TestJSONDecodeError(t *testing.T) {
	c := qt.New(t)
	for i, test := range jsonDecodeErrorTests {
		c.Logf("test %d: %v", i, test.about)
		var m macaroon.Macaroon
		err := json.Unmarshal([]byte(test.data), &m)
		c.Assert(err, qt.ErrorMatches, test.expectError)
	}
}

func TestFirstPartyCaveatWithInvalidUTF8(t *testing.T) {
	c := qt.New(t)
	rootKey := []byte("secret")
	badString := "foo\xff"

	m0 := MustNew(rootKey, []byte("some id"), "a location", macaroon.LatestVersion)
	err := m0.AddFirstPartyCaveat([]byte(badString))
	c.Assert(err, qt.Equals, nil)
}

func decodeB64(s string) []byte {
	data, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

type caveat struct {
	rootKey   string
	location  string
	condition string
}

type macaroonSpec struct {
	rootKey  string
	id       string
	caveats  []caveat
	location string
}

func makeMacaroons(mspecs []macaroonSpec) (rootKey []byte, macaroons macaroon.Slice) {
	for _, mspec := range mspecs {
		macaroons = append(macaroons, makeMacaroon(mspec))
	}
	primary := macaroons[0]
	for _, m := range macaroons[1:] {
		m.Bind(primary.Signature())
	}
	return []byte(mspecs[0].rootKey), macaroons
}

func makeMacaroon(mspec macaroonSpec) *macaroon.Macaroon {
	m := MustNew([]byte(mspec.rootKey), []byte(mspec.id), mspec.location, macaroon.LatestVersion)
	for _, cav := range mspec.caveats {
		if cav.location != "" {
			err := m.AddThirdPartyCaveat([]byte(cav.rootKey), []byte(cav.condition), cav.location)
			if err != nil {
				panic(err)
			}
		} else {
			m.AddFirstPartyCaveat([]byte(cav.condition))
		}
	}
	return m
}

func assertEqualMacaroons(c *qt.C, m0, m1 *macaroon.Macaroon) {
	m0json, err := m0.MarshalJSON()
	c.Assert(err, qt.IsNil)
	m1json, err := m1.MarshalJSON()
	var m0val, m1val interface{}
	err = json.Unmarshal(m0json, &m0val)
	c.Assert(err, qt.IsNil)
	err = json.Unmarshal(m1json, &m1val)
	c.Assert(err, qt.IsNil)
	c.Assert(m0val, qt.DeepEquals, m1val)
}

func TestBinaryRoundTrip(t *testing.T) {
	c := qt.New(t)
	// Test the binary marshalling and unmarshalling of a macaroon with
	// first and third party caveats.
	rootKey := []byte("secret")
	m0 := MustNew(rootKey, []byte("some id"), "a location", macaroon.LatestVersion)
	err := m0.AddFirstPartyCaveat([]byte("first caveat"))
	c.Assert(err, qt.IsNil)
	err = m0.AddFirstPartyCaveat([]byte("second caveat"))
	c.Assert(err, qt.IsNil)
	err = m0.AddThirdPartyCaveat([]byte("shared root key"), []byte("3rd party caveat"), "remote.com")
	c.Assert(err, qt.IsNil)
	data, err := m0.MarshalBinary()
	c.Assert(err, qt.IsNil)
	var m1 macaroon.Macaroon
	err = m1.UnmarshalBinary(data)
	c.Assert(err, qt.IsNil)
	assertEqualMacaroons(c, m0, &m1)
}

func TestBinaryMarshalingAgainstLibmacaroon(t *testing.T) {
	c := qt.New(t)
	// Test that a libmacaroon marshalled macaroon can be correctly unmarshaled
	data, err := base64.RawURLEncoding.DecodeString(
		"MDAxY2xvY2F0aW9uIGh0dHA6Ly9teWJhbmsvCjAwMmNpZGVudGlmaWVyIHdlIHVzZWQgb3VyIG90aGVyIHNlY3JldCBrZXkKMDAxZGNpZCBhY2NvdW50ID0gMzczNTkyODU1OQowMDMwY2lkIHRoaXMgd2FzIGhvdyB3ZSByZW1pbmQgYXV0aCBvZiBrZXkvcHJlZAowMDUxdmlkIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAANNuxQLgWIbR8CefBV-lJVTRbRbBsUB0u7g_8P3XncL-CY8O1KKwkRMOa120aiCoawowMDFiY2wgaHR0cDovL2F1dGgubXliYW5rLwowMDJmc2lnbmF0dXJlINJ9sv0fInYOTD2ugTfi2Pwd9sB0HBiu1LlyVr940fVcCg")
	c.Assert(err, qt.IsNil)
	var m0 macaroon.Macaroon
	err = m0.UnmarshalBinary(data)
	c.Assert(err, qt.IsNil)
	jsonData := []byte(`{"caveats":[{"cid":"account = 3735928559"},{"cid":"this was how we remind auth of key\/pred","vid":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA027FAuBYhtHwJ58FX6UlVNFtFsGxQHS7uD_w_dedwv4Jjw7UorCREw5rXbRqIKhr","cl":"http:\/\/auth.mybank\/"}],"location":"http:\/\/mybank\/","identifier":"we used our other secret key","signature":"d27db2fd1f22760e4c3dae8137e2d8fc1df6c0741c18aed4b97256bf78d1f55c"}`)
	var m1 macaroon.Macaroon
	err = m1.UnmarshalJSON(jsonData)
	c.Assert(err, qt.IsNil)
	assertEqualMacaroons(c, &m0, &m1)
}

var binaryFieldBase64ChoiceTests = []struct {
	id           string
	expectBase64 bool
}{
	{"x", false},
	{"\x00", true},
	{"\x03\x00", true},
	{"a longer id with more stuff", false},
	{"a longer id with more stuff and one invalid \xff", true},
	{"a longer id with more stuff and one encoded \x00", false},
}

func TestBinaryFieldBase64Choice(t *testing.T) {
	c := qt.New(t)
	for i, test := range binaryFieldBase64ChoiceTests {
		c.Logf("test %d: %q", i, test.id)
		m := MustNew([]byte{0}, []byte(test.id), "", macaroon.LatestVersion)
		data, err := json.Marshal(m)
		c.Assert(err, qt.Equals, nil)
		var x struct {
			Id   *string `json:"i"`
			Id64 *string `json:"i64"`
		}
		err = json.Unmarshal(data, &x)
		c.Assert(err, qt.Equals, nil)
		if test.expectBase64 {
			c.Assert(x.Id64, qt.Not(qt.IsNil))
			c.Assert(x.Id, qt.IsNil)
			idDec, err := base64.RawURLEncoding.DecodeString(*x.Id64)
			c.Assert(err, qt.Equals, nil)
			c.Assert(string(idDec), qt.Equals, test.id)
		} else {
			c.Assert(x.Id64, qt.IsNil)
			c.Assert(x.Id, qt.Not(qt.IsNil))
			c.Assert(*x.Id, qt.Equals, test.id)
		}
	}
}
