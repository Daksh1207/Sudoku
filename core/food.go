package core

import "math/rand"

type Food struct {
	X Point
}

func NewFoodAtRandom(xMax int, yMax int) Food {
	return Food{
		X: Point{
			X: rand.Intn(xMax),
			Y: rand.Intn(yMax),
		},
	}
}
