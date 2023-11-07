package server

import (
	"time"
)

var dateParser = map[string]func() (from time.Time, to time.Time){
	"today": func() (from time.Time, to time.Time) {
		now := time.Now()

		from = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		to = from.AddDate(0, 0, 1).Add(-time.Second)

		return
	},
	"yesterday": func() (from time.Time, to time.Time) {
		now := time.Now()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

		from = today.AddDate(0, 0, -1)
		to = today.Add(-time.Second)

		return
	},
	"week": func() (from time.Time, to time.Time) {
		from = dateOfPrevWeekday(time.Monday, true)
		to = from.AddDate(0, 0, 7).Add(-time.Second)

		return
	},
	"last-week": func() (from time.Time, to time.Time) {
		to = dateOfPrevWeekday(time.Sunday, false)
		from = time.Date(to.Year(), to.Month(), to.Day(), 0, 0, 0, 0, to.Location()).AddDate(0, 0, -6)

		to = to.AddDate(0, 0, 1).Add(-time.Second)

		return
	},
	"month": func() (from time.Time, to time.Time) {
		now := time.Now()

		from = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		to = from.AddDate(0, 1, 0).Add(-time.Second)

		return
	},
	"last-month": func() (from time.Time, to time.Time) {
		now := time.Now()

		currMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		from = currMonth.AddDate(0, -1, 0)

		to = currMonth.Add(-time.Second)

		return
	},
	"monday": func() (from time.Time, to time.Time) {
		from = dateOfPrevWeekday(time.Monday, false)
		to = from.AddDate(0, 0, 1).Add(-time.Second)

		return
	},
	"tuesday": func() (from time.Time, to time.Time) {
		from = dateOfPrevWeekday(time.Tuesday, false)
		to = from.AddDate(0, 0, 1).Add(-time.Second)

		return
	},
	"wednesday": func() (from time.Time, to time.Time) {
		from = dateOfPrevWeekday(time.Wednesday, false)
		to = from.AddDate(0, 0, 1).Add(-time.Second)

		return
	},
	"thursday": func() (from time.Time, to time.Time) {
		from = dateOfPrevWeekday(time.Thursday, false)
		to = from.AddDate(0, 0, 1).Add(-time.Second)

		return
	},
	"friday": func() (from time.Time, to time.Time) {
		from = dateOfPrevWeekday(time.Friday, false)
		to = from.AddDate(0, 0, 1).Add(-time.Second)

		return
	},
	"saturday": func() (from time.Time, to time.Time) {
		from = dateOfPrevWeekday(time.Saturday, false)
		to = from.AddDate(0, 0, 1).Add(-time.Second)

		return
	},
	"sunday": func() (from time.Time, to time.Time) {
		from = dateOfPrevWeekday(time.Sunday, false)
		to = from.AddDate(0, 0, 1).Add(-time.Second)

		return
	},
}

func dateOfPrevWeekday(weekday time.Weekday, includeToday bool) time.Time {
	currDate := time.Now()
	daysToSubtract := currDate.Weekday() - weekday

	if daysToSubtract < 0 || (daysToSubtract == 0 && includeToday) {
		daysToSubtract = 7 + daysToSubtract
	}

	day := time.Date(currDate.Year(), currDate.Month(), currDate.Day(), 0, 0, 0, 0, currDate.Location()).AddDate(0, 0, -int(daysToSubtract))

	return day
}
