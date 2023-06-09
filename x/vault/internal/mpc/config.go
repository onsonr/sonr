package mpc

import (
	"github.com/sonrhq/core/internal/crypto"
	"github.com/sonrhq/core/x/vault/internal/mpc/algorithm"
	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
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

// WithHandlers sets the list of handlers that are called when a new account is generated.
func WithHandlers(handlers ...OnConfigGenerated) KeygenOption {
	return func(c *KeygenOpts) {
		c.Handlers = handlers
	}
}

// WithErrorChan sets the channel that is used to send errors.
func WithErrorChan(errorsChan chan error) KeygenOption {
	return func(c *KeygenOpts) {
		c.errorsChan = errorsChan
	}
}

// KeygenOpts is the configuration of an account.
type KeygenOpts struct {
	// Network is the network that is used to communicate with other parties.
	Network crypto.Network

	// Threshold is the number of required signatures to authorize a transaction.
	Threshold int

	// Self SelfID is the SelfID of the account that is used to sign transactions.
	SelfID crypto.PartyID

	// Group is the list of all parties that can sign transactions.
	Peers []crypto.PartyID

	// Handlers is the list of handlers that are called when a new account is generated.
	Handlers []OnConfigGenerated

	// finalConfigMap is a map that is used to store the configuration of all parties.
	group []party.ID

	current crypto.PartyID

	errorsChan chan error
}

func defaultKeygenOpts(current crypto.PartyID) *KeygenOpts {
	return &KeygenOpts{
		Threshold: kDefaultThreshold,
		Peers:     kDefaultGroup,
		current:   current,
	}
}

func (o *KeygenOpts) Apply(opts ...KeygenOption) {
	for _, opt := range opts {
		opt(o)
	}
	o.group = algorithm.EnsureSelfIDInGroup(o.current, o.Peers)
}

func (o *KeygenOpts) getOfflineNetwork() crypto.Network {
	closed := make(chan *protocol.Message)
	close(closed)
	c := &offlineNetwork{
		parties:          o.group,
		listenChannels:   make(map[party.ID]chan *protocol.Message, 2*len(o.group)),
		closedListenChan: closed,
	}
	return c
}

func (o *KeygenOpts) handleRoutineErr(err error) {
	if err == nil {
		return
	}
	if o.errorsChan != nil {
		o.errorsChan <- err
	}
}

func (o *KeygenOpts) handleConfigGeneration(config *cmp.Config) error {
	if o.Handlers == nil {
		return nil
	}

	for _, handler := range o.Handlers {
		if err := handler(config); err != nil {
			return err
		}
	}
	return nil
}
