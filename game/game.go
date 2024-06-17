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
	lasers           []*Laser
	laserTimer       *Timer
}

func NewGame() *Game {
	g := &Game{}
	g.meteorSpawnTimer = NewTimer(2 * time.Second)
	g.laserTimer = NewTimer(1 * time.Second)
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
	g.laserTimer.Update()
	if g.laserTimer.IsReady() && ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.laserTimer.Reset()
		l := NewLaser(g.player.position.X + float64(g.player.sprite.Bounds().Dx())/2)
		g.lasers = append(g.lasers, l)
	}
	for i, meteor := range g.meteors {
		meteor.Update()
		if meteor.position.Y > 480 {
			g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
		}
	}
	for i, laser := range g.lasers {
		laser.Update()
		if laser.position.Y < 0 {
			g.lasers = append(g.lasers[:i], g.lasers[i+1:]...)
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
	for _, laser := range g.lasers {
		laser.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
