package platformer

import (

	"fmt"
	
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 672
	ScreenHeight = 400
)

const (
	DT          = 0.1
	WINDOWSIZE  = 256
	WINDOWSCALE = 3
	TILESIZE    = 16
)

type Game struct {
	sceneManager *SceneManager
	input        Input
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	if g.sceneManager == nil {
		g.sceneManager = &SceneManager{}
		g.sceneManager.GoTo(&TitleScene{})
	}

	g.input.Update()
	if err := g.sceneManager.Update(&g.input); err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
	msg := fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS())
	msg2 := fmt.Sprintf("FPS: %0.2f", ebiten.ActualFPS())
	ebitenutil.DebugPrint(screen, msg)
	ebitenutil.DebugPrintAt(screen, msg2, 0, 20)
}
