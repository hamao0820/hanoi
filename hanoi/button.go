package hanoi

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Button struct {
	img  *ebiten.Image
	text string
	x, y int
}

const (
	ButtonWidth  = 100
	ButtonHeight = 50
)

var (
	mplusNormalFont font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func NewButton(t string, x, y int) *Button {
	img := ebiten.NewImage(ButtonWidth, ButtonHeight)
	img.Fill(color.White)
	textImage := ebiten.NewImage(20, 20)
	text.Draw(textImage, t, mplusNormalFont, textImage.Bounds().Dx()/2-5, textImage.Bounds().Dy()/2+10, color.Black)
	textImageX := (ButtonWidth - textImage.Bounds().Dx()) / 2
	textImageY := (ButtonHeight - textImage.Bounds().Dy()) / 2
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(textImageX), float64(textImageY))
	img.DrawImage(textImage, op)
	return &Button{
		img:  img,
		text: t,
		x:    x,
		y:    y,
	}
}

func (b *Button) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.x), float64(b.y))
	screen.DrawImage(b.img, op)
}

func (b *Button) DrawHover(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.x+1), float64(b.y-1))
	op.ColorScale.Scale(1, 1, 0.8, 1)
	screen.DrawImage(b.img, op)
}

func (b *Button) In(x, y int) bool {
	return b.x <= x && x < b.x+ButtonWidth && b.y <= y && y < b.y+ButtonHeight
}
