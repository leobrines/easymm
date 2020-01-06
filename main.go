package main

import (
	"os"

	"github.com/leobrines/easymm/pkg/gin"
	"github.com/leobrines/easymm/pkg/steam"
)

func main() {
	// db, queries := sql.Connect()
	// playerService := player.NewService(db, queries)
	steamService := steam.NewService(os.Getenv("STEAM_API_KEY"))
	httpServer := gin.NewHttpServer(steamService)
	httpServer.Start()
}
