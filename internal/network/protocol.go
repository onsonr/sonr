package network

import (
	"encoding/json"
	"net/http"

	md "github.com/sonr-io/core/pkg/models"
)

// ^ Returns Location from GeoIP ^ //
func Location(target *md.GeoIP) error {
	url := "https://find-any-ip-address-or-domain-location-world-wide.p.rapidapi.com/iplocation?apikey=873dbe322aea47f89dcf729dcc8f60e8"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("x-rapidapi-key", "a09329a7a8mshc7688c2dc3de4f9p197eb4jsn186162847bfb")
	req.Header.Add("x-rapidapi-host", "find-any-ip-address-or-domain-location-world-wide.p.rapidapi.com")

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	return json.NewDecoder(r.Body).Decode(target)
}
