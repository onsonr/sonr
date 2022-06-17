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
	"io/ioutil"

	"github.com/golang/protobuf/proto"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/sonr-io/alice/crypto/tss/ecdsa/gg18/reshare"
	"github.com/sonr-io/alice/example/utils"
	"github.com/sonr-io/alice/message/types"
)

type service struct {
	config *ReshareConfig
	pm     types.PeerManager

	reshare *reshare.Reshare
	done    chan struct{}
}

func NewService(config *ReshareConfig, pm types.PeerManager) (*service, error) {
	s := &service{
		config: config,
		pm:     pm,
		done:   make(chan struct{}),
	}

	// Reshare needs results from DKG.
	dkgResult, err := utils.ConvertDKGResult(config.Pubkey, config.Share, config.BKs)
	if err != nil {

		return nil, err
	}

	// Create reshare
	reshare, err := reshare.NewReshare(pm, config.Threshold, dkgResult.PublicKey, dkgResult.Share, dkgResult.Bks, s)
	if err != nil {

		return nil, err
	}
	s.reshare = reshare
	return s, nil
}

func (p *service) Handle(s network.Stream) {
	data := &reshare.Message{}
	buf, err := ioutil.ReadAll(s)
	if err != nil {
		return
	}
	s.Close()

	// unmarshal it
	err = proto.Unmarshal(buf, data)
	if err != nil {
		return
	}

	err = p.reshare.AddMessage(data)
	if err != nil {
		return
	}
}

func (p *service) Process() {
	// 1. Start a reshare process.
	p.reshare.Start()
	defer p.reshare.Stop()

	// 2. Wait the reshare is done or failed
	<-p.done
}

func (p *service) OnStateChanged(oldState types.MainState, newState types.MainState) {
	if newState == types.StateFailed {

		close(p.done)
		return
	} else if newState == types.StateDone {
		result, err := p.reshare.GetResult()
		if err == nil {
			writeReshareResult(p.pm.SelfID(), result)
		} else {

		}
		close(p.done)
		return
	}
}
