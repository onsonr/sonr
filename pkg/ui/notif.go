package ui

import (
	"log"

	"github.com/gen2brain/beeep"
)

func TestNotif() {
	err := beeep.Notify("Title", "Message body", "assets/information.png")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Sent Test Notification")
}

func TestAlert() {
	err := beeep.Alert("Title", "Message body", "assets/warning.png")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Sent Test Alert")
}

func TestBeep() {
	err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Sent Test Beep")
}
