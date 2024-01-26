package keeper

import (
	"context"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"

	modulev1 "github.com/sonrhq/sonr/api/service/module/v1"
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

	k.db.ServiceRecordTable().Insert(ctx, &modulev1.ServiceRecord{
		Id:          0,
		Origin:      "localhost",
		Name:        "Sonr LocalAuth",
		Description: "Sonr authentication service",
		Permissions: modulev1.ServicePermissions_SERVICE_PERMISSIONS_OWN,
	})

	// Set default permissions for the base, read, write and own modules
	k.db.BaseParamsTable().Save(ctx, &modulev1.BaseParams{
		Permissions:              modulev1.ServicePermissions_SERVICE_PERMISSIONS_BASE,
		Algorithm:                -7,
		AuthenticationAttachment: "platform",
	})
	k.db.ReadParamsTable().Save(ctx, &modulev1.ReadParams{
		Permissions:              modulev1.ServicePermissions_SERVICE_PERMISSIONS_READ,
		Algorithm:                -7,
		AuthenticationAttachment: "platform",
	})
	k.db.WriteParamsTable().Save(ctx, &modulev1.WriteParams{
		Permissions:              modulev1.ServicePermissions_SERVICE_PERMISSIONS_WRITE,
		Algorithm:                -8,
		ResidentKey:              "preferred",
		AuthenticationAttachment: "cross-platform",
	})
	k.db.OwnParamsTable().Save(ctx, &modulev1.OwnParams{
		Permissions:              modulev1.ServicePermissions_SERVICE_PERMISSIONS_OWN,
		Algorithm:                -8,
		ResidentKey:              "preferred",
		AuthenticationAttachment: "cross-platform",
	})
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
