package hanoi

import "github.com/hajimehoshi/ebiten/v2"

type Game struct{}

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
