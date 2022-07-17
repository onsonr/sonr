package schemas

import (
	"github.com/ipfs/go-cid"
	ipfsApi "github.com/ipfs/go-ipfs-api"
)

type Event struct {
	name     string
	previous cid.Cid
}

func fetchEvent(shell ipfsApi.Shell, eventCid cid.Cid) (*Event, error) {
	/*
		builder := eventNodeBuilder{}
		lnk := cidlink.Link{Cid: eventCid}
		err := lnk.Load(
			context.Background(),
			ipld.LinkContext{},
			&builder,
			func(lnk ipld.Link, ctx ipld.LinkContext) (io.Reader, error) {
				theCid, ok := lnk.(cidlink.Link)
				if !ok {
					return nil, fmt.Errorf("Attempted to load a non CID link: %v", lnk)
				}
				block, err := shell.BlockGet(theCid.String())
				if err != nil {
					return nil, fmt.Errorf("error loading %v: %v", theCid.String(), err)
				}
				return bytes.NewBuffer(block), nil
			},
		)

		if err != nil {
			return nil, err
		}

		node := builder.Build()
		event := node.(*Event)
		return event, nil
	*/

	return nil, nil
}
