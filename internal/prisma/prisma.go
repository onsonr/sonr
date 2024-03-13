package prisma

import (
	"github.com/steebchen/prisma-client-go/engine"

	"github.com/sonrhq/sonr/internal/prisma/hway"
	"github.com/sonrhq/sonr/internal/prisma/indexer"
	"github.com/sonrhq/sonr/internal/prisma/matrix"
)

type Client struct {
	Hway    engine.Engine
	Indexer engine.Engine
	Matrix  engine.Engine
}

func NewClient() *Client {
	return &Client{
		Hway:    hway.NewClient(),
		Indexer: indexer.NewClient(),
		Matrix:  matrix.NewClient(),
	}
}
