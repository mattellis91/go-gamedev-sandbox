package marchingsquares

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/solarlune/resolv"
)

const (
	ScreenWidth = 1200
	ScreenHeight = 800
	Scale = 20
	Cols = 1 + ScreenWidth / Scale
	Rows = 1 + ScreenHeight / Scale
)

type Game struct {
	field [][]float64
}

func NewGame() *Game {
	g := &Game{}
	g.field = make([][]float64, Cols)
	for i := range g.field {
		g.field[i] = make([]float64, Rows)
	}
	for i := 0; i < Cols; i++ {
		for j := 0; j < Rows; j++ {
			g.field[i][j] = rand.Float64()
		}
	}
	return g
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	//draw debug nodes
	// for i := 0; i < Cols; i++ {
	// 	for j := 0; j < Rows; j++ {
	// 		gScale := g.field[i][j] * 255
	// 		vector.DrawFilledCircle(screen, float32(i*Scale), float32(j*Scale), (Scale * 0.2), color.Gray{uint8(gScale)}, false)
	// 	}
	// }

	for i := 0; i < Cols - 1; i++ {
		for j := 0; j < Rows - 1; j++ {

			x := i * Scale
			y := j * Scale
			a := resolv.NewVector(float64(x + Scale / 2), float64(y)) // top left
			b := resolv.NewVector(float64(x + Scale), float64(y + Scale / 2)) // top right
			c := resolv.NewVector(float64(x + Scale / 2), float64(y + Scale)) // bottom right
			d := resolv.NewVector(float64(x), float64(y + Scale / 2)) // bottom left

			state := getSquareState(g.field[i][j], g.field[i + 1][j], g.field[i + 1][j + 1], g.field[i][j + 1])
			//fmt.Println("%d", state)
			switch state {
				case 1:
					drawLineBetweenPoints(screen, c, d)
				case 2:	
					drawLineBetweenPoints(screen, b, c)
				case 3:
					drawLineBetweenPoints(screen, b, d)
				case 4:
					drawLineBetweenPoints(screen, a, b)
				case 5:
					drawLineBetweenPoints(screen, a, d)
					drawLineBetweenPoints(screen, b, c)
				case 6:
					drawLineBetweenPoints(screen, a, c)
				case 7:
					drawLineBetweenPoints(screen, a, d)
				case 8:
					drawLineBetweenPoints(screen, a, d)
				case 9:
					drawLineBetweenPoints(screen, a, c)
				case 10:
					drawLineBetweenPoints(screen, a, b)
					drawLineBetweenPoints(screen, c, d)
				case 11:
					drawLineBetweenPoints(screen, a, b)
				case 12:
					drawLineBetweenPoints(screen, b, d)
				case 13:
					drawLineBetweenPoints(screen, b, c)
				case 14:
					drawLineBetweenPoints(screen, c, d)
			}
		}
	}
}

func getSquareState(a, b, c, d float64) int {
	s := getRoundedValue(a) * 8 + getRoundedValue(b) * 4 + getRoundedValue(c) * 2 + getRoundedValue(d) * 1
	return s
}

func getRoundedValue(v float64) int {
	if v < 0.5 {
		return 0
	} else {
		return 1
	}
}

func drawLineBetweenPoints(screen *ebiten.Image, a, b resolv.Vector) {
	vector.StrokeLine(screen, float32(a.X), float32(a.Y), float32(b.X), float32(b.Y), 3, color.White, false)
}

