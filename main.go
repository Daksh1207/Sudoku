package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"snake/game"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	snake     game.Snake
	food      game.Point
	keys      []ebiten.Key
	direction string
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	if len(g.keys) == 1 {
		switch g.keys[0] {
		case ebiten.KeyArrowRight:
			g.direction = "R"
		case ebiten.KeyArrowLeft:
			g.direction = "L"
		case ebiten.KeyArrowUp:
			g.direction = "U"
		case ebiten.KeyArrowDown:
			g.direction = "D"
		}
	}

	g.snake.Move(g.direction)

	snakeHead := g.snake.Head()
	xDiff := float64(snakeHead.X - g.food.X)
	yDiff := float64(snakeHead.Y - g.food.Y)
	dist := math.Sqrt(xDiff*xDiff + yDiff*yDiff)

	if dist < 5 {
		g.snake.Grow()
		g.food.X = rand.Intn(320)
		g.food.Y = rand.Intn(240)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, sp := range g.snake.X {
		ebitenutil.DrawRect(
			screen,
			float64(sp.X),
			float64(sp.Y),
			10,
			10,
			color.RGBA{R: 255, G: 255, B: 255, A: 255},
		)
	}
	ebitenutil.DrawRect(
		screen,
		float64(g.food.X),
		float64(g.food.Y),
		10,
		10,
		color.RGBA{R: 255, G: 0, B: 0, A: 255},
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	snakeGame := Game{
		snake: game.Snake{
			X: []game.Point{{X: 0, Y: 0}},
		},
		food:      game.Point{X: 50, Y: 50},
		keys:      []ebiten.Key{},
		direction: "R",
	}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Snake")

	if err := ebiten.RunGame(&snakeGame); err != nil {
		log.Fatal(err)
	}
}
