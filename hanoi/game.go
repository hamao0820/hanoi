package hanoi

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Mode int

const (
	ModeSelect Mode = iota
	ModeGame
	ModeResult
)

type Game struct {
	hanoi    [3]*Hanoi
	stand    *Stand
	selected *Hanoi
	hovered  *Hanoi
	mode     Mode
	count    int

	selectPage *SelectPage
	level      Level
}

const (
	ScreenWidth  = 960
	ScreenHeight = 540
)

func NewGame() *Game {
	return &Game{
		hanoi: [3]*Hanoi{
			NewHanoi(0),
			NewHanoi(1),
			NewHanoi(2),
		},
		stand:      NewStand(),
		mode:       ModeSelect,
		selectPage: NewSelectPage(),
		level:      Level1,
	}
}

func (g *Game) initGame() {
	for i := 0; i < g.level.Int()+2; i++ {
		g.hanoi[0].Push(NewDisk(g.level.Int()+2-i, g.level)) // 1から順にディスクを積む
	}
}

func (g *Game) Update() error {
	switch g.mode {
	case ModeSelect:
		g.selectPage.Update()

		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			for _, b := range g.selectPage.buttons {
				if b.In(x, y) {
					g.level = b.level
					g.mode = ModeGame
					g.initGame()
					break
				}
			}
		}
		return nil
	case ModeGame:
		{
			x, y := ebiten.CursorPosition()
			exist := false
			for _, h := range g.hanoi {
				if h.tower.IsAround(x, y) {
					g.hovered = h
					exist = true
					break
				}
			}
			if !exist {
				g.hovered = nil
			}

			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
				x, y := ebiten.CursorPosition()
				for _, h := range g.hanoi {
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
		}
	case ModeResult:
		return nil
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.mode {
	case ModeSelect:
		g.selectPage.Draw(screen)
	case ModeGame:
		{
			for _, h := range g.hanoi {
				h.Draw(screen, h == g.selected, h == g.hovered)
			}

			if g.selected != nil {
				x, y := ebiten.CursorPosition()
				g.selected.Top().Draw(screen, x-g.selected.Top().width/2, y-DiskHeight/2, 1)
			}

			g.stand.Draw(screen)

			ebitenutil.DebugPrint(screen, fmt.Sprintf("count: %d", g.count))
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
