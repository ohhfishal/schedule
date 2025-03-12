package cmd

import (
	"context"
	"errors"
	"time"

	"github.com/ohhfishal/schedule/db"
)

type Edit struct {
	ID          int64     `arg:"" required:"" help:"ID of event to edit"`
	Name        string    `help:"Name of event."`
	Start       time.Time `short:"s" format:"2006-01-02 15:04" help:"Date event starts."`
	Description string    `short:"d" default:"" help:"Description for event."`
}

func (cmd Edit) Run(ctx context.Context, stdout Stdout, queries *db.Queries, now func() time.Time) error {
	// TODO: Implement. Requires editing the SQL
	// event, err := queries.CreateEvent(ctx, db.CreateEventParams{
	// 	Name:        cmd.Name,
	// 	Description: cmd.Description,
	// 	StartTime:   start.Unix(),
	// })
	// fmt.Fprintf(stdout, "%+v\n", event)
	return errors.New("edit is not implemented")
}
