package common

import "strings"

type AccountId string

func AccountIdFromString(id string) (AccountId, error) {
	return AccountId(id), nil
}

// Checks if the AccountId is a AccountAddress on the Cosmos Blockchain
func (id AccountId) IsAddress() bool {
	return strings.HasPrefix(id.String(), "snr") && len(id.String()) == 44
}

func (id AccountId) IsAlias() bool {
	if strings.Contains(id.String(), ".snr") {
		spts := strings.Split(id.String(), ".")
		if len(spts) == 2 && spts[1] == "snr" {
			return true
		}
	}
	return false
}

func (id AccountId) IsDid() bool {
	if strings.Contains(string(id), "did:snr:") {
		spts := strings.Split(string(id), ".")
		if len(spts) == 2 && spts[1] == "snr" {
			return true
		}
	}
	return false
}

func (id AccountId) IsPeerId() bool {
	if strings.Contains(string(id), "Qm") {
		return true
	}
	return false
}

func (id AccountId) String() string {
	return string(id)
}
