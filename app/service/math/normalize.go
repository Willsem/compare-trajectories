// Package math ...
package math

import (
	m "math"

	"github.com/Willsem/compare-trajectories/app/model"
)

func Normalize(array []float64) []float64 {
	min := array[0]
	max := array[0]
	n := len(array)

	for i := 1; i < n; i++ {
		if array[i] < min {
			min = array[i]
		}

		if array[i] > max {
			max = array[i]
		}
	}

	result := make([]float64, n)
	for i := 0; i < n; i++ {
		if array[i] < 0 {
			result[i] = array[i] / m.Abs(min)
		} else {
			result[i] = array[i] / max
		}
	}

	return result
}

func PointNormalize(array []model.FloatPoint) []model.FloatPoint {
	min := array[0].X
	max := array[0].X

	n := len(array)

	for i := 0; i < n; i++ {
		if array[i].X < min {
			min = array[i].X
		}

		if array[i].X > max {
			max = array[i].X
		}

		if array[i].Y < min {
			min = array[i].Y
		}

		if array[i].Y > max {
			max = array[i].Y
		}

		if array[i].Z < min {
			min = array[i].Z
		}

		if array[i].Z > max {
			max = array[i].Z
		}
	}

	result := make([]model.FloatPoint, n)
	for i := 0; i < n; i++ {
		if array[i].X < 0 {
			result[i].X = array[i].X / m.Abs(min)
		} else {
			result[i].X = array[i].X / max
		}

		if array[i].Y < 0 {
			result[i].Y = array[i].Y / m.Abs(min)
		} else {
			result[i].Y = array[i].Y / max
		}

		if array[i].Z < 0 {
			result[i].Z = array[i].Z / m.Abs(max)
		} else {
			result[i].Z = array[i].Z / max
		}
	}

	return result
}
