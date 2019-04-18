package websocketServer

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	readBufferSize  = 1024
	writeBufferSize = 1024
)

type WebsocketServer struct {
	Addr        string
	clientSlice []*websocket.Conn
}

func (w *WebsocketServer) Start() {
	http.Handle("/ws", http.HandlerFunc(w.wsHandler))
	go http.ListenAndServe(w.Addr, nil)
}

func (ws *WebsocketServer) wsHandler(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Upgrade(w, r, w.Header(), readBufferSize, writeBufferSize)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}
	ws.clientSlice = append(ws.clientSlice, c)
}

func (ws *WebsocketServer) Send(v interface{}) {
	for _, c := range ws.clientSlice {
		if err := c.WriteJSON(v); err != nil {
			log.Println(err)
		}
	}
}
