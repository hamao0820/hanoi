package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hamao0820/hanoi/hanoi"
)

func main() {
	ebiten.SetWindowSize(hanoi.ScreenWidth, hanoi.ScreenHeight)
	ebiten.SetWindowTitle("Hanoi")
	if err := ebiten.RunGame(hanoi.NewGame()); err != nil {
		panic(err)
	}
}
