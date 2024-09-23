// Code generated from Pkl module `browser`. DO NOT EDIT.
package browser

type PublicKeyCredentialRequestOptions struct {
	Challenge string `pkl:"challenge"`

	Timeout int `pkl:"timeout"`

	RpId string `pkl:"rpId"`

	AllowCredentials []*PublicKeyCredentialDescriptor `pkl:"allowCredentials"`

	UserVerification string `pkl:"userVerification"`

	Extensions []*PublicKeyCredentialParameters `pkl:"extensions"`
}
