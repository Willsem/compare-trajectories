// Package model ...
package model

import "errors"

type Accelerometer struct {
	Date []ParsedDate `json:"date"`
	Acc  []Point      `json:"acc"`
	Gyro []Point      `json:"gyro"`
}

func (acc *Accelerometer) Check() error {
	if len(acc.Date) != len(acc.Acc) || len(acc.Date) != len(acc.Gyro) || len(acc.Acc) != len(acc.Gyro) {
		return errors.New("various len of arrays in gps")
	}
	return nil
}

func (acc *Accelerometer) Len() int {
	return len(acc.Acc)
}
