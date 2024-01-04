package hanoi

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type SelectPage struct {
	image *ebiten.Image

	buttons []*LevelSelectButton
	hovered *LevelSelectButton
}

func NewSelectPage() *SelectPage {
	img := ebiten.NewImage(ScreenWidth, ScreenHeight)

	messageImage := ebiten.NewImage(140, 30) // 140はSelect Levelの文字列の幅
	text.Draw(messageImage, "Select Level", mplusNormalFont, 0, messageImage.Bounds().Dy()/2+10, color.White)

	messageOP := &ebiten.DrawImageOptions{}
	messageOP.GeoM.Translate(float64((ScreenWidth-messageImage.Bounds().Dx()))/2, 10)
	img.DrawImage(messageImage, messageOP)

	buttons := []*LevelSelectButton{}
	for i := 1; i <= 9; i++ {
		var l Level
		switch i {
		case 1:
			l = Level1
		case 2:
			l = Level2
		case 3:
			l = Level3
		case 4:
			l = Level4
		case 5:
			l = Level5
		case 6:
			l = Level6
		case 7:
			l = Level7
		case 8:
			l = Level8
		case 9:
			l = Level9
		}
		b := NewButton(l, 0, 0)
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
