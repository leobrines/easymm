package gin

import (
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/leobrines/easymm/core"
)

type Server struct {
	router *gin.Engine
	app    *core.App
}

func NewServer(app *core.App) *Server {
	return &Server{
		router: gin.Default(),
		app:    app,
	}
}

func (s *Server) Start() {
	store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))
	s.router.Use(sessions.Sessions("player", store))
	s.createEndpoints()
	s.router.Run(":8080")
}

func (s *Server) createEndpoints() {
	s.router.GET("/steam/login", SteamLogin)
	s.router.GET("/steam/login/callback", SteamCallback)
	s.router.GET("/steam/register/callback", SteamCallback, s.RegisterPlayer)
	s.router.GET("/player/:id", s.IsAuthenticated, s.GetPlayer)
	s.router.Use(errorHandler)
}

func (s *Server) RegisterPlayer(c *gin.Context) {
	session := sessions.Default(c)
	steamid := session.Get("steamid")

	_, err := s.app.PlayerService.Create(c, steamid.(string))
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}

func (s *Server) GetPlayer(c *gin.Context) {
	session := sessions.Default(c)
	steamid := session.Get("steamid").(string)

	player, err := s.app.PlayerService.GetBySteamID(c, steamid)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, player)
}
