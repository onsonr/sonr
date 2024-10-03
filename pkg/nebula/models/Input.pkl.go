// Code generated from Pkl module `models`. DO NOT EDIT.
package models

import "github.com/onsonr/sonr/pkg/nebula/models/inputtype"

type Input struct {
	Label string `pkl:"label"`

	Type inputtype.InputType `pkl:"type"`

	Placeholder string `pkl:"placeholder"`

	Value *string `pkl:"value"`

	Error *string `pkl:"error"`

	Help *string `pkl:"help"`

	Required *bool `pkl:"required"`
}
