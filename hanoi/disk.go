package hanoi

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hamao0820/hanoi/hsl"
)

const (
	DiskHeight = 20
)

type Disk struct {
	size  int
	width int
	image *ebiten.Image
}

func NewDisk(size int) *Disk {
	baseWidth := 50
	width := baseWidth * (size)
	image := ebiten.NewImage(width, DiskHeight)
	image.Fill(hsl.NewHSL(rand.Float64(), 1, 0.5).ToRGB())
	return &Disk{
		size:  size,
		width: width,
		image: image,
	}
}

func (d *Disk) Size() int {
	return d.size
}

func (d *Disk) Draw(screen *ebiten.Image, x, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(d.image, op)
}
