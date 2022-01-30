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
	newKeys        []ebiten.Key
	paused         bool
	lastPauseEvent time.Time
	difficulty     int
	tick           int
}

func NewGame(settings core.Settings) GameUi {
	return GameUi{
		settings: settings,
		game: core.NewGame(
			settings.Width-settings.SquareSize*3,
			settings.Height-settings.SquareSize*3-settings.TopBarHeight,
			settings.SquareSize,
		),
		keys:       []ebiten.Key{},
		paused:     false,
		difficulty: 3,
		tick:       0,
	}
}

func (g *GameUi) pause() {
	if time.Since(g.lastPauseEvent).Milliseconds() > 250 {
		g.lastPauseEvent = time.Now()
		g.paused = !g.paused
	}
}

func (g *GameUi) readLastKey() ebiten.Key {
	g.newKeys = inpututil.AppendPressedKeys(g.newKeys[:0])
	keyPressed := ebiten.Key(-1)

	if len(g.newKeys) == 1 {
		keyPressed = g.newKeys[0]
	} else if len(g.keys) == 1 && len(g.newKeys) >= 1 {
		lastKeyPressed := g.keys[0]

		for _, k := range g.newKeys {
			if k != lastKeyPressed {
				keyPressed = k
			}
		}
	}

	g.keys = g.keys[:0]
	g.keys = append(g.keys, g.newKeys...)

	return keyPressed
}

func (g *GameUi) Update() error {
	direction := g.game.Snake.GetDirection()
	keyPressed := g.readLastKey()

	switch keyPressed {
	case ebiten.KeyArrowRight:
		direction = core.DIRECTION_RIGHT
	case ebiten.KeyD:
		direction = core.DIRECTION_RIGHT
	case ebiten.KeyArrowLeft:
		direction = core.DIRECTION_LEFT
	case ebiten.KeyA:
		direction = core.DIRECTION_LEFT
	case ebiten.KeyArrowUp:
		direction = core.DIRECTION_UP
	case ebiten.KeyW:
		direction = core.DIRECTION_UP
	case ebiten.KeyArrowDown:
		direction = core.DIRECTION_DOWN
	case ebiten.KeyS:
		direction = core.DIRECTION_DOWN
	case ebiten.KeySpace:
		g.pause()
	case ebiten.KeyEscape:
		g.pause()
	}

	if g.paused {
		return nil
	}

	g.game.Snake.SetDirection(direction)

	if g.tick == g.difficulty {
		g.tick = 0
		g.game.Update()
	}

	g.tick += 1

	return nil
}

func (g *GameUi) Draw(screen *ebiten.Image) {
	xOffset := g.settings.SquareSize
	yOffset := g.settings.SquareSize + g.settings.TopBarHeight

	g.drawBorder(screen)
	g.drawSnake(screen, xOffset, yOffset)
	g.drawFood(screen, xOffset, yOffset)

	if g.paused {
		ebitenutil.DebugPrintAt(screen, "Game paused", 120, 5)
	}

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Points %d", g.game.Points), 10, 5)
}

func (g *GameUi) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.settings.Width, g.settings.Height
}

func (g *GameUi) drawBorder(screen *ebiten.Image) {
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
}

func (g *GameUi) drawSnake(screen *ebiten.Image, xOffset int, yOffset int) {
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
}

func (g *GameUi) drawFood(screen *ebiten.Image, xOffset int, yOffset int) {
	ebitenutil.DrawRect(
		screen,
		float64(g.game.Food.X.X+xOffset),
		float64(g.game.Food.X.Y+yOffset),
		float64(g.settings.SquareSize),
		float64(g.settings.SquareSize),
		g.settings.FoodColor,
	)
}
