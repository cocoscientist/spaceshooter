package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

type Game struct {
	player           *Player
	bg               *Background
	meteorSpawnTimer *Timer
	meteors          []*Meteor
}

func NewGame() *Game {
	g := &Game{}
	g.meteorSpawnTimer = NewTimer(2 * time.Second)
	g.player = NewPlayer()
	g.bg = NewBackground()
	return g
}

func (g *Game) Update() error {
	g.bg.Update()
	g.player.Update()
	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()
		m := NewMeteor()
		g.meteors = append(g.meteors, m)
	}
	for i, meteor := range g.meteors {
		meteor.Update()
		if meteor.position.Y > 480 {
			g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.bg.Draw(screen)
	g.player.Draw(screen)
	for _, meteor := range g.meteors {
		meteor.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
