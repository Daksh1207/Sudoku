package core

import "math/rand"

type Food struct {
	X Point
}

func NewFoodAtRandom(xMax int, yMax int) Food {
	x := rand.Intn(xMax)
	y := rand.Intn(yMax)

	return Food{
		X: Point{
			X: x - x%SQUARE_SIZE,
			Y: y - y%SQUARE_SIZE,
		},
	}
}
