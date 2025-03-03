package cmd

import (
	"context"
	"fmt"
	"github.com/ohhfishal/schedule/db"
	"time"
)

type New struct {
	Name        string    `arg:"" help:"Name of event."`
	Description string    `default:"" help:"Description for event."`
	StartTime   time.Time `arg:"" format:"2006-01-02 15:04" help:"Time event starts."`
}

func Normalize(s, now time.Time) time.Time {
	return time.Date(s.Year(), s.Month(), s.Day(), s.Hour(), s.Minute(), int(0), int(0), now.Location())

}

func (cmd New) Run(ctx context.Context, queries *db.Queries, now func() time.Time) error {

	event, err := queries.CreateEvent(ctx, db.CreateEventParams{
		Name:        cmd.Name,
		Description: cmd.Description,
		StartTime:   Normalize(cmd.StartTime, now()).Unix(),
	})
	fmt.Println(event)
	return err
}
