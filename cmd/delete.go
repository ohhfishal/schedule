package cmd

import (
	"context"
	"fmt"

	"github.com/ohhfishal/schedule/db"
)

type Delete struct {
	ID []int64 `arg:"" required:"" help:"ID of event to delete."`
}

func (cmd Delete) Run(ctx context.Context, stdout Stdout, queries *db.Queries) error {
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
		fmt.Fprintf(stdout.Verbose(), "deleted %d\n", id)
	}
	return nil
}
