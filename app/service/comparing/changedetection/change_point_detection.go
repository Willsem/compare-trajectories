// Package changedetection ...
package changedetection

import (
	"pgregory.net/changepoint"

	"github.com/Willsem/compare-trajectories/app/model"
	"github.com/Willsem/compare-trajectories/app/service/math"
)

func ChangePointDetection(trajectory *model.ComparedTrajectory) []model.ComparedTrajectory {
	data := trajectory.Backlog

	indexes := changepoint.NonParametric(data, 1)
	indexes = append([]int{0}, indexes...)
	indexes = append(indexes, len(trajectory.Backlog)-1)

	trajectory.Backlog = math.Normalize(trajectory.Backlog)
	trajectory.DeltaLong = math.Normalize(trajectory.DeltaLong)
	trajectory.DeltaLat = math.Normalize(trajectory.DeltaLat)
	trajectory.DeltaAcc = math.PointNormalize(trajectory.DeltaAcc)
	trajectory.DeltaGyro = math.PointNormalize(trajectory.DeltaGyro)

	result := make([]model.ComparedTrajectory, 0)
	for i := 0; i < len(indexes)-1; i++ {
		start := indexes[i]
		finish := indexes[i+1]

		part := model.ComparedTrajectory{}
		for j := start; j <= finish; j++ {
			part.Backlog = append(part.Backlog, trajectory.Backlog[j])
			part.Long = append(part.Long, trajectory.Long[j])
			part.Lat = append(part.Lat, trajectory.Lat[j])
			part.DeltaLat = append(part.DeltaLat, trajectory.DeltaLat[j])
			part.DeltaLong = append(part.DeltaLong, trajectory.DeltaLong[j])
			part.DeltaAcc = append(part.DeltaAcc, trajectory.DeltaAcc[j])
			part.DeltaGyro = append(part.DeltaGyro, trajectory.DeltaGyro[j])
		}

		result = append(result, part)
	}

	return result
}
