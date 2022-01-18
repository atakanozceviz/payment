package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewGRPCServer,
	NewHTTPServer,
	wire.Bind(new(http.Handler), new(*chi.Mux)),
)
