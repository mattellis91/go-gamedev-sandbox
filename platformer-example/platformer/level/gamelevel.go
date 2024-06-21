package level

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/mattellis91/go-gamedev-sandbox/platformer-example/platformer/ldtk"
	"github.com/mattellis91/go-gamedev-sandbox/platformer-example/platformer/resources/tiles"
	"github.com/mattellis91/go-gamedev-sandbox/platformer-example/platformer/util"
	"github.com/solarlune/resolv"
)

const (
	Friction = 0.5
	Accel = 0.5 + Friction
	MaxSpeed = 4.0
	JumpSpeed = 10.0
	Gravity = 0.75
)

type GameLevel struct {
	levelSpace *resolv.Space
	levelTiles []*ldtk.Tile
	levelEntities []Entity
	tileSet    *ebiten.Image
	drawDebug  bool
}

type Entity interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type LevelTile interface {
	Draw(screen *ebiten.Image)
}

func NewGameLevel(gameData *ldtk.Project) *GameLevel {
	s := resolv.NewSpace(util.ScreenWidth, util.ScreenHeight, util.TileSize, util.TileSize)
	g := &GameLevel{
		levelSpace: s,
	}
	g.init(gameData)
	return g
}

func (g *GameLevel) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		g.drawDebug = !g.drawDebug
	}

	for _, entity := range g.levelEntities {
		entity.Update()
	}
	return nil
}

func (g *GameLevel) Draw(screen *ebiten.Image) {

	if g.drawDebug {

		for _, obj := range g.levelSpace.Objects() {
			// Draw the solid tile
			ops := &ebiten.DrawImageOptions{}
			ops.GeoM.Translate(float64(obj.Position.X), float64(obj.Position.Y))
			vector.DrawFilledRect(screen, float32(obj.Position.X), float32(obj.Position.Y), float32(util.TileSize), float32(util.TileSize), color.RGBA{255, 50, 100, 0}, false)
		}

	}

	// Draw the level tiles
	for _, tile := range g.levelTiles {
		ops := &ebiten.DrawImageOptions{}
		ops.GeoM.Translate(float64(tile.Position[0]), float64(tile.Position[1]))
		tile := g.tileSet.SubImage(image.Rect(tile.Src[0], tile.Src[1], tile.Src[0]+util.TileSize, tile.Src[1]+util.TileSize)).(*ebiten.Image)
		screen.DrawImage(tile, ops)
	}

	//Draw the entities
	for _, entity := range g.levelEntities {
		entity.Draw(screen)
	}
}

func (g *GameLevel) init(gameData *ldtk.Project) {

	levelData := gameData.Levels[0]

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

	layerEntities := levelData.LayerByIdentifier(util.EntityLayerIdentifier)

	for _, entity := range layerEntities.Entities {

		fmt.Printf("Entity: %v\n", entity.Identifier)
		
		switch entity.Identifier {
		case "Player":
			playerEntity := gameData.EntityDefinitionByIdentifier(entity.Identifier)
			playerImg := g.tileSet.SubImage(image.Rect(
				playerEntity.TileRect.X, 
				playerEntity.TileRect.Y, 
				playerEntity.TileRect.X + playerEntity.Width, 
				playerEntity.TileRect.Y + playerEntity.Height)).(*ebiten.Image)
			g.levelEntities = append(g.levelEntities, NewPlayer(playerImg, float64(entity.Position[0]), float64(entity.Position[1]), g.levelSpace))
		case "Spike_Up", 
			 "Spike_Down", 
			 "Spike_Left", 
			 "Spike_Right":
			spikeEntity := gameData.EntityDefinitionByIdentifier(entity.Identifier)
			fmt.Printf("SpikeEntity: %v\n", spikeEntity)
			spikeImg := g.tileSet.SubImage(image.Rect(
				spikeEntity.TileRect.X,
				spikeEntity.TileRect.Y,
				spikeEntity.TileRect.X + spikeEntity.Width,
				spikeEntity.TileRect.Y + spikeEntity.Height)).(*ebiten.Image)
			g.levelEntities = append(g.levelEntities, NewSpike(spikeImg, entity.Position[0], entity.Position[1], g.levelSpace))
		}
	}
}
