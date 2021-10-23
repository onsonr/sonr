package api

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/sonr-io/core/pkg/common"
)

const (
	// LOCATION_API_HOST is the hostname of the location api
	LOCATION_API_URL = "https://find-any-ip-address-or-domain-location-world-wide.p.rapidapi.com/iplocation?"

	// LOCATION_API_URL is the url to perform Location request
	LOCATION_API_HOST = "find-any-ip-address-or-domain-location-world-wide.p.rapidapi.com"
)

// GetIPLocation returns location of the IP address
func GetLocation() *common.Location {
	// Lookup Location API Key
	locKey, ok := os.LookupEnv("LOCATION_KEY")
	if !ok {
		logger.Warn("Location or RapidAPI Keys missing from Env, using default...")
		return defaultLocation()
	}

	// Lookup Rapid API Key
	rapidKey, ok := os.LookupEnv("RAPID_KEY")
	if !ok {
		logger.Warn("Location or RapidAPI Keys missing from Env, using default...")
		return defaultLocation()
	}

	// Get the IP address of the client.
	url := LOCATION_API_URL + locKey

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Errorf("Error creating request: %s", err)
		return defaultLocation()
	}

	// Set the header to specify the content type
	req.Header.Add("x-rapidapi-host", LOCATION_API_HOST)
	req.Header.Add("x-rapidapi-key", rapidKey)

	// Await a response from the api
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Errorf("Error getting response: %s", err)
		return defaultLocation()
	}

	// Get Location Object
	obj, err := createLocationObject(res)
	if err != nil {
		logger.Errorf("Error decoding JSON to Location Object: %s", err)
		return defaultLocation()
	}
	return obj.ToLocation()
}

// defaultLocation is the default location if the location is not found
func defaultLocation() *common.Location {
	return &common.Location{
		Latitude:  34.102920,
		Longitude: -118.394190,
		Placemark: &common.Location_Placemark{
			Country:            "US",
			PostalCode:         "90210",
			IsoCountryCode:     "US",
			AdministrativeArea: "CA",
			Locality:           "Los Angeles",
		},
	}
}

// LocationObject is the object that is returned by the API
type LocationObject struct {
	Continent            string  `json:"continent"`
	Country              string  `json:"country"`
	ZipCode              string  `json:"zipCode"`
	AccuracyRadius       int     `json:"accuracyRadius"`
	Flag                 string  `json:"flag"`
	City                 string  `json:"city"`
	Timezone             string  `json:"timezone"`
	Latitude             float64 `json:"latitude"`
	CountryGeoNameID     int     `json:"countryGeoNameId"`
	Gmt                  string  `json:"gmt"`
	Network              string  `json:"network"`
	CurrencyName         string  `json:"currencyName"`
	CountryNativeName    string  `json:"countryNativeName"`
	StateGeoNameID       int     `json:"stateGeoNameId"`
	PhoneCode            string  `json:"phoneCode"`
	State                string  `json:"state"`
	ContinentCode        string  `json:"continentCode"`
	Longitude            float64 `json:"longitude"`
	CurrencyNamePlural   string  `json:"currencyNamePlural"`
	CityGeoNameID        int     `json:"cityGeoNameId"`
	Languages            string  `json:"languages"`
	NumOfCities          int     `json:"numOfCities"`
	Org                  string  `json:"org"`
	IP                   string  `json:"ip"`
	CurrencySymbol       string  `json:"currencySymbol"`
	CurrencySymbolNative string  `json:"currencySymbolNative"`
	IsEU                 string  `json:"isEU"`
	CountryTLD           string  `json:"countryTLD"`
	CountryCapital       string  `json:"countryCapital"`
	MetroCode            int     `json:"metroCode"`
	ContinentGeoNameID   int     `json:"continentGeoNameId"`
	StateCode            string  `json:"stateCode"`
	CountryISO2          string  `json:"countryISO2"`
	NumOfStates          int     `json:"numOfStates"`
	CountryISO3          string  `json:"countryISO3"`
	CurrencyCode         string  `json:"currencyCode"`
	AsNo                 int     `json:"asNo"`
	Status               int     `json:"status"`
}

// createLocationObject is a helper function to create a LocationObject from the response
func createLocationObject(r *http.Response) (LocationObject, error) {
	defer r.Body.Close()
	// Declare a new Person struct.
	var p LocationObject

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		return p, err
	}
	return p, nil
}

// ToLocation converts LocationObject to common.Location
func (l LocationObject) ToLocation() *common.Location {
	return &common.Location{
		Latitude:  l.Latitude,
		Longitude: l.Longitude,
		Placemark: &common.Location_Placemark{
			Country:            l.Country,
			PostalCode:         l.ZipCode,
			IsoCountryCode:     l.CountryISO2,
			AdministrativeArea: l.State,
			Locality:           l.City,
		},
	}
}
