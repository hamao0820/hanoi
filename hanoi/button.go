package hanoi

import (
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Level int

const (
	Level1 Level = 1
	Level2 Level = 2
	Level3 Level = 3
	Level4 Level = 4
	Level5 Level = 5
	Level6 Level = 6
	Level7 Level = 7
	Level8 Level = 8
	Level9 Level = 9
)

func (l Level) String() string {
	return strconv.Itoa(int(l))
}

func (l Level) Int() int {
	return int(l)
}

type LevelSelectButton struct {
	img   *ebiten.Image
	level Level
	x, y  int
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

func NewButton(level Level, x, y int) *LevelSelectButton {
	img := ebiten.NewImage(ButtonWidth, ButtonHeight)
	img.Fill(color.White)
	textImage := ebiten.NewImage(20, 20)
	text.Draw(textImage, level.String(), mplusNormalFont, textImage.Bounds().Dx()/2-5, textImage.Bounds().Dy()/2+10, color.Black)
	textImageX := (ButtonWidth - textImage.Bounds().Dx()) / 2
	textImageY := (ButtonHeight - textImage.Bounds().Dy()) / 2
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(textImageX), float64(textImageY))
	img.DrawImage(textImage, op)
	return &LevelSelectButton{
		img:   img,
		level: level,
		x:     x,
		y:     y,
	}
}

func (b *LevelSelectButton) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.x), float64(b.y))
	screen.DrawImage(b.img, op)
}

func (b *LevelSelectButton) DrawHover(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.x+1), float64(b.y-1))
	op.ColorScale.Scale(1, 1, 0.8, 1)
	screen.DrawImage(b.img, op)
}

func (b *LevelSelectButton) In(x, y int) bool {
	return b.x <= x && x < b.x+ButtonWidth && b.y <= y && y < b.y+ButtonHeight
}
