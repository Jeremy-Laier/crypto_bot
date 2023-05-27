package main

import (
	"bot/db"
	"bot/server"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func main() {
	db := db.New()
	server := server.New(db)

	e := echo.New()
	interactionsGroup := e.Group("/interactions")
	interactionsGroup.POST("", server.HandlePing)

	log.Info().Msg("listening to 8080...")
	e.Logger.Fatal(e.Start(":8080"))
}
