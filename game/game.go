package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

type Game struct {
	player *Player
	bg     *Background
}

func NewGame() *Game {
	g := &Game{}
	g.player = NewPlayer()
	g.bg = NewBackground()
	return g
}

func (g *Game) Update() error {
	g.bg.Update()
	g.player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.bg.Draw(screen)
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
