package game

import "github.com/hajimehoshi/ebiten/v2"

const (
	screenWidth  = 800
	screenHeight = 600
)

type Game struct{}

func NewGame() *Game {
	g := &Game{}
	return g
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
