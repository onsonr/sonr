package did

// VerificationMethod is an object that represents a verification method.
type VerificationMethod struct {
	PublicKeyJwk map[string]interface{} `json:"publicKeyJwk"`
	ID           string                 `json:"id"`
	Type         string                 `json:"type"`
	Controller   string                 `json:"controller"`
}
