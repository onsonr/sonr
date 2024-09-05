package types

// ByteArray is a list of byte arrays
type ByteArray = [][]byte

// ToVerificationMethod converts a Profile to a VerificationMethod
func (p Profile) ToVerificationMethod() VerificationMethod {
	return VerificationMethod{
		Id:         p.Id,
		Controller: p.Controller,
	}
}
