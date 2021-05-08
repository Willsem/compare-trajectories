package model

import "time"

type Acceletometer struct {
	Date []time.Time `json:"date"`
	Acc  []Point     `json:"acc"`
	Gyro []Point     `json:"gyro"`
}
