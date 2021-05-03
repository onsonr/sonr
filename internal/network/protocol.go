package network

import (
	"encoding/json"
	"log"
	"net/http"

	md "github.com/sonr-io/core/pkg/models"
)

// ^ Returns Location from GeoIP ^ //
func Location(target *md.GeoIP) error {
	r, err := http.Get("https://api.ipgeolocationapi.com/geolocate")
	if err != nil {
		log.Fatalln(err)
	}
	return json.NewDecoder(r.Body).Decode(target)
}
