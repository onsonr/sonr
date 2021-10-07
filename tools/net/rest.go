package net

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/kataras/golog"
	"github.com/pkg/errors"
)

// Constant for the Namebase API
const (
	DOMAIN              = "snr"
	API_ADDRESS         = "https://www.namebase.io/api/v0/"
	DNS_ENDPOINT        = "dns/domains/snr/nameserver"
	FINGERPRINT_DIVIDER = "v=0;fingerprint="
	AUTH_DIVIDER        = "._auth."
	NB_DNS_API_URL      = API_ADDRESS + DNS_ENDPOINT
)

// Error Definitions
var (
	logger = golog.Child("Tools/Net")
	ErrGetNamebase = errors.New("Failed to perform GET Request on Namebase API")
	ErrPutNamebase = errors.New("Failed to perform PUT Request on Namebase API")
)

// NamebaseAPIClient handles DNS Table Registration and Verification.
type NamebaseAPIClient struct {
	ctx          context.Context // Context of Protocol
	restClient   *http.Client    // REST Client
	apiClientKey string          // Namebase API Client Key
	apiSecretKey string          // Namebase API Secret Key
}

// NewNamebaseClient returns a new NamebaseAPIClient
func NewNamebaseClient(ctx context.Context, apiKey, apiSecret string) *NamebaseAPIClient {
	return &NamebaseAPIClient{
		ctx:          ctx,
		restClient:   &http.Client{},
		apiClientKey: apiKey,
		apiSecretKey: apiSecret,
	}
}

// NewNBGetRequest returns a GET Request for the Namebase REST API call
func (p *NamebaseAPIClient) NewNBGetRequest() (*http.Request, error) {
	// Create the request
	req, err := http.NewRequest("GET", NB_DNS_API_URL, nil)
	if err != nil {
		logger.Error("Failed to create Namebase HTTP Request", err)
		return nil, err
	}

	// Set the headers
	req.SetBasicAuth(p.apiClientKey, p.apiSecretKey)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

// NewNBPutRequest returns a PUT Request for the Namebase REST API call
func (p *NamebaseAPIClient) NewNBPutRequest(nbReq NamebaseRequest) (*http.Request, error) {
	// Marshal the request
	jsonreq, err := json.Marshal(nbReq)
	if err != nil {
		logger.Error("Failed to marshal Namebase Request", err)
		return nil, err
	}
	bytes := bytes.NewBuffer(jsonreq)

	// Create the request
	req, err := http.NewRequest("PUT", NB_DNS_API_URL, bytes)
	if err != nil {
		logger.Error("Failed to create Namebase HTTP Request", err)
		return nil, err
	}

	// Set the headers
	req.SetBasicAuth(p.apiClientKey, p.apiSecretKey)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

// FindRecords returns a list of records matching the given query
func (p *NamebaseAPIClient) FindRecords(query string) ([]Record, error) {
	// Fetch records from Namebase
	recs, err := p.GetRecords()
	if err != nil {
		return nil, err
	}

	// Filter records
	filtered := []Record{}
	for _, rec := range recs {
		// Clean up query
		val := strings.ToLower(rec.Host)
		q := strings.ToLower(strings.TrimSpace(query))

		// Check if query matches
		if strings.Contains(val, q) {
			filtered = append(filtered, rec)
		}
	}

	// Return filtered records
	return filtered, nil
}

// GetRecords returns a list of all records on Root TLD
func (p *NamebaseAPIClient) GetRecords() ([]Record, error) {
	// Create new GET request
	req, err := p.NewNBGetRequest()
	if err != nil {
		return nil, err
	}

	// Execute request
	resp, err := p.restClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Read Response Body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal Response Body
	nbResponse := &NamebaseResponse{}
	err = json.Unmarshal(respBody, nbResponse)
	if err != nil {
		return nil, err
	}

	// Handle Response
	if nbResponse.Success {
		nbResponse.Print()
		return nbResponse.Records, nil
	} else {
		return nil, ErrGetNamebase
	}
}

// HasRecords returns true if the given hostname has a record on Root TLD
func (p *NamebaseAPIClient) HasRecords(query string) (bool, error) {
	// Fetch records from Namebase
	recs, err := p.FindRecords(query)
	if err != nil {
		return false, err
	}

	// Return true if records were found
	return len(recs) > 0, nil
}

// PutRecords adds/deletes records to/from Root TLD
func (p *NamebaseAPIClient) PutRecords(nbReq NamebaseRequest) (bool, error) {
	// Create new GET request
	req, err := p.NewNBPutRequest(nbReq)
	if err != nil {
		return false, err
	}

	// Execute request
	resp, err := p.restClient.Do(req)
	if err != nil {
		return false, err
	}

	// Read Response Body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	// Unmarshal Response Body
	nbResponse := &NamebaseResponse{}
	err = json.Unmarshal(respBody, nbResponse)
	if err != nil {
		return false, err
	}

	// Return Response
	return nbResponse.Success, nil
}
