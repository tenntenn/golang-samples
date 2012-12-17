package chat

import (
	"code.google.com/p/go.net/websocket"
	"log"
)

const BUF_SIZE = 1000

type Client struct {
	ws *websocket.Conn
	server *Server
	ch chan *Message
	done chan bool
}

func NewClient(ws *websocket.Conn, server *Server) *Client {

	if ws == nil {
		panic("ws cannot be nil")
	} else if server == nil {
		panic("server cannot be nil")
	}

	ch := make(chan *Message)
	done := make(chan bool)

	return &Client{ws, server, ch, done}
}

func (self *Client) Conn() *websocket.Conn {
	return self.ws
}

func (self *Client) Write() chan<-*Message {
	return (chan<-*Message)(self.ch)
}

func (self *Client) Done() chan<-bool {
	return (chan<-bool)(self.done)
}

func (self *Client) ListenWrite() {
	log.Println("Listening write to client")
	for {
		select {
		case msg := <-self.ch:
			log.Println("Send:", msg)
			websocket.JSON.Send(self.ws, msg)
		case <-self.done:
			self.server.RemoveClient() <- self
			return
		}
	}
}

func (self *Client) ListenRead() {
	log.Println("Listening read from client")
	for {
		select {
		case <-self.done:
			self.server.RemoveClient() <- self
			return
		default:
			var msg Message
			err := websocket.JSON.Receive(self.ws, &msg)
			if err != nil {
				self.done<-true
			} else {
				self.server.SendAll() <- &msg
			}
		}
	}
}
