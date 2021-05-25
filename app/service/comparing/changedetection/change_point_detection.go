// Package changedetection ...
package changedetection

import (
	"github.com/Willsem/compare-trajectories/app/model"
	"pgregory.net/changepoint"
)

func ChangePointDetection(trajectory *model.ComparedTrajectory) []model.ComparedTrajectory {
	data := trajectory.Backlog

	indexes := changepoint.NonParametric(data, 1)
	indexes = append([]int{0}, indexes...)
	indexes = append(indexes, len(trajectory.Backlog)-1)

	result := make([]model.ComparedTrajectory, 0)
	for i := 0; i < len(indexes)-1; i++ {
		start := indexes[i]
		finish := indexes[i+1]

		part := model.ComparedTrajectory{}
		for j := start; j <= finish; j++ {
			part.Backlog = append(part.Backlog, trajectory.Backlog[j])
			part.Long = append(part.Long, trajectory.Long[j])
			part.Lat = append(part.Lat, trajectory.Lat[j])
			part.DeltaLat = append(part.Backlog, trajectory.DeltaLat[j])
			part.DeltaLong = append(part.Backlog, trajectory.DeltaLong[j])
			part.DeltaAcc = append(part.DeltaAcc, trajectory.DeltaAcc[j])
			part.DeltaGyro = append(part.DeltaAcc, trajectory.DeltaGyro[j])
		}

		result = append(result, part)
	}

	return result
}
