
package cmd

import (
  "context"
  "fmt"
  "io"
  "time"
  "github.com/ohhfishal/schedule/db"
)

const DAY = 24 * time.Hour
var TIME_FORMAT = "15:04"

type Get struct {
  Date time.Time `arg:"" optional:"" format:"2006-01-02" help:"Date to get (Default=today)"`
  Format string `enum:"markdown" default:"markdown" help:"Output format (markdown)"`
}

func Midnight(now time.Time) time.Time {
  return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

func Markdown(e db.Event) string {
  var description string
  timeStr := Normalize(time.Unix(e.StartTime, 0), time.Now()).Format(TIME_FORMAT)
  if e.EndTime.Valid && e.EndTime.Int64 != 0  {
    timeStr += "-" + Normalize(time.Unix(e.EndTime.Int64, 0), time.Now()).Format(TIME_FORMAT)
  }
  if e.Description != `` {
    description = ": " + e.Description

  }
  return fmt.Sprintf("- %s - %s%s", timeStr, e.Name, description)
}


func (cmd Get) Run(ctx context.Context, stdout io.Writer, queries *db.Queries, now func()time.Time) error {
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

  events, err := queries.GetEvents(ctx, db.GetEventsParams{
    Start: today.Unix(),
    End: today.Add(DAY).Unix(),
  })

  if err != nil {
    return fmt.Errorf(`getting events: %w`, err)
  }

  if len(events) == 0 {
    fmt.Fprintln(stdout, `No events for today :(`)
    return nil
  }
  // TODO: Validate format / switch here
  fmt.Fprintf(stdout, "# %s\n", today.Format(time.DateOnly))
  for _, event := range events {
    fmt.Fprintln(stdout, Markdown(event))
  }
  return err
}
