package hanoi

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	TowerWidth  = 10
	TowerHeight = 300
)

type Tower struct {
	image *ebiten.Image
	x, y  int
}

func NewTower(x, y int) *Tower {
	image := ebiten.NewImage(TowerWidth, TowerHeight)
	image.Fill(color.White)

	return &Tower{
		image: image,
		x:     x,
		y:     y,
	}
}

func (t *Tower) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(t.x), float64(t.y))
	screen.DrawImage(t.image, op)
}

func (t *Tower) IsAround(x, y int) bool {
	radius := TowerWidth * 5
	return t.x-radius <= x+TowerWidth/2 && x+TowerWidth/2 <= t.x+radius
}
