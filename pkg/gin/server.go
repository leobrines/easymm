package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leobrines/easymm/pkg"
)

type HTTPServer struct {
	router       *gin.Engine
	steamService pkg.SteamService
}

func NewHttpServer(s pkg.SteamService) *HTTPServer {
	return &HTTPServer{
		router:       gin.Default(),
		steamService: s,
	}
}

func (s *HTTPServer) Start() {
	s.createEndpoints()
	s.router.Run(":8080")
}

func (s *HTTPServer) createEndpoints() {
	s.router.GET("/steam/login", s.steamLogin)
	s.router.Use(errorHandler)
}

func (s *HTTPServer) steamLogin(c *gin.Context) {
	result, err := s.steamService.UserLogin(c.Request)
	if err != nil {
		c.Error(err)
		return
	}

	//http.Redirect(c.Writer, c.Request, result["steam_login_url"].(string), 302)
	fmt.Println(result)
	c.JSON(http.StatusOK, result)
}
