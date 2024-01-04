package hanoi

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
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
	start    time.Time
	end      time.Time

	selectPage *SelectPage
	level      Level
}

const (
	ScreenWidth  = 960
	ScreenHeight = 540
)

func NewGame() *Game {
	return &Game{
		stand:      NewStand(),
		mode:       ModeSelect,
		selectPage: NewSelectPage(),
		level:      Level1,
	}
}

func (g *Game) initGame() {
	g.hanoi = [3]*Hanoi{
		NewHanoi(0),
		NewHanoi(1),
		NewHanoi(2),
	}
	for i := 0; i < g.level.Int()+2; i++ {
		g.hanoi[0].Push(NewDisk(g.level.Int()+2-i, g.level)) // 1から順にディスクを積む
	}
	g.count = 0
	g.start = time.Now()
	g.end = time.Time{}
}

func (g *Game) isCleared() bool {
	return len(g.hanoi[2].disks) == g.level.Int()+2
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

			if g.isCleared() {
				g.mode = ModeResult
				g.end = time.Now()
			}
		}
	case ModeResult:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.mode = ModeSelect
			g.initGame()
		}
		return nil
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.mode == ModeSelect {
		g.selectPage.Draw(screen)
		return
	}

	for _, h := range g.hanoi {
		h.Draw(screen, h == g.selected, h == g.hovered)
	}

	if g.selected != nil {
		x, y := ebiten.CursorPosition()
		g.selected.Top().Draw(screen, x-g.selected.Top().width/2, y-DiskHeight/2, 1)
	}

	g.stand.Draw(screen)

	fontSize := 36
	if g.mode == ModeResult {
		yellow := color.RGBA{0xff, 0xff, 0x00, 0xff}
		text.Draw(screen, "Congratulation!", mplusNormalFont, 1*ScreenWidth/4, fontSize, yellow)
		text.Draw(screen, fmt.Sprintf("count: %d", g.count), mplusNormalFont, 1*ScreenWidth/4, 2*fontSize, yellow)
		text.Draw(screen, fmt.Sprintf("time: %.3f", float64(g.end.Sub(g.start).Milliseconds())/1000), mplusNormalFont, 1*ScreenWidth/4, 3*fontSize, yellow)
		text.Draw(screen, "Press Space to return to the title", mplusNormalFont, 1*ScreenWidth/4, 4*fontSize, yellow)
	} else {
		text.Draw(screen, fmt.Sprintf("level: %d", g.level.Int()), mplusNormalFont, 10, fontSize, color.White)
		text.Draw(screen, fmt.Sprintf("count: %d", g.count), mplusNormalFont, 10, 2*fontSize, color.White)
		text.Draw(screen, fmt.Sprintf("time: %.3f", float64(time.Since(g.start).Milliseconds())/1000), mplusNormalFont, 10, 3*fontSize, color.White)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
