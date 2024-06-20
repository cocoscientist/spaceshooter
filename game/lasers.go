package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Lasers struct {
	lasers             []*Laser
	laserCooldownTimer *Timer
}

func (l *Lasers) ResetCooldown() {
	l.laserCooldownTimer.currentTicks = l.laserCooldownTimer.targetTicks
}

func NewLasers(resetTime time.Duration) *Lasers {
	e := &Lasers{}
	e.laserCooldownTimer = NewTimer(resetTime)
	return e

}

func (l *Lasers) UpdateAllLasers(player *Player) {
	l.laserCooldownTimer.Update()
	if l.laserCooldownTimer.IsReady() && ebiten.IsKeyPressed(ebiten.KeySpace) {
		l.laserCooldownTimer.Reset()
		l.AddRandomLaser(player)
	}
	for i, laser := range l.lasers {
		laser.Update()
		if laser.position.Y < 0 {
			l.RemoveLaser(i)
		}
	}
}

func (l *Lasers) DrawAllLasers(screen *ebiten.Image) {
	for _, laser := range l.lasers {
		laser.Draw(screen)
	}
}

func (l *Lasers) AddRandomLaser(player *Player) {
	l.lasers = append(l.lasers, NewLaser(player.position.X+float64(player.sprite.Bounds().Dx())/2))
}

func (l *Lasers) RemoveLaser(position int) {
	if position < 0 || position > len(l.lasers) {
		return
	}
	l.lasers = append(l.lasers[:position], l.lasers[position+1:]...)
}
