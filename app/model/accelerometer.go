package model

import "time"

type Acceletometer struct {
	Date time.Time
	Acc  Point
	Gyro Point
}

func NewAccelerometer(date string, a Point, g Point) (acc *Acceletometer, err error) {
	acc.Date, err = parseDate(date)
	acc.Acc = a
	acc.Gyro = g
	return
}
