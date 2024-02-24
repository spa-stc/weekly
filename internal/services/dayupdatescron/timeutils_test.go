package dayupdatescron_test

import (
	"testing"
	"time"

	"github.com/spa-stc/newsletter/internal/services/dayupdatescron"
	"gotest.tools/assert"
)

func TestGetSunday(t *testing.T) {
	expected := time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)
	initial := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	actual := dayupdatescron.GetSunday(initial)

	assert.Equal(t, expected, actual, "times were not equal")
}

func TestGetWeek(t *testing.T) {
	expected := []time.Time{
		time.Date(2023, 12, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 12, 4, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 12, 5, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 12, 6, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 12, 7, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 12, 8, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 12, 9, 0, 0, 0, 0, time.UTC),
	}

	sunday := time.Date(2023, 12, 3, 0, 0, 0, 0, time.UTC)

	actual := dayupdatescron.GetWeek(sunday)

	assert.DeepEqual(t, expected, actual)
}
