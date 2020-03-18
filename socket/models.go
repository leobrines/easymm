package socket

import (
	"sync"

	"github.com/gorilla/websocket"
	"github.com/leobrines/easymm/player"
)

type WebsocketMessage struct {
	WSPlayer *WebsocketPlayer
	Message  string
}

type WebsocketPlayer struct {
	Player *player.Player
	Conn   *websocket.Conn
}

type GameQueue struct {
	Players []*WebsocketPlayer
	Mode    struct {
		MaxPlayers int
		MaxRounds  int
	}
	mutex sync.Mutex
}

func (s *GameQueue) AddPlayer(p *WebsocketPlayer) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.Players = append(s.Players, p)
}

func (s *GameQueue) CheckQueue() []*WebsocketPlayer {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.Players) > 2 {
		players := s.Players[0:2]
		s.Players = s.Players[2:]
		return players
	}

	return nil
}
