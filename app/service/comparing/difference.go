// Package comparing ...
package comparing

import (
	m "math"

	"github.com/Willsem/compare-trajectories/app/model"
	"github.com/Willsem/compare-trajectories/app/service/comparing/interpolation"
	"github.com/Willsem/compare-trajectories/app/service/comparing/speed"
	"github.com/Willsem/compare-trajectories/app/service/math"
)

func difference(perfect interpolation.InterpolatedTrajectory, compared speed.SpeedTrajectory, comparedAcc model.Accelerometer) (ct model.ComparedTrajectory, err error) {
	ct = model.ComparedTrajectory{
		Backlog:   make([]float64, compared.Gps.Len()),
		DeltaLong: make([]float64, compared.Gps.Len()),
		DeltaLat:  make([]float64, compared.Gps.Len()),
		DeltaAcc:  make([]model.Point, compared.Gps.Len()),
		DeltaGyro: make([]model.Point, compared.Gps.Len()),
	}
	for i := 0; i < compared.Gps.Len(); i++ {
		var minDist float64 = -1
		var minIndex int
		x1 := compared.Gps.Lat[i]
		y1 := compared.Gps.Long[i]

		for j := 0; j < perfect.Trajectory.Gps.Len(); j++ {
			x2 := perfect.Trajectory.Gps.Lat[j]
			y2 := perfect.Trajectory.Gps.Long[j]
			dist := math.Distance(x1, y1, x2, y2)
			if minDist == -1 || dist < minDist {
				minDist = dist
				minIndex = j
			}
		}

		var secondIndex int
		if minIndex == 0 {
			secondIndex = 1
		} else if minIndex == perfect.Trajectory.Gps.Len()-1 {
			secondIndex = minIndex - 1
		} else {
			x2 := perfect.Trajectory.Gps.Lat[minIndex-1]
			y2 := perfect.Trajectory.Gps.Long[minIndex-1]
			x3 := perfect.Trajectory.Gps.Lat[minIndex+1]
			y3 := perfect.Trajectory.Gps.Long[minIndex+1]

			dist1 := math.Distance(x1, y1, x2, y2)
			dist2 := math.Distance(x1, y1, x3, y3)

			if dist1 < dist2 {
				secondIndex = minIndex - 1
			} else {
				secondIndex = minIndex + 1
			}
		}

		speed, lat, long, acc, gyro := perfect.TakeValues(minIndex, secondIndex)

		minDist = -1
		for j := 0; j < len(lat); j++ {
			x2 := lat[j]
			y2 := long[j]
			dist := math.Distance(x1, y1, x2, y2)

			if minDist == -1 || dist < minDist {
				minDist = dist
				minIndex = j
			}
		}

		accIndex := 0
		var minDateDiff float64 = -1
		for j := 0; j < comparedAcc.Len(); j++ {
			dateDiff := m.Abs(model.DateDiffSeconds(compared.Gps.Date[i], comparedAcc.Date[j]))
			if minDateDiff == -1 || dateDiff < minDateDiff {
				minDateDiff = dateDiff
				accIndex = j
			} else {
				break
			}
		}

		ct.Backlog[i] = speed[minIndex] - compared.Speed[i]
		ct.DeltaLat[i] = lat[minIndex] - compared.Gps.Lat[i]
		ct.DeltaLong[i] = long[minIndex] - compared.Gps.Long[i]

		ct.DeltaAcc[i].X = acc[minIndex].X - comparedAcc.Acc[accIndex].X
		ct.DeltaAcc[i].Y = acc[minIndex].Y - comparedAcc.Acc[accIndex].Y
		ct.DeltaAcc[i].Z = acc[minIndex].Z - comparedAcc.Acc[accIndex].Z

		ct.DeltaGyro[i].X = gyro[minIndex].X - comparedAcc.Gyro[accIndex].X
		ct.DeltaGyro[i].Y = gyro[minIndex].Y - comparedAcc.Gyro[accIndex].Y
		ct.DeltaGyro[i].Z = gyro[minIndex].Z - comparedAcc.Gyro[accIndex].Z
	}

	ct.Backlog = math.Normalize(ct.Backlog)

	return
}
