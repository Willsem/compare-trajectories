package model

import "time"

type Gps struct {
	Date time.Time `json:"date"`
	Long float32   `json:"long"`
	Lat  float32   `json:"lat"`
}

func NewGps(date string, long, lat float32) (gps *Gps, err error) {
	gps.Date, err = parseDate(date)
	gps.Long = long
	gps.Lat = lat
	return
}
