package model

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

func NewPoint(x, y, z int) (p *Point) {
	p.X = x
	p.Y = y
	p.Z = z
	return
}
