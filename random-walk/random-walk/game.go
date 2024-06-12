package randomwalk

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	ScreenWidth  = 1024
	ScreenHeight = 1024
	Cols 	   = 64	
	Rows 	   = 64
	CellWidth    = ScreenWidth / Cols
	CellHeight    = ScreenHeight / Rows	
)

var TilesMax int = int(math.Floor((Cols * Rows) / 2))

const (
	Up = iota
	Down
	Left
	Right
)

func RandomDirection() int {
	return rand.Intn(4)
}

type Game struct {
	grid [][]int
	tilesPlaced int
	lastPlaced []int
}

func NewGame() *Game {

	grid := make([][]int, Rows)
	for i := range grid {
		grid[i] = make([]int, Cols)
	}

	for i := 0; i < Rows; i++ {
		for j := 0; j < Cols; j++ {
			grid[i][j] = 0
		}
	}

	randomX := rand.Intn(Cols)
	randomY := rand.Intn(Rows)

	grid[randomY][randomX] = 1

	return &Game{
		grid: grid,
		tilesPlaced: 1,
		lastPlaced: []int{randomX, randomY},
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	if g.tilesPlaced < TilesMax {
		randDir := RandomDirection()
		switch randDir {
			case Up:
				if g.lastPlaced[1] - 1 >= 0 {
					g.grid[g.lastPlaced[1] - 1][g.lastPlaced[0]] = 1
					g.lastPlaced[1] = g.lastPlaced[1] - 1
					g.tilesPlaced++
				}
			case Down:
				if g.lastPlaced[1] + 1 < Rows {
					g.grid[g.lastPlaced[1] + 1][g.lastPlaced[0]] = 1
					g.lastPlaced[1] = g.lastPlaced[1] + 1
					g.tilesPlaced++
				}
			case Left:
				if g.lastPlaced[0] - 1 >= 0 {
					g.grid[g.lastPlaced[1]][g.lastPlaced[0] - 1] = 1
					g.lastPlaced[0] = g.lastPlaced[0] - 1
					g.tilesPlaced++
				}
			case Right:
				if g.lastPlaced[0] + 1 < Cols {
					g.grid[g.lastPlaced[1]][g.lastPlaced[0] + 1] = 1
					g.lastPlaced[0] = g.lastPlaced[0] + 1
					g.tilesPlaced++
				}
		}
	} else {
		fmt.Println("All tiles placed")
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{108, 122, 137, 255})
	for i := 0; i < Rows; i++ {
		for j := 0; j < Cols; j++ {
			if g.grid[i][j] == 1 {
				vector.DrawFilledRect(screen, float32(j*CellWidth), float32(i*CellHeight), float32(CellWidth), float32(CellHeight), color.RGBA{255, 255, 255, 255}, false)
			}
		}
	}
}

