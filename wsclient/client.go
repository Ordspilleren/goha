package wsclient

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

type Client struct {
	url       string
	conn      net.Conn
	command   chan []byte
	onMessage func([]byte)
}

func New(url string, onMessage func([]byte)) *Client {
	client := &Client{
		url:       url,
		command:   make(chan []byte),
		onMessage: onMessage,
	}

	return client
}

func (c *Client) Start() error {
	c.reconnect()

	go c.readMessages()
	go c.sendCommands()

	return nil
}

func (c *Client) reconnect() error {
	if c.conn != nil {
		c.conn.Close()
		time.Sleep(time.Second * 2)
	}

	conn, _, _, err := ws.DefaultDialer.Dial(context.Background(), c.url)
	if err != nil {
		log.Fatal(err)
	}
	c.conn = conn

	return nil
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
			c.reconnect()
		}
		go c.onMessage(msg)
	}
}

func (c *Client) sendCommands() {
	for {
		command := <-c.command
		wsutil.WriteClientText(c.conn, command)
	}
}
