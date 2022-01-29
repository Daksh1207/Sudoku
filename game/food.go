package game

import "math/rand"

type Food struct {
	X Point
}

func NewFoodAtRandom() Food {
	return Food{
		X: Point{
			X: rand.Intn(320),
			Y: rand.Intn(240),
		},
	}
}
