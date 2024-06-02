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

var (
	tilesImage *ebiten.Image
	ldtkProject *ldtk.Project
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(tiles.Tiles_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)
	
	ldtkProject, err = ldtk.Read(levels.Platformer_ldtk)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Project: %v\n", ldtkProject)
}

type GameScene struct {
	
}


const (
	EMPTY = -1
	GRASS_TL = 6
	GRASS_TM = 7
	GRASS_TR = 8
	GROUND_TL = 28
	GROUND_TM = 29
	GROUND_TR = 30
)

func NewGameScene() *GameScene {
	return &GameScene{
	}
}

func (s *GameScene) Update(state *GameState) error {
	return nil
}

func (s *GameScene) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}