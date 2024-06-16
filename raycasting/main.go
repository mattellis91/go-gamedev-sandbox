package main

import (

	"github.com/mattellis91/go-gamedev-sandbox/raycasting/raycasting"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(raycasting.ScreenWidth, raycasting.ScreenHeight)
	ebiten.SetWindowTitle("Raycasting")
	if err := ebiten.RunGame(raycasting.NewGame()); err != nil {
		panic(err)
	}
}