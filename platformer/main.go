package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mattellis91/go-gamedev-sandbox/platformer/platformer"
)

var cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal(err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()
	}

	ebiten.SetWindowSize(platformer.ScreenWidth*2, platformer.ScreenHeight*2)
	ebiten.SetWindowTitle("Platformer Project")
	if err := ebiten.RunGame(&platformer.Game{}); err != nil {
		log.Fatal(err)
	}
}
