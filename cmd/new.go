package cmd

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/ohhfishal/schedule/db"
)

type New struct {
	Name        string    `arg:"" help:"Name of event."`
	StartDate   time.Time `arg:"" format:"2006-01-02" help:"Date event starts."`
	StartTime   time.Time `arg:"" optional:"" format:"15:04" help:"Time event starts."`
	Description string    `short:"d" default:"" help:"Description for event."`
}

func (cmd New) Run(ctx context.Context, stdout io.Writer, queries *db.Queries, now func() time.Time) error {
	date := cmd.StartDate
	var start time.Time
	if cmd.StartTime.IsZero() {
		start = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, now().Location())
	} else {
		t := cmd.StartTime
		start = time.Date(date.Year(), date.Month(), date.Day(), t.Hour(), t.Minute(), 0, 0, now().Location())
	}
	event, err := queries.CreateEvent(ctx, db.CreateEventParams{
		Name:        cmd.Name,
		Description: cmd.Description,
		StartTime:   start.Unix(),
	})
	fmt.Fprintf(stdout, "%+v\n", event)
	return err
}
