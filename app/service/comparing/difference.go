// Package comparing ...
package comparing

import (
	"github.com/Willsem/compare-trajectories/app/model"
	"github.com/Willsem/compare-trajectories/app/service/comparing/interpolation"
	"github.com/Willsem/compare-trajectories/app/service/comparing/speed"
)

func difference(perfect interpolation.InterpolatedTrajectory, compared speed.SpeedTrajectory) (ct model.ComparedTrajectory, err error) {
	return
}
