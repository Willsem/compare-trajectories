// Package interpolation ...
package interpolation

import (
	"github.com/Willsem/compare-trajectories/app/model"
	"github.com/cnkei/gospline"
)

func BSplineInterpolation(gps *model.Gps) gospline.Spline {
	return gospline.NewCubicSpline(gps.Lat, gps.Long)
}
