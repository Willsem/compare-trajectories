package model

type Point struct {
	X int
	Y int
	Z int
}

func NewPoint(x, y, z int) (p *Point) {
	p.X = x
	p.Y = y
	p.Z = z
	return
}
