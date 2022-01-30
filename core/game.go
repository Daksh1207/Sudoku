package core

import (
	"math"
)

type Game struct {
	Width  int
	Height int
	Snake  Snake
	Food   Food
	Points int
	Radius int
}

func NewGame(width int, height int, radius int) Game {
	return Game{
		Width:  width,
		Height: height,
		Snake:  NewSnake(width/2, height/2),
		Food:   NewFoodAtRandom(width, height),
		Points: 0,
		Radius: radius,
	}
}

func (g *Game) Update() bool {
	snakePos := g.Snake.Head()

	if snakePos.X == 0 || snakePos.X == g.Width || snakePos.Y == 0 || snakePos.Y == g.Height {
		return false
	}

	// TODO check self collision
	g.Snake.Move()

	snakeHead := g.Snake.Head()

	radius := g.Radius / 2.0
	xDiff := float64(snakeHead.X + radius - (g.Food.X.X + radius))
	yDiff := float64(snakeHead.Y + radius - (g.Food.X.Y + radius))
	dist := math.Sqrt(xDiff*xDiff + yDiff*yDiff)

	if dist < float64(g.Radius) {
		g.Snake.Grow()
		g.Food = NewFoodAtRandom(g.Width, g.Height)
		g.Points += 1
	}

	return true
}
