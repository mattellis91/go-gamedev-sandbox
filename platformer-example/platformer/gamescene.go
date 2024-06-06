package platformer

import (

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mattellis91/go-gamedev-sandbox/platformer-example/platformer/ldtk"
	"github.com/mattellis91/go-gamedev-sandbox/platformer-example/platformer/level"
	"github.com/mattellis91/go-gamedev-sandbox/platformer-example/platformer/resources/levels"
)

type GameScene struct {
	CurrentLevel *level.GameLevel
}

func NewGameScene(startingLevelId string) *GameScene {

	gameData, err := ldtk.Read(levels.Dungeon_ldtk)
	if err != nil {
		panic(err)
	}

	g := &GameScene{
		CurrentLevel: level.NewGameLevel(gameData.Levels[0]),
	}
	return g
}


func (s *GameScene) Update(state *GameState) error {
	s.CurrentLevel.Update()
	return nil
}

func (s *GameScene) Draw(screen *ebiten.Image) {
	s.CurrentLevel.Draw(screen)
}