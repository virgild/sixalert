package websocket

import (
	"fmt"
	"log"

	"goji.io/pat"

	"goji.io"

	"golang.org/x/net/websocket"
)

type Server struct {
	clients   map[int]*Client
	addCh     chan *Client
	delCh     chan *Client
	sendAllCh chan *Message
	doneCh    chan bool
	errCh     chan error
}

func NewServer() *Server {
	clients := make(map[int]*Client)
	addCh := make(chan *Client)
	delCh := make(chan *Client)
	sendAllCh := make(chan *Message)
	doneCh := make(chan bool)
	errCh := make(chan error)
	return &Server{
		clients,
		addCh,
		delCh,
		sendAllCh,
		doneCh,
		errCh,
	}
}

func (s *Server) Add(c *Client) {
	s.addCh <- c
}

func (s *Server) Del(c *Client) {
	s.delCh <- c
}

func (s *Server) SendAll(msg *Message) {
	s.sendAllCh <- msg
}

func (s *Server) Done() {
	s.doneCh <- true
}

func (s *Server) Err(err error) {
	s.errCh <- err
}

func (s *Server) sendAll(msg *Message) {
	for _, c := range s.clients {
		c.Write(msg)
	}
}

func (s *Server) Listen(mux *goji.Mux) {
	log.Printf("Socket server listening...\n")

	onConnected := func(ws *websocket.Conn) {
		defer func() {
			err := ws.Close()
			if err != nil {
				s.errCh <- err
			}
		}()

		client := NewClient(ws, s)
		s.Add(client)
		client.Listen()
	}

	mux.Handle(pat.Get("/echo"), websocket.Handler(onConnected))

	for {
		select {
		case c := <-s.addCh:
			s.clients[c.id] = c
			log.Printf("Client %d connected", c.id)
		case c := <-s.delCh:
			delete(s.clients, c.id)
			log.Printf("Client %d disconnected", c.id)
		case msg := <-s.sendAllCh:
			fmt.Printf("Got message: %#v\n", msg)
			s.sendAll(msg)
		case err := <-s.errCh:
			log.Printf("Error: %s\n", err.Error())
		case <-s.doneCh:
			return
		}
	}
}
