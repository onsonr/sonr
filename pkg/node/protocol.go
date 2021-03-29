package node

import (
	"log"

	brprot "berty.tech/berty/v2/go/pkg/bertyprotocol"
	brtypes "berty.tech/berty/v2/go/pkg/protocoltypes"
)

func (n *Node) StartProtocol() error {
	client, err := brprot.New(n.ctx, brprot.Opts{})
	if err != nil {
		return err
	}

	ret, err := client.InstanceGetConfiguration(n.ctx, &brtypes.InstanceGetConfiguration_Request{})
	if err != nil {
		return err
	}

	for _, listener := range ret.Listeners {
		log.Println(listener)
	}
	return nil
}
