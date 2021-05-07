package model

import "time"

type Gps struct {
	Date time.Time
	Long float32
	Lat  float32
}

func NewGps(date string, long, lat float32) (gps *Gps, err error) {
	gps.Date, err = parseDate(date)
	gps.Long = long
	gps.Lat = lat
	return
}
