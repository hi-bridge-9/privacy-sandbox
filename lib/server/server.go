package server

import (
	"log"
	"net/http"
)

type Server struct {
	path map[string]func(http.ResponseWriter, *http.Request)
}

func New(path map[string]func(http.ResponseWriter, *http.Request)) *Server {
	return &Server{
		path: path,
	}
}

func (s *Server) Start(port string) error {
	if len(port) == 0 {
		port = "80"
	}
	for path, handler := range s.path {
		log.Printf("go server start: %s\n", path)
		http.HandleFunc(path, handler)
	}

	return http.ListenAndServe(":"+port, nil)
}
