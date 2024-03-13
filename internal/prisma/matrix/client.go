package matrix

import "github.com/sonrhq/sonr/internal/prisma/matrix/db"

func NewClient() *db.PrismaClient {
	return db.NewClient()
}
