package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Lasers struct {
	lasers             []*Laser
	laserCooldownTimer *Timer
}

func NewLasers(resetTime time.Duration) *Lasers {
	e := &Lasers{}
	e.laserCooldownTimer = NewTimer(resetTime)
	return e

}

func (e *Lasers) UpdateAllLasers(player *Player) {
	e.laserCooldownTimer.Update()
	if e.laserCooldownTimer.IsReady() && ebiten.IsKeyPressed(ebiten.KeySpace) {
		e.laserCooldownTimer.Reset()
		e.AddRandomLaser(player)
	}
	for i, laser := range e.lasers {
		laser.Update()
		if laser.position.Y < 0 {
			e.RemoveLaser(i)
		}
	}
}

func (e *Lasers) DrawAllLasers(screen *ebiten.Image) {
	for _, laser := range e.lasers {
		laser.Draw(screen)
	}
}

func (e *Lasers) AddRandomLaser(player *Player) {
	e.lasers = append(e.lasers, NewLaser(player.position.X+float64(player.sprite.Bounds().Dx())/2))
}

func (e *Lasers) RemoveLaser(position int) {
	e.lasers = append(e.lasers[:position], e.lasers[position+1:]...)
}
