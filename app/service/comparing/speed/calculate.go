// Package speed ...
package speed

import (
	"math"

	"github.com/Willsem/compare-trajectories/app/model"
)

type SpeedTrajectory struct {
	Speed []float64
	Gps   model.Gps
}

func Create(gps model.Gps) SpeedTrajectory {
	st := SpeedTrajectory{
		Gps:   gps,
		Speed: make([]float64, gps.Len()),
	}

	st.Speed[0] = 0
	for i := 1; i < gps.Len(); i++ {
		x1 := st.Gps.Long[i-1]
		y1 := st.Gps.Lat[i-1]
		x2 := st.Gps.Long[i]
		y2 := st.Gps.Lat[i]

		dist := math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2)
		time := model.DateDiffSeconds(st.Gps.Date[i], st.Gps.Date[i-1])

		st.Speed[i] = dist / time
	}

	return st
}
