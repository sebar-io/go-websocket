package ws

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
)

type Topic struct {
	Clients map[*Client]bool
	Join    chan *Client
	Leave   chan *Client
	Forward chan []byte
}

func NewTopic() *Topic {
	t := &Topic{
		Clients: make(map[*Client]bool),
		Join:    make(chan *Client),
		Leave:   make(chan *Client),
		Forward: make(chan []byte),
	}
	go t.Run()
	return t
}

func (t *Topic) Run() {
	for {
		select {
		case client := <-t.Join:
			t.Clients[client] = true
		case client := <-t.Leave:
			delete(t.Clients, client)
			close(client.Receive)
		case msg := <-t.Forward:
			for client := range t.Clients {
				client.Receive <- msg
			}
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (t *Topic) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		slog.Error("ServeHTTP:", err)
		return
	}
	client := &Client{
		Socket:  socket,
		Receive: make(chan []byte, upgrader.ReadBufferSize),
		Topic:   t,
	}
	t.Join <- client
	defer func() { t.Leave <- client }()
	go client.Write()
	client.Read()
}
