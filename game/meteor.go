package game

import (
	"math/rand"
	"spaceshooter/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	rotationSpeedMin = -2
	rotationSpeedMax = 2
)

type Meteor struct {
	position      Vector
	rotation      float64
	velocity      float64
	rotationSpeed float64
	sprite        *ebiten.Image
}

func NewMeteor(baseSpeed float64) *Meteor {
	sprite := assets.MeteorSprites[rand.Intn(len(assets.MeteorSprites))]
	position := Vector{
		X: rand.Float64() * (800 - float64(sprite.Bounds().Dx())),
		Y: -1 * float64(sprite.Bounds().Dy()),
	}
	velocity := baseSpeed + rand.Float64()*50
	rotationSpeed := rotationSpeedMin + rand.Float64()*(rotationSpeedMax-rotationSpeedMin)
	return &Meteor{
		sprite:        sprite,
		position:      position,
		velocity:      velocity, 
		rotationSpeed: rotationSpeed,
	}
}

func (m *Meteor) Update() {
	m.position.Y += m.velocity / float64(ebiten.TPS())
	m.rotation += m.rotationSpeed / float64(ebiten.TPS())
}

func (m *Meteor) Draw(screen *ebiten.Image) {
	bounds := m.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(m.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(m.position.X, m.position.Y)

	screen.DrawImage(m.sprite, op)
}

func (m *Meteor) CheckCollision(x float64, y float64, w float64, h float64) bool {
	return m.position.X < x+w && m.position.X+float64(m.sprite.Bounds().Dx()) > x && m.position.Y < y+h && m.position.Y+float64(m.sprite.Bounds().Dy()) > y
}
