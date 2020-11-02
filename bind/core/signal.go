package core

// Update enters the room with given OLC(Open-Location-Code)
func update(updateJSON string) string {
	nodeProfile = updateJSON
	lobbyRef.Publish(updateJSON)
	return lobbyRef.ID
}

// Exit informs lobby and closes host
func exit() string {
	lobbyRef.Publish(nodeProfile)
	hostNode.Close()
	return lobbyRef.ID
}

// Offer informs peer about file
func offer() string {
	lobbyRef.Publish(nodeProfile)
	return lobbyRef.ID
}

// Answer and accept Peers offer
func answer() string {
	lobbyRef.Publish(nodeProfile)
	return lobbyRef.ID
}

// Decline Peers offer
func decline() string {
	lobbyRef.Publish(nodeProfile)
	return lobbyRef.ID
}

// Failed informs Peers transfer was unsuccesful
func failed() string {
	lobbyRef.Publish(nodeProfile)
	return lobbyRef.ID
}
