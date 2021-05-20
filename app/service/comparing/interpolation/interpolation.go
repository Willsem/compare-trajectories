// Package interpolation ...
package interpolation

import (
	"github.com/Willsem/compare-trajectories/app/model"
	"github.com/Willsem/compare-trajectories/app/service/comparing/speed"
)

type InterpolatedTrajectory struct {
	trajectory speed.SpeedTrajectory
	acc        model.Accelerometer
}

func CreateTrajectory(trajectory *speed.SpeedTrajectory, acc *model.Accelerometer) InterpolatedTrajectory {
	return InterpolatedTrajectory{
		trajectory: *trajectory,
		acc:        *acc,
	}
}

// func TakeValues
