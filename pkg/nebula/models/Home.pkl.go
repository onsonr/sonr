// Code generated from Pkl module `models`. DO NOT EDIT.
package models

type Home struct {
	Hero *Hero `pkl:"hero"`

	Highlights []*Highlight `pkl:"highlights"`

	Features []*Features `pkl:"features"`

	Bento *Bento `pkl:"bento"`

	Lowlights []*Lowlights `pkl:"lowlights"`

	CallToAction *CallToAction `pkl:"callToAction"`

	Footer *Footer `pkl:"footer"`
}
