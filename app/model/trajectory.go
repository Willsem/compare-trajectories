// Package model ...
package model

type Trajectory struct {
	Gps           Gps           `json:"gps"`
	Accelerometer Acceletometer `json:"acc"`
}
