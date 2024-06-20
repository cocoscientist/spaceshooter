package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemies struct {
	enemies         []*Enemy
	enemySpawnTimer *Timer
}

func NewEnemies() *Enemies {
	e := &Enemies{}
	e.enemySpawnTimer = NewTimer(2500 * time.Millisecond)
	return e

}

func (e *Enemies) UpdateAllEnemies(p *Player, baseSpeed float64) {
	e.enemySpawnTimer.Update()
	if e.enemySpawnTimer.IsReady() {
		e.addRandomEnemy(baseSpeed)
		e.enemySpawnTimer.Reset()
	}
	for i, enemy := range e.enemies {
		enemy.Update(p)
		if enemy.position.Y > 600 {
			e.removeEnemy(i)
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

func (e *Enemies) DrawAllEnemies(screen *ebiten.Image) {
	for _, enemy := range e.enemies {
		enemy.Draw(screen)
	}
}

func (e *Enemies) addRandomEnemy(baseSpeed float64) {
	e.enemies = append(e.enemies, NewEnemy(baseSpeed))
}

func (e *Enemies) removeEnemy(position int) {
	e.enemies = append(e.enemies[:position], e.enemies[position+1:]...)
}
