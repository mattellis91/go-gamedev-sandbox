package astar

import (
	_ "fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
	Cols 	   = 5
	Rows 	   = 5
	CellWidth    = ScreenWidth / Cols
	CellHeight    = ScreenHeight / Rows
)

type Game struct {
	Grid [][]*Node
	OpenSet []*Node
	ClosedSet []*Node 
	Start *Node
	End *Node
}

type Node struct {
	x int
	y int
	f int // f = g + h
	g int // cost from start node
	h int // cost to end node
}

func NewGame() *Game {
	grid := make([][]*Node, Cols)
	for i := range grid {
		grid[i] = make([]*Node, Rows)
		for j := range grid[i] {
			grid[i][j] = &Node{x: i, y: j}
		}
	}

	start := grid[0][0]
	end := grid[Cols - 1][Rows - 1]

	g := &Game {
		Grid: grid,
		Start: start,
		End: end,
		OpenSet: []*Node{start},
		ClosedSet: []*Node{},
	}

	return g
}


func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {


	if(len(g.OpenSet) > 0) {
		// find the node in the open set with the lowest f value

	} else {
		// no solution
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	//fill nodes
	for _, col := range g.Grid {
		for _, node := range col {
			color := color.RGBA{0x40, 0x80, 0x80, 0xff}
			if node == g.Start {
				color.R = 0xff
			}
			vector.DrawFilledRect(
				screen, 
				float32(node.x * CellWidth), 
				float32(node.y * CellHeight), 
				float32(CellWidth), 
				float32(CellHeight), 
				color, 
				false,
			)
		}
	}

	//draw grid
	for i := 0; i < Cols; i++ {
		startX := float32(i * CellWidth)
		startY := float32(0)
		endX := float32(i * CellWidth)
		endY := float32(ScreenHeight)
		vector.StrokeLine(screen, startX, startY, endX, endY, 1, color.RGBA{0xff, 0xff, 0xff, 0xff}, false)
	}

	for i := 0; i < Rows; i++ {
		startX := float32(0)
		startY := float32(i * CellHeight)
		endX := float32(ScreenWidth)
		endY := float32(i * CellHeight)
		vector.StrokeLine(screen, startX, startY, endX, endY, 1, color.RGBA{0xff, 0xff, 0xff, 0xff}, false)
	}

}
