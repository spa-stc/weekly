package dayupdatescron_test

import (
	"testing"

	"github.com/spa-stc/newsletter/internal/services/dayupdatescron"
	"github.com/spa-stc/newsletter/internal/types/models"
	"gotest.tools/assert"
)

func TestMapCsvDataOnJobDay(t *testing.T) {
	data := dayupdatescron.CsvData{
		Date:     "1",
		Rday:     "2",
		Location: "3",
		Event:    "4",
		Grade9:   "5",
		Grade10:  "6",
		Grade11:  "7",
		Grade12:  "8",
		ApInfo:   "9",
		CcInfo:   "10",
	}

	day := models.Day{
		RotationDay: "2",
		Location:    "3",
		XPeriod:     "4",
		Grade9:      "5",
		Grade10:     "6",
		Grade11:     "7",
		Grade12:     "8",
		ApInfo:      "9",
		CCInfo:      "10",
	}

	expected := dayupdatescron.JobDay{
		Day: day,
	}

	assert.DeepEqual(t, expected, dayupdatescron.MapCSVDataOnJobDay(data, dayupdatescron.JobDay{}))
}
