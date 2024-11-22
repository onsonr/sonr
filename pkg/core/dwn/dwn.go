package dwn

import (
	"encoding/json"
	"os"
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
