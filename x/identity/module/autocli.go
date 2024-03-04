package module

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	identityv1 "github.com/sonrhq/sonr/api/sonr/identity/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service:           identityv1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:           identityv1.Msg_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{},
		},
	}
}
