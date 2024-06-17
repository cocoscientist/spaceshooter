package game

import (
	"spaceshooter/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	speedPerSecond = 350
)

type Player struct {
	position Vector
	sprite   *ebiten.Image
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
	}
}

func (p *Player) Update() {
	speed := speedPerSecond / float64(ebiten.TPS())
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.X -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.X += speed
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.sprite, op)
}

func (p *Player) getWidth() float64 {
	return float64(p.sprite.Bounds().Dx())
}

func (p *Player) getHeight() float64 {
	return float64(p.sprite.Bounds().Dy())
}
