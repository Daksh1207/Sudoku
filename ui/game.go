package ui

import (
	"fmt"
	"snake/core"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GameUi struct {
	settings       core.Settings
	game           core.Game
	keys           []ebiten.Key
	paused         bool
	lastPauseEvent time.Time
}

func NewGame(settings core.Settings) GameUi {
	return GameUi{
		settings: settings,
		game: core.NewGame(
			settings.Width-settings.SquareSize*3,
			settings.Height-settings.SquareSize*3-settings.TopBarHeight,
			settings.SquareSize,
		),
		keys:   []ebiten.Key{},
		paused: false,
	}
}

func (g *GameUi) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	direction := g.game.Snake.GetDirection()

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

	g.game.Update(direction)

	return nil
}

func (g *GameUi) Draw(screen *ebiten.Image) {
	// Borders
	ebitenutil.DrawRect(
		screen,
		0,
		float64(g.settings.TopBarHeight),
		float64(g.settings.Width),
		float64(g.settings.SquareSize),
		g.settings.SnakeColor,
	)
	ebitenutil.DrawRect(
		screen,
		0,
		float64(g.settings.Height-g.settings.SquareSize),
		float64(g.settings.Width),
		float64(g.settings.SquareSize),
		g.settings.SnakeColor,
	)
	ebitenutil.DrawRect(
		screen,
		0,
		float64(g.settings.TopBarHeight),
		float64(g.settings.SquareSize),
		float64(g.settings.Height),
		g.settings.SnakeColor,
	)
	ebitenutil.DrawRect(
		screen,
		float64(g.settings.Width-g.settings.SquareSize),
		float64(g.settings.TopBarHeight),
		float64(g.settings.SquareSize),
		float64(g.settings.Height),
		g.settings.SnakeColor,
	)

	// Snake
	xOffset := g.settings.SquareSize
	yOffset := g.settings.SquareSize + g.settings.TopBarHeight

	for _, sp := range g.game.Snake.X {
		ebitenutil.DrawRect(
			screen,
			float64(sp.X+xOffset),
			float64(sp.Y+yOffset),
			float64(g.settings.SquareSize),
			float64(g.settings.SquareSize),
			g.settings.SnakeColor,
		)
	}

	// Food
	ebitenutil.DrawRect(
		screen,
		float64(g.game.Food.X.X+xOffset),
		float64(g.game.Food.X.Y+yOffset),
		float64(g.settings.SquareSize),
		float64(g.settings.SquareSize),
		g.settings.FoodColor,
	)

	// Menu
	if g.paused {
		ebitenutil.DebugPrintAt(screen, "Game paused", 120, 5)
	}

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Points %d", g.game.Points), 10, 5)
}

func (g *GameUi) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.settings.Width, g.settings.Height
}
