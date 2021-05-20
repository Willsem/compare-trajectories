// Package model ...
package model

type Trajectory struct {
	Gps           Gps           `json:"gps"`
	Accelerometer Accelerometer `json:"acc"`
}

func (t *Trajectory) Check() error {
	if err := t.Gps.Check(); err != nil {
		return err
	}

	if err := t.Accelerometer.Check(); err != nil {
		return err
	}

	return nil
}
