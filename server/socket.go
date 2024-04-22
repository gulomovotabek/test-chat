package server

import (
	"cae_chat/repository"
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"github.com/gorilla/websocket"
)

type socket struct {
	server  *socketio.Server
	repo    repository.Repository
	clients map[string]socketio.Conn
}

func newSocket(server *socketio.Server, repo repository.Repository, clients map[string]socketio.Conn) *socket {
	return &socket{
		server:  server,
		repo:    repo,
		clients: clients,
	}
}

func InitializeServer(clients map[string]socketio.Conn) *socketio.Server {
	server := socketio.NewServer(nil)

	s := newSocket(server, repository.NewRepository(), clients)

	server.OnConnect("/", s.onConnect)
	server.OnEvent("/", "notice", s.onNotice)
	server.OnError("/", s.onError)
	server.OnDisconnect("/", s.onDisconnect)

	server.OnEvent("/chat", "msg", s.onChat)

	return server
}

func (s *socket) onConnect(c socketio.Conn) error {
	// should check token
	s.clients[c.ID()] = c
	return nil
}

func (s *socket) onDisconnect(c socketio.Conn, reason string) {
	fmt.Println("closing: ", c.ID())
	c.Close()
}

func (s *socket) onNotice(c socketio.Conn, msg string) {
	c.Emit("reply", "bro, it's working, chill")
}

func (s *socket) onChat(c socketio.Conn, msg string) {
	c.Emit("reply", "bro, it's working, chill")
}

func (s *socket) onError(c socketio.Conn, e error) {
	if websocket.IsCloseError(e, 1001) {
		return
	}
	fmt.Println("meet error:", e)
}
