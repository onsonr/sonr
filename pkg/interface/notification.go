package ui

import (
	"log"
	"runtime"

	macNotif "github.com/deckarep/gosx-notifier"
	"github.com/gen2brain/beeep"
	md "github.com/sonr-io/core/internal/models"
)

const kCompatibleFileTypes = "*.png *.jpg *.jpeg *.mp4 *.avi *.pdf *.doc *.docx *.ttf *.mp3 *.xml *.csv *.key *.ppt *.pptx *.xls *.xlsm *.xlsx *.rtf *.txt"

func PushInvited(inv *md.AuthInvite) {
	if runtime.GOOS == "darwin" {
		//At a minimum specifiy a message to display to end-user.
		note := macNotif.NewNotification("Check your Apple Stock!")
		note.Title = "It's money making time 💰"
		note.Subtitle = "My subtitle"
		note.Sound = macNotif.Basso
		note.Group = "io.sonr.mac"
		note.Sender = "io.sonr.mac"
		note.AppIcon = "gopher.png"
		note.ContentImage = "gopher.png"
		err := note.Push()
		if err != nil {
			log.Println("Uh oh!")
		}
	} else {
		err := beeep.Notify("Invited", inv.From.Profile.FirstName+" has sent an invite to share "+inv.Payload.String(), "assets/information.png")
		if err != nil {
			log.Println(err)
		}
	}

}

func BeepCompleted() {
	err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if err != nil {
		log.Println(err)
	}
}
