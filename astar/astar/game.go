package astar

import (
	_ "fmt"
	"image/color"
	"math/rand/v2"

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
	WallChance = 30
)

var path []*Node

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
	neighbors []*Node
	prev *Node
	wall bool
}

func (n *Node) AddNeighbors(grid [][]*Node) {
	x := n.x
	y := n.y
	if x < Cols - 1 {
		n.neighbors = append(n.neighbors, grid[x + 1][y])
	}
	if x > 0 {
		n.neighbors = append(n.neighbors, grid[x - 1][y])
	}
	if y < Rows - 1 {
		n.neighbors = append(n.neighbors, grid[x][y + 1])
	}
	if y > 0 {
		n.neighbors = append(n.neighbors, grid[x][y - 1])
	}
}

func EuclideanDistance(x1, y1, x2, y2 int) int {
	return (x1 - x2) * (x1 - x2) + (y1 - y2) * (y1 - y2)
}

func NewGame() *Game {
	grid := make([][]*Node, Cols)
	for i := range grid {
		grid[i] = make([]*Node, Rows)
		for j := range grid[i] {
			grid[i][j] = &Node{x: i, y: j}

			//chance to make node a wall
			if rand.IntN(100) < WallChance {
				grid[i][j].wall = true
			}
		}
	}

	for i := range grid {
		for j := range grid[i] {
			grid[i][j].AddNeighbors(grid)
		}
	}

	start := grid[0][0]

	g := &Game {
		Grid: grid,
		Start: start,
		End: nil,
		OpenSet: []*Node{start},
		ClosedSet: []*Node{},
	}

	endX, endY := g.GetRandColAndRow()
	g.End = g.Grid[endX][endY]

	if g.Start.wall {
		g.Start.wall = false
	}
	
	if g.End.wall {
		g.End.wall = false
	}

	return g
}

func (g *Game) Reset() {
	g.OpenSet = []*Node{g.Start}
	g.ClosedSet = []*Node{}
	endX, endY := g.GetRandColAndRow()
	g.End = g.Grid[endX][endY]
}


func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func NodeInSlice(node *Node, slice []*Node) bool {
	for _, n := range slice {
		if n == node {
			return true
		}
	}
	return false
}

func (g *Game) GetRandColAndRow() (int, int) {
	randX := rand.IntN(Cols) + 1
	randY := rand.IntN(Rows) + 1
	return randX, randY
}

func (g *Game) Update() error {


	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.Reset()
		path = nil
	}


	if(len(g.OpenSet) > 0) {
		// find the node in the open set with the lowest f value
		lowestIndex := 0
		for i, node := range g.OpenSet {
			if node.f < g.OpenSet[lowestIndex].f {
				lowestIndex = i
			}
		}
		current := g.OpenSet[lowestIndex]

		if current == g.End {

			// reconstruct path
			path = []*Node{current}
			temp := current

			for temp.prev != nil {
				path = append(path, temp.prev)
				temp = temp.prev
			}

			return nil
		}

		g.ClosedSet = append(g.ClosedSet, current)
		g.OpenSet = append(g.OpenSet[:lowestIndex], g.OpenSet[lowestIndex+1:]...)

		for _, neighbor := range current.neighbors {
			
			neighborInClosedSet := NodeInSlice(neighbor, g.ClosedSet)

			if !neighborInClosedSet && !neighbor.wall {
				
				tempG := current.g + 1
				neighborInOpenSet := NodeInSlice(neighbor, g.OpenSet)
				if neighborInOpenSet {
					if tempG < neighbor.g {
						neighbor.g = tempG
					}
				} else {
					neighbor.g = tempG
					g.OpenSet = append(g.OpenSet, neighbor)
				}

				neighbor.h = EuclideanDistance(neighbor.x, neighbor.y, g.End.x, g.End.y)
				neighbor.f = neighbor.g + neighbor.h
				neighbor.prev = current

			}

		}

	} else {
		// no solution
		return nil
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	for _, col := range g.Grid {
		for _, node := range col {
			color := color.RGBA{0x40, 0x80, 0x80, 0xff}

			//draw end node in blue
			if node == g.End {
				color.G = 0x00 
				color.B = 0xff
			}

			//draw wall node in black
			if node.wall {
				color.R = 0x00
				color.G = 0x00
				color.B = 0x00
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

	//split open set and closed set drawing for better visualization of the algorithm
	for _, node := range g.OpenSet {
		color := color.RGBA{0x00, 0xff, 0x00, 0xff}

		//draw end node in blue
		if node == g.End {
			color.G = 0x00 
			color.B = 0xff
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

	for _, node := range g.ClosedSet {
		color := color.RGBA{0xff, 0x00, 0x00, 0xff}
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

	//draw grid
	gridColor := color.RGBA{0xff, 0xff, 0xff, 0xff}
	for i := 0; i < Cols; i++ {
		startX := float32(i * CellWidth)
		startY := float32(0)
		endX := float32(i * CellWidth)
		endY := float32(ScreenHeight)
		vector.StrokeLine(screen, startX, startY, endX, endY, 1, gridColor, false)
	}

	for i := 0; i < Rows; i++ {
		startX := float32(0)
		startY := float32(i * CellHeight)
		endX := float32(ScreenWidth)
		endY := float32(i * CellHeight)
		vector.StrokeLine(screen, startX, startY, endX, endY, 1, gridColor, false)
	}

	//draw path
	for _, node := range path {
		color := color.RGBA{0x00, 0x00, 0xff, 0xff}
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
