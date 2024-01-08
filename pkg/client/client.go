package client

import "github.com/sonrhq/sonr/pkg/client/modules"

type BlockchainClient struct {
    Bank modules.BankClient
    Gov modules.GovClient
    Identity modules.IdentityClient
    Service modules.ServiceClient
    Staking modules.StakingClient
    Tendermint modules.TendermintClient
}
