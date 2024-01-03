package hanoi

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	honoi [3]*Hanoi
	stand *Stand
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

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		var d *Disk
		var c *Hanoi
		for _, h := range g.honoi {
			if h.tower.IsAround(x, y) {
				if !h.IsEmpty(x, y) {
					d = h.Top()
					c = h
					break
				}
			}
		}

		if d == nil {
			return nil
		}

		var n *Hanoi
		for _, h := range g.honoi {
			if h.CanPush(d) {
				n = h
				break
			}
		}

		if n == nil {
			return nil
		}

		c.Pop()
		n.Push(d)

	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.stand.Draw(screen)
	for _, h := range g.honoi {
		h.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
