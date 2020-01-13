package gin

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/leobrines/easymm/apierrors"
	"github.com/leobrines/easymm/steam"
)

func SteamLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"steam_auth_url": steam.AuthURL(c.Request),
	})
}

func SteamCallback(c *gin.Context) {
	steamID, err := steam.ValidateAuth(c.Request)
	if err != nil {
		c.Error(err)
		return
	}

	session := sessions.Default(c)
	session.Set("steamid", steamID)
	session.Save()

	c.Next()
}

func (s *Server) IsAuthenticated(c *gin.Context) {
	session := sessions.Default(c)
	steamid := session.Get("steamid")

	if str, _ := steamid.(string); str == "" {
		c.Error(apierrors.NewUnauthorizedApiError("Player unauthorized"))
		return
	}

	c.Next()
}
