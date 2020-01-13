package core

import (
	"context"
	"time"
)

type App struct {
	PlayerService PlayerService
}

type HTTPServer interface {
	Start()
}

type PlayerService interface {
	Create(ctx context.Context, steamid string) (*Player, error)
	GetBySteamID(ctx context.Context, steamid string) (*Player, error)
}

type Player struct {
	ID        string
	SteamID   string
	CreatedAt time.Time
}
