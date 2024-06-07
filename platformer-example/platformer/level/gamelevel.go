package level

import (
	"bytes"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/mattellis91/go-gamedev-sandbox/platformer-example/platformer/ldtk"
	"github.com/mattellis91/go-gamedev-sandbox/platformer-example/platformer/resources/tiles"
	"github.com/mattellis91/go-gamedev-sandbox/platformer-example/platformer/util"
	"github.com/solarlune/resolv"
)

type GameLevel struct {
	levelSpace *resolv.Space
	levelTiles []*ldtk.Tile
	tileSet    *ebiten.Image
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

	// Draw the level tiles
	for _, tile := range g.levelTiles {
		ops := &ebiten.DrawImageOptions{}
		ops.GeoM.Translate(float64(tile.Position[0]), float64(tile.Position[1]))
		tile := g.tileSet.SubImage(image.Rect(tile.Src[0], tile.Src[1], tile.Src[0]+util.TileSize, tile.Src[1]+util.TileSize)).(*ebiten.Image)
		screen.DrawImage(tile, ops)
	}
}

func (g *GameLevel) init(levelData *ldtk.Level) {

	// Decode an image from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(tiles.Dungeon_tiles_png))
	if err != nil {
		log.Fatal(err)
	}
	g.tileSet = ebiten.NewImageFromImage(img)

	solidLayer := levelData.LayerByIdentifier(util.SolidTileLayerIdentifier)

	g.levelTiles = solidLayer.Tiles

	for _, tile := range solidLayer.Tiles {
		g.levelSpace.Add(resolv.NewObject(
			float64(tile.Position[0]),
			float64(tile.Position[1]),
			float64(util.TileSize),
			float64(util.TileSize), util.SolidTileSpaceIdentifier),
		)
	}
}
