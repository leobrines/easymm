package pkg

import (
	"net/http"
	"time"
)

type HTTPServer interface {
	Start()
}

type SteamService interface {
	UserLogin(r *http.Request) (map[string]interface{}, error)
}

type PlayerService interface {
	Create(steamid string) (*Player, error)
}

type Player struct {
	ID        string
	SteamID   string
	CreatedAt time.Time
}
