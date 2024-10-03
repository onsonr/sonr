// Code generated from Pkl module `models`. DO NOT EDIT.
package models

type Hero struct {
	TitleFirst string `pkl:"titleFirst"`

	TitleEmphasis string `pkl:"titleEmphasis"`

	TitleSecond string `pkl:"titleSecond"`

	Subtitle string `pkl:"subtitle"`

	PrimaryButton *Button `pkl:"primaryButton"`

	SecondaryButton *Button `pkl:"secondaryButton"`

	Image *Image `pkl:"image"`

	Stats []*Stat `pkl:"stats"`
}
