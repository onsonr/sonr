package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	net "github.com/sonr-io/core/pkg/net"
)

type IPLocation struct {
	Continent           string   `json:"continent"`
	Alpha2              string   `json:"alpha2"`
	CountryCode         string   `json:"country_code"`
	InternationalPrefix string   `json:"international_prefix"`
	Name                string   `json:"name"`
	CurrencyCode        string   `json:"currency_code"`
	LanguagesSpoken     []string `json:"languages_spoken"`
	Geo                 struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"geo"`
}

func GetLocation() IPLocation {
	address := net.IPv4()

	// There is also /xml/ and /csv/ formats available
	response, err := http.Get("https://api.ipgeolocationapi.com/geolocate/" + address)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	// response.Body() is a reader type. We have
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	// Unmarshal the JSON byte slice to a GeoIP struct
	geo := IPLocation{}
	err = json.Unmarshal(body, &geo)
	if err != nil {
		fmt.Println(err)
	}

	// Everything accessible in struct now
	fmt.Println("Country Code:\t", geo.CountryCode)
	fmt.Println("Country Name:\t", geo.Name)
	fmt.Println("Languages:\t", geo.LanguagesSpoken)
	fmt.Println("Latitude:\t", geo.Geo.Latitude)
	fmt.Println("Longitude:\t", geo.Geo.Longitude)
	fmt.Println("Continent:\t", geo.Continent)
	return geo
}
