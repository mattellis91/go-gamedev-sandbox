package platformer

import (
	"github.com/hajimehoshi/ebiten/v2"
)


type Player struct {

}

func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) Update() {

}

func (p *Player) Draw(screen *ebiten.Image) {
	println("Drawing player")
}

