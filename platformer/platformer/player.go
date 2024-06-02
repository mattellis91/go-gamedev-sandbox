package platformer

import (
	"bytes"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mattellis91/go-gamedev-sandbox/platformer/platformer/resources/sprites"
)

const (
	FrameOX = 0
	FrameOY = 32
	FrameWidth = 32
	FrameHeight = 32
	FrameCount = 8
)

var (
	playerRunImage *ebiten.Image
)

type Player struct {
	playerFrameCount int
	x int
	y int
}

func NewPlayer(x, y int) *Player {

	// Load the player run image
	img, _, err := image.Decode(bytes.NewReader(sprites.Player_run_png))
	if err != nil {
		log.Fatal(err)
	
	}
	playerRunImage = ebiten.NewImageFromImage(img)

	return &Player{
		x: x,
		y: y,
	}
}

func (p *Player) Update() error {
	p.playerFrameCount++
	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.x), float64(p.y))
	screen.DrawImage(playerRunImage.SubImage(image.Rect(0, 0, 32, 32)).(*ebiten.Image), op)
}

