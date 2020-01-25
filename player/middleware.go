package player

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginPlayerHandler(c *gin.Context) {
	session := sessions.Default(c)
	steamid := session.Get("steamid")

	_, err := LoginPlayer(c, steamid.(string))
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}

func GetPlayerHandler(c *gin.Context) {
	session := sessions.Default(c)
	steamid := session.Get("steamid").(string)

	player, err := GetPlayerBySteamID(c, steamid)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, player)
}
