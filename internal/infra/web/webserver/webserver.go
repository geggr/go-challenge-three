package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	return &WebServer{
		Router:        r,
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	switch method {
	case http.MethodGet:
		s.Router.Get(path, handler)
	case http.MethodPost:
		s.Router.Post(path, handler)
	case http.MethodPut:
		s.Router.Put(path, handler)
	case http.MethodDelete:
		s.Router.Delete(path, handler)
	default:
		s.Router.Handle(path, handler)
	}
}

func (s *WebServer) Start() {
	http.ListenAndServe(s.WebServerPort, s.Router)
}
