// Code generated from Pkl module `models`. DO NOT EDIT.
package models

import "github.com/onsonr/sonr/pkg/nebula/models/formstate"

type RegistrationForm struct {
	Title string `pkl:"title"`

	Description string `pkl:"description"`

	State formstate.FormState `pkl:"state"`

	Inputs []*Input `pkl:"inputs"`
}
