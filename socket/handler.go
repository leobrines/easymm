package socket

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

var queues map[string]*GameQueue
var server *socketio.Server

func init() {
	var err error
	server, err = socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
		return
	})
	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})
	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})
	server.OnError("/", func(e error) {
		fmt.Println("meet error:", e)
		return
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
		return
	})

	go server.Serve()
}

func Handler(c *gin.Context) {
	server.ServeHTTP(c.Writer, c.Request)
}

/*
// Solution with Gorilla library

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     originHandler,
}

func gorillaHandle(c *gin.Context, w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to upgrade ws: %+v\n", err)
		return
	}

	playerws, err := getPlayerWebsocket(c, conn)
	if err != nil {
		fmt.Printf("Failed get player websocket")
		return
	}

	for {
		t, b, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("error leyendo mensaje : %s", err)
			break
		}

		msg := string(b)
		pm := &WebsocketMessage{
			WSPlayer: playerws,
			Message:  msg,
		}

		fmt.Printf("Type %d, Msg: %s\n", t, msg)

		HandleMessage(pm)
	}

	fmt.Println("Conexion finalizada!")
}

func HandleMessage(wsmsg *WebsocketMessage) {
	switch wsmsg.Message {
	case "game-1v1-search":
		OnGameSearch(wsmsg)
	}
}

func OnGameSearch(wsmsg *WebsocketMessage) {
	if queues[wsmsg.Message] == nil {
		log.Println("Creando cola para buscar partida...")

		queues[wsmsg.Message] = &GameQueue{
			Mode: struct {
				MaxPlayers int
				MaxRounds  int
			}{
				MaxPlayers: 2,
				MaxRounds:  8,
			},
		}
	}

	gameQueue := queues[wsmsg.Message]

	log.Println("Añadiendo jugador a cola...")

	gameQueue.AddPlayer(wsmsg.WSPlayer)
	players := gameQueue.CheckQueue()

	if players != nil {
		EmitConfirmation(players)
	}
}

func EmitConfirmation(playerws []*WebsocketPlayer) {
	log.Println("Confirmando jugadores...")

	for _, p := range playerws {
		event := []byte("game-confirmation")
		if err := p.Conn.WriteMessage(websocket.TextMessage, event); err != nil {
			panic(err)
		}
	}

	log.Printf("Confirmacion enviada a %d jugadores...\n", len(playerws))

}

func originHandler(r *http.Request) bool {
	if !env.IsProduction() {
		return true
	}

	if r.Header.Get("Origin") != "http://"+r.Host {
		return false
	}

	return true
}

func getPlayerWebsocket(c *gin.Context, conn *websocket.Conn) (*WebsocketPlayer, error) {
	session := sessions.Default(c)
	steamid := session.Get("steamid").(string)

	p, err := player.GetPlayerBySteamID(c, steamid)
	if err != nil {
		return nil, err
	}

	pw := &WebsocketPlayer{
		//Player: p,
		Conn: conn,
	}
	return pw, nil
}
*/
