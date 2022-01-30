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

func (g *Game) checkBorderCollision() bool {
	snakePos := g.Snake.Head()
	return snakePos.X == 0 || snakePos.X == g.Width || snakePos.Y == 0 || snakePos.Y == g.Height
}

func (g *Game) checkSnakeSelfCollision() bool {
	head := g.Snake.Head()
	count := 0

	for _, sp := range g.Snake.X {
		// TODO add radius to collision detection
		if sp.X == head.X && sp.Y == head.Y {
			count += 1
		}
	}

	return count == 2
}

func (g *Game) checkFoodCollision() bool {
	snakeHead := g.Snake.Head()

	radius := g.Radius / 2.0
	xDiff := float64(snakeHead.X + radius - (g.Food.X.X + radius))
	yDiff := float64(snakeHead.Y + radius - (g.Food.X.Y + radius))
	dist := math.Sqrt(xDiff*xDiff + yDiff*yDiff)

	return dist < float64(g.Radius)
}

func (g *Game) Update() bool {

	if g.checkBorderCollision() || g.checkSnakeSelfCollision() {
		return false
	}

	g.Snake.Move()

	if g.checkFoodCollision() {
		g.Snake.Grow()
		g.Food = NewFoodAtRandom(g.Width, g.Height)
		g.Points += 1
	}

	return true
}
