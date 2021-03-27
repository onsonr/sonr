package network

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
)

// ^ Return Bootstrap List Address Info ^ //
func GetBootstrapAddrInfo() ([]peer.AddrInfo, error) {
	// Create Bootstrapper List
	var bootstrappers []multiaddr.Multiaddr
	for _, s := range []string{
		// Libp2p Default
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",
		"/ip4/104.131.131.82/tcp/4001/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
		"/ip4/104.131.131.82/udp/4001/quic/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
	} {
		ma, err := multiaddr.NewMultiaddr(s)
		if err != nil {
			return nil, err
		}
		bootstrappers = append(bootstrappers, ma)
	}

	ds := make([]peer.AddrInfo, 0, len(bootstrappers))
	for i := range bootstrappers {
		info, err := peer.AddrInfoFromP2pAddr(bootstrappers[i])
		if err != nil {
			continue
		}
		ds = append(ds, *info)
	}
	return ds, nil
}

// ^ Geographical Position from IP ^ //
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
