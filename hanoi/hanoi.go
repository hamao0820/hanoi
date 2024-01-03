package hanoi

import "github.com/hajimehoshi/ebiten/v2"

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
