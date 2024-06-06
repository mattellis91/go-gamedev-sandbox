package level

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/mattellis91/go-gamedev-sandbox/platformer-example/platformer/ldtk"
	"github.com/mattellis91/go-gamedev-sandbox/platformer-example/platformer/util"
	"github.com/solarlune/resolv"
)

type GameLevel struct {
	levelSpace *resolv.Space
}

func NewGameLevel(levelData *ldtk.Level) *GameLevel {
	s := resolv.NewSpace(util.ScreenWidth, util.ScreenHeight, util.TileSize, util.TileSize)
	g := &GameLevel{
		levelSpace: s,
	}
	g.init(levelData)
	return g
}


func (g *GameLevel) Update() error {
	return nil
}

func (g *GameLevel) Draw(screen *ebiten.Image) {
	for _, obj := range g.levelSpace.Objects() {
		// Draw the solid tile
		ops := &ebiten.DrawImageOptions{}
		ops.GeoM.Translate(float64(obj.Position.X), float64(obj.Position.Y))
		vector.DrawFilledRect(screen, float32(obj.Position.X), float32(obj.Position.Y), float32(util.TileSize), float32(util.TileSize), color.RGBA{255, 50, 100, 0}, false)
	
	}
}

func (g *GameLevel) init(levelData *ldtk.Level) {
	solidLayer := levelData.LayerByIdentifier(util.SolidTileLayerIdentifier)

	for _, tile := range solidLayer.Tiles {
		g.levelSpace.Add(resolv.NewObject(
			float64(tile.Position[0]), 
			float64(tile.Position[1]), 
			float64(util.TileSize), 
			float64(util.TileSize), util.SolidTileSpaceIdentifier),
		)
	}
}