// Package interpolation ...
package interpolation

import (
	"github.com/Willsem/compare-trajectories/app/model"
	"github.com/Willsem/compare-trajectories/app/service/comparing/speed"
	"github.com/Willsem/compare-trajectories/app/service/comparing/spline"
)

type InterpolatedTrajectory struct {
	Trajectory speed.SpeedTrajectory
	Acc        model.Accelerometer

	splineTrajectory *spline.Bspline
	splineAcc        *spline.Bspline
	splineGyro       *spline.Bspline
}

func CreateTrajectory(trajectory *speed.SpeedTrajectory, acc *model.Accelerometer) InterpolatedTrajectory {
	pointsTrajectory := make([][]float64, 0)
	pointsAcc := make([][]float64, 0)
	pointsGyro := make([][]float64, 0)

	for i := 0; i < trajectory.Gps.Len(); i++ {
		pointsTrajectory = append(pointsTrajectory, []float64{
			trajectory.Speed[i],
			trajectory.Gps.Lat[i],
			trajectory.Gps.Long[i],
		})
	}

	for i := 0; i < acc.Len(); i++ {
		pointsAcc = append(pointsAcc, []float64{
			float64(acc.Acc[i].X),
			float64(acc.Acc[i].Y),
			float64(acc.Acc[i].Z),
		})

		pointsGyro = append(pointsAcc, []float64{
			float64(acc.Gyro[i].X),
			float64(acc.Gyro[i].Y),
			float64(acc.Gyro[i].Z),
		})
	}

	it := InterpolatedTrajectory{
		Trajectory:       *trajectory,
		Acc:              *acc,
		splineTrajectory: spline.NewBspline(pointsTrajectory, 3, false),
		splineAcc:        spline.NewBspline(pointsAcc, 3, false),
		splineGyro:       spline.NewBspline(pointsGyro, 3, false),
	}

	it.splineTrajectory.Init()
	it.splineAcc.Init()
	it.splineGyro.Init()

	return it
}

func (it *InterpolatedTrajectory) TakeValues(start, finish int) ([]float64, []float64, []float64, []model.Point, []model.Point) {
	if start > finish {
		start, finish = finish, start
	}

	speed := make([]float64, 0)
	lat := make([]float64, 0)
	long := make([]float64, 0)
	acc := make([]model.Point, 0)
	gyro := make([]model.Point, 0)

	startT := float64(start) / float64(it.Trajectory.Gps.Len())
	finishT := float64(finish) / float64(it.Trajectory.Gps.Len())
	delta := (finishT - startT) / 10.0

	for t := startT; t <= finishT; t += delta {
		traj := it.splineTrajectory.Interpolate(t)
		speed = append(speed, traj[0])
		lat = append(lat, traj[1])
		long = append(long, traj[2])

		accInter := it.splineAcc.Interpolate(t)
		acc = append(acc, model.Point{
			X: int(accInter[0]),
			Y: int(accInter[1]),
			Z: int(accInter[2]),
		})

		gyroInter := it.splineGyro.Interpolate(t)
		gyro = append(gyro, model.Point{
			X: int(gyroInter[0]),
			Y: int(gyroInter[1]),
			Z: int(gyroInter[2]),
		})
	}

	return speed, lat, long, acc, gyro
}
