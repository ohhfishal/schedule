package cmd

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/ohhfishal/schedule/db"
	"github.com/ohhfishal/schedule/lib/recurrence"
)

type New struct {
	Name        string    `arg:"" help:"Name of event."`
	StartDate   time.Time `arg:"" format:"2006-01-02" help:"Date event starts."`
	StartTime   time.Time `arg:"" optional:"" format:"15:04" help:"Time event starts."`
	Description string    `short:"d" default:"" help:"Description for event."`
	Recurrence  string    `short:"r" default:"" help:"RRULE for event."`
}

func (cmd New) Run(ctx context.Context, stdout io.Writer, queries *db.Queries, location *time.Location) error {
	date := cmd.StartDate
	var start time.Time
	if cmd.StartTime.IsZero() {
		start = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, location)
	} else {
		t := cmd.StartTime
		start = time.Date(date.Year(), date.Month(), date.Day(), t.Hour(), t.Minute(), 0, 0, location)
	}
	var err error
	var r *recurrence.Rule
	if cmd.Recurrence != `` {
		r, err = recurrence.ParseRRule(cmd.Recurrence)
		if err != nil {
			return fmt.Errorf(`rrule invalid: %w`, err)
		}
	}

	event, err := queries.CreateEvent(ctx, db.CreateEventParams{
		Name:        cmd.Name,
		Description: cmd.Description,
		StartTime:   start.Unix(),
		Recurrence:  r,
	})
	fmt.Fprintf(stdout, "%+v\n", event)
	return err
}
