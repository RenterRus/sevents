package v1

import (
	"context"
	"net/http"
	"storing_events/internal/db"
)

type HTTPServer struct {
	s     http.Server
	Mongo *db.Mongo
}

func NewServer(addr string, mongo *db.Mongo) *HTTPServer {
	s := new(HTTPServer)
	s.s = http.Server{
		Addr: addr,
	}
	s.Mongo = mongo

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
