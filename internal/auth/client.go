package auth

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	b64 "encoding/base64"

	md "github.com/sonr-io/core/pkg/models"
)

// Base Client Properties
const baseUrl = "https://www.namebase.io"

// Endpoint for Nameserver DNS Settings
const nameDnsPoint = "/api/v0/dns/domains/snr/nameserver"

type NamebaseClient interface {
	AddRecord(record *md.HSRecord) (bool, error)
	DeleteRecord(record *md.HSRecord) (bool, error)
	Refresh() ([]*md.HSRecord, error)
}

type namebaseClient struct {
	NamebaseClient
	apiKeys         *md.APIKeys
	client          *http.Client
	restrictedNames []string
	blockedNames    []string
}

// ^ Method to Create Namebase Client ^ //
func newNambaseClient(keys *md.APIKeys) NamebaseClient {
	return &namebaseClient{
		apiKeys:         keys,
		client:          &http.Client{},
		restrictedNames: []string{"elon", "vitalik", "prad", "rishi", "brax", "vt", "isa"},
		blockedNames:    []string{"root", "admin", "mail", "auth", "crypto", "id", "app", "beta", "alpha", "code", "ios", "android", "test", "node", "sonr"},
	}
}

// ^ Method to Add a HSRecord ^ //
func (nc *namebaseClient) AddRecord(record *md.HSRecord) (bool, error) {
	// Create Body
	nbreq := md.NewNamebaseRequest(record, true)
	json, err := nbreq.JSON()
	if err != nil {
		return false, err
	}

	// Create Request
	req, err := http.NewRequest("PUT", baseUrl+nameDnsPoint, bytes.NewBuffer(json))
	if err != nil {
		return false, err
	}

	// Set Headers and Perform Action
	req = nc.setHeaders(req)
	resp, err := nc.client.Do(req)
	if err != nil {
		return false, err
	}
	// Parse Response Body
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	// Unmarshal Data
	namebaseResponse, err := md.NewNamebaseResponse(bodyBytes)
	if err != nil {
		return false, err
	}

	// Return Records
	return namebaseResponse.GetSuccess(), nil
}

// ^ Method to Add a HSRecord ^ //
func (nc *namebaseClient) DeleteRecord(record *md.HSRecord) (bool, error) {
	// Create Body
	nbreq := md.NewNamebaseRequest(record, true)
	json, err := nbreq.JSON()
	if err != nil {
		return false, err
	}

	// Create Request
	req, err := http.NewRequest("PUT", baseUrl+nameDnsPoint, bytes.NewBuffer(json))
	if err != nil {
		return false, err
	}

	// Set Headers and Perform Action
	req = nc.setHeaders(req)
	resp, err := nc.client.Do(req)
	if err != nil {
		return false, err
	}
	// Parse Response Body
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	// Unmarshal Data
	namebaseResponse, err := md.NewNamebaseResponse(bodyBytes)
	if err != nil {
		return false, err
	}

	// Return Records
	return namebaseResponse.GetSuccess(), nil
}

// ^ Method Returns all known Records ^ //
func (nc *namebaseClient) Refresh() ([]*md.HSRecord, error) {
	// Create Request
	req, err := http.NewRequest("GET", baseUrl+nameDnsPoint, nil)
	if err != nil {
		return nil, err
	}

	// Set Headers and Perform Action
	req = nc.setHeaders(req)
	resp, err := nc.client.Do(req)
	if err != nil {
		return nil, err
	}

	// Parse Response Body
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal Data
	namebaseResponse, err := md.NewNamebaseResponse(bodyBytes)
	if err != nil {
		return nil, err
	}

	// Return Records
	return namebaseResponse.GetRecords(), nil
}

// @ Helper: Sets Request HTTP Headers
func (nc *namebaseClient) setHeaders(req *http.Request) *http.Request {
	// Get Authorization
	data := []byte(fmt.Sprintf("%s:%s", nc.apiKeys.GetHandshakeKey(), nc.apiKeys.GetHandshakeSecret()))
	auth := b64.StdEncoding.EncodeToString(data)

	// Set Headers
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", auth))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	return req
}
