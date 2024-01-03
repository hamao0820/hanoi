package hanoi

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	StandHeight = 30
)

type Stand struct {
	img *ebiten.Image
}

func NewStand() *Stand {
	img := ebiten.NewImage(ScreenWidth, StandHeight)
	img.Fill(color.White)

	return &Stand{
		img: img,
	}
}

func (s *Stand) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, float64(ScreenHeight-StandHeight))
	screen.DrawImage(s.img, op)
}
