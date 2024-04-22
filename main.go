package main

import (
	"cae_chat/repository"
	"cae_chat/server"
	socketio "github.com/googollee/go-socket.io"
)

func main() {
	_ = repository.NewRepository()

	clients := map[string]socketio.Conn{}

	s := server.InitializeServer(clients)

	server.Run(s)
}
