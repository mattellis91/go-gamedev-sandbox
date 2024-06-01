package blankproject

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func init() {

}

type GameScene struct {

}

func NewGameScene() *GameScene {
	return &GameScene{}
}

func (s *GameScene) Update(state *GameState) error {
	return nil
}

func (s *GameScene) Draw(screen *ebiten.Image) {
	message := "GAME SCENE"
	x := ScreenWidth / 2
	y := ScreenHeight - 48
	drawTextWithShadow(screen, message, x, y, 1, color.RGBA{0x80, 0, 0, 0xff}, text.AlignCenter, text.AlignStart)
}