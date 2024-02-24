package dayupdatescron_test

import (
	"strings"
	"testing"
	"time"

	"github.com/apognu/gocal"
	"github.com/spa-stc/newsletter/internal/services/dayupdatescron"
	"gotest.tools/assert"
)

const ical = `
BEGIN:VEVENT
DTSTAMP:20231229T204822Z
UID:event_184301_20231222@api.veracross.com
DTSTART;VALUE=DATE:20231222
DESCRIPTION:Chef's Choice
SEQUENCE:783913291
SUMMARY:Randolph Campus Lunch
END:VEVENT`

type testVeracrossLunchParse struct {
	Time     string
	Expected string
}

func TestLunchParse(t *testing.T) {
	var veracrossLunchParseTests = []testVeracrossLunchParse{
		{Time: "2023-12-22", Expected: "Chef's Choice"},
		{Time: "2023-12-18", Expected: "Not Available"},
	}
	cal := gocal.NewParser(strings.NewReader(ical))
	cal.SkipBounds = true
	if err := cal.Parse(); err != nil {
		t.Fatal(err)
	}

	for _, v := range veracrossLunchParseTests {
		te, err := time.Parse(time.DateOnly, v.Time)
		if err != nil {
			t.Fatal(err)
		}
		lunch := dayupdatescron.GetLunch(cal, te)
		assert.Equal(t, lunch, v.Expected)
	}
}
