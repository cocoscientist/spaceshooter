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
	screenWidth         = 800
	screenHeight        = 600
	initialBaseSpeed    = 200
	startingLives       = 3
	meteorSpawnTime     = 1500 * time.Millisecond
	laserCooldownTime   = 1000 * time.Millisecond
	enemySpawnTime      = 3000 * time.Millisecond
	speedupIntervalTime = 5000 * time.Millisecond
)

type Game struct {
	player       *Player
	bg           *Background
	meteors      *Meteors
	lasers       *Lasers
	enemies      *Enemies
	speedupTimer *Timer
	baseSpeed    float64
	score        int
	lives        int
	gameOver     bool
}

func NewGame() *Game {
	g := &Game{}
	g.meteors = NewMeteors(meteorSpawnTime)
	g.lasers = NewLasers(laserCooldownTime)
	g.enemies = NewEnemies(enemySpawnTime)
	g.speedupTimer = NewTimer(speedupIntervalTime)
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
		g.enemies.UpdateAllEnemies(g.player, g.baseSpeed)
		g.meteors.UpdateAllMeteors(g.baseSpeed)
		g.lasers.UpdateAllLasers(g.player)
		g.speedupTimer.Update()
		if g.speedupTimer.IsReady() {
			g.speedupTimer.Reset()
			g.baseSpeed += 10
		}
		for i, laser := range g.lasers.lasers {
			for j, meteor := range g.meteors.meteors {
				if meteor.CheckCollision(laser.position.X, laser.position.Y, laser.getWidth(), laser.getHeight()) {
					g.score++
					g.lasers.RemoveLaser(i)
					g.meteors.RemoveMeteor(j)
					g.lasers.ResetCooldown()
				}
			}
		}
		hittingBullets := g.enemies.CheckBulletHit(g.lasers.lasers)
		if len(hittingBullets) > 0 {
			g.lasers.ResetCooldown()
		}
		g.score += len(hittingBullets)
		for _, index := range hittingBullets {
			g.lasers.RemoveLaser(index)
		}

		// Check for collision between player and meteors
		collidingMeteors := g.getPlayerCollidingMeteors()
		if len(collidingMeteors) > 0 && !g.player.isHit {
			g.player.SetHit()
			g.lives--
			if g.lives == 0 {
				g.gameOver = true
			} else {
				for _, meteorIndex := range collidingMeteors {
					g.meteors.RemoveMeteor(meteorIndex)
				}
			}
		}
		//Check for collision between player and enemies
		collidingEnemies := g.enemies.CheckCollisionWithPlayer(g.player)
		if len(collidingEnemies) > 0 && !g.player.isHit {
			g.player.SetHit()
			g.lives--
			if g.lives == 0 {
				g.gameOver = true
			} else {
				for _, enemyIndex := range collidingEnemies {
					g.enemies.RemoveEnemy(enemyIndex)
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
		g.enemies.DrawAllEnemies(screen)
		for _, meteor := range g.meteors.meteors {
			meteor.Draw(screen)
		}
		for _, laser := range g.lasers.lasers {
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
	g.speedupTimer.Reset()
	g.lives = startingLives
	g.baseSpeed = initialBaseSpeed
	g.meteors = NewMeteors(meteorSpawnTime)
	g.lasers = NewLasers(laserCooldownTime)
	g.enemies = NewEnemies(enemySpawnTime)
}

func (g *Game) getPlayerCollidingMeteors() []int {
	var collidingMeteors []int
	for i, meteor := range g.meteors.meteors {
		if meteor.CheckCollision(g.player.position.X, g.player.position.Y, g.player.getWidth(), g.player.getHeight()) {
			collidingMeteors = append(collidingMeteors, i)
		}
	}
	return collidingMeteors
}
