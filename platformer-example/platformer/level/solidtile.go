package level

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

type SolidTile struct {
	image *ebiten.Image
	Object * resolv.Object
}

func NewSolidTile(image *ebiten.Image, x, y, width, height int) *SolidTile {
	s := &SolidTile{
		image: image,
		
	}
	return s
}

func (s *SolidTile) Draw(screen *ebiten.Image) {
	// Draw the solid tile
	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Translate(float64(s.Object.Position.X), float64(s.Object.Position.Y))
	screen.DrawImage(s.image, ops)
}