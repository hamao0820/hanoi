package hanoi

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type SelectPage struct {
	image *ebiten.Image

	buttons []*Button
}

func NewSelectPage() *SelectPage {
	img := ebiten.NewImage(ScreenWidth, ScreenHeight)

	messageImage := ebiten.NewImage(140, 30) // 140はSelect Levelの文字列の幅
	text.Draw(messageImage, "Select Level", mplusNormalFont, 0, messageImage.Bounds().Dy()/2+10, color.White)

	messageOP := &ebiten.DrawImageOptions{}
	messageOP.GeoM.Translate(float64((ScreenWidth-messageImage.Bounds().Dx()))/2, 10)
	img.DrawImage(messageImage, messageOP)

	buttons := []*Button{}

	return &SelectPage{
		image:   img,
		buttons: buttons,
	}
}

func (p *SelectPage) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	for _, b := range p.buttons {
		op.GeoM.Translate(float64(b.x), float64(b.y))
		screen.DrawImage(b.img, op)
	}
	screen.DrawImage(p.image, op)
}
