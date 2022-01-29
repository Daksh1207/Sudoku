package main

import (
	"log"
	"snake/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := ui.NewGame()

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Snake")

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
