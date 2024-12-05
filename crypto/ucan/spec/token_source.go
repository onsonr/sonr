package spec

// import (
// 	"context"
// 	"testing"

// 	"github.com/golang-jwt/jwt"
// 	cfgtest "github.com/qri-io/qri/config/test"
// 	"github.com/qri-io/qri/repo/profile"
// 	"github.com/ucan-wg/go-ucan"
// )

// // AssertTokenSourceSpec ensures a TokenSource implementation behaves as
// // expected
// func AssertTokenSourceSpec(t *testing.T, newTokenSource func(ctx context.Context) ucan.TokenSource) {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	source := newTokenSource(ctx)

// 	p1 := &profile.Profile{
// 		ID:       profile.IDB58DecodeOrEmpty(cfgtest.GetTestPeerInfo(1).EncodedPeerID),
// 		Peername: "username",
// 	}

// 	raw, err := source.CreateToken(p1, 0)
// 	if err != nil {
// 		t.Errorf("source should allow creating key with valid profile & zero duration. got: %q", err)
// 	}

// 	p := &jwt.Parser{
// 		UseJSONNumber:        true,
// 		SkipClaimsValidation: false,
// 	}
// 	if _, _, err := p.ParseUnverified(raw, &ucan.TokenClaims{}); err != nil {
// 		t.Errorf("created token must parse with acces.TokenClaims. got: %q", err)
// 	}

// 	if _, err := ucan.ParseToken(raw, source); err != nil {
// 		t.Errorf("source must create tokens that parse with it's own verification keys. error: %q", err)
// 	}
// }
