package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct{}

const (
	screenWidth  = 512
	screenHeight = 512
)

func (g *Game) Update() error {
	// Update the logical stage
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Render the screen
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Return the game screen size
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Snake")
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
