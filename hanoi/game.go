package hanoi

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	honoi    [3]*Hanoi
	stand    *Stand
	selected *Hanoi
	hovered  *Hanoi
	count    int
}

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

func NewGame() *Game {
	h1 := NewHanoi(0)
	h1.disks = append(h1.disks, NewDisk(3))
	h1.disks = append(h1.disks, NewDisk(2))
	h1.disks = append(h1.disks, NewDisk(1))
	return &Game{
		honoi: [3]*Hanoi{
			h1,
			NewHanoi(1),
			NewHanoi(2),
		},
		stand: NewStand(),
	}
}

// TODO: レベル選択画面を作る
func (g *Game) Update() error {
	x, y := ebiten.CursorPosition()
	for _, h := range g.honoi {
		if h.tower.IsAround(x, y) {
			g.hovered = h
			break
		}
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		for _, h := range g.honoi {
			if h.tower.IsAround(x, y) {
				switch g.selected {
				case nil:
					if !h.IsEmpty(x, y) {
						g.selected = h
					}
				case h:
					g.selected = nil
				default:
					if h.CanPush(g.selected.Top()) {
						h.Push(g.selected.Pop())
						g.selected = nil
						g.count++
					}
				}
				break
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, h := range g.honoi {
		h.Draw(screen, h == g.selected, h == g.hovered)
	}

	if g.selected != nil {
		x, y := ebiten.CursorPosition()
		g.selected.Top().Draw(screen, x-g.selected.Top().width/2, y-DiskHeight/2, 1)
	}

	g.stand.Draw(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("count: %d", g.count))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
