package http_server

import (
	"github.com/gorilla/mux"
	"net/http"
	service "patreon-statistics/internal/controller/patreon-user"
	"time"
)

type HttpServer interface {
	Listen() error
}

type httpServer struct {
	srv *http.Server
}

func (s *httpServer) Listen() error {
	return s.srv.ListenAndServe()
}

func NewHttpServer(patreonUserController service.PatreonUserController) HttpServer {
	r := mux.NewRouter()

	r.HandleFunc(`/patreon-user/{userId}`, patreonUserController.GetOne)

	server := &httpServer{
		srv: &http.Server{
			Handler:      r,
			Addr:         "127.0.0.1:8000",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		},
	}

	return server
}
