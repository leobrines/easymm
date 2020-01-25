package player

import (
	"context"
	"strconv"

	"github.com/leobrines/easymm/sql/query"

	"github.com/leobrines/easymm/sql"
)

func LoginPlayer(ctx context.Context, steamid string) (*Player, error) {
	var player *Player
	var err error

	// If exist player
	player, err = GetPlayerBySteamID(ctx, steamid)
	if err == nil {
		return player, nil
	}

	// else create player
	player, err = CreatePlayer(ctx, steamid)
	if err != nil {
		return nil, err
	}

	return player, nil
}

func CreatePlayer(ctx context.Context, steamid string) (*Player, error) {
	tx, err := sql.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	queriestx := sql.Query.WithTx(tx)

	userdb, err := queriestx.CreateUser(ctx)
	if err != nil {
		return nil, err
	}
	createPlayerParams := query.CreatePlayerParams{
		SteamID: steamid,
		UserID:  userdb.ID,
	}
	playerdb, err := queriestx.CreatePlayer(ctx, createPlayerParams)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	player := &Player{
		ID:        strconv.Itoa(int(userdb.ID)),
		CreatedAt: userdb.CreatedAt,
		SteamID:   playerdb.SteamID,
	}

	return player, nil
}

func GetPlayerBySteamID(ctx context.Context, steamid string) (*Player, error) {
	playerdb, err := sql.Query.GetPlayerBySteamID(ctx, steamid)
	if err != nil {
		return nil, err
	}

	userdb, err := sql.Query.GetUser(ctx, playerdb.UserID)
	if err != nil {
		return nil, err
	}

	return &Player{
		ID:        strconv.Itoa(int(userdb.ID)),
		SteamID:   playerdb.SteamID,
		CreatedAt: userdb.CreatedAt,
	}, nil
}
