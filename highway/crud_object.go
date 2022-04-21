package highway

import (
	context "context"
	"fmt"

	otv1 "github.com/sonr-io/blockchain/x/object/types"
	ot "go.buf.build/grpc/go/sonr-io/blockchain/object"
	registry "go.buf.build/grpc/go/sonr-io/blockchain/registry"
)

// CreateObject creates a new object.
func (s *HighwayServer) CreateObject(ctx context.Context, req *ot.MsgCreateObject) (*ot.MsgCreateObjectResponse, error) {
	// Broadcast the Transaction
	resp, err := s.cosmos.BroadcastCreateObject(otv1.NewMsgCreateObjectFromBuf(req))
	if err != nil {
		return nil, err
	}
	fmt.Println(resp.String())

	// Upload Object Schema to IPFS
	cid, err := s.ipfsProtocol.PutObjectSchema(resp.GetWhatIs().GetObjectDoc())
	if err != nil {
		return nil, err
	}
	fmt.Println(cid)
	return &ot.MsgCreateObjectResponse{}, nil
}

// UpdateObject updates an object.
func (s *HighwayServer) UpdateObject(ctx context.Context, req *ot.MsgUpdateObject) (*ot.MsgUpdateObjectResponse, error) {
	// Broadcast the Transaction
	resp, err := s.cosmos.BroadcastUpdateObject(otv1.NewMsgUpdateObjectFromBuf(req))
	if err != nil {
		return nil, err
	}
	fmt.Println(resp.String())
	return &ot.MsgUpdateObjectResponse{
		Code:    resp.Code,
		Message: resp.Message,
		WhatIs:  otv1.NewWhatIsToBuf(resp.WhatIs),
	}, nil
}

// -----------------
// Helper Functions
// -----------------

// Translate ot (buf) fields to ot_v1 (blockchain) //TODO rename this
func (s *HighwayServer) bufToBlockchain(bufFields []*ot.ObjectField) []*ot_v1.ObjectField {
	var blockchainFields []*ot_v1.ObjectField
	for _, v := range bufFields {
		item := &ot_v1.ObjectField{
			Label: v.Label,
			Type:  ot_v1.ObjectFieldType(v.Type),
			Did:   v.Did,
			// Value: v.GetValue(), //TODO maybe switch based on specific implementation of type?
		}
		blockchainFields = append(blockchainFields, item)
	}

	return blockchainFields
}

// Translate the objectDoc ot_v1 (blockchain) fields to ot (buf) //TODO rename this
func (s *HighwayServer) decodeObjectDocFields(blockchainFields map[string]*ot_v1.ObjectField) map[string]*ot.ObjectField {
	// decode ot_v1 fields to ot

	if blockchainFields == nil || len(blockchainFields) < 1 {
		return map[string]*ot.ObjectField{}
	}

	outputFields := make(map[string]*ot.ObjectField, len(blockchainFields))
	for k, v := range blockchainFields {
		outputFields[k] = &ot.ObjectField{
			Label: v.GetLabel(),
			Type:  ot.ObjectFieldType(v.GetType()),
			Did:   v.GetDid(),
			//Value:    v.GetValue(), //TODO this should work
			Metadata: v.GetMetadata(),
		}
	}

	return outputFields
}

// Translate Registery session to types session
func (s *HighwayServer) regSessToTypeSess(regSess registry.Session) *types.Session {
	return &types.Session{
		BaseDid: regSess.GetBaseDid(),
		Whois: &types.WhoIs{
			Name:        regSess.Whois.GetName(),
			Did:         regSess.Whois.GetDid(),
			Document:    regSess.Whois.GetDocument(),
			Creator:     regSess.Whois.GetCreator(),
			Credentials: s.regCredtoTypeCred(regSess.Whois.GetCredentials()),
			Type:        types.WhoIs_Type(regSess.Whois.GetType()),
			Metadata:    regSess.Whois.GetMetadata(),
			Timestamp:   regSess.Whois.GetTimestamp(),
			IsActive:    regSess.Whois.GetIsActive(),
		},
		Credential: &types.Credential{
			ID:              regSess.Credential.GetID(),
			PublicKey:       regSess.Credential.GetPublicKey(),
			AttestationType: regSess.Credential.GetAttestationType(),
			// Authenticator: &types.Authenticator{  //TODO this causes nil dereference, figure out why
			// 	Aaguid:       regSess.Credential.Authenticator.Aaguid,
			// 	SignCount:    regSess.Credential.Authenticator.SignCount,
			// 	CloneWarning: regSess.Credential.Authenticator.CloneWarning,
			// },
		},
	}
}

// Translate registry credential to types credential
func (s *HighwayServer) regCredtoTypeCred(regCred []*registry.Credential) []*types.Credential {
	var typesCred []*types.Credential
	for _, v := range regCred {
		typesCred = append(typesCred, &types.Credential{
			ID:              v.GetID(),
			PublicKey:       v.GetPublicKey(),
			AttestationType: v.GetAttestationType(),
			Authenticator: &types.Authenticator{
				Aaguid:       v.Authenticator.GetAaguid(),
				SignCount:    v.Authenticator.GetSignCount(),
				CloneWarning: v.Authenticator.GetCloneWarning(),
			},
		})
	}

	return typesCred
}
