// Package model ...
package model

type Point struct {
	X int `json:"x,string"`
	Y int `json:"y,string"`
	Z int `json:"z,string"`
}

type FloatPoint struct {
	X float64 `json:"x,string"`
	Y float64 `json:"y,string"`
	Z float64 `json:"z,string"`
}

func NewPoint(x, y, z int) (p *Point) {
	p.X = x
	p.Y = y
	p.Z = z
	return
}
