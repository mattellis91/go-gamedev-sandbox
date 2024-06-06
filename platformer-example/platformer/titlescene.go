package platformer

import (
	_ "image/png"
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type TitleScene struct {
	count int
}

func anyGamepadVirtualButtonJustPressed(i *Input) bool {
	if !i.gamepadConfig.IsGamepadIDInitialized() {
		return false
	}

	for _, b := range virtualGamepadButtons {
		if i.gamepadConfig.IsButtonJustPressed(b) {
			return true
		}
	}
	return false
}

func (s *TitleScene) Update(state *GameState) error {
	s.count++
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		state.SceneManager.GoTo(NewGameScene("level1"))
		println("Space pressed")
		return nil
	}

	if anyGamepadVirtualButtonJustPressed(state.Input) {
		state.SceneManager.GoTo(NewGameScene("level1"))
		return nil
	}

	if state.Input.gamepadConfig.IsGamepadIDInitialized() {
		return nil
	}

	// If 'virtual' gamepad buttons are not set and any gamepad buttons are pressed,
	// go to the gamepad configuration scene.
	id := state.Input.GamepadIDButtonPressed()
	if id < 0 {
		return nil
	}
	state.Input.gamepadConfig.SetGamepadID(id)
	if state.Input.gamepadConfig.NeedsConfiguration() {
		g := &GamepadScene{}
		g.gamepadID = id
		state.SceneManager.GoTo(g)
	}
	return nil
}

func (s *TitleScene) Draw(r *ebiten.Image) {
	
	message := "PRESS SPACE TO START"
	x := ScreenWidth / 2
	y := ScreenHeight / 2 - 16
	drawTextWithShadow(r, message, x, y, 1, color.RGBA{0x80, 0, 0, 0xff}, text.AlignCenter, text.AlignStart)
	
}
