package ui

import (
	"log"

	"github.com/gen2brain/beeep"
)

func PushInvited() {
	err := beeep.Notify("Title", "Message body", "assets/information.png")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Sent Test Notification")
}
