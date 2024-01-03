package hanoi

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

/*
Tower構造体とDisk構造体のsliceをセットで扱うための構造体
idは0, 1, 2のいずれか
disksは大きい順に並んでいるものとする
*/
type Hanoi struct {
	id    int
	tower *Tower
	disks []*Disk
}

func NewHanoi(id int) *Hanoi {
	return &Hanoi{
		id:    id,
		tower: NewTower(id*200+100, ScreenHeight-TowerHeight),
		disks: []*Disk{},
	}
}

func (h *Hanoi) Draw(screen *ebiten.Image) {
	h.tower.Draw(screen)
	for i, d := range h.disks {
		d.Draw(screen, h.id*200+100-d.width/2+TowerWidth/2, ScreenHeight-(i+1)*DiskHeight-StandHeight)
	}
}

func (h *Hanoi) Top() *Disk {
	if len(h.disks) == 0 {
		return nil
	}
	return h.disks[len(h.disks)-1]
}

func (h *Hanoi) Pop() *Disk {
	if len(h.disks) == 0 {
		return nil
	}
	d := h.disks[len(h.disks)-1]
	h.disks = h.disks[:len(h.disks)-1]
	return d
}

func (h *Hanoi) Push(d *Disk) {
	h.disks = append(h.disks, d)
}

func (h *Hanoi) IsEmpty(x, y int) bool {
	return len(h.disks) == 0
}

func (h *Hanoi) CanPush(d *Disk) bool {
	return h.IsEmpty(0, 0) || h.Top().size > d.size
}

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
