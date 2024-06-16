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
	meteor *Meteor
}

func NewGame() *Game {
	g := &Game{}
	g.player = NewPlayer()
	g.bg = NewBackground()
	g.meteor = NewMeteor()
	return g
}

func (g *Game) Update() error {
	g.bg.Update()
	g.player.Update()
	g.meteor.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.bg.Draw(screen)
	g.player.Draw(screen)
	g.meteor.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
