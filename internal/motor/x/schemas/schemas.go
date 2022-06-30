

type AppSchemaInternal interface {
	GetAllWhatIs()
	GetAllSchemaDefinitions()
	BuildSchemaFromDefinition()
	VerifyShema()
	VerifyObject()
}

type vaultImpl struct {
	vaultEndpoint     string
	schemaDefinitions map[string]interface{}
	WhatIs            []string
}

func New() AppSchemaInternal {
	return &vaultImpl{
		storageEndpoint: "https://vault.sonr.ws",
	}
}