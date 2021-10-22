package api

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/errors"
)

func initClient() (*namebaseAPIClient, error) {
	// Define logger warning message
	logWarning := func() error {
		logger.Warn("Could not fetch Namebase API Keys from Env, skipping API client...")
		return ErrNBKeys
	}

	// Get API Key
	key, ok := os.LookupEnv("NB_KEY")
	if !ok {
		return nil, logWarning()
	}

	// Get API Secret
	secret, ok := os.LookupEnv("NB_SECRET")
	if !ok {
		return nil, logWarning()
	}

	nbClient := &namebaseAPIClient{
		ctx:          context.TODO(),
		restClient:   &http.Client{},
		apiClientKey: key,
		apiSecretKey: secret,
	}
	return nbClient, nil
}

// NamebaseRequest for either Adding or Removing DNS Records
type NamebaseRequest struct {
	// Records to be added to DNS Table
	Records []Record `json:"records"`

	// DeleteRecords are to be deleted from DNS Table
	DeleteRecords []DeleteRecord `json:"deleteRecords"`
}

// NewNamebaseRequest creates a new NamebaseRequest for adding records
func newNBAddRequest(records ...Record) NamebaseRequest {
	return NamebaseRequest{
		Records:       records,
		DeleteRecords: make([]DeleteRecord, 0),
	}
}

// newNBDeleteRequest creates a new NamebaseRequest for deleting records
func newNBDeleteRequest(records ...DeleteRecord) NamebaseRequest {
	return NamebaseRequest{
		Records:       make([]Record, 0),
		DeleteRecords: records,
	}
}

// NamebaseResponse is JSON Response for NamebaseRequest
type NamebaseResponse struct {
	// Success is true if the request was successful
	Success bool `json:"success"`

	// Records is the list of records from GET request
	Records []Record `json:"records"`
}

// Print prints the NamebaseResponse
func (nr *NamebaseResponse) Print() {
	// Loop through all records
	for _, record := range nr.Records {
		record.Print()
	}
}

// DeleteRecord is for Removing Records in Request
type DeleteRecord struct {
	// Type is the type of record to be deleted
	Type string `json:"type"`

	// Host is the hostname of the record to be deleted
	Host string `json:"host"`
}

// NewNamebaseDeleteRecord creates a new DeleteRecord
func NewNBDeleteRecord(host string) DeleteRecord {
	return DeleteRecord{
		Type: "TXT",
		Host: host,
	}
}

// LookupTXT looks up the TXT record for the given SName.
func LookupTXT(ctx context.Context, name string) (Records, error) {
	// Create a new net.Resolver
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			// Create Dialer
			d := net.Dialer{
				Timeout: DIAL_TIMEOUT,
			}

			// Dial First Resolver
			c, err := d.DialContext(ctx, network, HDNS_RESOLVER_ONE)
			if err == nil {
				return c, nil
			}

			// Dial Second Resolver
			c, err = d.DialContext(ctx, network, HDNS_RESOLVER_TWO)
			if err == nil {
				return c, nil
			}

			// Dial Third Resolver
			c, err = d.DialContext(ctx, network, HDNS_RESOLVER_THREE)
			if err == nil {
				return c, nil
			}

			// Return Error if we failed to dial all three resolvers
			return nil, ErrHDNSResolve
		},
	}
	// Call internal resolver
	recs, err := r.LookupTXT(ctx, strings.ToLower(name))
	if err != nil {
		return nil, err
	}

	// Check Record count
	if len(recs) == 0 {
		return nil, ErrEmptyTXT
	} else if len(recs) > 1 {
		return nil, ErrMultipleRecords
	} else {
		// Create NB records
		records := make([]Record, len(recs))
		for _, rec := range recs {
			records = append(records, NewNBRecord(name, rec))
		}
		return records, nil
	}
}

// namebaseAPIClient handles DNS Table Registration and Verification.
type namebaseAPIClient struct {
	ctx          context.Context // Context of Protocol
	restClient   *http.Client    // REST Client
	apiClientKey string          // Namebase API Client Key
	apiSecretKey string          // Namebase API Secret Key
}

// newNBGetRequest returns a GET Request for the Namebase REST API call
func (p *namebaseAPIClient) newNBGetRequest() (*http.Request, error) {
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

// newNBPutRequest returns a PUT Request for the Namebase REST API call
func (p *namebaseAPIClient) newNBPutRequest(nbReq NamebaseRequest) (*http.Request, error) {
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
func FindRecords(ctx context.Context, query string) ([]Record, error) {
	// Fetch records from Namebase
	recs, err := GetRecords(ctx)
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
func GetRecords(ctx context.Context) ([]Record, error) {
	// Init Client
	nbClient, err := initClient()
	if err != nil {
		return nil, err
	}

	// Create new GET request
	req, err := nbClient.newNBGetRequest()
	if err != nil {
		return nil, err
	}

	// Execute request
	resp, err := nbClient.restClient.Do(req)
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
		return nil, errors.Wrap(err, "Failed to Get Namebase Records")
	}
}

// HasRecords returns true if the given hostname has a record on Root TLD
func HasRecords(ctx context.Context, query string) (bool, error) {
	// Fetch records from Namebase
	recs, err := FindRecords(ctx, query)
	if err != nil {
		return false, err
	}

	// Return true if records were found
	return len(recs) > 0, nil
}

// PutRecords adds/deletes records to/from Root TLD
func PutRecords(ctx context.Context, recs ...Record) (bool, error) {
	// Init Client
	nbClient, err := initClient()
	if err != nil {
		return false, err
	}
	// Create new GET request
	req, err := nbClient.newNBPutRequest(newNBAddRequest(recs...))
	if err != nil {
		return false, err
	}

	// Execute request
	resp, err := nbClient.restClient.Do(req)
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

// RemoveRecords removes records from Root TLD
func RemoveRecords(ctx context.Context, recs ...Record) (bool, error) {
	// Init Client
	nbClient, err := initClient()
	if err != nil {
		return false, err
	}

	// Convert records to DeleteRecord
	deleteRecords := make([]DeleteRecord, len(recs))
	for _, rec := range recs {
		deleteRecords = append(deleteRecords, rec.ToDeleteRecord())
	}

	// Create new PUT request
	req, err := nbClient.newNBPutRequest(newNBDeleteRequest(deleteRecords...))
	if err != nil {
		return false, err
	}

	// Execute request
	resp, err := nbClient.restClient.Do(req)
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
