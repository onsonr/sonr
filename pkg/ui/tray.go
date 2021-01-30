package ui

import (
	"fmt"
	"log"

	"github.com/gen2brain/dlgs"
	"github.com/getlantern/systray"
	sonr "github.com/sonr-io/core/bind"
	md "github.com/sonr-io/core/internal/models"
)

type SystemMenu struct {
	mPeers     *systray.MenuItem
	MLink      *systray.MenuItem
	mCount     *systray.MenuItem
	mQuit      *systray.MenuItem
	mPeersList []*systray.MenuItem
	peerCount  int32
	lobbySize  int32
}

// ^ StartTray Starts System tray with Library ^ //
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
func (sm *SystemMenu) HandleMenuInput() {
	go func() {
		<-sm.mQuit.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()
}

// ^ Method to Rebuild Menu for Lobby Refresh ^ //
func (sm *SystemMenu) UpdatePeers(node *sonr.Node, newLob *md.Lobby) {
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

			// Add Peer to Menu
			log.Println(p)
			peerItem := sm.mPeers.AddSubMenuItem(itemTitle, "Nearby Available Peer")
			peerItem.SetTemplateIcon(GetDeviceIcon(p.Device), GetDeviceIcon(p.Device))

			// Add Peer Send Options
			linkItem := peerItem.AddSubMenuItem("Send Link", "Send a Link to "+itemTitle)
			fileItem := peerItem.AddSubMenuItem("Send File", "Send a File to "+itemTitle)

			// Spawn Routine to handle Peer Item Actions
			go func(fileItem *systray.MenuItem, linkItem *systray.MenuItem, peer *md.Peer) {
				for {
					select {
					case <-fileItem.ClickedCh:
						// Load File
						filename, _, err := dlgs.File("Select File", ".png .jpg .jpeg .mp4 .avi", false)
						if err != nil {
							log.Println(err)
						}

						// Process File
						node.Process(filename)

						// Invite Peer
						node.InviteWithFile(peer.Id)

					case <-linkItem.ClickedCh:
						// Load File
						link, _, err := dlgs.Entry("URL Link", "Enter a URL Here: ", "")
						if err != nil {
							panic(err)
						}

						// Invite Peer
						node.InviteWithURL(peer.Id, link)
					}
				}
			}(fileItem, linkItem, p)

			// Add to Menu List
			sm.mPeersList = append(sm.mPeersList, peerItem)
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
