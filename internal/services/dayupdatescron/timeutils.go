package dayupdatescron

import "time"

// Get the 7 Job Days Representing a Week.
func generateJobDays(t time.Time) []JobDay {
	sunday := GetSunday(t)
	week := GetWeek(sunday)

	var days []JobDay
	for _, v := range week {
		days = append(days, JobDay{RelatedTime: v})
	}

	return days
}

// Get the Last Sunday Before The Provided Time.
func GetSunday(t time.Time) time.Time {
	weekday := int(t.Weekday())
	time := t.AddDate(0, 0, -weekday)
	return time
}

// Get all of the tiem values from the next week given a Sunday.
func GetWeek(sunday time.Time) []time.Time {
	var times []time.Time
	for v := 0; v < 7; v++ {
		times = append(times, sunday.AddDate(0, 0, v))
	}

	return times
}

// Get the year, week, and weekday of any particular time.
func getYearWeekDay(t time.Time) (int, int, int) {
	year, week := t.ISOWeek()
	day := int(t.Weekday())

	return year, week, day
}
