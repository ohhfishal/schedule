package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ohhfishal/schedule/db"
)

type Edit struct {
	ID          int64      `arg:"" required:"" help:"ID of event to edit"`
	Name        *string    `help:"Name of event."`
	Start       *time.Time `short:"s" format:"2006-01-02 15:04" help:"Date event starts."`
	End         *time.Time `short:"e" format:"2006-01-02 15:04" help:"Date event ends."`
	Description *string    `short:"d" default:"" help:"Description for event."`
}

func (cmd Edit) Run(ctx context.Context, stdout Stdout, queries *db.Queries) error {
	params := db.UpdateEventParams{
		ID: cmd.ID,
	}

	if cmd.Name != nil {
		params.Name = sql.NullString{String: *cmd.Name, Valid: true}
	}
	if cmd.Description != nil {
		params.Description = sql.NullString{String: *cmd.Description, Valid: true}
	}
	if cmd.Start != nil {
		params.StartTime = sql.NullInt64{Int64: cmd.Start.Unix(), Valid: true}
	}
	if cmd.End != nil {
		params.EndTime = sql.NullInt64{Int64: cmd.End.Unix(), Valid: true}
	}

	event, err := queries.UpdateEvent(ctx, params)
	fmt.Fprintln(stdout, event)
	return err
}
