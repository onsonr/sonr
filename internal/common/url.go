package common

import (
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"github.com/sonr-io/core/tools/net"
)

// IsUrl returns true if the given string is a valid url
func IsUrl(s string) bool {
	_, err := url.ParseRequestURI(s)
	if err != nil {
		return false
	}

	u, err := url.Parse(s)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}
	return true
}

// NewUrlItem creates a new transfer url item
func NewUrlItem(address string) (*Payload_Item, error) {
	// Create UrlItem
	link := &UrlItem{
		Mime: DefaultUrlMIME(),
		Link: address,
	}

	// Fetch Data and Set Primary
	primary, err := link.FetchData()
	if err != nil {
		return nil, err
	}

	// Return Transfer Item
	return &Payload_Item{
		Size: 0,
		Mime: DefaultUrlMIME(),
		Data: &Payload_Item_Url{
			Url: link,
		},
		Preview: &Payload_Item_OpenGraph{
			OpenGraph: primary,
		},
	}, nil
}

// ** ─── URLLink MANAGEMENT ────────────────────────────────────────────────────────
// FetchData Sets URLLink Data
func (u *UrlItem) FetchData() (*OpenGraph_Primary, error) {
	// Set Mime
	if u.Mime == nil {
		u.Mime = DefaultUrlMIME()
	}

	// Create Request
	resp, err := http.Get(u.Link)
	if err != nil {
		return nil, err
	}

	// Get Info
	info, err := net.GetPageData(resp)
	if err != nil {
		return nil, err
	}

	// Set Link
	u.Title = info.Title
	u.Site = info.Site
	u.SiteName = info.SiteName
	u.Description = info.Description

	// Set OpenGraph
	primary, err := u.AddOpenGraph(info)
	if err != nil {
		return nil, err
	}
	return primary, nil
}

// AddOpenGraph method adds OpenGraph data to the UrlItem
func (u *UrlItem) AddOpenGraph(info *net.PageInfo) (*OpenGraph_Primary, error) {
	// Initialize OpenGraph
	ogImages := make([]*OpenGraph_Image, 0)
	ogVideos := make([]*OpenGraph_Video, 0)
	ogAudios := make([]*OpenGraph_Audio, 0)
	twitter := &OpenGraph_Twitter{}
	ogPrimary := &OpenGraph_Primary{
		Type: OpenGraph_NONE,
	}

	// Get Audios
	if info.Audios != nil {
		// Get Audios
		for _, v := range info.Audios {
			ogAudios = append(ogAudios, &OpenGraph_Audio{
				Url:       v.Url,
				SecureUrl: v.SecureUrl,
				Type:      v.Type,
			})
		}

		// Set OpenGraph Type
		ogPrimary = &OpenGraph_Primary{
			Type: OpenGraph_AUDIO,
			Data: &OpenGraph_Primary_Audio{
				Audio: ogAudios[0],
			},
		}
	}

	// Get Twitter
	if info.Twitter != nil {
		// Set Twitter
		twitter = &OpenGraph_Twitter{
			Card:        info.Twitter.Card,
			Site:        info.Twitter.Site,
			SiteId:      info.Twitter.SiteId,
			Creator:     info.Twitter.Creator,
			CreatorId:   info.Twitter.CreatorId,
			Description: info.Twitter.Description,
			Title:       info.Twitter.Title,
			Image:       info.Twitter.Image,
			ImageAlt:    info.Twitter.ImageAlt,
			Url:         info.Twitter.Url,
			Player: &OpenGraph_Twitter_Player{
				Url:    info.Twitter.Player.Url,
				Width:  int32(info.Twitter.Player.Width),
				Height: int32(info.Twitter.Player.Height),
				Stream: info.Twitter.Player.Stream,
			},
			Iphone: &OpenGraph_Twitter_IPhone{
				Name: info.Twitter.IPhone.Name,
				Id:   info.Twitter.IPhone.Id,
				Url:  info.Twitter.IPhone.Url,
			},
			Ipad: &OpenGraph_Twitter_IPad{
				Name: info.Twitter.IPad.Name,
				Id:   info.Twitter.IPad.Id,
				Url:  info.Twitter.IPad.Url,
			},
			GooglePlay: &OpenGraph_Twitter_GooglePlay{
				Name: info.Twitter.Googleplay.Name,
				Id:   info.Twitter.Googleplay.Id,
				Url:  info.Twitter.Googleplay.Url,
			},
		}

		// Set OpenGraph Type
		ogPrimary = &OpenGraph_Primary{
			Type: OpenGraph_TWITTER,
			Data: &OpenGraph_Primary_Twitter{
				Twitter: twitter,
			},
		}
	}

	// Get Videos
	if info.Videos != nil {
		// Get Videos
		for _, v := range info.Videos {
			ogVideos = append(ogVideos, &OpenGraph_Video{
				Url:       v.Url,
				SecureUrl: v.SecureUrl,
				Width:     int32(v.Width),
				Height:    int32(v.Height),
				Type:      v.Type,
			})
		}

		// Set OpenGraph Type
		ogPrimary = &OpenGraph_Primary{
			Type: OpenGraph_VIDEO,
			Data: &OpenGraph_Primary_Video{
				Video: ogVideos[0],
			},
		}
	}

	// Get Images
	if info.Images != nil {
		// Get Images
		for _, v := range info.Images {
			ogImages = append(ogImages, &OpenGraph_Image{
				Url:       v.Url,
				SecureUrl: v.SecureUrl,
				Width:     int32(v.Width),
				Height:    int32(v.Height),
				Type:      v.Type,
			})
		}

		// Set OpenGraph Type
		ogPrimary = &OpenGraph_Primary{
			Type: OpenGraph_IMAGE,
			Data: &OpenGraph_Primary_Image{
				Image: ogImages[0],
			},
		}
	}

	// Set OpenGraph Values
	u.OpenGraph = &OpenGraph{
		Primary: ogPrimary,
		Images:  ogImages,
		Videos:  ogVideos,
		Audios:  ogAudios,
		Twitter: twitter,
	}

	// Check Primary
	if u.OpenGraph.Primary.Type == OpenGraph_NONE {
		return nil, errors.New("No OpenGraph Primary Type")
	} else {
		return u.OpenGraph.Primary, nil
	}
}

// ToTransferItem Returns Transfer for URLLink
func (u *UrlItem) ToTransferItem() *Payload_Item {
	return &Payload_Item{
		Mime: DefaultUrlMIME(),
		Data: &Payload_Item_Url{
			Url: u,
		},
	}
}
