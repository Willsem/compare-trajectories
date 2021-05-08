package service

import (
	"github.com/konimarti/kalman"
	"github.com/konimarti/lti"
)

func FilterArray(array []float32) []float32 {
	filter := kalman.NewFilter(
		lti.Discrete{},
		kalman.Noise{},
	)
}
