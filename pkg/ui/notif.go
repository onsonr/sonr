package ui

import (
	"log"

	"github.com/gen2brain/beeep"
	md "github.com/sonr-io/core/internal/models"
)

func PushInvited(inv *md.AuthInvite) {
	err := beeep.Notify("Invited", inv.From.FirstName+" has sent an invite to share "+inv.Payload.Type.String(), "assets/information.png")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Pushed Invite Notification")
}
