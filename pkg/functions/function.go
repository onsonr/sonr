package functions

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
)

type Function struct {
	file     *bytes.Buffer
	callback string
}

func NewFunction(file *[]byte, callback string) *Function {
	return &Function{
		file:     bytes.NewBuffer(*file),
		callback: callback,
	}
}

func (f *Function) Marshal() ([]byte, error) {
	b := make([]byte, f.file.Len())
	_, err := f.file.Read(b)
	if err != nil {
		return nil, err
	}
	blob := map[string]interface{}{
		"file":     b,
		"callback": f.callback,
	}
	data, err := json.Marshal(blob)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (f *Function) Unmarshal(d []byte) error {
	blob := make(map[string]interface{})
	err := json.Unmarshal(d, &blob)
	if err != nil {
		return err
	}

	str := blob["file"].(string)
	b, _ := base64.StdEncoding.DecodeString(str)
	f.file = bytes.NewBuffer(b)
	f.callback = blob["callback"].(string)
	return nil
}
