// Package model ...
package model

type ComparedTrajectory struct {
	Backlog   []float64 `json:"backlog"`
	Long      []float64 `json:"long"`
	Lat       []float64 `json:"lat"`
	DeltaLong []float64 `json:"dlong"`
	DeltaLat  []float64 `json:"dlat"`
	DeltaAcc  []Point   `json:"dacc"`
	DeltaGyro []Point   `json:"dgyro"`
}
