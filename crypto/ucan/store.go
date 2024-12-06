package ucan

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"sync"

	"github.com/golang-jwt/jwt"
	"github.com/ipfs/go-cid"
)

// ErrTokenNotFound is returned by stores that cannot find an access token
// for a given key
var ErrTokenNotFound = errors.New("access token not found")

// TokenStore is a store intended for clients, who need to persist jwts.
// It deals in raw, string-formatted json web tokens, which are more useful
// when working with APIs, but validates the tokens are well-formed when placed
// in the store
//
// implementations of TokenStore must conform to the assertion test defined
// in the spec subpackage
type TokenStore interface {
	PutToken(ctx context.Context, key, rawToken string) error
	RawToken(ctx context.Context, key string) (rawToken string, err error)
	DeleteToken(ctx context.Context, key string) (err error)
	ListTokens(ctx context.Context, offset, limit int) (results []RawToken, err error)
}

// RawToken is a struct that binds a key to a raw token string
type RawToken struct {
	Key string
	Raw string
}

// RawTokens is a list of tokens that implements sorting by keys
type RawTokens []RawToken

func (rts RawTokens) Len() int           { return len(rts) }
func (rts RawTokens) Less(a, b int) bool { return rts[a].Key < rts[b].Key }
func (rts RawTokens) Swap(i, j int)      { rts[i], rts[j] = rts[j], rts[i] }

type memTokenStore struct {
	toksLk sync.Mutex
	toks   map[string]string
}

var (
	_ TokenStore       = (*memTokenStore)(nil)
	_ CIDBytesResolver = (*memTokenStore)(nil)
)

// NewMemTokenStore creates an in-memory token store
func NewMemTokenStore() TokenStore {
	return &memTokenStore{
		toks: map[string]string{},
	}
}

// MarshalJSON implements the json.Marshaller interface
func (st *memTokenStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(st.toRawTokens())
}

func (st *memTokenStore) PutToken(ctx context.Context, key string, raw string) error {
	p := &jwt.Parser{
		UseJSONNumber:        true,
		SkipClaimsValidation: false,
	}
	if _, _, err := p.ParseUnverified(raw, jwt.MapClaims{}); err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidToken, err)
	}

	st.toksLk.Lock()
	defer st.toksLk.Unlock()

	st.toks[key] = raw
	return nil
}

func (st *memTokenStore) ResolveCIDBytes(ctx context.Context, id cid.Cid) ([]byte, error) {
	rt, err := st.RawToken(ctx, id.String())
	if err != nil {
		return nil, err
	}
	return []byte(rt), nil
}

func (st *memTokenStore) RawToken(ctx context.Context, key string) (rawToken string, err error) {
	t, ok := st.toks[key]
	if !ok {
		return "", ErrTokenNotFound
	}
	return t, nil
}

func (st *memTokenStore) DeleteToken(ctx context.Context, key string) (err error) {
	st.toksLk.Lock()
	defer st.toksLk.Unlock()

	if _, ok := st.toks[key]; !ok {
		return ErrTokenNotFound
	}
	delete(st.toks, key)
	return nil
}

func (st *memTokenStore) ListTokens(ctx context.Context, offset, limit int) ([]RawToken, error) {
	var results []RawToken

	toks := st.toRawTokens()
	for i := 0; i < len(toks); i++ {
		if offset > 0 {
			offset--
			continue
		}
		results = append(results, toks[i])
		if limit > 0 && len(results) == limit {
			break
		}
	}

	return results, nil
}

func (st *memTokenStore) toRawTokens() RawTokens {
	toks := make(RawTokens, len(st.toks))
	i := 0
	for key, t := range st.toks {
		toks[i] = RawToken{
			Key: key,
			Raw: t,
		}
		i++
	}
	sort.Sort(toks)
	return toks
}
