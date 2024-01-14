package module

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	identityv1 "github.com/sonrhq/sonr/api/identity/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: identityv1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Get the current module parameters",
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: identityv1.Msg_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "IncrementCounter",
					Use:       "counter [sender]",
					Short:     "Increments the counter by 1 for the sender",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "sender"},
					},
				},
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // This is a authority gated tx, so we skip it.
				},
			},
		},
	}
}
