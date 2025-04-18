package get

import (
	"context"
	"fmt"
	"io"

	"github.com/ohhfishal/schedule/db"
)

type All struct {
	// TODO: Implement
	// Out string `short:"o" enum:"markdown,raw" default:"raw" help:"Output format (markdown,raw)"`
	// TODO: JSON output? Markdown? YAML?
}

func (cmd All) Run(ctx context.Context, stdout io.Writer, queries *db.Queries) error {
	if stdout == nil {
		return fmt.Errorf("no stdout")
	}

	events, err := queries.GetAllEvents(ctx)
	if err != nil {
		return fmt.Errorf(`getting events: %w`, err)
	}

	if len(events) == 0 {
		fmt.Fprintln(stdout, `No events`)
		return nil
	}

	for _, event := range events {
		fmt.Fprintf(stdout, "%+v\n", event)
	}
	return nil
}
