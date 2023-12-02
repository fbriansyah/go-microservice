package echo

import (
	"net/http"

	"github.com/fbriansyah/go-microservice/internal/application"
	"github.com/rs/zerolog"
)

// Server represents an HTTP server that serves requests to the payment service.
type Server struct {
	// app is the payment service application.
	app application.App

	// address is the address the server listens on.
	address string

	logger zerolog.Logger
}

func NewServer(app application.App, address string, logger zerolog.Logger) *Server {
	return &Server{
		app:     app,
		address: address,
		logger:  logger,
	}
}

func (s *Server) Start() error {
	// srv is the HTTP server.
	srv := &http.Server{
		Addr:    s.address,
		Handler: nil,
	}

	// ListenAndServe starts the server and listens for incoming requests.
	// If an error occurs, it is logged and the program is terminated.
	err := srv.ListenAndServe()
	if err != nil {
		s.logger.Fatal().Msgf("error listen and serve: %s", err.Error())
	}

	return nil
}
