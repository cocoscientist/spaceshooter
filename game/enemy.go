package game

import (
	"math"
	"math/rand"
	"spaceshooter/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	BASE_ACCELERATION = 200
)

type Enemy struct {
	sprite      *ebiten.Image
	position    *Vector
	velocity    *Vector
	oldRotation float64
	newRotation float64
}

func NewEnemy(baseSpeed float64) *Enemy {
	sprite := assets.EnemySprites[rand.Intn(len(assets.MeteorSprites))]
	position := &Vector{
		X: rand.Float64() * (800 - float64(sprite.Bounds().Dx())),
		Y: -1 * float64(sprite.Bounds().Dy()),
	}
	velocity := &Vector{
		X: 0,
		Y: baseSpeed,
	}

	return &Enemy{
		sprite:      sprite,
		position:    position,
		velocity:    velocity,
		oldRotation: velocity.getAngle(),
		newRotation: velocity.getAngle(),
	}
}

func (e *Enemy) Update(p *Player) {
	acceleration := e.getAcceleration(p)
	e.velocity.X += getValuePerFrame(acceleration.X)
	newVelocityAngle := e.velocity.getAngle()
	e.oldRotation = e.newRotation
	e.newRotation = newVelocityAngle

	e.position.X += getValuePerFrame(e.velocity.X)
	e.position.Y += getValuePerFrame(e.velocity.Y)
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	bounds := e.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate((e.newRotation - e.oldRotation) * float64(ebiten.TPS()))
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(e.position.X, e.position.Y)

	screen.DrawImage(e.sprite, op)
}

func (m *Enemy) CheckCollision(x float64, y float64, w float64, h float64) bool {
	return m.position.X < x+w && m.position.X+float64(m.sprite.Bounds().Dx()) > x && m.position.Y < y+h && m.position.Y+float64(m.sprite.Bounds().Dy()) > y
}

func (e *Enemy) getAcceleration(p *Player) *Vector {
	theta := getAngle(e.position, p.position)
	return &Vector{
		X: BASE_ACCELERATION * math.Cos(theta),
		Y: BASE_ACCELERATION * math.Sin(theta),
	}
}

/*
* 	    dx
*    V1----
*	 |\ ) = Angle
*  dy| \
*    |  \
*    |	 \
*	     V2
 */
func getAngle(vector1 *Vector, vector2 *Vector) float64 {
	delX := vector2.X - vector1.X
	delY := vector2.Y - vector1.Y

	return math.Acos(delX / math.Hypot(delX, delY))
}

func getValuePerFrame(val float64) float64 {
	return val / float64(ebiten.TPS())
}
