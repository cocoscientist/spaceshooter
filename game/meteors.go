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

func (m *Meteors) UpdateAllMeteors(baseSpeed float64) {
	m.meteorSpawnTimer.Update()
	if m.meteorSpawnTimer.IsReady() {
		m.AddRandomMeteor(baseSpeed)
		m.meteorSpawnTimer.Reset()
	}
	for i, meteor := range m.meteors {
		meteor.Update()
		if meteor.position.Y > 600 {
			m.RemoveMeteor(i)
		}
	}
}

func (m *Meteors) DrawAllMeteors(screen *ebiten.Image) {
	for _, meteor := range m.meteors {
		meteor.Draw(screen)
	}
}

func (m *Meteors) AddRandomMeteor(baseSpeed float64) {
	m.meteors = append(m.meteors, NewMeteor(baseSpeed))
}

func (m *Meteors) RemoveMeteor(position int) {
	if position < 0 || position > len(m.meteors) {
		return
	}
	m.meteors = append(m.meteors[:position], m.meteors[position+1:]...)
}
