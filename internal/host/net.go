package host

import (
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
func ExtractURLData(link string) *md.URLLink {
	// Initialize
	titleFound := false
	ul := new(md.URLLink)
	ul.Url = link

	// Create Request
	resp, err := http.Get(link)
	if err != nil {
		return ul
	}
	defer resp.Body.Close()

	// Tokenize Response
	z := html.NewTokenizer(resp.Body)

	// Iterate through URL Elements
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			ul = analyzeData(ul)
			return ul
		case html.StartTagToken:
			t := z.Token()
			// Title Tag
			if t.Data == "title" {
				titleFound = true
				title, ok := extractMetaProperty(t, "title")
				if ok {
					ul.Title = title
				}
			}

			// Meta Tags
			if t.Data == "meta" {
				desc, ok := extractMetaProperty(t, "description")
				if ok {
					ul.Description = desc
				}

				ogTitle, ok := extractMetaProperty(t, "og:title")
				if ok {
					ul.OpenGraph.Title = ogTitle
				}

				ogDesc, ok := extractMetaProperty(t, "og:description")
				if ok {
					ul.OpenGraph.Description = ogDesc
				}

				ogImage, ok := extractMetaProperty(t, "og:image")
				if ok {
					ul.OpenGraph.Image = ogImage
				}

				ogSiteName, ok := extractMetaProperty(t, "og:site_name")
				if ok {
					ul.OpenGraph.SiteName = ogSiteName
				}
				twCard, ok := extractMetaProperty(t, "twitter:card")
				if ok {
					ul.Twitter.Card = twCard
				}

				twDomain, ok := extractMetaProperty(t, "twitter:domain")
				if ok {
					ul.Twitter.Domain = twDomain
				}

				twUrl, ok := extractMetaProperty(t, "twitter:url")
				if ok {
					ul.Twitter.Url = twUrl
				}
				twTitle, ok := extractMetaProperty(t, "twitter:title")
				if ok {
					ul.Twitter.Title = twTitle
				}

				twDesc, ok := extractMetaProperty(t, "twitter:description")
				if ok {
					ul.Twitter.Description = twDesc
				}

				twImage, ok := extractMetaProperty(t, "twitter:image")
				if ok {
					ul.Twitter.Image = twImage
				}
				ul = analyzeData(ul)
			}
		case html.TextToken:
			if titleFound {
				t := z.Token()
				ul.Title = t.Data
				titleFound = false
			}
		}
	}
}

// ^ Helper: Analyzes Extracted data to set type ^ //
func analyzeData(data *md.URLLink) *md.URLLink {
	if data.Twitter != nil {
		data.Type = md.URLLink_TWITTER
		return data
	} else if data.OpenGraph != nil {
		data.Type = md.URLLink_OPENGRAPH
		return data
	} else {
		data.Type = md.URLLink_DEFAULT
		return data
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
