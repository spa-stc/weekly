package dayupdatescron

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gocarina/gocsv"
)

type CsvData struct {
	Date     string `csv:"DATE"`
	Rday     string `csv:"R. DAY"`
	Location string `csv:"LOCATION"`
	Event    string `csv:"EVENT"`
	Grade9   string `csv:"9th GRADE"`
	Grade10  string `csv:"10th GRADE"`
	Grade11  string `csv:"11th GRADE"`
	Grade12  string `csv:"12th GRADE"`
	ApInfo   string `csv:"AP EXAMS"`
	CcInfo   string `csv:"CC TOPICS"`
}

// Get An io.Reader for the CSV from google, given the sheetid, and sheet.
func GetCSVReaderFromInfo(id, sheet string) (io.ReadCloser, error) {
	// Setup the url from where the csv will be downloaded.
	// I initially tried to use the google api, but the current go sdk
	// returns null values for empty sheet values, which is not acceptable.
	queryurl := fmt.Sprintf("https://docs.google.com/spreadsheets/d/%s/gviz/tq?tqx=out:csv&sheet=%s", id, sheet)

	req, err := http.Get(queryurl) //nolint:gosec,noctx // Due to lack of end user exposure.
	if err != nil {
		return nil, err
	}

	return req.Body, nil
}

// Read CSV Data Values From an IoReader.
func GetCSVDataFromCSV(r io.Reader) ([]CsvData, error) {
	var data []CsvData
	if err := gocsv.Unmarshal(r, &data); err != nil {
		return data, err
	}

	return data, nil
}

// Find A Date In the Data Array.
func FindCSVDataForDate(t time.Time, data []CsvData) CsvData {
	templ := t.Format("1/2/2006")
	for _, v := range data {
		if v.Date == templ {
			return v
		}
	}

	return CsvData{}
}

// Map A CSV Data Object Onto A Job Day.
func MapCSVDataOnJobDay(data CsvData, j JobDay) JobDay {
	day := j.Day

	day.RotationDay = data.Rday
	day.Location = data.Location
	day.XPeriod = data.Event
	day.Grade9 = data.Grade9
	day.Grade10 = data.Grade10
	day.Grade11 = data.Grade11
	day.Grade12 = data.Grade12
	day.CCInfo = data.CcInfo
	day.ApInfo = data.ApInfo

	j.Day = day
	return j
}

// Get X Info For An Array Of Job Days.
func populateOtherInfoForJobDays(j []JobDay, data []CsvData) []JobDay {
	var newjs []JobDay
	for _, day := range j {
		today := FindCSVDataForDate(day.RelatedTime, data)
		day = MapCSVDataOnJobDay(today, day)
		newjs = append(newjs, day)
	}

	return newjs
}
