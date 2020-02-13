package app

import (
	"github.com/leobrines/easymm/router"
	"github.com/leobrines/easymm/sql"
)

func Start() {
	sql.Connect()
	router.Init()
	router.Start()
}
