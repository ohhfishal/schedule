package cmd

import (
	"context"
	"log/slog"

	"github.com/ohhfishal/schedule/api"
	"github.com/ohhfishal/schedule/db"
)

type Server struct {
	Addr string `default:"localhost:8080"`
}

func (cmd *Server) Run(ctx context.Context, logger *slog.Logger, queries *db.Queries) error {
	server := api.NewServer(logger, queries)
	if err := server.Run(ctx, cmd.Addr); err != nil {
		return err
	}
	return nil
}
