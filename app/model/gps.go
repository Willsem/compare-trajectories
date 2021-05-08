package model

type Gps struct {
	Date []ParsedDate `json:"date"`
	Long []float32    `json:"long"`
	Lat  []float32    `json:"lat"`
}
