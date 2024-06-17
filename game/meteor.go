package game

import (
	"math"
	"math/rand"
	"spaceshooter/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	rotationSpeedMin = -2
	rotationSpeedMax = 2
	baseVelocity     = 200
)

type Meteor struct {
	position      Vector
	rotation      float64
	velocity      float64
	rotationSpeed float64
	sprite        *ebiten.Image
}

func NewMeteor() *Meteor {
	sprite := assets.MeteorSprites[rand.Intn(len(assets.MeteorSprites))]
	position := Vector{
		X: rand.Float64() * (800 - float64(sprite.Bounds().Dx())),
		Y: -1 * float64(sprite.Bounds().Dy()),
	}
	velocity := baseVelocity + rand.Float64()*50
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
	radius := math.Max(float64(m.sprite.Bounds().Dx())/2, float64(m.sprite.Bounds().Dy())/2)
	distanceX := math.Abs(m.position.X - x - (w / 2))
	distanceY := math.Abs(m.position.Y - y - (h / 2))
	if distanceX > (radius + w/2) {
		return false
	}
	if distanceY > (radius + h/2) {
		return false
	}
	if distanceX < w/2 {
		return true
	}
	if distanceY < h/2 {
		return true
	}

	squaredCornered := math.Pow(m.position.X-(w/2), 2) + math.Pow(m.position.Y-(h/2), 2)
	return squaredCornered <= math.Pow(radius, 2)
}
