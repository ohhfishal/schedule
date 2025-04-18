package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"time"

	"github.com/ohhfishal/schedule/db"
)

type Edit struct {
	ID          int64     `arg:"" required:"" help:"ID of event to edit"`
	Name        *string   `help:"Name of event."`
	Start       time.Time `short:"s" optional:"" format:"2006-01-02 15:04" help:"Date event starts."`
	End         time.Time `short:"e" optional:"" format:"2006-01-02 15:04" help:"Date event ends."`
	Description *string   `short:"d" default:"" help:"Description for event."`
}

func addLocation(date time.Time, location *time.Location) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), int(0), int(0), location)
}

func (cmd Edit) Run(ctx context.Context, stdout io.Writer, queries *db.Queries, location *time.Location) error {
	params := db.UpdateEventParams{
		ID: cmd.ID,
	}

	if cmd.Name != nil {
		params.Name = sql.NullString{String: *cmd.Name, Valid: true}
	}
	if cmd.Description != nil {
		params.Description = sql.NullString{String: *cmd.Description, Valid: true}
	}
	if !cmd.Start.IsZero() {
		start := addLocation(cmd.Start, location)
		params.StartTime = sql.NullInt64{Int64: start.Unix(), Valid: true}
	}
	if !cmd.End.IsZero() {
		end := addLocation(cmd.End, location)
		params.EndTime = sql.NullInt64{Int64: end.Unix(), Valid: true}
	}

	event, err := queries.UpdateEvent(ctx, params)
	fmt.Fprintln(stdout, event)
	return err
}
