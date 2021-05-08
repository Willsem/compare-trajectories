package model

type Gps struct {
	Date []ParsedDate `json:"date"`
	Long []float64    `json:"long"`
	Lat  []float64    `json:"lat"`
}
