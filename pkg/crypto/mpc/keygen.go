package mpc

import (
	"sync"

	"github.com/sonrhq/core/pkg/crypto"
	"github.com/sonrhq/core/pkg/crypto/mpc/internal/algorithm"
	"github.com/sonrhq/core/pkg/crypto/mpc/internal/network"
	"github.com/sonrhq/core/pkg/crypto/mpc/internal/utils"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	"github.com/taurusgroup/multi-party-sig/protocols/cmp"
)

var (
	// Default Account Name is the default name of the account.
	kDefaultAccName = "primary"

	// Default Threshold is the default number of required signatures to authorize a transaction.
	kDefaultThreshold = 1

	// Default CoinType is the default coin type of the account.
	kDefaultCoinType = crypto.SONRCoinType

	// Default Group is the default list of all parties that can sign transactions.
	kDefaultGroup = []crypto.PartyID{crypto.PartyID("vault")}
)

// OnConfigGenerated is a callback function that is called when a new account is generated.
type OnConfigGenerated func(*cmp.Config) error

// KeygenOption is a function that configures an account.
type KeygenOption func(*KeygenOpts)

// WithThreshold sets the number of required signatures to authorize a transaction.
func WithThreshold(threshold int) KeygenOption {
	return func(c *KeygenOpts) {
		c.Threshold = threshold
	}
}

// WithSelfID sets the ID of the account that is used to sign transactions.
func WithSelfID(selfID string) KeygenOption {
	return func(c *KeygenOpts) {
		c.SelfID = crypto.PartyID(selfID)
	}
}

// WithPeers sets the list of all parties that can sign transactions.
func WithPeers(peers ...string) KeygenOption {
	return func(c *KeygenOpts) {
		c.Peers = make([]crypto.PartyID, len(peers))
		for i, p := range peers {
			c.Peers[i] = crypto.PartyID(p)
		}
	}
}

// Keygen Generates a new ECDSA private key shared among all the given participants.
func Keygen(current crypto.PartyID, threshold int, peers []crypto.PartyID) ([]*cmp.Config, error) {
	group := utils.EnsureSelfIDInGroup(current, peers)
	net := network.NewOfflineNetwork(group...)
	var mtx sync.Mutex
	var wg sync.WaitGroup
	confs := make([]*cmp.Config, 0)
	for _, id := range net.Ls() {
		wg.Add(1)
		go func(id party.ID) {
			pl := pool.NewPool(0)
			defer pl.TearDown()
			conf, err := algorithm.CmpKeygen(id, net.Ls(), net, threshold, &wg, pl)
			if err != nil {
				return
			}
			mtx.Lock()
			confs = append(confs, conf)
			mtx.Unlock()
		}(id)
	}
	wg.Wait()
	return confs, nil
}

// KeygenOpts is the configuration of an account.
type KeygenOpts struct {
	// AccName is the name of the account.
	AccName string

	// Network is the network that is used to communicate with other parties.
	Network crypto.Network

	// Threshold is the number of required signatures to authorize a transaction.
	Threshold int

	// Self SelfID is the SelfID of the account that is used to sign transactions.
	SelfID crypto.PartyID

	// Group is the list of all parties that can sign transactions.
	Peers []crypto.PartyID

	// CoinType is the coin type of the account.
	CoinType crypto.CoinType

	// finalConfigMap is a map that is used to store the configuration of all parties.
	finalConfigMap map[party.ID]*cmp.Config

	// doneConfChan is a channel that is used to signal that the configuration of the current party is generated.
	doneConfChan chan *cmp.Config

	// mtx is a mutex that is used to synchronize access to the doneConfChan.
	sync.Mutex
}

func defaultKeygenOpts() *KeygenOpts {
	return &KeygenOpts{
		AccName:        kDefaultAccName,
		Threshold:      kDefaultThreshold,
		Peers:          kDefaultGroup,
		CoinType:       kDefaultCoinType,
		finalConfigMap: make(map[party.ID]*cmp.Config),
		doneConfChan:   make(chan *cmp.Config),
	}
}
