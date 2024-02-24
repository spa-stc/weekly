package dayupdatescron

import (
	"net/http"
	"time"

	"github.com/apognu/gocal"
)

// Get The Lunch Calendar From The Calendar URL.
func getLunchIcalFromURL(url string) (*gocal.Gocal, error) {
	res, err := http.Get(url) //nolint:gosec,noctx // Due to the fact that this is not exposed to end users.
	if err != nil {
		return &gocal.Gocal{}, err
	}
	defer res.Body.Close()

	cal := gocal.NewParser(res.Body)
	// Load All Calendar Events In The Calendar Into Memory.
	// This can be changed to optimize, once we load the get reqest output into a file.
	cal.SkipBounds = true

	err = cal.Parse()
	return cal, err
}

// Get A Lunch Value, Or "Not Available" From A Gocal Parser.
func GetLunch(parser *gocal.Gocal, today time.Time) string {
	todayonly := today.UTC().Format(time.DateOnly)
	for _, event := range parser.Events {
		if event.Start.UTC().Format(time.DateOnly) == todayonly {
			return event.Description
		}
	}

	return "Not Available"
}
