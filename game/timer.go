package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Timer struct {
	currentTicks float64
	targetTicks  float64
}

func NewTimer(d time.Duration) *Timer {
	return &Timer{
		currentTicks: 0,
		targetTicks:  float64(d.Milliseconds()) * float64(ebiten.TPS()) / 1000,
	}
}

func (t *Timer) Update() {
	if t.currentTicks < t.targetTicks {
		t.currentTicks++
	}
}

func (t *Timer) IsReady() bool {
	return t.currentTicks >= t.targetTicks
}

func (t *Timer) Reset() {
	t.currentTicks = 0
}
