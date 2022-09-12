package motor

import (
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
	objectcli "github.com/sonr-io/sonr/internal/object"
	"github.com/sonr-io/sonr/internal/schemas"
	"github.com/sonr-io/sonr/pkg/motor/x/object"
)

func (mtr *motorNodeImpl) NewObjectBuilder(did string) (*object.ObjectBuilder, error) {
	whatIs, _, found := mtr.Resources.GetSchema(did)
	if !found {
		return nil, fmt.Errorf("could not find WhatIs with did '%s'", did)
	}

	schemaImpl := schemas.NewWithClient(mtr.GetClient(), whatIs)
	objCli := objectcli.New(schemaImpl, shell.NewShell(mtr.Cosmos.GetIPFSApiAddress()))
	return object.NewBuilder(schemaImpl, objCli), nil
}
