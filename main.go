package main

import (
	"image/color"
	"log"
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

	dist := g.snake.X[0].X - g.food.X + g.snake.X[0].Y - g.food.Y - 5 - 5
	if dist < 0 {
		dist = -dist
	}

	if dist < 5 {
		g.snake.Grow()
		//g.food.X += 100
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	println(len(g.snake.X))
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
