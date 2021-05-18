// Package filtering ...
package filtering

import (
	"errors"

	"github.com/konimarti/kalman"
	"github.com/konimarti/lti"
	"gonum.org/v1/gonum/mat"
)

func KalmanFilter(x []float64, y []float64) (filteredX []float64, filteredY []float64, err error) {
	defer func() {
		rec := recover()
		if rec != nil {
			err = rec.(error)
		}
	}()

	if len(x) != len(y) {
		return nil, nil, errors.New("cannot filter arrays with different lens")
	}

	ctx := kalman.Context{
		X: mat.NewVecDense(2, []float64{x[0], y[0]}),
		P: mat.NewDense(2, 2, []float64{
			1, 0,
			0, 1,
		}),
	}

	dt := 0.00001

	lti := lti.Discrete{
		// prediction matrix
		Ad: mat.NewDense(2, 2, []float64{
			1, dt,
			0, 1,
		}),
		// no external influence
		Bd: mat.NewDense(2, 2, nil),
		// scaling matrix for measurement
		C: mat.NewDense(2, 2, []float64{
			1, 0,
			0, 1,
		}),
		// scaling matrix for control
		D: mat.NewDense(2, 2, nil),
	}

	// G
	G := mat.NewDense(2, 2, []float64{
		1, 0,
		0, 1,
	})
	var Gd mat.Dense
	Gd.Mul(lti.Ad, G)

	// process model covariance matrix
	qk := mat.NewDense(2, 2, []float64{
		0.01, 0,
		0, 0.01,
	})
	var Q mat.Dense
	Q.Product(&Gd, qk, Gd.T())

	// measurement errors
	corr := 0.5
	R := mat.NewDense(2, 2, []float64{1, corr, corr, 1})

	// create noise struct
	nse := kalman.Noise{
		Q: &Q,
		R: R,
	}

	// create Kalman filter
	filter := kalman.NewFilter(lti, nse)

	// no control
	control := mat.NewVecDense(2, nil)

	for i := 0; i < len(x); i++ {
		measurement := mat.NewVecDense(2, []float64{x[i], y[i]})
		filtered := filter.Apply(&ctx, measurement, control)

		filteredX = append(filteredX, filtered.AtVec(0))
		filteredY = append(filteredY, filtered.AtVec(1))
	}

	return filteredX, filteredY, nil
}
