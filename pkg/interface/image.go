package ui

import (
	"log"

	"github.com/gobuffalo/packr"
	md "github.com/sonr-io/core/internal/models"
)

// ** Icon Type for Image
type Icon int

const (
	SystemTray Icon = iota
	Close
	User
	Peer
	Invite
	iPhone
	Android
	Mac
	Windows
	Unknown
	Link
	URL
	File
)

func (d Icon) File() string {
	return [...]string{"systray.png", "close.png", "user.png", "peer.png", "invite.png", "iphone.png", "android.png", "mac.png", "windows.png", "unknown.png", "link.png", "url.png", "file.png"}[d]
}

// @ Struct for All Image Assets //
type ImageAssets struct {
	Device packr.Box
	Icon   packr.Box
	System packr.Box
}

// ^ Returns Buffer of Image by Icon Type
func (ai *ImageAssets) SystemIcon(i Icon) []byte {
	// Get Path
	data, err := ai.System.Find(i.File())
	if err != nil {
		log.Println(err)
	}
	return data
}

// ^ Returns Buffer of Image from Device Type
func (ai *ImageAssets) DeviceIcon(p md.Platform) []byte {
	if p == md.Platform_Android {
		data, err := ai.Device.Find("android.png")
		if err != nil {
			log.Println(err)
		}
		return data
	} else if p == md.Platform_iOS {
		data, err := ai.Device.Find("iphone.png")
		if err != nil {
			log.Println(err)
		}
		return data
	} else if p == md.Platform_MacOS {
		data, err := ai.Device.Find("mac.png")
		if err != nil {
			log.Println(err)
		}
		return data
	} else if p == md.Platform_Windows {
		data, err := ai.Device.Find("windows.png")
		if err != nil {
			log.Println(err)
		}
		return data
	}
	data, err := ai.Device.Find("unknown.png")
	if err != nil {
		log.Println(err)
	}
	return data
}
