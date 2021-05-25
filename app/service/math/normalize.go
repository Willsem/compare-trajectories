// Package math ...
package math

import (
	m "math"
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
