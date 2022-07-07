package schemas

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/sonr-io/sonr/pkg/tx"
	st "github.com/sonr-io/sonr/x/schema/types"
)

func (as *appSchemaInternalImpl) GetTopLevelNodeById(id string) (datamodel.Node, error) {
	if _, ok := as.nodes[id]; ok {
		return as.nodes[id], nil
	}

	return nil, errIdNotFound
}

func (as *appSchemaInternalImpl) GetPath(id string) (datamodel.ListIterator, error) {
	if _, ok := as.nodes[id]; !ok {
		return nil, errIdNotFound
	}
	node := as.nodes[id]
	return node.ListIterator(), nil
}

func (as *appSchemaInternalImpl) GetNodeMap() map[string]datamodel.Node {
	return as.nodes
}

func (as *appSchemaInternalImpl) GetWhatIsMap() map[string]*st.WhatIs {
	return as.WhatIs
}

func (as *appSchemaInternalImpl) Broadcast(addr string, msg sdk.Msg) ([]byte, error) {
	reqTx, err := tx.BuildTx(as.Acct, addr, msg)
	if err != nil {
		return nil, err
	}
	req, err := reqTx.Marshal()
	if err != nil {
		return nil, err
	}

	apiEndpoint := "http://127.0.0.1:1317/cosmos/tx/v1beta1/txs"
	values := map[string]interface{}{
		"description": addr,
		"tx_bytes":    req,
		"mode":        "BROADCAST_MODE_BLOCK",
	}

	jsonBytes, err := json.Marshal(values)

	if err != nil {
		return nil, err
	}

	res, err := http.Post(apiEndpoint, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	respJson := make(map[string]interface{})
	json.Unmarshal(bytes, &respJson)
	return bytes, nil
}
