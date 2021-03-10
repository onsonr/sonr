package interface

import (
	"log"

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

// ** Const UI Resource Path ** //
const RES_PATH = "/Users/prad/Sonr/core/pkg/res/"
const ICON_PATH = "/Users/prad/Sonr/core/pkg/res/systray.png"

func (d Icon) File() string {
	return [...]string{"systray.png", "close.png", "user.png", "peer.png", "invite.png", "iphone.png", "android.png", "mac.png", "windows.png", "unknown.png", "link.png", "url.png", "file.png"}[d]
}

type AssetImage struct {
	
}

// ^ Returns Buffer of Image by Icon Type
func GetIcon(i Icon) []byte {
	// Get Path
	data, err := box.Find(i.File())
	if err != nil {
		log.Println(err)
	}
	return data
}

// ^ Returns Buffer of Image from Device Type
func GetDeviceIcon(p md.Platform) []byte {
	if p == md.Platform_Android {
		return ai.GetIcon(Android)
	} else if p == md.Platform_iOS {
		return ai.GetIcon(iPhone)
	} else if p == md.Platform_MacOS {
		return ai.GetIcon(Mac)
	} else if p == md.Platform_Windows {
		return ai.GetIcon(Windows)
	}
	return ai.GetIcon(Unknown)
}
