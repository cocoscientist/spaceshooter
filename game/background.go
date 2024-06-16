package game

import (
	"spaceshooter/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Viewport struct {
	x int
	y int
}

func (p *Viewport) Move(y int) {
	p.y -= y * 10 / ebiten.TPS()
	p.y %= (16 * y)
}

func (p *Viewport) Position() (int, int) {
	return p.x, p.y
}

type Background struct {
	img      *ebiten.Image
	viewport Viewport
}

func NewBackground() *Background {
	img := assets.BackGroundSprite
	return &Background{
		img:      img,
		viewport: Viewport{},
	}
}

func (bg *Background) Update() {
	s := bg.img.Bounds().Size()
	bg.viewport.Move(s.Y)
}

func (bg *Background) Draw(screen *ebiten.Image) {
	x, y := bg.viewport.Position()
	offsetX, offsetY := float64(-x)/16, float64(-y)/16
	const repeat = 5
	w, h := bg.img.Bounds().Dx(), bg.img.Bounds().Dy()
	for j := 0; j < repeat; j++ {
		for i := 0; i < repeat; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(800-float64(w*i), 600-float64(h*j))
			op.GeoM.Translate(offsetX, offsetY)
			screen.DrawImage(bg.img, op)
		}
	}
}
