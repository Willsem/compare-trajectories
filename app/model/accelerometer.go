package model

type Acceletometer struct {
	Date []ParsedDate `json:"date"`
	Acc  []Point      `json:"acc"`
	Gyro []Point      `json:"gyro"`
}
