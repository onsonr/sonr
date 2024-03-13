package hway

import "github.com/sonrhq/sonr/internal/prisma/hway/db"

func NewClient() *db.PrismaClient {
	return db.NewClient()
}
