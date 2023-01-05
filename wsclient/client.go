package wsclient

import (
	"context"
	"log"
	"net"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

type Client struct {
	conn    net.Conn
	message chan []byte
	command chan []byte
}

func StartClient() *Client {
	conn, _, _, err := ws.DefaultDialer.Dial(context.Background(), "ws://localhost:8123/api/websocket")
	if err != nil {
		log.Fatal(err)
	}
	client := &Client{
		conn:    conn,
		message: make(chan []byte),
		command: make(chan []byte),
	}

	go client.readMessages()
	go client.sendCommands()

	return client
}

func (c *Client) OnMessage(action func([]byte)) {
	for {
		message := <-c.message
		go action(message)
	}
}

func (c *Client) SendCommand(command []byte) {
	c.command <- command
}

func (c *Client) readMessages() {
	defer c.conn.Close()

	for {
		msg, _, err := wsutil.ReadServerData(c.conn)
		if err != nil {
			log.Printf("error: %v", err)
			break
		}
		c.message <- msg
	}
}

func (c *Client) sendCommands() {
	for {
		command := <-c.command
		wsutil.WriteClientText(c.conn, command)
	}
}
