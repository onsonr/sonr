package motor

import (
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
	objectcli "github.com/sonr-io/sonr/internal/object"
	"github.com/sonr-io/sonr/internal/schemas"
	"github.com/sonr-io/sonr/pkg/motor/x/object"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	st "github.com/sonr-io/sonr/x/schema/types"
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

func (mtr *motorNodeImpl) GetDocument(req mt.GetDocumentRequest) (*mt.GetDocumentResponse, error) {
	obj, err := mtr.QueryObject(req.GetCid())
	if err != nil {
		return nil, err
	}

	doc := st.NewDocumentFromMap(req.GetCid(), obj)
	return &mt.GetDocumentResponse{
		Did:      doc.Did,
		Status:   200,
		Document: doc,
		Cid:      req.GetCid(),
	}, nil
}

func (mtr *motorNodeImpl) UploadDocument(req mt.UploadDocumentRequest) (*mt.UploadDocumentResponse, error) {
	builder, err := mtr.NewObjectBuilder(req.GetDefinition().GetDid())
	if err != nil {
		return nil, err
	}

	builder.SetLabel(req.GetLabel())
	builder.Set("@did", req.GetDefinition().GetDid())
	for _, field := range req.GetFields() {
		builder.Set(field.GetName(), field.GetValue())
	}

	resp, err := builder.Upload()
	if err != nil {
		return nil, err
	}
	return &mt.UploadDocumentResponse{
		Status: resp.Code,
		Cid:    resp.Reference.Cid,
		Did:    resp.Reference.Did,
		Document: &st.SchemaDocument{
			Did:        resp.Reference.Did,
			Cid:        resp.Reference.Cid,
			Creator:    mtr.GetAddress(),
			Definition: req.GetDefinition(),
			Fields:     req.GetFields(),
		},
	}, nil
}
