package local

import (
	"bytes"
	"io"
	"net/http"
	"time"

	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
)

// BroadcastTxResponse is a type alias for the BroadcastTxResponse type from the Cosmos SDK
type BroadcastTxResponse = txtypes.BroadcastTxResponse

// ! ||--------------------------------------------------------------------------------||
// ! ||                               Tendermint Node RPC                              ||
// ! ||--------------------------------------------------------------------------------||

// GetJSON makes a GET request to the given URL and returns the response body as bytes
func GetJSON(url string) ([]byte, error) {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, r.Body)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// PostJSON makes a POST request to the given URL and returns the response body as bytes
func PostJSON(url string, body []byte) ([]byte, error) {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, r.Body)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
