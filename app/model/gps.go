package model

import "time"

type Gps struct {
	Date []time.Time `json:"date"`
	Long []float32   `json:"long"`
	Lat  []float32   `json:"lat"`
}
