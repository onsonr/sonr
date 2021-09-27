package common

import (
	"net/http"
	"net/url"

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
func NewUrlItem(url string, m *Metadata) (*Payload_Item, error) {
	// Create UrlItem
	link := &UrlItem{
		Mime: DefaultUrlMIME(),
		Url:  url,
	}
	err := link.FetchData()
	if err != nil {
		return nil, err
	}

	return &Payload_Item{
		Size:     0,
		Metadata: m,
		Mime:     DefaultUrlMIME(),
		Data: &Payload_Item_Url{
			Url: link,
		},
	}, nil
}

// ** ─── URLLink MANAGEMENT ────────────────────────────────────────────────────────
// FetchData Sets URLLink Data
func (u *UrlItem) FetchData() error {
	// Create Request
	resp, err := http.Get(u.Url)
	if err != nil {
		return err
	}

	// Get Info
	info, err := net.GetPageData(resp)
	if err != nil {
		return err
	}

	// Set Link
	u.Title = info.Title
	u.Type = info.Type
	u.Site = info.Site
	u.SiteName = info.SiteName
	u.Description = info.Description

	// Get Images
	if info.Images != nil {
		for _, v := range info.Images {
			u.Images = append(u.Images, &UrlItem_OpenGraphImage{
				Url:       v.Url,
				SecureUrl: v.SecureUrl,
				Width:     int32(v.Width),
				Height:    int32(v.Height),
				Type:      v.Type,
			})
		}
	}

	// Get Videos
	if info.Videos != nil {
		for _, v := range info.Videos {
			u.Videos = append(u.Videos, &UrlItem_OpenGraphVideo{
				Url:       v.Url,
				SecureUrl: v.SecureUrl,
				Width:     int32(v.Width),
				Height:    int32(v.Height),
				Type:      v.Type,
			})
		}
	}

	// Get Audios
	if info.Audios != nil {
		for _, v := range info.Videos {
			u.Audios = append(u.Audios, &UrlItem_OpenGraphAudio{
				Url:       v.Url,
				SecureUrl: v.SecureUrl,
				Type:      v.Type,
			})
		}
	}

	// Get Twitter
	if info.Twitter != nil {
		u.Twitter = &UrlItem_TwitterCard{
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
			Player: &UrlItem_TwitterCard_Player{
				Url:    info.Twitter.Player.Url,
				Width:  int32(info.Twitter.Player.Width),
				Height: int32(info.Twitter.Player.Height),
				Stream: info.Twitter.Player.Stream,
			},
			Iphone: &UrlItem_TwitterCard_IPhone{
				Name: info.Twitter.IPhone.Name,
				Id:   info.Twitter.IPhone.Id,
				Url:  info.Twitter.IPhone.Url,
			},
			Ipad: &UrlItem_TwitterCard_IPad{
				Name: info.Twitter.IPad.Name,
				Id:   info.Twitter.IPad.Id,
				Url:  info.Twitter.IPad.Url,
			},
			GooglePlay: &UrlItem_TwitterCard_GooglePlay{
				Name: info.Twitter.Googleplay.Name,
				Id:   info.Twitter.Googleplay.Id,
				Url:  info.Twitter.Googleplay.Url,
			},
		}
	}

	// Set Mime
	if u.Mime == nil {
		u.Mime = DefaultUrlMIME()
	}
	return nil
}

// ToTransferItem Returns Transfer for URLLink
func (u *UrlItem) ToTransferItem(m *Metadata) *Payload_Item {
	return &Payload_Item{
		Mime:     DefaultUrlMIME(),
		Metadata: m,
		Data: &Payload_Item_Url{
			Url: u,
		},
	}
}
