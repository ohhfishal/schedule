package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ohhfishal/schedule/api"
	"github.com/ohhfishal/schedule/db"
)

type Server struct {
	Addr string `default:"localhost:8080"`
}

func (cmd *Server) Run(ctx context.Context, logger *slog.Logger, queries *db.Queries) error {
	server := api.NewServer(logger, queries)
	router := http.NewServeMux()
	handler := api.HandlerFromMux(server, router)
	s := &http.Server{
		Handler: handler,
		Addr:    cmd.Addr,
	}
	go func() {
		logger.Info(`starting server`, `addr`, cmd.Addr)
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error(`server died due to error`, `error`, err)
		}
	}()

	logger.Info(`main goroutine sleeping`)
	<-ctx.Done()

	err := ctx.Err()
	logger.Info(`context ended`, `error`, err)

	if err := s.Shutdown(context.TODO()); err != nil {
		logger.Error(`shutting down`, `error`, err)
		return fmt.Errorf(`shutting down: %w`, err)
	}
	return nil
}
