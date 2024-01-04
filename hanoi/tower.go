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
	image      *ebiten.Image
	hoverImage *ebiten.Image
	x, y       int
}

func NewTower(x, y int, c color.Color) *Tower {
	image := ebiten.NewImage(TowerWidth, TowerHeight)
	image.Fill(c)

	padding := 8
	hoverImage := ebiten.NewImage(TowerWidth+padding*2, TowerHeight+padding*2)
	hoverImage.Fill(color.Black)
	light := ebiten.NewImage(TowerWidth+4, TowerHeight+2)
	light.Fill(color.RGBA{0xff, 0xff, 0x00, 0xff})
	for i := -5; i <= 5; i++ {
		for j := -5; j <= 5; j++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(padding-2+i), float64(padding-2+j))
			op.ColorScale.ScaleAlpha(1 / float32(50))
			hoverImage.DrawImage(light, op)
		}
	}
	return &Tower{
		image:      image,
		hoverImage: hoverImage,
		x:          x,
		y:          y,
	}
}

func (t *Tower) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(t.x), float64(t.y))
	screen.DrawImage(t.image, op)
}

func (t *Tower) DrawHover(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(t.x-3-TowerWidth/2), float64(t.y-8))
	screen.DrawImage(t.hoverImage, op)
}

func (t *Tower) IsAround(x, y int) bool {
	radius := TowerWidth * 5
	return t.x-radius <= x+TowerWidth/2 && x+TowerWidth/2 <= t.x+radius
}
