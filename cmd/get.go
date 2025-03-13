package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/ohhfishal/schedule/db"
)

const DAY = 24 * time.Hour

var TIME_FORMAT = "15:04"

type Get struct {
	// TODO: Make this its own command
	All  bool      `short:"a" help:"Print all events (Enables -o raw)"`
	Date time.Time `arg:"" optional:"" format:"2006-01-02" help:"Date to get (Default=today)"`
	Out  string    `short:"o" enum:"markdown,raw" default:"markdown" help:"Output format (markdown,raw)"`
}

func Midnight(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

func Normalize(s, now time.Time) time.Time {
	return time.Date(s.Year(), s.Month(), s.Day(), s.Hour(), s.Minute(), int(0), int(0), now.Location())
}

func Markdown(e db.Event) string {
	var description string
	// TODO: Location coming from user preference would be nice
	timeStr := Normalize(time.Unix(e.StartTime, 0), time.Now()).Format(TIME_FORMAT)
	if e.EndTime != 0 {
		timeStr += "-" + Normalize(time.Unix(e.EndTime, 0), time.Now()).Format(TIME_FORMAT)
	}
	if e.Description != `` {
		description = ": " + e.Description

	}
	return fmt.Sprintf("- %s - %s%s", timeStr, e.Name, description)
}

func (cmd Get) Run(ctx context.Context, stdout Stdout, queries *db.Queries, now func() time.Time) error {
	if stdout == nil {
		return fmt.Errorf("no stdout")
	}

	var today time.Time
	if !cmd.Date.IsZero() {
		today = Midnight(cmd.Date)
	} else {
		if now == nil {
			return fmt.Errorf("no time")
		}
		today = Midnight(now())
	}

	var events []db.Event
	var err error
	if cmd.All {
		events, err = queries.GetAllEvents(ctx)
		// TODO: Make the logic better
		cmd.Out = `raw`
	} else {
		events, err = queries.GetEvents(ctx, db.GetEventsParams{
			Start: today.Unix(),
			End:   today.Add(DAY).Unix(),
		})
	}

	if err != nil {
		return fmt.Errorf(`getting events: %w`, err)
	}

	if len(events) == 0 {
		fmt.Fprintln(stdout, `No events for today :(`)
		return nil
	}

	switch cmd.Out {
	case `raw`:
		fmt.Fprintln(stdout, today)
		for _, event := range events {
			fmt.Fprintf(stdout, "%+v\n", event)
		}
	case `markdown`:
		fmt.Fprintf(stdout, "# %s\n", today.Format(time.DateOnly))
		for _, event := range events {
			fmt.Fprintln(stdout, Markdown(event))
		}
	default:
		return fmt.Errorf("unknown format: %s", cmd.Out)
	}
	return nil
}
