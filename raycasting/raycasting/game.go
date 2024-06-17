package raycasting

import (
	"image/color"
	"math"
	"math/rand"

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
	RayCollisionRadius = 8
	ParticleRadius = 10
)

type Game struct {
	walls []*Wall
	particle *Particle
}

type Wall struct {
	a resolv.Vector
	b resolv.Vector
}

type Ray struct {
	pos resolv.Vector
	dir resolv.Vector
}

type Particle struct {
	pos resolv.Vector
	rays  []*Ray
}

func NewGame() *Game {

	walls := make([]*Wall, 0)
	for i := 0; i < 5; i++ {
		x1 := rand.Float64() * ScreenWidth
		y1 := rand.Float64() * ScreenHeight
		x2 := rand.Float64() * ScreenWidth
		y2 := rand.Float64() * ScreenHeight
		walls = append(walls, NewWall(x1, y1, x2, y2))
	}

	return &Game{
		walls: walls,
		particle: NewParticle(200, ScreenHeight/2),
	}
}

func NewWall(x1, y1, x2, y2 float64) *Wall {
	return &Wall{
		a: resolv.NewVector(x1, y1),
		b: resolv.NewVector(x2, y2),
	}
}

func NewVectorFromAngle(angle float64) resolv.Vector {
	x := math.Cos(angle)
	y := math.Sin(angle)
	return resolv.NewVector(x, y)
}

func NewRay(pos resolv.Vector, angle float64) *Ray {
	return &Ray{
		pos: pos,
		dir: NewVectorFromAngle(angle),
	}
}

func NewParticle(x, y float64) *Particle {
	
	pos := resolv.NewVector(x, y)
	rays := make([]*Ray, 0)
	for i := 0; i < 360; i += 10 {
		rays = append(rays, NewRay(pos, float64(resolv.ToRadians(float64(i)))))
	}

	return &Particle{
		pos: resolv.NewVector(x, y),
		rays: rays,
	}
}
 
func (r *Ray) Cast(wall *Wall) *resolv.Vector {

	//line line intersection
	x1 := wall.a.X
	y1 := wall.a.Y
	x2 := wall.b.X
	y2 := wall.b.Y

	x3 := r.pos.X
	y3 := r.pos.Y
	x4 := r.pos.X + r.dir.X
	y4 := r.pos.Y + r.dir.Y

	den := (x1-x2)*(y3-y4) - (y1-y2)*(x3-x4)

	// If den is 0, the lines are parallel
	if den == 0 {
		return nil
	}

	t := ((x1-x3)*(y3-y4) - (y1-y3)*(x3-x4)) / den
	u := -((x1-x2)*(y1-y3) - (y1-y2)*(x1-x3)) / den

	//lines intersect
	if t > 0 && t < 1 && u > 0 {
		v := resolv.NewVector(float64(x1+t*(x2-x1)), float64(y1+t*(y2-y1)))
		return &v
	}

	return nil	
}

func (r *Ray) LookAt(x, y float64) {
	
	r.dir.X = x - r.pos.X
	r.dir.Y = y - r.pos.Y

	length := r.dir.Magnitude()
	r.dir.X /= length
	r.dir.Y /= length
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {


	x, y := ebiten.CursorPosition()
	g.particle.pos.X = float64(x)
	g.particle.pos.Y = float64(y)


	//update rays
	for _, ray := range g.particle.rays {
		ray.pos = g.particle.pos
	}

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

	//Draw particle
	vector.DrawFilledCircle(
		screen,
		float32(g.particle.pos.X),
		float32(g.particle.pos.Y),
		ParticleRadius,
		color.RGBA{255, 255, 255, 255},
		false,
	)

	//cast rays
	for _, ray := range g.particle.rays {
		var closestWall *resolv.Vector
		for _, wall := range g.walls {
			pt := ray.Cast(wall)
			if pt != nil {
				if closestWall == nil {
					closestWall = pt
				} else {
					d1 := ray.pos.Distance(*pt)
					d2 := ray.pos.Distance(*closestWall)
					if d1 < d2 {
						closestWall = pt
					}
				}
			}
		}
		if closestWall != nil {
			vector.StrokeLine(
				screen,
				float32(ray.pos.X),
				float32(ray.pos.Y),
				float32(closestWall.X),
				float32(closestWall.Y),
				RayWidth,
				color.RGBA{255, 255, 255, 255},
				false,
			)
		}
	}

}
