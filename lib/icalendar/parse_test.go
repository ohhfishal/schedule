package icalendar_test

import (
	_ "embed"
	"testing"

	// 	"fmt"
	// 	"testing"
	assert "github.com/alecthomas/assert/v2"
	"github.com/ohhfishal/schedule/lib/icalendar"
)

//go:embed test.ics
var icsExample string

func TestParse(t *testing.T) {
	cal, err := icalendar.Parse(icsExample)
	assert.NoError(t, err)
	assert.NotZero(t, cal)
}
