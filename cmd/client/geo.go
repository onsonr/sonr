package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	net "github.com/sonr-io/core/pkg/net"
)

type GeoIP struct {
	// The right side is the name of the JSON variable
	Ip          string  `json:"ip"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
	RegionCode  string  `json:"region_code"`
	RegionName  string  `json:"region_name"`
	City        string  `json:"city"`
	Zipcode     string  `json:"zipcode"`
	Lat         float32 `json:"latitude"`
	Lon         float32 `json:"longitude"`
	MetroCode   int     `json:"metro_code"`
	AreaCode    int     `json:"area_code"`
}

func GetLocation() GeoIP {
	address := net.IPv4()

	// There is also /xml/ and /csv/ formats available
	response, err := http.Get("https://freegeoip.net/json/" + address)
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
	geo := GeoIP{}
	err = json.Unmarshal(body, &geo)
	if err != nil {
		fmt.Println(err)
	}

	// Everything accessible in struct now
	fmt.Println("IP address:\t", geo.Ip)
	fmt.Println("Country Code:\t", geo.CountryCode)
	fmt.Println("Country Name:\t", geo.CountryName)
	fmt.Println("Zip Code:\t", geo.Zipcode)
	fmt.Println("Latitude:\t", geo.Lat)
	fmt.Println("Longitude:\t", geo.Lon)
	fmt.Println("Metro Code:\t", geo.MetroCode)
	fmt.Println("Area Code:\t", geo.AreaCode)
	return geo
}
