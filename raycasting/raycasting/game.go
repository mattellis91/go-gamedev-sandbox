package raycasting

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/solarlune/resolv"
)

const (
	ScreenWidth  = 820
	ScreenHeight = 640
	RayLength    = 100
	RayWidth     = 1
	WallWidth    = 5
)

type Game struct {
	walls []*Wall
	ray   *Ray
}

type Wall struct {
	a resolv.Vector
	b resolv.Vector
}

type Ray struct {
	pos resolv.Vector
	dir resolv.Vector
}

func NewGame() *Game {
	return &Game{
		walls: []*Wall{
			NewWall(500, 100, 500, 200),
		},
		ray: NewRay(100, 200, 1, 0),
	}
}

func NewWall(x1, y1, x2, y2 float64) *Wall {
	return &Wall{
		a: resolv.NewVector(x1, y1),
		b: resolv.NewVector(x2, y2),
	}
}

func NewRay(x, y, dirX, dirY float64) *Ray {
	return &Ray{
		pos: resolv.NewVector(x, y),
		dir: resolv.NewVector(dirX, dirY),
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	// Draw walls
	for _, wall := range g.walls {
		vector.StrokeLine(
			screen,
			float32(wall.a.X),
			float32(wall.a.Y),
			float32(wall.b.X),
			float32(wall.b.Y),
			WallWidth,
			color.White,
			false,
		)

	}

	vector.StrokeLine(
		screen,
		float32(g.ray.pos.X),
		float32(g.ray.pos.Y),
		float32(g.ray.pos.X+(g.ray.dir.X*RayLength)),
		float32(g.ray.pos.Y+(g.ray.dir.Y*RayLength)),
		RayWidth,
		color.White,
		false,
	)
}
