package server

import (
	"bot/db"
	"os"
)

type serverImpl struct {
	appPublicKey string
	db           db.DB
}

func New(db2 db.DB) Server {
	return serverImpl{
		appPublicKey: os.Getenv("DISCORD_APP_PUBLIC_KEY"),
		db:           db2,
	}
}
