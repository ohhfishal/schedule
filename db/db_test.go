package db_test

import (
	"fmt"
	"testing"
	"time"

	assert "github.com/alecthomas/assert/v2"
	"github.com/ohhfishal/schedule/db"
)

func New(t *testing.T) *db.Queries {
	q, err := db.Connect(t.Context(), `sqlite`, `:memory:`)
	assert.Equal(t, err, nil)
	return q
}

func TestRegression(t *testing.T) {
	q := New(t)
	for i := 0; i < 100; i++ {
		input := db.CreateEventParams{
			Name:      fmt.Sprintf(`event: %d`, i),
			StartTime: time.Now().Unix(),
		}
		event, err := q.CreateEvent(t.Context(), input)
		assert.Equal(t, err, nil)

		_, err = q.DeleteEvent(t.Context(), event.ID)
		assert.Equal(t, err, nil)
	}
}
