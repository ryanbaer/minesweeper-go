package lib

import "errors"

var ErrOutOfBounds = errors.New("CoordinateOutOfBounds")

type Coordinate struct {
	X int
	Y int
}

type Location struct {
	Coordinate
	Visited bool
}

func (l *Location) Visit() {
	l.Visited = true
}

func (c *Coordinate) Add(c1 *Coordinate) *Coordinate {
	return &Coordinate{
		X: c.X + c1.X,
		Y: c.Y + c1.Y,
	}
}

func (c *Coordinate) Perimeter() []*Coordinate {
	return []*Coordinate{
		&Coordinate{-1, -1}, // Upper Left
		&Coordinate{-1, 0},  // Up
		&Coordinate{-1, 1},  // Upper right
		&Coordinate{0, 1},   // Right
		&Coordinate{1, 1},   // Lower Right
		&Coordinate{1, 0},   // Bottom
		&Coordinate{1, -1},  // Lower Left
		&Coordinate{0, -1},  // Left
	}
}

func (c *Coordinate) RealPerimeter() []*Coordinate {
	p := c.Perimeter()
	coords := make([]*Coordinate, len(p))
	for i, c1 := range p {
		coords[i] = c.Add(c1)
	}

	return coords
}

var Perimeter = []Coordinate{}
