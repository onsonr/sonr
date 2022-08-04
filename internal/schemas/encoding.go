package schemas

import (
	"bytes"

	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/node/basicnode"
)

func (as *schemaImpl) EncodeDagJson() ([]byte, error) {
	if as.nodes == nil {
		return nil, errNodeNotFound
	}

	buffer := bytes.Buffer{}
	err := dagjson.Encode(as.nodes, &buffer)

	return buffer.Bytes(), err
}

func (as *schemaImpl) EncodeDagCbor() ([]byte, error) {
	if as.nodes == nil {
		return nil, nil
	}

	buffer := &bytes.Buffer{}
	err := dagcbor.Encode(as.nodes, buffer)
	return buffer.Bytes(), err
}

func (as *schemaImpl) DecodeDagJson(buffer []byte) error {
	np := basicnode.Prototype.Any
	nb := np.NewBuilder()

	reader := bytes.NewReader(buffer)
	err := dagjson.Decode(nb, reader)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	node := nb.Build()

	as.nodes = node

	return nil
}

func (as *schemaImpl) DecodeDagCbor(buffer []byte) error {
	np := basicnode.Prototype.Any
	nb := np.NewBuilder()

	reader := bytes.NewReader(buffer)
	err := dagcbor.Decode(nb, reader)
	if err != nil {
		return err
	}
	as.nodes = nb.Build()

	return nil
}
