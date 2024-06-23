package http

import (
	"context"
	"net/http"
	"sync"

	"github.com/Roman-Otus-Learning/roman-otus-go-2024-01/hw12_13_14_15_calendar/internal/app"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type Server struct {
	server *http.Server
	events app.EventsUseCaseInterface
}

func CreateHTTPServer(events app.EventsUseCaseInterface, addr string) *Server {
	return &Server{
		server: &http.Server{ //nolint:gosec
			Addr:    addr,
			Handler: loggingMiddleware(simpleHandler()),
		},
		events: events,
	}
}

func (s *Server) Start(wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()

		log.Info().Msg("starting http server on " + s.server.Addr)

		if err := s.server.ListenAndServe(); err != nil {
			log.Error().Err(err).Send()
		}
	}()
}

func (s *Server) Stop(ctx context.Context) {
	log.Info().Msg("stopping http server")

	if err := s.server.Shutdown(ctx); err != nil {
		log.Error().Err(err).Send()
	}

	log.Info().Msg("http server stopped")
}

func simpleHandler() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc(
		"/",
		func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("Hello, World"))
		}).Methods("GET")

	return router
}
