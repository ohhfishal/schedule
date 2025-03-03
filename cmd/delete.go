package cmd

import (
	"context"
	"time"

	"github.com/ohhfishal/schedule/db"
)

type Delete struct {
	ID int64 `arg:"" required:"" help:"ID of event to delete."`
}

func (cmd Delete) Run(ctx context.Context, queries *db.Queries, now func() time.Time) error {
	return queries.DeleteEvent(ctx, cmd.ID)
}
