//nolint
package ui

import (
	"fmt"
	"log"

	"github.com/getlantern/systray"
	sonr "github.com/sonr-io/core/bind"
	md "github.com/sonr-io/core/internal/models"
	"github.com/sqweek/dialog"
)

type SystemMenu struct {
	mPeers     *systray.MenuItem
	MLink      *systray.MenuItem
	mCount     *systray.MenuItem
	mQuit      *systray.MenuItem
	mPeersList []*systray.MenuItem
	node       *sonr.Node
	peerCount  int32
	lobbySize  int32
}

func StartTray() SystemMenu {
	// Set Initial Menu Vars
	systray.SetTemplateIcon(GetIcon(SystemTray), GetIcon(SystemTray))
	systray.SetTitle("")
	systray.SetTooltip("Sonr")

	sm := SystemMenu{}
	sm.peerCount = 0
	sm.lobbySize = 1

	// Quit Sonr
	sm.MLink = systray.AddMenuItem("Link Device", "Link a Device to Sonr")
	sm.MLink.SetTemplateIcon(GetIcon(Link), GetIcon(Link))

	// Quit Sonr
	sm.mQuit = systray.AddMenuItem("Quit", "Quit the whole app")
	sm.mQuit.SetTemplateIcon(GetIcon(Close), GetIcon(Close))
	systray.AddSeparator()

	// Pers Label
	sm.mCount = systray.AddMenuItem("Available Peers", "Peers Near You")
	sm.mCount.Disable()

	// Handle Menu Events
	go sm.HandleMenuInput()
	return sm
}

// ^ Routine Handles Menu Input ^ //
func (sm *SystemMenu) SetNode(n *sonr.Node) {
	sm.node = n
}

// ^ Routine Handles Menu Input ^ //
func (sm *SystemMenu) HandleMenuInput() {
	go func() {
		<-sm.mQuit.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()
}

// ^ Method to Rebuild Menu for Lobby Refresh ^ //
func (sm *SystemMenu) UpdatePeers(newLob *md.Lobby) {
	// Check if Lobby Updated
	if newLob.Size != sm.lobbySize {
		// Change Lobby
		sm.lobbySize = newLob.Size
		sm.peerCount = newLob.Size - 1

		// Reset Menu
		sm.ResetPeers()

		// Add Peers
		for _, p := range newLob.Peers {
			// Build Item
			itemTitle := p.FirstName

			// Add Item to Menu
			log.Println(p)
			item := sm.mPeers.AddSubMenuItem(itemTitle, "Nearby Available Peer")
			item.SetTemplateIcon(GetDeviceIcon(p.Device), GetDeviceIcon(p.Device))
			log.Println(p)

			// Spawn Routine to handle Item
			go func(item *systray.MenuItem, peer *md.Peer) {
				for {
					select {
					case <-item.ClickedCh:
						// Load File
						filename, err := dialog.File().Filter("PNG Image file", "png").Load()

						// Process File
						sm.node.Process(filename)

						// Invite Peer
						sm.node.InviteWithFile(peer.Id)

						// Log Error 
						if err != nil {
							log.Fatalln(err)
						}
					}
				}
			}(item, p)

			// Add to Menu List
			sm.mPeersList = append(sm.mPeersList, item)
		}
	}
}

// ^ Method that resets peers list ^ //
func (sm *SystemMenu) ResetPeers() {
	for _, mi := range sm.mPeersList {
		mi.Hide()
	}
	sm.mPeersList = nil
}
