package producer

import (
	"github.com/labstack/echo/v4"
	"github.com/onsonr/sonr/crypto/ucan"
	"github.com/onsonr/sonr/crypto/ucan/store"
	"github.com/onsonr/sonr/pkg/common/ipfs"
)

type ProducerContext struct {
	echo.Context
	// TokenParser is the attentuations assigned to the producer service
	TokenParser *ucan.TokenParser

	// TokenStore is the token store used to store and retrieve tokens
	TokenStore store.IPFSTokenStore

	// IPFSClient is the IPFS client used to resolve the UCAN
	IPFSClient ipfs.Client
}
