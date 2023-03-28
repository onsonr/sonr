package resolver

import (
	"context"
	"strings"

	"github.com/sonrhq/core/internal/local"
	"github.com/sonrhq/core/pkg/node"
)

type KVStoreItem interface {
	Bytes() []byte
	Did() string
}

type BasicStoreItem struct {
	did  string
	data []byte
}

func (i *BasicStoreItem) Bytes() []byte {
	return i.data
}

func (i *BasicStoreItem) Did() string {
	return i.did
}

func NewBasicStoreItem(did string, data []byte) *BasicStoreItem {
	return &BasicStoreItem{
		did:  did,
		data: data,
	}
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                          Global Resolver Store Methods                         ||
// ! ||--------------------------------------------------------------------------------||

// InsertKeyShare inserts a record into the IPFS store for the given controller
func InsertKeyShare(i KVStoreItem) error {
	err := setupKeyshareStore()
	if err != nil {
		return err
	}
	_, err = ksStore.KsTable.Put(context.Background(), keysharePrefix(i.Did()), i.Bytes())
	if err != nil {
		return err
	}
	return nil
}

// InsertKeyShare inserts a record into the IPFS store for the given controller
func InsertKSItem(did string, bz []byte) error {
	err := setupKeyshareStore()
	if err != nil {
		return err
	}
	_, err = ksStore.KsTable.Put(context.Background(), keysharePrefix(did), bz)
	if err != nil {
		return err
	}
	return nil
}


// InsertAccountInfo inserts a record into the IPFS store for the given controller
func InsertAccountInfo(accDid string, keyShareDids []string) error {
	err := setupKeyshareStore()
	if err != nil {
		return err
	}
	val := strings.Join(keyShareDids, ",")
	_, err = ksStore.KsTable.Put(context.Background(), accountPrefix(accDid), []byte(val))
	if err != nil {
		return err
	}
	return nil
}

// GetKeyShare gets a record from the IPFS store for the given controller
func GetKeyShare(key string) ([]byte, error) {
	err := setupKeyshareStore()
	if err != nil {
		return nil, err
	}
	vBiz, err := ksStore.KsTable.Get(context.Background(), keysharePrefix(key))
	if err != nil {
		return nil, err
	}
	return vBiz, nil
}

// GetAccountInfo gets a record from the IPFS store for the given controller
func GetAccountInfo(key string) ([]string, error) {
	err := setupKeyshareStore()
	if err != nil {
		return nil, err
	}
	vBiz, err := ksStore.KsTable.Get(context.Background(), accountPrefix(key))
	if err != nil {
		return nil, err
	}
	return strings.Split(string(vBiz), ","), nil
}

// DeleteKeyShare deletes a record from the IPFS store for the given controller
func DeleteKeyShare(key string) error {
	err := setupKeyshareStore()
	if err != nil {
		return err
	}
	_, err = ksStore.KsTable.Delete(context.Background(), keysharePrefix(key))
	if err != nil {
		return err
	}
	return nil
}

// DeleteAccountInfo deletes a record from the IPFS store for the given controller
func DeleteAccountInfo(key string) error {
	err := setupKeyshareStore()
	if err != nil {
		return err
	}
	_, err = ksStore.KsTable.Delete(context.Background(), accountPrefix(key))
	if err != nil {
		return err
	}
	return nil
}

// ListKeyShares lists all records in the IPFS store for the given controller
func ListKeyShares(accDid string) (map[string][]byte, error) {
	err := setupKeyshareStore()
	if err != nil {
		return nil, err
	}
	accMap, err := ListAccountInfo()
	if err != nil {
		return nil, err
	}
	keyShares := make(map[string][]byte, 0)
	for _, v := range accMap[accDid] {
		vBiz, err := ksStore.KsTable.Get(context.Background(), keysharePrefix(v))
		if err != nil {
			return nil, err
		}
		keyShares[v] = vBiz
	}
	return keyShares, nil
}

// ListAccountInfo lists all records in the IPFS store for the given controller
func ListAccountInfo() (map[string][]string, error) {
	err := setupKeyshareStore()
	if err != nil {
		return nil, err
	}
	m := make(map[string][]string)
	for k, v := range ksStore.KsTable.All() {
		if !strings.HasPrefix(k, "acc/") {
			continue
		}
		m[k] = strings.Split(string(v), ",")
	}
	return m, nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||              IPFS Based Wallet Store Implementation using OrbitDB              ||
// ! ||--------------------------------------------------------------------------------||

type vaultImpl struct {
	KsTable  node.IPFSKVStore
}

func makeControllerVault(store node.IPFSKVStore) *vaultImpl {
	return &vaultImpl{
		KsTable:  store,
	}
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                         Helper Methods for Module Setup                        ||
// ! ||--------------------------------------------------------------------------------||
var (
	ksStore *vaultImpl
	inStore node.IPFSDocsStore
)

// setupKeyshareStore initializes the global keyshare store
func setupKeyshareStore() error {
	if ksStore != nil {
		return nil
	}
	snrctx := local.NewContext()
	kv, err := node.OpenKeyValueStore(context.Background(), snrctx.GlobalKvKsStore)
	if err != nil {
		return err
	}
	ksStore = makeControllerVault(kv)
	return nil
}

func keysharePrefix(v string) string {
	return "ks/" + v
}

func accountPrefix(v string) string {
	return "acc/" + v
}
