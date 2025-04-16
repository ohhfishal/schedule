package get

import (
	"context"
	_ "embed"
	"fmt"
	"io"
	"text/template"
	"time"

	"github.com/ohhfishal/schedule/db"
)

//go:embed markdown.template
var _MARKDOWN_TEMPLATE string

const DAY = 24 * time.Hour

var TIME_FORMAT = "15:04"
var funcMap = template.FuncMap{
	"Markdown": Markdown,
}

type CMD struct {
	// TODO: Make this its own command
	All     All     `cmd:""`
	Default Default `cmd:"" default:"withargs"`
}

type Default struct {
	Date time.Time `arg:"" optional:"" format:"2006-01-02" help:"Date to get (Default=today)"`
}

type TemplateInput struct {
	Date   string
	Events []db.Event
}

func (cmd Default) Run(ctx context.Context, stdout io.Writer, queries *db.Queries) error {
	// TODO: A *time.Location is now in the kong.Context
	if stdout == nil {
		return fmt.Errorf("no stdout")
	}

	var today time.Time
	if !cmd.Date.IsZero() {
		today = Midnight(cmd.Date)
	} else {
		today = Midnight(time.Now())
	}

	events, err := queries.GetEvents(ctx, db.GetEventsParams{
		Start: today.Unix(),
		End:   today.Add(DAY).Unix(),
	})

	if err != nil {
		return fmt.Errorf(`getting events: %w`, err)
	}

	tmpl, err := template.New("markdown-template").
		Funcs(funcMap).
		Parse(_MARKDOWN_TEMPLATE)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	if err = tmpl.Execute(stdout, TemplateInput{
		Date:   today.Format(time.DateOnly),
		Events: events,
	}); err != nil {
		return fmt.Errorf("printing template: %w", err)
	}
	return nil
}

func Midnight(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

func Markdown(event db.Event) string {
	return fmt.Sprintf("- %s - %s%s",
		time.Unix(event.StartTime, 0).Format("15:04"),
		event.Name,
		event.Description,
	)
}
