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
	img *ebiten.Image
}

func NewTower() *Tower {
	img := ebiten.NewImage(TowerWidth, TowerHeight)
	img.Fill(color.White)

	return &Tower{
		img: img,
	}
}

func (p *Tower) Draw(screen *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(p.img, op)
}
