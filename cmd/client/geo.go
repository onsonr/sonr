package client

import (
	"github.com/ip2location/ip2location-go"
	net "github.com/sonr-io/core/pkg/net"
)

type IPLocation struct {
	CountryShort string
	CountryLong  string
	Latitude     float64
	Longitude    float64
}

func GetLocation() (*IPLocation, error) {
	address := net.IPv4()
	db, err := ip2location.OpenDB("./IP-COUNTRY-REGION-CITY-LATITUDE-LONGITUDE-ZIPCODE-TIMEZONE-ISP-DOMAIN-NETSPEED-AREACODE-WEATHER-MOBILE-ELEVATION-USAGETYPE.BIN")

	if err != nil {
		return nil, err
	}

	results, err := db.Get_all(address)
	if err != nil {
		return nil, err
	}
	
	return &IPLocation{
		CountryShort: results.Country_short,
		CountryLong:  results.Country_long,
		Latitude:     float64(results.Latitude),
		Longitude:    float64(results.Longitude),
	}, nil
}
