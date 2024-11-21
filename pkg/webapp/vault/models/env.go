package models

import "github.com/onsonr/sonr/pkg/core/dwn"

type Environment struct {
	KeyShare string
	Address  string
	ChainID  string
	Schema   *dwn.Schema
}
