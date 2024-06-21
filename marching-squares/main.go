package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mattellis91/go-gamedev-sandbox/marching-squares/marching-squares"
)

func main() {
	game := marchingsquares.NewGame()
	ebiten.SetWindowSize(marchingsquares.ScreenWidth, marchingsquares.ScreenHeight)
	ebiten.SetWindowTitle("Marchin Squares")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}