package models

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// ** ─── CALLBACK MANAGEMENT ────────────────────────────────────────────────────────
// Define Function Types
type GetStatus func() Status
type SetStatus func(s Status)
type GetContact func() *Contact
type OnProtobuf func([]byte)
type OnInvite func(data []byte)
type OnProgress func(data float32)
type OnReceived func(data *TransferCard)
type OnTransmitted func(data *TransferCard)
type OnError func(err *SonrError)
type NodeCallback struct {
	Contact     GetContact
	Invited     OnInvite
	Refreshed   OnProtobuf
	Event       OnProtobuf
	RemoteStart OnProtobuf
	Responded   OnProtobuf
	Progressed  OnProgress
	Received    OnReceived
	Status      SetStatus
	Transmitted OnTransmitted
	Error       OnError
	GetStatus   GetStatus
}

// @ Binary State Management ** //
type state struct {
	flag uint64
	chn  chan bool
}

var (
	instance *state
	once     sync.Once
)

func GetState() *state {
	once.Do(func() {
		chn := make(chan bool)
		close(chn)

		instance = &state{chn: chn}
	})

	return instance
}

// Checks rather to wait or does not need
func (c *state) NeedsWait() {
	<-c.chn
}

// Says all of goroutines to resume execution
func (c *state) Resume() {
	if atomic.LoadUint64(&c.flag) == 1 {
		close(c.chn)
		atomic.StoreUint64(&c.flag, 0)
	}
}

// Says all of goroutines to pause execution
func (c *state) Pause() {
	if atomic.LoadUint64(&c.flag) == 0 {
		atomic.StoreUint64(&c.flag, 1)
		c.chn = make(chan bool)
	}
}

// ** ─── GEOLACTION FROM IP ADDRESS ─────────────────────────────────────────────────
// Geographical Position from IP ^ //
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

// Convert to String
func (g *GeoIP) String() string {
	lat := g.Geo.Latitude
	lon := g.Geo.Longitude
	maxlat := g.Geo.MaxLatitude
	maxlon := g.Geo.MaxLongitude
	minlat := g.Geo.MinLatitude
	minlon := g.Geo.MinLongitude
	return fmt.Sprintf("Latitude: %f \n Longitude: %f \n Max Latitude: %f \n Max Longitude: %f \n Min Latitude: %f \n Min Longitude: %f", lat, lon, maxlat, maxlon, minlat, minlon)
}
