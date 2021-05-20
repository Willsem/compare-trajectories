// Package model ...
package model

type ComparedTrajectory struct {
	Backlog   []float64 `json:"backlog"`
	DeltaLong []float64 `json:"dlong"`
	DeltaLat  []float64 `json:"dlat"`
	DeltaAcc  []Point   `json:"dacc"`
	DeltaGyro []Point   `json:"dgyro"`
}
