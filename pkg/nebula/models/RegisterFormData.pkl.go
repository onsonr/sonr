// Code generated from Pkl module `models`. DO NOT EDIT.
package models

type RegisterFormData interface {
	Form
}

var _ RegisterFormData = (*RegisterFormDataImpl)(nil)

type RegisterFormDataImpl struct {
	Title string `pkl:"title"`

	Description string `pkl:"description"`

	Inputs []*Input `pkl:"inputs"`
}

func (rcv *RegisterFormDataImpl) GetTitle() string {
	return rcv.Title
}

func (rcv *RegisterFormDataImpl) GetDescription() string {
	return rcv.Description
}

func (rcv *RegisterFormDataImpl) GetInputs() []*Input {
	return rcv.Inputs
}
