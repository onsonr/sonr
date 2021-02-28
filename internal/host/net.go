package host

import (
	"log"
	"net"
	"net/http"
	"os"

	md "github.com/sonr-io/core/internal/models"
	"golang.org/x/net/html"
)

// ^ IPv4 returns the non loopback local IP of the host as IPv4 ^
func IPv4() string {
	// @1. Find IPv4 Address
	osHost, _ := os.Hostname()
	addrs, _ := net.LookupIP(osHost)
	var ipv4Ref string

	// Iterate through addresses
	for _, addr := range addrs {
		// @ Set IPv4
		if ipv4 := addr.To4(); ipv4 != nil {
			ipv4Ref = ipv4.String()
		} else {
			ipv4Ref = "0.0.0.0"
		}
	}
	// No IPv4 Found
	return ipv4Ref
}

// ^ IPv4 returns the non loopback local IP of the host as IPv4 ^
func IPv6() string {
	// @1. Find IPv4 Address
	osHost, _ := os.Hostname()
	addrs, _ := net.LookupIP(osHost)
	var ipv6Ref string

	// Iterate through addresses
	for _, addr := range addrs {
		// @ Set IPv4
		if ipv6 := addr.To16(); ipv6 != nil {
			ipv6Ref = ipv6.String()
		} else {
			ipv6Ref = "::"
		}
	}
	// No IPv4 Found
	return ipv6Ref
}

// ^ Retreives URL Metadata ^ //
func ExtractURL(link string) *md.URLLink {
	// Create Request
	resp, err := http.Get(link)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()

	// Tokenize Response
	z := html.NewTokenizer(resp.Body)
	titleFound := false
	hm := new(md.URLLink)

	// Iterate through URL Elements
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return hm
		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()
			if t.Data == `body` {
				return hm
			}
			if t.Data == "title" {
				titleFound = true
			}
			if t.Data == "meta" {
				desc, ok := extractMetaProperty(t, "description")
				if ok {
					hm.Description = desc
				}

				ogTitle, ok := extractMetaProperty(t, "og:title")
				if ok {
					hm.Title = ogTitle
				}

				ogDesc, ok := extractMetaProperty(t, "og:description")
				if ok {
					hm.Description = ogDesc
				}

				ogImage, ok := extractMetaProperty(t, "og:image")
				if ok {
					hm.Image = ogImage
				}

				ogSiteName, ok := extractMetaProperty(t, "og:site_name")
				if ok {
					hm.SiteName = ogSiteName
				}
				return hm
			}
		case html.TextToken:
			if titleFound {
				t := z.Token()
				hm.Title = t.Data
				return hm
			}
		}
	}
}

// ^ Helper: Extracts a Meta Property ^ //
func extractMetaProperty(t html.Token, prop string) (content string, ok bool) {
	for _, attr := range t.Attr {
		if attr.Key == "property" && attr.Val == prop {
			ok = true
		}

		if attr.Key == "content" {
			content = attr.Val
		}
	}
	return
}
