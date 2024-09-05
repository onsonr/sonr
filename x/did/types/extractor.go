package types

import didv1 "github.com/onsonr/sonr/api/did/v1"

func (m *MsgRegisterService) ExtractServiceRecord() (*didv1.ServiceRecord, error) {
	return &didv1.ServiceRecord{
		Controller:  m.Controller,
		OriginUri:   m.OriginUri,
		Description: m.Description,
	}, nil
}

func convertPermissions(permissions *Permissions) *didv1.Permissions {
	if permissions == nil {
		return nil
	}
	return &didv1.Permissions{}
}
