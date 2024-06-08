package level

import (

	"github.com/mattellis91/go-gamedev-sandbox/platformer-example/platformer/util"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

type Spike struct {
	image *ebiten.Image
	x, y int
	object *resolv.Object
	space *resolv.Space
}

func NewSpike(image *ebiten.Image, x, y int, space *resolv.Space) *Spike {
	space.Add(resolv.NewObject(float64(x), float64(y), util.TileSize, util.TileSize, util.PlayerEntityIdentifier))
	return &Spike{
		image: image,
		x: x,
		y: y,
		object: resolv.NewObject(float64(x), float64(y), util.TileSize, util.TileSize),
		space: space,
	}
}

func (s *Spike) Update() error{
	return nil
}

func (s *Spike) Draw(screen *ebiten.Image) {
	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Translate(float64(s.x), float64(s.y))
	screen.DrawImage(s.image, ops)
}
