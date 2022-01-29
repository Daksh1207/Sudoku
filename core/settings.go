package core

import "image/color"

type Settings struct {
	SnakeColor color.RGBA
	FoodColor  color.RGBA
	SquareSize int
}

func NewSettings() Settings {
	return Settings{
		SnakeColor: color.RGBA{R: 255, G: 255, B: 255, A: 255},
		FoodColor:  color.RGBA{R: 255, G: 0, B: 0, A: 255},
		SquareSize: 10,
	}
}
