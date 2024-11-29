package mpc

import (
	"encoding/json"
	"io"
)

type exportData struct {
	Address  string `json:"addr"`
	PubKey   []byte `json:"pubKey"`
	ValData  []byte `json:"val"`
	UserData []byte `json:"user"`
}

func (e *exportData) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

func (e *exportData) Unmarshal(data []byte) error {
	return json.Unmarshal(data, e)
}

func ImportKeyset(secret []byte, dat File) (Keyset, error) {
	data, err := io.ReadAll(dat)
	if err != nil {
		return nil, err
	}
	var ed exportData
	err = ed.Unmarshal(data)
	if err != nil {
		return nil, err
	}
	user, val, err := loadShareFromExportData(&ed)
	if err != nil {
		return nil, err
	}
	k := keyset{
		user: user,
		val:  val,
	}
	return k, nil
}

func (k keyset) Export(client IPFSClient, secret []byte) (ExportedKeyset, error) {
	valData, err := k.val.Marshal()
	if err != nil {
		return nil, err
	}
	userData, err := k.user.Marshal()
	if err != nil {
		return nil, err
	}
	addr, err := ComputeSonrAddr(k.val.PublicKey)
	if err != nil {
		return nil, err
	}
	ed := exportData{
		Address:  addr,
		PubKey:   k.val.PublicKey,
		ValData:  valData,
		UserData: userData,
	}
	return ed.Marshal()
}

func loadShareFromExportData(data *exportData) (*UserKeyshare, *ValKeyshare, error) {
	var (
		valMsg  Message
		userMsg Message
	)
	err := json.Unmarshal(data.UserData, &userMsg)
	if err != nil {
		return nil, nil, err
	}
	err = json.Unmarshal(data.ValData, &valMsg)
	if err != nil {
		return nil, nil, err
	}
	user := &UserKeyshare{
		Message:   userMsg,
		Role:      2,
		PublicKey: data.PubKey,
	}
	val := &ValKeyshare{
		Message:   valMsg,
		Role:      1,
		PublicKey: data.PubKey,
	}
	return user, val, nil
}
