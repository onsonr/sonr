package exchange

type HDNSRecord struct {
	Host  string `json:"host"`
	TTL   int    `json:"ttl"`
	Value string `json:"value"`
}

type HDNSRequest struct {
	Domain  string       `json:"domain"`
	Records []HDNSRecord `json:"records"`
}

type HDNSResponse struct {
	Status  string   `json:"status"`
	Records []string `json:"records"`
}
