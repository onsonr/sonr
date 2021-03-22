package net

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"
)

type GeoIP struct {
	Continent                      string   `json:"continent"`
	AddressFormat                  string   `json:"address_format"`
	Alpha2                         string   `json:"alpha2"`
	Alpha3                         string   `json:"alpha3"`
	CountryCode                    string   `json:"country_code"`
	InternationalPrefix            string   `json:"international_prefix"`
	Ioc                            string   `json:"ioc"`
	Gec                            string   `json:"gec"`
	Name                           string   `json:"name"`
	NationalDestinationCodeLengths []int    `json:"national_destination_code_lengths"`
	NationalNumberLengths          []int    `json:"national_number_lengths"`
	NationalPrefix                 string   `json:"national_prefix"`
	Number                         string   `json:"number"`
	Region                         string   `json:"region"`
	Subregion                      string   `json:"subregion"`
	WorldRegion                    string   `json:"world_region"`
	UnLocode                       string   `json:"un_locode"`
	Nationality                    string   `json:"nationality"`
	PostalCode                     bool     `json:"postal_code"`
	UnofficialNames                []string `json:"unofficial_names"`
	LanguagesOfficial              []string `json:"languages_official"`
	LanguagesSpoken                []string `json:"languages_spoken"`
	Geo                            struct {
		Latitude     float64 `json:"latitude"`
		LatitudeDec  string  `json:"latitude_dec"`
		Longitude    float64 `json:"longitude"`
		LongitudeDec string  `json:"longitude_dec"`
		MaxLatitude  float64 `json:"max_latitude"`
		MaxLongitude float64 `json:"max_longitude"`
		MinLatitude  float64 `json:"min_latitude"`
		MinLongitude float64 `json:"min_longitude"`
		Bounds       struct {
			Northeast struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"northeast"`
			Southwest struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"southwest"`
		} `json:"bounds"`
	} `json:"geo"`
	CurrencyCode string `json:"currency_code"`
	StartOfWeek  string `json:"start_of_week"`
}

// @ Returns Node Public IPv4 Address
func IPv4() string {
	osHost, _ := os.Hostname()
	addrs, _ := net.LookupIP(osHost)
	ipv4Ref := "0.0.0.0"
	// Iterate through addresses
	for _, addr := range addrs {
		// @ Set IPv4
		if ipv4 := addr.To4(); ipv4 != nil {
			ipv4Ref = ipv4.String()
		}
	}
	return ipv4Ref
}

// @ Returns Node Public IPv6 Address
func IPv6() string {
	osHost, _ := os.Hostname()
	addrs, _ := net.LookupIP(osHost)
	ipv6Ref := "::"

	// Iterate through addresses
	for _, addr := range addrs {
		// @ Set IPv4
		if ipv6 := addr.To16(); ipv6 != nil {
			ipv6Ref = ipv6.String()
		}
	}
	return ipv6Ref
}

func Location(target *GeoIP) error {
	r, err := http.Get("https://api.ipgeolocationapi.com/geolocate")
	if err != nil {
		log.Fatalln(err)
	}
	return json.NewDecoder(r.Body).Decode(target)
}
