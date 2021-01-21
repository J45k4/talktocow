package main

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/gorilla/websocket"
)

func main() {
	m := martini.Classic()

	m.Get("/socket", func(w http.ResponseWriter, r *http.Request) {
		println("Hello new client")

		ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)

		if _, ok := err.(websocket.HandshakeError); ok {
			http.Error(w, "Not websocket handshake", 400)
		} else if err != nil {
			return
		}

		ws.WriteMessage(1, []byte("Hello"))
	})

	m.Get("/", func() string {
		return "Hello world"
	})

	m.RunOnAddr(":12001")
}
