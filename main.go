package main

import (
	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

type Game struct{}


func (g *Game) Update() error {
	// Update the logical stage
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Render the screen
}

func (g *Game) Layout(
	outsideWidth, outsideHeight, int)
	(screenWidth, screenHeight, int) {
		// Return the game screen size
		return 128, 128
	}
)

func main() {
	ebiten.SetWindowSize(128, 128)
	ebiten.SetWindowTitle("Snake")
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}