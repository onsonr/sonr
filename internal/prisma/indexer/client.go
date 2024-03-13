package indexer

import "github.com/sonrhq/sonr/internal/prisma/indexer/db"

func NewClient() *db.PrismaClient {
	return db.NewClient()
}
