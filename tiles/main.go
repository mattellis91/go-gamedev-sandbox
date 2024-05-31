package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 240
	screenHeight = 240
)

const (
	tileSize = 16
)

var (
	tilesImage *ebiten.Image
)

func init() {
	// Decode an image from the image file's byte slice.
	img, _, err := ebitenutil.NewImageFromFile("tiles.png")
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = img
}

type Game struct {
	layers [][]int
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	w := tilesImage.Bounds().Dx()
	tileXCount := w / tileSize

	// Draw each tile with each DrawImage call.
	// As the source images of all DrawImage calls are always same,
	// this rendering is done very efficiently.
	// For more detail, see https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#Image.DrawImage
	const xCount = screenWidth / tileSize
	for _, l := range g.layers {
		for i, t := range l {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%xCount)*tileSize), float64((i/xCount)*tileSize))

			sx := (t % tileXCount) * tileSize
			sy := (t / tileXCount) * tileSize
			screen.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
		}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	g := &Game{
		layers: [][]int{
			{
				64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64,
				64, 0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  64,
				64, 0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  64,
				64, 0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  64,
				64, 0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  64,
				64, 0,  0,  0,  0,  64,  0,  64,  0,  64,  0,  0,  0,  0,  64,
				64, 0,  0,  0,  0,  0,  64,  64,  64,  0,  0,  0,  0,  0,  64,
				64, 0,  0,  0,  0,  0,  64,  64,  64,  0,  0,  0,  0,  0,  64,
				64, 0,  0,  0,  0,  64,  0,  64,  0,  64,  0,  0,  0,  0,  64,
				64, 0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  64,
				64, 0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  64,
				64, 0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  64,
				64, 0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  64,
				64, 0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  64,
				64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64,
			},
		},
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Tiles (Ebitengine Demo)")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}