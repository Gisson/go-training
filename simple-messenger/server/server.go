package server

import (
	"github.com/Gisson/simple-messenger/message"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Server struct {
	http    *http.Server
	router  *httprouter.Router
	manager *message.MessageManager
}

type ServerHandler func(*Server) httprouter.Handle

func New(m *message.MessageManager, router *httprouter.Router) (*Server, error) {
	server := &Server{}
	server.router = router
	server.http = &http.Server{}
	server.http.Addr = "0.0.0.0:8080" // FIXME
	server.manager = m

	server.http.Handler = router
	return server, nil
}

func (server *Server) Start() error {
	err := server.http.ListenAndServe()
	return err
}

func (server *Server) Manager() *message.MessageManager {
	return server.manager
}

func (s *Server) AddServerHandler(method, path string, handler ServerHandler) {
	s.router.Handle(method, path, handler(s))
}
