package ui

import (
	"math"
	"snake/core"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	settings       core.Settings
	snake          core.Snake
	food           core.Food
	keys           []ebiten.Key
	paused         bool
	lastPauseEvent time.Time
}

func NewGame() Game {
	return Game{
		settings: core.NewSettings(),
		snake:    core.NewSnake(),
		food:     core.NewFoodAtRandom(),
		keys:     []ebiten.Key{},
		paused:   false,
	}
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	direction := g.snake.GetDirection()

	if len(g.keys) == 1 {
		switch g.keys[0] {
		case ebiten.KeyArrowRight:
			direction = core.DIRECTION_RIGHT
		case ebiten.KeyArrowLeft:
			direction = core.DIRECTION_LEFT
		case ebiten.KeyArrowUp:
			direction = core.DIRECTION_UP
		case ebiten.KeyArrowDown:
			direction = core.DIRECTION_DOWN
		case ebiten.KeySpace:
			if time.Since(g.lastPauseEvent).Milliseconds() > 250 {
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

	snakeHead := g.snake.Head()
	radius := g.settings.SquareSize / 2.0
	xDiff := float64(snakeHead.X + radius - (g.food.X.X + radius))
	yDiff := float64(snakeHead.Y + radius - (g.food.X.Y + radius))
	dist := math.Sqrt(xDiff*xDiff + yDiff*yDiff)

	if dist < float64(g.settings.SquareSize) {
		g.snake.Grow()
		g.food = core.NewFoodAtRandom()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, sp := range g.snake.X {
		ebitenutil.DrawRect(
			screen,
			float64(sp.X),
			float64(sp.Y),
			float64(g.settings.SquareSize),
			float64(g.settings.SquareSize),
			g.settings.SnakeColor,
		)
	}

	ebitenutil.DrawRect(
		screen,
		float64(g.food.X.X),
		float64(g.food.X.Y),
		float64(g.settings.SquareSize),
		float64(g.settings.SquareSize),
		g.settings.FoodColor,
	)

	if g.paused {
		ebitenutil.DebugPrintAt(screen, "Game paused", 50, 50)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
