// Package comparing ...
package comparing

import (
	"github.com/Willsem/compare-trajectories/app/model"
)

type speedTrajectory struct {
	Speed []float64
}

func Compare(perfect model.Trajectory, compared model.Trajectory) (ct model.ComparedTrajectory, err error) {
	err = perfect.Check()
	if err != nil {
		return
	}

	err = compared.Check()
	if err != nil {
		return
	}

	return
}
