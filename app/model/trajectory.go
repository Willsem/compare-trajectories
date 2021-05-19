// Package model ...
package model

type Trajectory struct {
	Gps           Gps           `json:"gps"`
	Accelerometer Accelerometer `json:"acc"`
}
