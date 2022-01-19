package server

import (
	"net/http"
	"payment/internal/config"
	"time"

	"github.com/go-chi/chi/v5"
)

func NewHTTPServer(c config.Server) *http.Server {
	r := chi.NewRouter()
	r.Get("/", indexHandler)

	httpServer := &http.Server{
		Addr:         c.HTTP.Addr,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return httpServer
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello World!"))
}
