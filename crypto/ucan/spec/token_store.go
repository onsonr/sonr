package spec

// import (
// 	"context"
// 	"errors"
// 	"testing"
// 	"time"

// 	"github.com/google/go-cmp/cmp"
// 	cfgtest "github.com/qri-io/qri/config/test"
// 	"github.com/qri-io/qri/repo/profile"
// 	"github.com/ucan-wg/go-ucan"
// )

// // AssertTokenStoreSpec ensures an ucan.TokenStore implementation behaves as
// // expected
// func AssertTokenStoreSpec(t *testing.T, newTokenStore func(context.Context) ucan.TokenStore) {
// 	prevTs := ucan.Timestamp
// 	ucan.Timestamp = func() time.Time { return time.Time{} }
// 	defer func() { ucan.Timestamp = prevTs }()

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	pk := cfgtest.GetTestPeerInfo(0).PrivKey
// 	tokens, err := ucan.NewPrivKeyTokenSource(pk)
// 	if err != nil {
// 		t.Fatalf("creating local tokens: %q", err)
// 	}
// 	store := newTokenStore(ctx)

// 	results, err := store.ListTokens(ctx, 0, -1)
// 	if err != nil {
// 		t.Errorf("listing all tokens of an empty store shouldn't error. got: %q ", err)
// 	}
// 	if len(results) > 0 {
// 		t.Errorf("new store should return no results. got: %d", len(results))
// 	}

// 	_, err = store.RawToken(ctx, "this doesn't exist")
// 	if !errors.Is(err, ucan.ErrTokenNotFound) {
// 		t.Errorf("expected store.RawToken(nonexistent key) to return a wrap of ucan.ErrTokenNotFound. got: %q", err)
// 	}
// 	err = store.DeleteToken(ctx, "this also doesn't exist")
// 	if !errors.Is(err, ucan.ErrTokenNotFound) {
// 		t.Errorf("expected store.D key to return a wrap of ucan.ErrTokenNotFound. got: %q", err)
// 	}
// 	if err := store.PutToken(ctx, "_bad_key", "not.a.key"); err == nil {
// 		t.Errorf("putting an invalid json web token should error. got nil")
// 	}

// 	p1 := &profile.Profile{
// 		ID:       profile.IDB58DecodeOrEmpty(cfgtest.GetTestPeerInfo(1).EncodedPeerID),
// 		Peername: "local_user",
// 	}
// 	t1Raw, err := tokens.CreateToken(p1, 0)
// 	if err != nil {
// 		t.Fatalf("creating token: %q", err)
// 	}

// 	if err := store.PutToken(ctx, "_root", t1Raw); err != nil {
// 		t.Errorf("putting root key shouldn't error. got: %q", err)
// 	}

// 	results, err = store.ListTokens(ctx, 0, -1)
// 	if err != nil {
// 		t.Errorf("listing all tokens of an empty store shouldn't error. got: %q ", err)
// 	}
// 	if len(results) != 1 {
// 		t.Errorf("result length mismatch listing keys after adding `root` key. expected 1, got: %d", len(results))
// 	}

// 	p2 := &profile.Profile{
// 		ID:       profile.IDB58DecodeOrEmpty(cfgtest.GetTestPeerInfo(1).EncodedPeerID),
// 		Peername: "user_2",
// 	}
// 	t2Raw, err := tokens.CreateToken(p2, time.Millisecond*10)
// 	if err != nil {
// 		t.Fatalf("creating token: %q", err)
// 	}

// 	secondKey := "http://registry.qri.cloud"
// 	if err := store.PutToken(ctx, secondKey, t2Raw); err != nil {
// 		t.Errorf("putting a second token with key=%q shouldn't error. got: %q", secondKey, err)
// 	}

// 	results, err = store.ListTokens(ctx, 0, -1)
// 	if err != nil {
// 		t.Errorf("listing all tokens of an empty store shouldn't error. got: %q ", err)
// 	}
// 	if len(results) != 2 {
// 		t.Errorf("result length mismatch listing keys after adding second key. expected 2, got: %d", len(results))
// 	}

// 	expect := []ucan.RawToken{
// 		{
// 			Key: "_root",
// 			Raw: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJRbVdZZ0Q0OXI5SG51WEVwcFFFcTFhN1NVVXJ5amE0UU5zOUU2WENIMlBheUNEIiwidXNlcm5hbWUiOiJsb2NhbF91c2VyIn0.hu1B92X8cLBRNtNNiwm_qn4T-s8WlDlsa0swNgeyUPJ921LfojmHobkuW4oRvNEjkq_OP2gkaZ_F0YyUgAM8K-pVg30L-jNG9cqA1EUx4cQ90ZSbMxvXzRmBevBa3Wq-RHErnGw-K7EvtZfuPrp60LuDBKkGCuAwfKV8D9O-6U4lrragFgfw3zWRdovnb28fO2W6sqP8azGDcY8klpysjx7W4V-qVynJ981_ex_G1wPbk1dov59MDlY6yoxt1rucyF5-f4oo9jv6k194Tigw3Uv6JR889kK5x87ruiApghfQIBosAd-hm79Xz0RmLahykoZZTbVASW6NcIPvqvZ5TA",
// 		},
// 		{
// 			Key: secondKey,
// 			Raw: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOi02MjEzNTU5NjgwMCwic3ViIjoiUW1XWWdENDlyOUhudVhFcHBRRXExYTdTVVVyeWphNFFOczlFNlhDSDJQYXlDRCIsInVzZXJuYW1lIjoidXNlcl8yIn0.V4ZlinPrjRBf6EiZPylrigXHEQ-5k8e54-G5o1OQEu9hAWmZAGlm5OdcAfWRLueWR8bk_9NFYwAzg7CaSvM0suPqHlFbCXS0tQynG3m9ptu_bucfEAaDdoNmBMyPrB7w6rLauFVh0TcwiTRV3KkrzbeKdo0Q6fDoRFy5ZuP5zPDxt_UOC50zIpdw5F4MH3OBbfjjfpR0XA8Q1tslOg5JWh21pCR-C8gyC5jJ9ilNsdwggXY0mkHl9f0utwMewPfC6b7i01t9kVtknI5Wg6dq_CuZlDuOs83LcL6xThRBiosmFZ-2I3b8nSPxfg_gouOBOSa5yFOcYxJsHHDKe-RIfw",
// 		},
// 	}

// 	if diff := cmp.Diff(expect, results); diff != "" {
// 		t.Errorf("mistmatched list keys results. (-want +got):\n%s", diff)
// 	}

// 	results, err = store.ListTokens(ctx, 1, 1)
// 	if err != nil {
// 		t.Errorf("listing all tokens of an empty store shouldn't error. got: %q ", err)
// 	}
// 	if len(results) != 1 {
// 		t.Errorf("result length mismatch listing keys after adding `root` key. expected 1, got: %d", len(results))
// 	}

// 	if diff := cmp.Diff(expect[1:], results); diff != "" {
// 		t.Errorf("mistmatched list keys with offset=1, limit=1. results. (-want +got):\n%s", diff)
// 	}

// 	if err := store.DeleteToken(ctx, secondKey); err != nil {
// 		t.Errorf("store.DeleteToken shouldn't error for existing key. got: %q", err)
// 	}

// 	_, err = store.RawToken(ctx, secondKey)
// 	if !errors.Is(err, ucan.ErrTokenNotFound) {
// 		t.Errorf("store.RawToken() for a just-deleted key must return a wrap of ucan.ErrTokenNotFound. got: %q", err)
// 	}
// }
