// Package math ...
package math

import "math"

func Distance(x1, y1, x2, y2 float64) float64 {
	return math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2)
}
