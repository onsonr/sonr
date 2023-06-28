// Class to interact with the underlying blockchain to operate on the highway
package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	identitytypes "github.com/sonrhq/core/x/identity/types"
	servicetypes "github.com/sonrhq/core/x/service/types"
	vaulttypes "github.com/sonrhq/core/x/vault/types"
)

const (
	kBaseApiURL = "http://localhost:1317"
	kBaseRpcURL = "http://localhost:26657"
)

// CheckAlias checks if an alias is already registered
func CheckAlias(alias string) (bool, error) {
	endpoint := fmt.Sprintf("%s/core/id/alias/%s/check", kBaseApiURL, alias)
	resp := new(identitytypes.QueryAliasAvailableResponse)
	bz, err := getJsonBytes(endpoint)
	if err != nil {
		return false, err
	}
err = json.Unmarshal(bz, resp)
	if err != nil {
		return false, err
	}
	return resp.Available, nil
}

// GetDID returns the DIDDocument of a given DID or Alias
func GetDID(alias string) (*identitytypes.DIDDocument, error) {
	endpoint := fmt.Sprintf("%s/core/id/alias/%s", kBaseApiURL, alias)
	resp := new(identitytypes.QueryDidByAlsoKnownAsResponse)
	bz, err := getJsonBytes(endpoint)
	if err != nil {
		return nil, err
	}
err = json.Unmarshal(bz, resp)
	if err != nil {
		return nil, err
	}
	return &resp.DidDocument, nil
}

// GetServiceRecord returns the ServiceRecord of a given origin
func GetServiceRecord(origin string) (*servicetypes.ServiceRecord, error) {
	endpoint := fmt.Sprintf("%s/core/service/%s", kBaseApiURL, origin)
	resp := new(servicetypes.QueryGetServiceRecordResponse)
	bz, err := getJsonBytes(endpoint)
	if err != nil {
		return nil, err
	}
err = json.Unmarshal(bz, resp)
	if err != nil {
		return nil, err
	}
	return &resp.ServiceRecord, nil
}
type QueryGetClaimableWalletResponse struct {
	ClaimableWallet ClaimableWallet `protobuf:"bytes,1,opt,name=ClaimableWallet,proto3" json:"ClaimableWallet"`
}

// GetUnclaimedWallet returns the UnclaimedWallet of a given id
func GetUnclaimedWallet(id uint64) (*vaulttypes.ClaimableWallet, error) {
	idstr := strconv.FormatUint(id, 10)
	endpoint := fmt.Sprintf("%s/core/vault/claims/%s", kBaseApiURL, idstr)
	resp := new(QueryGetClaimableWalletResponse)
	bz, err := getJsonBytes(endpoint)
	if err != nil {
		return nil, err
	}
err = json.Unmarshal(bz, resp)
	if err != nil {
		return nil, err
	}
	strucw := resp.ClaimableWallet
	index, err := strconv.ParseUint(strucw.Index, 10, 64)
	if err != nil {
		return nil, err
	}

	ucw := vaulttypes.ClaimableWallet{
		Index:   index,
		Creator: strucw.Creator,
		Address: strucw.Address,
	}
	return &ucw, nil
}

type ClaimableWallet struct {
	Index   string `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

type QueryAllClaimableWalletResponse struct {
	ClaimableWallet []ClaimableWallet   `protobuf:"bytes,1,rep,name=ClaimableWallet,proto3" json:"ClaimableWallet"`
}

// NextUnclaimedWallet returns the next UnclaimedWallet from the queue
func NextUnclaimedWallet() (*vaulttypes.ClaimableWallet, error) {
	endpoint := fmt.Sprintf("%s/core/vault/claims", kBaseApiURL)
	resp := new(QueryAllClaimableWalletResponse)
	bz, err := getJsonBytes(endpoint)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bz, resp)
	if err != nil {
		return nil, err
	}
	strucw := resp.ClaimableWallet[0]
	index, err := strconv.ParseUint(strucw.Index, 10, 64)
	if err != nil {
		return nil, err
	}

	ucw := vaulttypes.ClaimableWallet{
		Index:   index,
		Creator: strucw.Creator,
		Address: strucw.Address,
	}
	return &ucw, nil
}

// ! ||--------------------------------------------------------------------------------||
// ! ||                                Utility Functions                               ||
// ! ||--------------------------------------------------------------------------------||

// getJsonBytes makes a GET request to the given URL and returns the response body as bytes
func getJsonBytes(url string) ([]byte, error) {
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
