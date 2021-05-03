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
	Continent            string  `json:"continent"`
	Country              string  `json:"country"`
	Zipcode              string  `json:"zipCode"`
	Accuracyradius       int     `json:"accuracyRadius"`
	Flag                 string  `json:"flag"`
	City                 string  `json:"city"`
	Timezone             string  `json:"timezone"`
	Latitude             float64 `json:"latitude"`
	Countrygeonameid     int     `json:"countryGeoNameId"`
	Gmt                  string  `json:"gmt"`
	Network              string  `json:"network"`
	Currencyname         string  `json:"currencyName"`
	Countrynativename    string  `json:"countryNativeName"`
	Stategeonameid       int     `json:"stateGeoNameId"`
	Phonecode            string  `json:"phoneCode"`
	State                string  `json:"state"`
	Continentcode        string  `json:"continentCode"`
	Longitude            float64 `json:"longitude"`
	Currencynameplural   string  `json:"currencyNamePlural"`
	Citygeonameid        int     `json:"cityGeoNameId"`
	Languages            string  `json:"languages"`
	Numofcities          int     `json:"numOfCities"`
	Org                  string  `json:"org"`
	IP                   string  `json:"ip"`
	Currencysymbol       string  `json:"currencySymbol"`
	Currencysymbolnative string  `json:"currencySymbolNative"`
	Iseu                 string  `json:"isEU"`
	Countrytld           string  `json:"countryTLD"`
	Countrycapital       string  `json:"countryCapital"`
	Metrocode            int     `json:"metroCode"`
	Continentgeonameid   int     `json:"continentGeoNameId"`
	Statecode            string  `json:"stateCode"`
	Countryiso2          string  `json:"countryISO2"`
	Numofstates          int     `json:"numOfStates"`
	Countryiso3          string  `json:"countryISO3"`
	Currencycode         string  `json:"currencyCode"`
	Asno                 int     `json:"asNo"`
	Status               int     `json:"status"`
}

// Convert to String
func (g *GeoIP) String() string {
	lat := g.Latitude
	lon := g.Longitude
	return fmt.Sprintf("Latitude: %f \n Longitude: %f \n", lat, lon)
}
