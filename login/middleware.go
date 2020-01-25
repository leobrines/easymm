package login

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/leobrines/easymm/apierrors"
	"github.com/leobrines/easymm/steam"
)

func SteamLoginHandler(c *gin.Context) {
	c.PureJSON(http.StatusOK, gin.H{
		"steam_auth_url": steam.AuthURL(c.Request),
	})
}

func SteamCallbackHandler(c *gin.Context) {
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

func IsAuthenticated(c *gin.Context) {
	session := sessions.Default(c)
	steamid := session.Get("steamid")

	if str, _ := steamid.(string); str == "" {
		c.Error(apierrors.NewUnauthorizedApiError("Player unauthorized"))
		return
	}

	c.Next()
}
