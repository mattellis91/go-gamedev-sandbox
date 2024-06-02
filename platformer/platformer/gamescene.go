package platformer

import (
	"bytes"
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/mattellis91/go-gamedev-sandbox/platformer/platformer/resources/tiles"
	"github.com/mattellis91/go-gamedev-sandbox/platformer/platformer/resources/ldtk"
	"github.com/mattellis91/go-gamedev-sandbox/platformer/platformer/ldtk"
)

const (
	tileSize = 16
)

const (
	EMPTY = -1
	GRASS_TL = 6
	GRASS_TM = 7
	GRASS_TR = 8
	GROUND_TL = 28
	GROUND_TM = 29
	GROUND_TR = 30
)

var (
	tilesImage *ebiten.Image
	ldtkProject *ldtk.Project
)

type GameScene struct {
	testLevel *ldtk.Level
	player *Player
}

func loadResources() {
	img, _, err := image.Decode(bytes.NewReader(tiles.Tiles_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)
	
	ldtkProject, err = ldtk.Read(levels.Platformer_ldtk)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Project: %+v\n", ldtkProject.Levels[0].Layers[0])
}

const (
	TILES_LAYER = "Game tiles"
	ENTITIES_LAYER = "Entities"
)

func NewGameScene() *GameScene {
	loadResources()
	sceneLevel := ldtkProject.Levels[0]
	return &GameScene{
		testLevel: sceneLevel,
		player: NewPlayer(),
	}
}

func (s *GameScene) Update(state *GameState) error {
	return nil
}

func (s *GameScene) Draw(screen *ebiten.Image) {

	for _, layer := range s.testLevel.Layers {
		if len(layer.Tiles) > 0 {
			for _, tile := range layer.Tiles {
				op := &ebiten.DrawImageOptions{}
				tilePos := tile.Position
				tileSubImagePos :=  tile.Src
				op.GeoM.Translate(float64(tilePos[0]), float64(tilePos[1]))
				screen.DrawImage(tilesImage.SubImage(
					image.Rect(tileSubImagePos[0], tileSubImagePos[1], tileSubImagePos[0] + tileSize, tileSubImagePos[1] + tileSize)).(*ebiten.Image),
				op)
			}
		}
	}

	s.player.Draw(screen)


	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))

}