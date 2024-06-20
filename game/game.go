package game

import (
	"fmt"
	"image/color"
	"spaceshooter/assets"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	screenWidth      = 800
	screenHeight     = 600
	initialBaseSpeed = 200
	startingLives    = 3
)

type Game struct {
	player           *Player
	bg               *Background
	meteorSpawnTimer *Timer
	meteors          []*Meteor
	lasers           []*Laser
	laserTimer       *Timer
	speedupTimer     *Timer
	baseSpeed        float64
	score            int
	lives            int
	gameOver         bool
}

func NewGame() *Game {
	g := &Game{}
	g.meteorSpawnTimer = NewTimer(1250 * time.Millisecond)
	g.laserTimer = NewTimer(1000 * time.Millisecond)
	g.speedupTimer = NewTimer(3500 * time.Millisecond)
	g.player = NewPlayer()
	g.bg = NewBackground()
	g.baseSpeed = initialBaseSpeed
	g.lives = startingLives
	g.gameOver = false
	return g
}

func (g *Game) Update() error {
	g.bg.Update()
	if g.gameOver {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.resetGame()
		}
	} else {
		g.player.Update()
		g.meteorSpawnTimer.Update()
		if g.meteorSpawnTimer.IsReady() {
			g.meteorSpawnTimer.Reset()
			m := NewMeteor(g.baseSpeed)
			g.meteors = append(g.meteors, m)
		}
		g.laserTimer.Update()
		if g.laserTimer.IsReady() && ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.laserTimer.Reset()
			l := NewLaser(g.player.position.X + float64(g.player.sprite.Bounds().Dx())/2)
			g.lasers = append(g.lasers, l)
		}
		g.speedupTimer.Update()
		if g.speedupTimer.IsReady() {
			g.speedupTimer.Reset()
			g.baseSpeed += 10
		}
		for i, meteor := range g.meteors {
			meteor.Update()
			if meteor.position.Y > 600 {
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
			}
		}
		for i, laser := range g.lasers {
			laser.Update()
			if laser.position.Y < 0 {
				g.lasers = append(g.lasers[:i], g.lasers[i+1:]...)
			}
		}
		for i, laser := range g.lasers {
			for j, meteor := range g.meteors {
				if meteor.CheckCollision(laser.position.X, laser.position.Y, laser.getWidth(), laser.getHeight()) {
					g.score++
					g.meteors = append(g.meteors[:j], g.meteors[j+1:]...)
					g.lasers = append(g.lasers[:i], g.lasers[i+1:]...)
				}
			}
		}
		// Check for collision between player and meteors
		collidingMeteors := g.getPlayerCollidingMeteors()
		if len(collidingMeteors) > 0 {
			g.lives--
			if g.lives == 0 {
				g.gameOver = true
			} else {
				for _, meteorIndex := range collidingMeteors {
					g.meteors = append(g.meteors[:meteorIndex], g.meteors[meteorIndex+1:]...)
				}
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.bg.Draw(screen)
	if !g.gameOver {
		g.player.Draw(screen)
		for _, meteor := range g.meteors {
			meteor.Draw(screen)
		}
		for _, laser := range g.lasers {
			laser.Draw(screen)
		}
		text.Draw(screen, fmt.Sprintf("SCORE:%04d", g.score), assets.ScoreFont, 20, 60, color.White)
		text.Draw(screen, fmt.Sprintf("LIVES:%02d", g.lives), assets.ScoreFont, screenWidth-186, 60, color.White)
	} else {
		text.Draw(screen, fmt.Sprintf("FINAL SCORE: %04d", g.score), assets.GameOverFont, screenWidth/2-200, 90, color.White)
		text.Draw(screen, fmt.Sprintf("GAME OVER!"), assets.GameOverFont, screenWidth/2-114, screenHeight/2-100, color.White)
		text.Draw(screen, fmt.Sprintf("Press Space to Play Again!"), assets.GameOverFont, screenWidth/2-302, screenHeight/2, color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) resetGame() {
	g.gameOver = false
	g.score = 0
	g.player.Reset()
	g.meteorSpawnTimer.Reset()
	g.laserTimer.Reset()
	g.speedupTimer.Reset()
	g.lives = startingLives
	g.baseSpeed = initialBaseSpeed
	g.lasers = make([]*Laser, 0)
	g.meteors = make([]*Meteor, 0)
}

func (g *Game) getPlayerCollidingMeteors() []int {
	var collidingMeteors []int
	for i, meteor := range g.meteors {
		if meteor.CheckCollision(g.player.position.X, g.player.position.Y, g.player.getWidth(), g.player.getHeight()) {
			collidingMeteors = append(collidingMeteors, i)
		}
	}
	return collidingMeteors
}
