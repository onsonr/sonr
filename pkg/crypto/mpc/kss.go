package mpc

import (
	"fmt"
	"time"

	"github.com/ucan-wg/go-ucan"
)

type KeyshareSource interface {
	ucan.Source
}

type keyshareSource struct {
	userShare Share
	valShare  Share
}

func KeyshareSetFromArray(arr []Share) (KeyshareSource, error) {
	if len(arr) != 2 {
		return nil, fmt.Errorf("invalid keyshare array length")
	}
	return keyshareSource{
		userShare: arr[0],
		valShare:  arr[1],
	}, nil
}

func (k keyshareSource) NewOriginToken(audienceDID string, att ucan.Attenuations, fct []ucan.Fact, notBefore, expires time.Time) (*ucan.Token, error) {
}

func (k keyshareSource) NewAttenuatedToken(parent *ucan.Token, audienceDID string, att ucan.Attenuations, fct []ucan.Fact, notBefore, expires time.Time) (*ucan.Token, error)
