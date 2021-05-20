// Package interpolation ...
package interpolation

import (
	"github.com/Willsem/compare-trajectories/app/model"
	"github.com/Willsem/compare-trajectories/app/service/comparing/speed"
)

type InterpolatedTrajectory struct {
	Trajectory speed.SpeedTrajectory
	Acc        model.Accelerometer
}

func CreateTrajectory(trajectory *speed.SpeedTrajectory, acc *model.Accelerometer) InterpolatedTrajectory {
	return InterpolatedTrajectory{
		Trajectory: *trajectory,
		Acc:        *acc,
	}
}

func (it *InterpolatedTrajectory) TakeValues(start, finish int) ([]float64, []float64, []float64, []model.Point, []model.Point) {
	if start > finish {
		start, finish = finish, start
	}

	return nil, nil, nil, nil, nil
}
