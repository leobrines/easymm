package app

import (
	"log"

	"github.com/leobrines/easymm/router"
	"github.com/leobrines/easymm/sql"
)

func Start() {
	sql.Connect()
	router.Init()

	if err := router.Start(); err != nil {
		log.Fatal("failed to initialize the server %v", err)
	}
}
