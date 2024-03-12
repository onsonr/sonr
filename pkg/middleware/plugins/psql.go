package plugins

import "github.com/sonrhq/sonr/internal/prisma/db"

func New() {
	client := db.NewClient()
	client.Connect()
}
