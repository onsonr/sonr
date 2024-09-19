// Code generated from Pkl module `browser`. DO NOT EDIT.
package browser

type PublicKeyCredentialCreationOptions struct {
	Rp *RpEntity `pkl:"rp"`

	User *UserEntity `pkl:"user"`

	Challenge string `pkl:"challenge"`

	PubKeyCredParams []*PublicKeyCredentialParameters `pkl:"pubKeyCredParams"`

	Timeout int `pkl:"timeout"`

	ExcludeCredentials []*PublicKeyCredentialDescriptor `pkl:"excludeCredentials"`

	AuthenticatorSelection *AuthenticatorSelectionCriteria `pkl:"authenticatorSelection"`

	Attestation string `pkl:"attestation"`

	Extensions []*PublicKeyCredentialParameters `pkl:"extensions"`
}
