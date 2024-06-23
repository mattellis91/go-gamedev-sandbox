package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mattellis91/go-gamedev-sandbox/raycasting3d/raycasting3d"
)

func main() {
	ebiten.SetWindowSize(raycasting3d.ScreenWidth, raycasting3d.ScreenHeight)
	ebiten.SetWindowTitle("Raycasting 3D")
	if err := ebiten.RunGame(raycasting3d.NewGame()); err != nil {
		panic(err)
	}
}