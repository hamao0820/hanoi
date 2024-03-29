package hanoi

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
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
	var c color.Color = color.White
	if id == 2 {
		c = color.RGBA{0xff, 0x00, 0x00, 0xff}
	}
	return &Hanoi{
		id:    id,
		tower: NewTower(id*ScreenWidth/3+ScreenWidth/6-TowerWidth/2, ScreenHeight-TowerHeight, c),
		disks: []*Disk{},
	}
}

func (h *Hanoi) Draw(screen *ebiten.Image, isSelected bool, isHovered bool) {
	if isHovered {
		h.tower.DrawHover(screen)
	}
	h.tower.Draw(screen)
	for i, d := range h.disks {
		var alpha float32 = 1.0
		if isSelected && d == h.Top() {
			alpha = 0.5
		}
		d.Draw(screen, h.id*ScreenWidth/3+ScreenWidth/6-d.width/2, ScreenHeight-(i+1)*DiskHeight-StandHeight, alpha)
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
