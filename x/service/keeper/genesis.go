package keeper

import (
	"context"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"

	modulev1 "github.com/sonrhq/sonr/api/sonr/service/module/v1"
	"github.com/sonrhq/sonr/x/service"
)

// serviceSchema is the schema for the service module.
var serviceSchema = &ormv1alpha1.ModuleSchemaDescriptor{
	SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{
			Id:            1,
			ProtoFileName: modulev1.File_sonr_service_module_v1_state_proto.Path(),
		},
	},
}

// InitGenesis initializes the module state from a genesis state.
func (k *Keeper) InitGenesis(ctx context.Context, data *service.GenesisState) error {
	// Set the params
	if err := k.Params.Set(ctx, data.Params); err != nil {
		return err
	}
	return nil
}

// ExportGenesis exports the module state to a genesis state.
func (k *Keeper) ExportGenesis(ctx context.Context) (*service.GenesisState, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	return &service.GenesisState{
		Params: params,
	}, nil
}
