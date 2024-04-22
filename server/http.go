package server

import (
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

func Run(server *socketio.Server) {
	defer server.Close()

	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/socket.io/", func(w http.ResponseWriter, r *http.Request) {
		server.ServeHTTP(w, r)
	})
	log.Println("Serving at localhost:8000...")
	defer server.Close()
	go func() {
		panic(server.Serve())
	}()

	log.Fatal(http.ListenAndServe(":8000", nil))
}
