package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Meteors struct {
	meteors          []*Meteor
	meteorSpawnTimer *Timer
}

func NewMeteors(resetTime time.Duration) *Meteors {
	e := &Meteors{}
	e.meteorSpawnTimer = NewTimer(resetTime)
	return e

}

func (e *Meteors) UpdateAllMeteors(baseSpeed float64) {
	e.meteorSpawnTimer.Update()
	if e.meteorSpawnTimer.IsReady() {
		e.AddRandomMeteor(baseSpeed)
		e.meteorSpawnTimer.Reset()
	}
	for i, meteor := range e.meteors {
		meteor.Update()
		if meteor.position.Y > 600 {
			e.RemoveMeteor(i)
		}
	}
}

func (e *Meteors) DrawAllMeteors(screen *ebiten.Image) {
	for _, meteor := range e.meteors {
		meteor.Draw(screen)
	}
}

func (e *Meteors) AddRandomMeteor(baseSpeed float64) {
	e.meteors = append(e.meteors, NewMeteor(baseSpeed))
}

func (e *Meteors) RemoveMeteor(position int) {
	e.meteors = append(e.meteors[:position], e.meteors[position+1:]...)
}
