package dwn

import (
	"encoding/json"
	"io"
	"os"

	"github.com/ipfs/boxo/files"
)

const dwnJSONFileName = "dwn.json"

func LoadJSONConfig() (*Config, error) {
	// Read dwn.json config
	dwnBz, err := os.ReadFile(dwnJSONFileName)
	if err != nil {
		return nil, err
	}
	dwnConfig := new(Config)
	err = json.Unmarshal(dwnBz, dwnConfig)
	if err != nil {
		return nil, err
	}
	return dwnConfig, nil
}

func (c *Config) MarshalJSON() ([]byte, error) {
	return json.Marshal(c)
}

func (c *Config) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, c)
}

func (c *Config) UnmarshalFileNode(node files.File) error {
	cnfgBz, err := io.ReadAll(node)
	if err != nil {
		return err
	}
	err = c.UnmarshalJSON(cnfgBz)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) MarshalFileNode() (files.Node, error) {
	cnfgBz, err := c.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return files.NewBytesFile(cnfgBz), nil
}
