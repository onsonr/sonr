// Code generated from Pkl module `common.types.Ctx`. DO NOT EDIT.
package types

// User agent information
type UserAgent struct {
	Architecture string `pkl:"architecture"`

	Bitness string `pkl:"bitness"`

	Browser *BrowserInfo `pkl:"browser"`

	Model string `pkl:"model"`

	PlatformName string `pkl:"platformName"`

	PlatformVersion string `pkl:"platformVersion"`

	IsMobile bool `pkl:"isMobile"`
}
