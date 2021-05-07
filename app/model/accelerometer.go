package model

import "time"

type Acceletometer struct {
	Date time.Time `json:"date"`
	Acc  Point     `json:"acc"`
	Gyro Point     `json:"gyro"`
}

func NewAccelerometer(date string, a Point, g Point) (acc *Acceletometer, err error) {
	acc.Date, err = parseDate(date)
	acc.Acc = a
	acc.Gyro = g
	return
}
