// Package model ...
package model

import "errors"

type Gps struct {
	Date []ParsedDate `json:"date"`
	Long []float64    `json:"long"`
	Lat  []float64    `json:"lat"`
}

func (gps *Gps) Check() error {
	if len(gps.Date) != len(gps.Long) || len(gps.Date) != len(gps.Lat) || len(gps.Long) != len(gps.Lat) {
		return errors.New("various len of arrays in gps")
	}
	return nil
}

func (gps *Gps) Len() int {
	return len(gps.Long)
}
