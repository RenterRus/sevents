package v1

import (
	"context"
	"net/http"
)

type HTTPServer struct {
	s http.Server
}

func NewServer(addr string) *HTTPServer {
	s := new(HTTPServer)
	s.s = http.Server{
		Addr: addr,
	}

	return s
}

func (s *HTTPServer) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/start", s.start)
	mux.HandleFunc("/v1/finish", s.finish)

	s.s.Handler = mux
	return s.s.ListenAndServe()
}

func (s *HTTPServer) GraceShutdown() {
	s.s.Shutdown(context.Background())
}
