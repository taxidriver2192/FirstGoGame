package main

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Dir int

type Game struct {
	moveDirection int
	snakeBody     []Position
	timer         int
	moveTime      int
}

// Create constants we can use continuously
const (
	screenWidth  = 500
	screenHeight = 500
	gridSize     = 25
	xNumInScreen = screenWidth / gridSize
	yNumInScreen = screenHeight / gridSize
)

const (
	dirNone = iota
	dirLeft
	dirRight
	dirDown
	dirUp
)

type Position struct {
	X int
	Y int
}

func (g *Game) needsToMoveSnake() bool {
	return g.timer%g.moveTime == 0
}

func (g *Game) Update() error {
	if g.needsToMoveSnake() {
		if inpututil.IsKeyJustPressed(ebiten.KeyD) {
			g.moveDirection = dirLeft
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyA) {
			g.moveDirection = dirRight
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyS) {
			g.moveDirection = dirDown
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyW) {
			g.moveDirection = dirUp
		}
		for i := int64(len(g.snakeBody)) - 1; i > 0; i-- {
			g.snakeBody[i].X = g.snakeBody[i-1].X
			g.snakeBody[i].Y = g.snakeBody[i-1].Y
		}
		switch g.moveDirection {
		case dirLeft:
			g.snakeBody[0].X--
		case dirRight:
			g.snakeBody[0].X++
		case dirDown:
			g.snakeBody[0].Y++
		case dirUp:
			g.snakeBody[0].Y--
		}
	}
	time.Sleep(300)
	g.timer++
	return nil
}

// Dir returns a currently pressed direction.
// Dir returns false if no direction key is pressed.

func (g *Game) Draw(screen *ebiten.Image) {
	for _, v := range g.snakeBody {
		ebitenutil.DrawRect(screen, float64(v.X*gridSize), float64(v.Y*gridSize), gridSize, gridSize, color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	}
}

func newGame() *Game {
	g := &Game{
		snakeBody: make([]Position, 1),
		//moveDirection
	}
	g.snakeBody[0].X = xNumInScreen / 2
	g.snakeBody[0].Y = yNumInScreen / 2
	return g
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	//ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Snake (Ebiten Demo)")
	//game := &Game{}
	if err := ebiten.RunGame(newGame()); err != nil {
		panic(err)
	}
}
