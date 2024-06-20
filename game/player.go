package game

import (
	"math"
	"spaceshooter/assets"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
)

const (
	speedPerSecond = 350
)

type Player struct {
	position Vector
	sprite   *ebiten.Image
	hitTimer *Timer
	isHit    bool
}

func NewPlayer() *Player {
	sprite := assets.PlayerSprite
	position := Vector{
		X: 350,
		Y: 500,
	}
	return &Player{
		sprite:   sprite,
		position: position,
		isHit:    false,
	}
}

func (p *Player) Update() {
	deltaXPerTick := speedPerSecond / float64(ebiten.TPS())
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.X -= deltaXPerTick
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.X += deltaXPerTick
	}

	if p.isHit {
		if p.hitTimer == nil {
			p.hitTimer = NewTimer(2 * time.Second)
		} else if p.hitTimer.IsReady() {
			p.isHit = false
			p.hitTimer = nil
		} else {
			p.hitTimer.Update()
		}
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	cp := &colorm.DrawImageOptions{}
	cp.GeoM.Translate(p.position.X, p.position.Y)
	var cm colorm.ColorM
	var alpha float64 = 1
	if p.hitTimer != nil {
		alpha = 0.5*math.Sin(p.hitTimer.currentTicks) + 0.5
	}
	cm.Scale(1, 1, 1, alpha)
	colorm.DrawImage(screen, p.sprite, cm, cp)
}

func (p *Player) getWidth() float64 {
	return float64(p.sprite.Bounds().Dx())
}

func (p *Player) getHeight() float64 {
	return float64(p.sprite.Bounds().Dy())
}

func (p *Player) Reset() {
	print(p.position.X, p.position.Y)
	p.position.X = 350
	p.position.Y = 500
	print(p.position.X, p.position.Y)
	p.hitTimer = nil
	p.isHit = false
}

func (p *Player) SetHit() {
	p.isHit = true
}
