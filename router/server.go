package router

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/leobrines/easymm/login"
	"github.com/leobrines/easymm/player"
)

var Router *gin.Engine

func Init() {
	store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))

	Router = gin.Default()
	Router.Use(sessions.Sessions("player", store))

	Router.GET("/steam/login", login.SteamLoginHandler)
	Router.GET("/steam/login/callback", login.SteamCallbackHandler, player.LoginPlayerHandler)
	Router.GET("/player/:id", login.IsAuthenticated, player.GetPlayerHandler)
	Router.Use(errorHandler)
}

func Start() error {
	return Router.Run(":8080")
}
