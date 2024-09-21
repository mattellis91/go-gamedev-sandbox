package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"log"

	"golang.org/x/image/font/gofont/goregular"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/joelschutz/stagehand"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type State int

type BaseScene struct {
	bounds image.Rectangle
	count  State
	sm     *stagehand.SceneManager[State]
}

func (s *BaseScene) Layout(w, h int) (int, int) {
	s.bounds = image.Rect(0, 0, w, h)
	return w, h
}

func (s *BaseScene) Load(st State, sm stagehand.SceneController[State]) {
	s.count = st
	s.sm = sm.(*stagehand.SceneManager[State])
}

func (s *BaseScene) Unload() State {
	return s.count
}

type FirstScene struct {
	BaseScene
	ui *ebitenui.UI
}

func (s *FirstScene) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.count++
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		s.sm.SwitchTo(&SecondScene{})
	}
	s.ui.Update()
	return nil
}

func (s *FirstScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{255, 0, 0, 255}) // Fill Red
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Count: %v, WindowSize: %s", s.count, s.bounds.Max), s.bounds.Dx()/2, s.bounds.Dy()/2)
	s.ui.Draw(screen)
}

type SecondScene struct {
	BaseScene
}

func (s *SecondScene) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.count--
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		s.sm.SwitchWithTransition(&ThirdScene{}, stagehand.NewFadeTransition[State](.05))
	}
	return nil
}

func (s *SecondScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 255, 255}) // Fill Blue
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Count: %v, WindowSize: %s", s.count, s.bounds.Max), s.bounds.Dx()/2, s.bounds.Dy()/2)
}

type ThirdScene struct {
	BaseScene
}

func (s *ThirdScene) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.count *= 2
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		s.sm.SwitchWithTransition(&FirstScene{}, stagehand.NewSlideTransition[State](stagehand.RightToLeft, .05))
	}
	return nil
}

func (s *ThirdScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 255, 0, 255}) // Fill Green
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Count: %v, WindowSize: %s", s.count, s.bounds.Max), s.bounds.Dx()/2, s.bounds.Dy()/2)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("My Game")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	state := State(10)

	rootContainer := widget.NewContainer()

	eui := &ebitenui.UI{
		Container: rootContainer,
	}

	so, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
	}

	fontFace := &text.GoTextFace{
		Source: so,
		Size:   32,
	}
	// This creates a text widget that says "Hello World!"
	helloWorldLabel := widget.NewText(
		widget.TextOpts.Text("Hello World!", fontFace, color.White),
	)

	// To display the text widget, we have to add it to the root container.
	rootContainer.AddChild(helloWorldLabel)

	s := &FirstScene{
		ui: eui,
	}
	
	sm := stagehand.NewSceneManager[State](s, state)

	if err := ebiten.RunGame(sm); err != nil {
		log.Fatal(err)
	}
}