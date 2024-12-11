package ui

// Helper function to shorten address
func shortenAddress(address string) string {
	if len(address) <= 20 {
		return address
	}
	return address[:16] + "..." + address[len(address)-4:]
}
