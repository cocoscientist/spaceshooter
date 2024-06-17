package game

import (
	"spaceshooter/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Laser struct {
	position Vector
	sprite   *ebiten.Image
}

func NewLaser(x float64) *Laser {
	sprite := assets.LaserSprite
	return &Laser{
		position: Vector{
			X: x - float64(sprite.Bounds().Dx())/2,
			Y: 500 - float64(sprite.Bounds().Dy()),
		},
		sprite: sprite,
	}
}

func (l *Laser) Update() {
	l.position.Y -= 250 / float64(ebiten.TPS())
}

func (l *Laser) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(l.position.X, l.position.Y)
	screen.DrawImage(l.sprite, op)
}

func (l *Laser) getWidth() float64 {
	return float64(l.sprite.Bounds().Dx())
}

func (l *Laser) getHeight() float64 {
	return float64(l.sprite.Bounds().Dy())
}
