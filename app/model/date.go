package model

import "time"

func parseDate(date string) (t time.Time, err error) {
	const layout = "01/02/06 15:04:05.000000"
	t, err = time.Parse(layout, date)
	return
}
