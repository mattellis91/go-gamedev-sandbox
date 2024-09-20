package rpg

import (

	"fmt"
	
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 672
	ScreenHeight = 400
)

type Game struct {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	msg := fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS())
	msg2 := fmt.Sprintf("FPS: %0.2f", ebiten.ActualFPS())
	ebitenutil.DebugPrint(screen, msg)
	ebitenutil.DebugPrintAt(screen, msg2, 0, 20)
}

func NewGame() *Game {
	return &Game{
	
	}
}
