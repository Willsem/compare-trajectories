// Package comparing ...
package comparing

import (
	"github.com/Willsem/compare-trajectories/app/model"
	"github.com/Willsem/compare-trajectories/app/service/comparing/changedetection"
	"github.com/Willsem/compare-trajectories/app/service/comparing/interpolation"
	"github.com/Willsem/compare-trajectories/app/service/comparing/speed"
)

func Compare(perfect model.Trajectory, compared model.Trajectory) (cts []model.ComparedTrajectory, err error) {
	defer func() {
		rec := recover()
		if rec != nil {
			err = rec.(error)
		}
	}()

	err = perfect.Check()
	if err != nil {
		return
	}

	err = compared.Check()
	if err != nil {
		return
	}

	perfectSpeed := speed.Create(perfect.Gps)
	comparedSpeed := speed.Create(compared.Gps)

	perfectInterpolate := interpolation.CreateTrajectory(&perfectSpeed, &perfect.Accelerometer)
	ct, err := difference(perfectInterpolate, comparedSpeed, compared.Accelerometer)

	cts = changedetection.ChangePointDetection(&ct)
	return
}
