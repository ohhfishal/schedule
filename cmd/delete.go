package cmd

import (
	"context"
	"fmt"
	"io"

	"github.com/ohhfishal/schedule/db"
)

type Delete struct {
	ID []int64 `arg:"" required:"" help:"ID of event to delete."`
}

func (cmd Delete) Run(ctx context.Context, stdout io.Writer, verbose bool, queries *db.Queries) error {
	for _, id := range cmd.ID {
		result, err := queries.DeleteEvent(ctx, id)
		if err != nil {
			return fmt.Errorf(`deleting: %d: %w`, id, err)
		}
		count, err := result.RowsAffected()
		if err != nil {
			return fmt.Errorf("deleting %d: failed to get result: %w", id, err)
		}

		if count == 0 {
			return fmt.Errorf("deleting %d: event not found", id)
		}
		if verbose {
			fmt.Fprintf(stdout, "deleted %d\n", id)
		}
	}
	return nil
}
