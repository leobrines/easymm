package main

import (
	"github.com/leobrines/easymm/core"
	"github.com/leobrines/easymm/gin"
	"github.com/leobrines/easymm/player"
	"github.com/leobrines/easymm/sql"
)

func main() {
	db, queries := sql.Connect()
	app := &core.App{
		PlayerService: player.NewService(db, queries),
	}
	httpServer := gin.NewServer(app)
	httpServer.Start()
}
