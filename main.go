package main

import (
	"image/color"
	"log"
	"math"
	"snake/game"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	snake          game.Snake
	food           game.Food
	keys           []ebiten.Key
	paused         bool
	lastPauseEvent time.Time
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	direction := g.snake.GetDirection()

	if len(g.keys) == 1 {
		switch g.keys[0] {
		case ebiten.KeyArrowRight:
			direction = game.DIRECTION_RIGHT
		case ebiten.KeyArrowLeft:
			direction = game.DIRECTION_LEFT
		case ebiten.KeyArrowUp:
			direction = game.DIRECTION_UP
		case ebiten.KeyArrowDown:
			direction = game.DIRECTION_DOWN
		case ebiten.KeySpace:
			if time.Since(g.lastPauseEvent).Milliseconds() > 500 {
				g.lastPauseEvent = time.Now()
				g.paused = !g.paused
			}
		}
	}

	if g.paused {
		return nil
	}

	g.snake.SetDirection(direction)
	g.snake.Move()

	// TODO improve collision system
	snakeHead := g.snake.Head()
	xDiff := float64(snakeHead.X - g.food.X.X)
	yDiff := float64(snakeHead.Y - g.food.X.Y)
	dist := math.Sqrt(xDiff*xDiff + yDiff*yDiff)

	if dist < 5 {
		g.snake.Grow()
		g.food = game.NewFoodAtRandom()
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
		float64(g.food.X.X),
		float64(g.food.X.Y),
		10,
		10,
		color.RGBA{R: 255, G: 0, B: 0, A: 255},
	)

	if g.paused {
		ebitenutil.DebugPrintAt(screen, "Game paused", 50, 50)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	snakeGame := Game{
		snake:  game.NewSnake(),
		food:   game.NewFoodAtRandom(),
		keys:   []ebiten.Key{},
		paused: false,
	}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Snake")

	if err := ebiten.RunGame(&snakeGame); err != nil {
		log.Fatal(err)
	}
}
