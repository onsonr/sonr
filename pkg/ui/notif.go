package ui

import "github.com/gen2brain/beeep"

func PushInvited() {
	err := beeep.Notify("Title", "Message body", "assets/information.png")
	if err != nil {
		panic(err)
	}
}
