package player

import (
	"context"
	"time"
)

type Service interface {
	CreatePlayer(ctx context.Context, steamid string) (*Player, error)
	GetPlayerBySteamID(ctx context.Context, steamid string) (*Player, error)
}

type Player struct {
	ID        string    `json:"id"`
	SteamID   string    `json:"steam_id"`
	CreatedAt time.Time `json:"created_at"`
}
