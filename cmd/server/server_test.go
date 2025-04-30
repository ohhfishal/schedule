package server_test

import (
	"context"
	"io"
	"log/slog"
	"testing"

	assert "github.com/alecthomas/assert/v2"
	"github.com/ohhfishal/schedule/cmd/server"
	"github.com/ohhfishal/schedule/db"
)

func DB(t *testing.T) *db.Queries {
	t.Helper()
	q, err := db.Connect(t.Context(), "sqlite", ":memory:")
	assert.NoError(t, err)
	return q
}

func NewLogger(w io.Writer) *slog.Logger {
	// handler := slog.NewJSONHandler(w, &slog.HandlerOptions{
	// 	Level: slog.LevelDebug,
	// })
	return slog.New(slog.DiscardHandler)
}

func TestResponses(t *testing.T) {
	tests := []struct {
		Name     string
		Requests []struct {
			Status int
			Route  string
			Method string
		}
	}{}

	logger := NewLogger(nil)
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			// TODO: Make configurable
			s := server.Server{"localhost:8080"}

			ctx, cancel := context.WithCancel(t.Context())
			defer cancel()

			db := DB(t)
			go func() {
				err := s.Run(ctx, logger, db)
				assert.NoError(t, err)
			}()

			for _, _ = range test.Requests {
				// TODO: Implement main test logic
			}
		})
	}
}

func TestContext(t *testing.T) {

}
