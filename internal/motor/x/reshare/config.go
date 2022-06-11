// Copyright © 2020 AMIS Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package reshare

import (
	"fmt"
	"io/ioutil"

	"github.com/getamis/sirius/log"
	"github.com/sonr-io/alice/crypto/tss/ecdsa/gg18/reshare"
	"github.com/sonr-io/sonr/internal/motor/x/config"
	"gopkg.in/yaml.v2"
)

type ReshareConfig struct {
	Port      int64                `yaml:"port"`
	Threshold uint32               `yaml:"threshold"`
	Share     string               `yaml:"share"`
	Pubkey    config.Pubkey        `yaml:"pubkey"`
	BKs       map[string]config.BK `yaml:"bks"`
	Peers     []int64              `yaml:"peers"`
}

type ReshareResult struct {
	Share string `yaml:"share"`
}

func readReshareConfigFile(filaPath string) (*ReshareConfig, error) {
	c := &ReshareConfig{}
	yamlFile, err := ioutil.ReadFile(filaPath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func writeReshareResult(id string, result *reshare.Result) error {
	reshareResult := &ReshareResult{
		Share: result.Share.String(),
	}
	err := config.WriteYamlFile(reshareResult, getFilePath(id))
	if err != nil {
		log.Error("Cannot write YAML file", "err", err)
		return err
	}
	return nil
}

func getFilePath(id string) string {
	return fmt.Sprintf("reshare/%s-output.yaml", id)
}
