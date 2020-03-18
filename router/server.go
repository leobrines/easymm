package router

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/leobrines/easymm/login"
	"github.com/leobrines/easymm/player"
	"github.com/leobrines/easymm/socket"
)

var router *gin.Engine

func init() {
	gin.SetMode(gin.DebugMode)

	router = gin.Default()

	store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))
	loggerHandler := gin.Logger()
	sessionHandler := sessions.Sessions("player", store)

	router.Use(loggerHandler)
	router.Use(sessionHandler)

	router.Use(cors.Default())

	router.Use(handleOrigin)

	router.GET("/steam/login", login.SteamLoginHandler)
	router.GET("/steam/login/callback", login.SteamCallbackHandler, player.LoginPlayerHandler)
	router.GET("/player/:id", login.IsAuthenticated, player.GetPlayerHandler)
	router.GET("/socket.io/", login.IsAuthenticated, socket.Handler)
}

func Start() {
	err := router.Run(":8000")
	log.Fatal(err)
}

func handleOrigin(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Request.Header.Del("Origin")

	c.Next()
}
