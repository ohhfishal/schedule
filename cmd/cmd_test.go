package cmd_test

import (
	"strings"
	"testing"
	"time"

	assert "github.com/alecthomas/assert/v2"
	"github.com/ohhfishal/schedule/cmd"
	"github.com/ohhfishal/schedule/db"
)

func Query(t *testing.T) *db.Queries {
	t.Helper()
	q, err := db.Connect(t.Context(), "sqlite", ":memory:")
	assert.NoError(t, err)
	return q
}

func Run(t *testing.T, args []string) (string, error) {
	var stdout strings.Builder
	err := cmd.Run(t.Context(), &stdout, args)
	return stdout.String(), err
}

func TestNew(t *testing.T) {
	tests := []struct {
		Name string
		Args []string
		Err  bool
	}{
		{Name: "No Args", Args: []string{}, Err: true},
		{Name: "No Date", Args: []string{"Test Event"}, Err: true},
		{Name: "Bad Date", Args: []string{"Test Event", "202-03-14"}, Err: true},
		{Name: "Bad Time", Args: []string{"Test Event", "2025-03-14", "11"}, Err: true},
		{Name: "Ok", Args: []string{"Test Event", "2025-03-14"}},
		{Name: "Time", Args: []string{"Test Event", "2025-03-14", "12:00"}},
		{Name: "Time & Description/short", Args: []string{"Test Event", "2025-03-14", "12:00", "-d", "description"}},
		{Name: "Description/short", Args: []string{"Test Event", "2025-03-14", "-d", "description"}},
		{Name: "Time & Description/short", Args: []string{"Test Event", "2025-03-14", "12:00", "--description", "description"}},
		{Name: "Description/short", Args: []string{"Test Event", "2025-03-14", "--description", "description"}},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			_, err := Run(t, append([]string{"new"}, test.Args...))
			if test.Err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	t.Run("ID not found", func(t *testing.T) {
		_, err := Run(t, []string{"delete", "99"})
		assert.Error(t, err)
	})
	t.Run("ID found", func(t *testing.T) {
		stdout := cmd.StdoutWriter{}
		q := Query(t)

		// Add some test events
		for range 5 {
			err := cmd.New{
				Name:      "Test Event",
				StartDate: time.Now(),
			}.Run(t.Context(), &stdout, q, time.Now)
			assert.NoError(t, err)
		}

		// Delete them all
		err := cmd.Delete{ID: []int64{1, 2, 3, 4, 5}}.Run(t.Context(), &stdout, q)
		assert.NoError(t, err)

		// Confirm no events are left
		events, err := q.GetAllEvents(t.Context())
		assert.Equal(t, 0, len(events), "events still in DB")
		assert.NoError(t, err)
	})
}
