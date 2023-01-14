package fs

import (
	"strconv"
)

// Storing the share of the wallet.
// Index at -1 is for root share, otherwise it is for a derived share.
func (c *Config) StoreShare(share []byte, partyId string, index int) error {
	if index == -1 {
		err := c.authDir.WriteFile(partyId, share)
		if err != nil {
			return err
		}
	} else {
		newDir, err := c.authDir.CreateFolder(strconv.Itoa(index))
		if err != nil {
			return err
		}
		err = newDir.WriteFile(partyId, share)
		if err != nil {
			return err
		}
	}
	return nil
}
