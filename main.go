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
	Snake     game.Snake
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

	g.Snake.Move(g.direction)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(
		screen,
		float64(g.Snake.X[0].X),
		float64(g.Snake.X[0].Y),
		10,
		10,
		color.RGBA{R: 255, G: 255, B: 255, A: 255},
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	snakeGame := Game{
		Snake: game.Snake{
			X: []game.Point{{X: 0, Y: 0}},
		},
		keys:      []ebiten.Key{},
		direction: "R",
	}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(&snakeGame); err != nil {
		log.Fatal(err)
	}
}
