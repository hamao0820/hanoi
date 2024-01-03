package hanoi

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hamao0820/hanoi/hsl"
)

const (
	DiskHeight = 20
)

type Disk struct {
	size       int
	width      int
	image      *ebiten.Image
	alphaImage *image.Alpha
}

func NewDisk(size int) *Disk {
	baseWidth := 50
	width := baseWidth * (size)
	img := ebiten.NewImage(width, DiskHeight)
	img.Fill(hsl.NewHSL(rand.Float64(), 1, 0.5).ToRGB())
	alphaImage := image.NewAlpha(image.Rect(0, 0, width, DiskHeight))
	for i := 0; i < width; i++ {
		for j := 0; j < DiskHeight; j++ {
			alphaImage.SetAlpha(i, j, color.Alpha{0xff})
		}
	}

	return &Disk{
		size:       size,
		width:      width,
		image:      img,
		alphaImage: alphaImage,
	}
}

func (d *Disk) Size() int {
	return d.size
}

func (d *Disk) Draw(screen *ebiten.Image, x, y int, alpha float32) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorScale.ScaleAlpha(alpha)
	screen.DrawImage(d.image, op)
}
