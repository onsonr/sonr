// Code generated from Pkl module `common.types.Ctx`. DO NOT EDIT.
package types

type Session struct {
	Id string `pkl:"id" json:"id,omitempty"`

	Challenge string `pkl:"challenge" json:"challenge,omitempty"`

	BrowserName string `pkl:"browserName" json:"browserName,omitempty"`

	BrowserVersion string `pkl:"browserVersion" json:"browserVersion,omitempty"`

	UserArchitecture string `pkl:"userArchitecture" json:"userArchitecture,omitempty"`

	Platform string `pkl:"platform" json:"platform,omitempty"`

	PlatformVersion string `pkl:"platformVersion" json:"platformVersion,omitempty"`

	DeviceModel string `pkl:"deviceModel" json:"deviceModel,omitempty"`

	IsMobile bool `pkl:"isMobile" json:"isMobile,omitempty"`

	VaultAddress string `pkl:"vaultAddress" json:"vaultAddress,omitempty"`
}
