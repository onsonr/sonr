package types

import "encoding/json"

type InfoFile struct {
	Creds      Credentials `json:"credentials"`
	Properties Properties  `json:"properties"`
}

func (i *InfoFile) Marshal() ([]byte, error) {
	return json.Marshal(i)
}

func (i *InfoFile) Unmarshal(data []byte) error {
	return json.Unmarshal(data, i)
}
