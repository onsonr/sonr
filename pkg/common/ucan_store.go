package common

import (
	"context"
	"fmt"
	"sort"
	"sync"

	"github.com/golang-jwt/jwt"
	"github.com/ipfs/go-cid"
	"github.com/onsonr/sonr/crypto/keys"
	"github.com/onsonr/sonr/crypto/ucan"
)

type IPFSTokenStore interface {
	ucan.TokenStore
	ResolveCIDBytes(ctx context.Context, id cid.Cid) ([]byte, error)
	ResolveDIDKey(ctx context.Context, did string) (keys.DID, error)
}

// ipfsUCANStore is a token store that uses IPFS to store tokens. It uses the memory store as a cache
// for CID strings to be used as keys for retrieving tokens.
type ipfsUCANStore struct {
	sync.Mutex
	ipfs  IPFS
	cache map[string]string
}

// NewUCANStore creates a new IPFS-backed token store
func NewUCANStore(ipfsClient IPFS) IPFSTokenStore {
	return &ipfsUCANStore{
		ipfs:  ipfsClient,
		cache: make(map[string]string),
	}
}

func (st *ipfsUCANStore) PutToken(ctx context.Context, key string, raw string) error {
	// Validate token format
	p := &jwt.Parser{
		UseJSONNumber:        true,
		SkipClaimsValidation: false,
	}
	if _, _, err := p.ParseUnverified(raw, jwt.MapClaims{}); err != nil {
		return fmt.Errorf("%w: %s", ucan.ErrInvalidToken, err)
	}

	// Store token in IPFS
	cid, err := st.ipfs.Add([]byte(raw))
	if err != nil {
		return fmt.Errorf("failed to store token in IPFS: %w", err)
	}

	// Update cache
	st.Lock()
	defer st.Unlock()
	st.cache[key] = cid
	return nil
}

func (st *ipfsUCANStore) RawToken(ctx context.Context, key string) (string, error) {
	st.Lock()
	cid, exists := st.cache[key]
	st.Unlock()

	if !exists {
		return "", ucan.ErrTokenNotFound
	}

	// Retrieve token from IPFS
	data, err := st.ipfs.Get(cid)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve token from IPFS: %w", err)
	}

	return string(data), nil
}

func (st *ipfsUCANStore) DeleteToken(ctx context.Context, key string) error {
	st.Lock()
	defer st.Unlock()

	cid, exists := st.cache[key]
	if !exists {
		return ucan.ErrTokenNotFound
	}

	// Unpin from IPFS
	if err := st.ipfs.Unpin(cid); err != nil {
		return fmt.Errorf("failed to unpin token from IPFS: %w", err)
	}

	delete(st.cache, key)
	return nil
}

func (st *ipfsUCANStore) ListTokens(ctx context.Context, offset, limit int) ([]ucan.RawToken, error) {
	st.Lock()
	defer st.Unlock()

	tokens := make(ucan.RawTokens, 0, len(st.cache))
	for key, cid := range st.cache {
		data, err := st.ipfs.Get(cid)
		if err != nil {
			continue // Skip invalid tokens
		}
		tokens = append(tokens, ucan.RawToken{
			Key: key,
			Raw: string(data),
		})
	}

	// Sort tokens
	sort.Sort(tokens)

	// Apply pagination
	if offset >= len(tokens) {
		return []ucan.RawToken{}, nil
	}

	end := offset + limit
	if end > len(tokens) || limit <= 0 {
		end = len(tokens)
	}

	return tokens[offset:end], nil
}

func (st *ipfsUCANStore) ResolveCIDBytes(ctx context.Context, id cid.Cid) ([]byte, error) {
	data, err := st.ipfs.Get(id.String())
	if err != nil {
		return nil, fmt.Errorf("failed to resolve CID bytes: %w", err)
	}
	return data, nil
}

func (st *ipfsUCANStore) ResolveDIDKey(ctx context.Context, did string) (keys.DID, error) {
	id, err := keys.Parse(did)
	if err != nil {
		return keys.DID{}, fmt.Errorf("failed to parse DID: %w", err)
	}
	return id, nil
}
