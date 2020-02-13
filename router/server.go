package router

import (
	"log"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/leobrines/easymm/login"
	"github.com/leobrines/easymm/player"
)

var wsserver *socketio.Server
var router *gin.Engine

func Init() {
	store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))

	gin.SetMode(gin.DebugMode)
	router = gin.Default()
	router.Use(gin.Logger())

	router.Use(sessions.Sessions("player", store))

	server := socket.WebsocketServer()
	defer server.Close()

	router.GET("/steam/login", login.SteamLoginHandler)
	router.GET("/steam/login/callback", login.SteamCallbackHandler, player.LoginPlayerHandler)
	router.GET("/player/:id", login.IsAuthenticated, player.GetPlayerHandler)
	router.GET("/socket.io/", gin.WrapH(server))
}

func Start() {
	defer wsserver.Close()
	err := router.Run(":8080")
	log.Fatal(err)
}
