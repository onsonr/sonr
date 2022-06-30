package schemas

import (
	"bytes"

	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/datamodel"
)

func (as *appSchemaInternalImpl) EncodeDagJson(node datamodel.Node) ([]byte, error) {
	buffer := &bytes.Buffer{}
	err := dagjson.Encode(node, buffer)

	return buffer.Bytes(), err
}

func (as *appSchemaInternalImpl) EncodeDagCbor(node datamodel.Node) ([]byte, error) {
	buffer := &bytes.Buffer{}
	err := dagcbor.Encode(node, buffer)
	return buffer.Bytes(), err
}

func (as *appSchemaInternalImpl) DecodeDagJson(buffer []byte) (datamodel.Node, error) {
	var asmblr datamodel.NodeAssembler
	reader := bytes.NewReader(buffer)
	err := dagjson.Decode(asmblr, reader)
	if err != nil {
		return nil, err
	}
	builder := asmblr.Prototype().NewBuilder()
	node := builder.Build()

	return node, nil
}

func (as *appSchemaInternalImpl) DecodeDagCbor(buffer []byte) (datamodel.Node, error) {
	var asmblr datamodel.NodeAssembler
	reader := bytes.NewReader(buffer)
	err := dagcbor.Decode(asmblr, reader)
	if err != nil {
		return nil, err
	}
	builder := asmblr.Prototype().NewBuilder()
	node := builder.Build()

	return node, nil
}
