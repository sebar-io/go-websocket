package ws

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Receive chan []byte
	Topic   *Topic
	Socket  *websocket.Conn
}

func (c *Client) Write() {
	defer c.Socket.Close()
	for msg := range c.Receive {
		err := c.Socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}

func (c *Client) Read() {
	defer c.Socket.Close()
	for {
		_, msg, err := c.Socket.ReadMessage()
		if err != nil {
			return
		}
		c.Topic.Forward <- msg
	}
}
