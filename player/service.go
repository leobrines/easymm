package player

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/leobrines/easymm/core"
	"github.com/leobrines/easymm/sql/query"
)

var _ core.PlayerService = &Service{}

type Service struct {
	db      *sql.DB
	queries *query.Queries
}

func NewService(db *sql.DB, queries *query.Queries) *Service {
	return &Service{
		db:      db,
		queries: queries,
	}
}

func (s *Service) Create(ctx context.Context, steamid string) (*core.Player, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	queriestx := s.queries.WithTx(tx)

	userdb, err := queriestx.CreateUser(ctx)
	if err != nil {
		return nil, err
	}
	playerdb, err := queriestx.CreatePlayer(ctx, steamid)
	if err != nil {
		return nil, err
	}

	tx.Commit()

	player := &core.Player{
		ID:        strconv.Itoa(int(userdb.ID)),
		CreatedAt: userdb.CreatedAt,
		SteamID:   playerdb.Steamid,
	}

	return player, nil
}

func (s *Service) GetBySteamID(ctx context.Context, steamid string) (*core.Player, error) {
	playerdb, err := s.queries.GetPlayerBySteamID(ctx, steamid)
	if err != nil {
		return nil, err
	}

	userdb, err := s.queries.GetUser(ctx, playerdb.UserID)
	if err != nil {
		return nil, err
	}

	return &core.Player{
		ID:        strconv.Itoa(int(userdb.ID)),
		SteamID:   playerdb.Steamid,
		CreatedAt: userdb.CreatedAt,
	}, nil
}
