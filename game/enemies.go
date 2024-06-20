package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemies struct {
	enemies         []*Enemy
	enemySpawnTimer *Timer
}

func NewEnemies(resetTime time.Duration) *Enemies {
	e := &Enemies{}
	e.enemySpawnTimer = NewTimer(resetTime)
	return e

}

func (e *Enemies) UpdateAllEnemies(p *Player, baseSpeed float64) {
	e.enemySpawnTimer.Update()
	if e.enemySpawnTimer.IsReady() {
		e.AddRandomEnemy(baseSpeed)
		e.enemySpawnTimer.Reset()
	}
	for i, enemy := range e.enemies {
		enemy.Update(p)
		if enemy.position.Y > 600 {
			e.RemoveEnemy(i)
		}
	}
}

func (e *Enemies) CheckCollisionWithPlayer(p *Player) []int {
	var collidingEnemies []int
	for i, enemy := range e.enemies {
		if enemy.CheckCollision(p.position.X, p.position.Y, p.getWidth(), p.getHeight()) {
			collidingEnemies = append(collidingEnemies, i)
		}
	}
	return collidingEnemies
}

func (e *Enemies) CheckBulletHit(l []*Laser) []int {
	var collidingBullets []int
	for i, laser := range l {
		for j, enemy := range e.enemies {
			if enemy.CheckCollision(laser.position.X, laser.position.Y, laser.getWidth(), laser.getHeight()) {
				collidingBullets = append(collidingBullets, i)
				e.RemoveEnemy(j)
			}
		}
	}
	return collidingBullets
}

func (e *Enemies) DrawAllEnemies(screen *ebiten.Image) {
	for _, enemy := range e.enemies {
		enemy.Draw(screen)
	}
}

func (e *Enemies) AddRandomEnemy(baseSpeed float64) {
	e.enemies = append(e.enemies, NewEnemy(baseSpeed))
}

func (e *Enemies) RemoveEnemy(position int) {
	if position < 0 || position > len(e.enemies) {
		return
	}
	e.enemies = append(e.enemies[:position], e.enemies[position+1:]...)
}
