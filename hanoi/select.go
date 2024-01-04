package hanoi

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type SelectPage struct {
	image *ebiten.Image

	buttons []*Button
	hovered *Button
}

func NewSelectPage() *SelectPage {
	img := ebiten.NewImage(ScreenWidth, ScreenHeight)

	messageImage := ebiten.NewImage(140, 30) // 140はSelect Levelの文字列の幅
	text.Draw(messageImage, "Select Level", mplusNormalFont, 0, messageImage.Bounds().Dy()/2+10, color.White)

	messageOP := &ebiten.DrawImageOptions{}
	messageOP.GeoM.Translate(float64((ScreenWidth-messageImage.Bounds().Dx()))/2, 10)
	img.DrawImage(messageImage, messageOP)

	buttons := []*Button{}
	for i := 1; i <= 9; i++ {
		b := NewButton(strconv.Itoa(i), 0, 0)
		b.x = (ScreenWidth/3 - ButtonWidth) / 2
		b.y = (ScreenHeight/3 - ButtonHeight) / 2
		b.x += (i - 1) % 3 * (ScreenWidth / 3)
		b.y += (i-1)/3*(ScreenHeight/3) + 30*(1-(i-1)/3) // 少し中央よりに配置する
		buttons = append(buttons, b)
	}

	return &SelectPage{
		image:   img,
		buttons: buttons,
	}
}

func (p *SelectPage) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	for _, b := range p.buttons {
		if b == p.hovered {
			b.DrawHover(screen)
		} else {
			b.Draw(screen)
		}
	}
	screen.DrawImage(p.image, op)
}

func (p *SelectPage) Update() {
	x, y := ebiten.CursorPosition()
	for _, b := range p.buttons {
		if b.In(x, y) {
			p.hovered = b
			return
		}
	}
	p.hovered = nil
}
