package level

import (
	"fmt"
	"math"

	"github.com/mattellis91/go-gamedev-sandbox/platformer-example/platformer/util"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/solarlune/resolv"
)

type Player struct {
	image *ebiten.Image
	object *resolv.Object
	space *resolv.Space
	onGround *resolv.Object
	slidingOnWall *resolv.Object
	FacingRight bool
	IgnorePlatform *resolv.Object
	Speed *resolv.Vector
}

func NewPlayer(image *ebiten.Image, x, y float64, space *resolv.Space) *Player {
	fmt.Printf("space: %v\n", space)

	playerObject := resolv.NewObject(x, y, util.TileSize, util.TileSize)

	space.Add(playerObject)
	
	return &Player{
		image: image,
		object: playerObject,
		space: space,
		FacingRight: true,
		Speed: &resolv.Vector{X: 0, Y: 0},
	}
}

func (p *Player) Update() error{

	p.Speed.Y += Gravity

	if p.slidingOnWall != nil && p.Speed.Y > 1 {
		p.Speed.Y = 1
	}

	// Horizontal movement is only possible when not wallsliding.
	if p.slidingOnWall == nil {
		if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.GamepadAxisValue(0,0) > 0.1 {
			p.Speed.X += Accel
			p.FacingRight = true
		}
		if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.GamepadAxisValue(0,0) < -0.1 {
			p.Speed.X -= Accel
			p.FacingRight = false
		}
	}

	//Apply friction and horizontal speed limit.
	if p.Speed.X > Friction {
		p.Speed.X -= Friction
	} else if p.Speed.X < -Friction {
		p.Speed.X += Friction
	} else {
		p.Speed.X = 0
	}

	if p.Speed.X > MaxSpeed {
		p.Speed.X = MaxSpeed
	} else if p.Speed.X < -MaxSpeed {
		p.Speed.X = -MaxSpeed
	}

	//check for jumping
	if inpututil.IsKeyJustPressed(ebiten.KeyX) || inpututil.IsGamepadButtonJustPressed(0,0) {

		//drop through platforms
		if(false)  { //TODO: check for platform below
			
		} else {
			if p.onGround != nil {
				p.Speed.Y = -JumpSpeed
			} else if p.slidingOnWall != nil {
				//Jump off wall
				p.Speed.Y = -JumpSpeed

				if p.slidingOnWall.Position.X > p.object.Position.X {
					p.Speed.X = -4
				} else {
					p.Speed.X = 4
				}

				p.slidingOnWall = nil
			}
		}
	}

	dx := p.Speed.X


	if check := p.object.Check(p.Speed.X, 0, util.SolidTileSpaceIdentifier); check != nil {
		
		dx = check.ContactWithCell(check.Cells[0]).X
		p.Speed.X = 0

		if p.onGround == nil {
			p.slidingOnWall = check.Objects[0]
		}

	}

	p.object.Position.X += dx

	p.onGround = nil

	dy := p.Speed.Y
	dy = math.Max(math.Min(dy, util.TileSize), -util.TileSize)

	checkDistance := dy
	if dy >= 0 {
		checkDistance++
	}

	fmt.Printf("checkDistance: %v\n", checkDistance)

	if check := p.object.Check(0, checkDistance, util.SolidTileSpaceIdentifier, "platform"); check != nil {

		fmt.Printf("check: %v\n", check.Objects[0])

		slide , slideOk := check.SlideAgainstCell(check.Cells[0], util.SolidTileSpaceIdentifier)

		if dy < 0 && check.Cells[0].ContainsTags(util.SolidTileSpaceIdentifier) && slideOk && math.Abs(slide.X) <= util.TileSize / 2 {
			
			p.object.Position.X += slide.X

		} else {

			if solids := check.ObjectsByTags(util.SolidTileSpaceIdentifier); len(solids) > 0 && (p.onGround == nil || p.onGround.Position.Y >= solids[0].Position.Y) {
				
				dy = check.ContactWithObject(solids[0]).Y
				p.Speed.Y = 0

				if solids[0].Position.Y > p.object.Position.Y {
					p.onGround = solids[0]
				}
			}

			if p.onGround != nil {
				p.slidingOnWall = nil
			}
		}
	}

	p.object.Position.Y += dy

	wallNext := 1.0
	if !p.FacingRight {
		wallNext = -1.0
	}

	if c := p.object.Check(wallNext, 0, util.SolidTileSpaceIdentifier); p.slidingOnWall != nil && c == nil {
		p.slidingOnWall = nil
	}

	p.object.Update()

	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Translate(float64(p.object.Position.X), float64(p.object.Position.Y))
	screen.DrawImage(p.image, ops)
}