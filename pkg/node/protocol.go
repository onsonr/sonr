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
	n.client = client

	resp, err := client.MultiMemberGroupCreate(n.ctx, &brtypes.MultiMemberGroupCreate_Request{})
	if err != nil {
		return err
	}
	log.Println(string(resp.GroupPK))

	if err != nil {
		return err
	}
	return nil
}
