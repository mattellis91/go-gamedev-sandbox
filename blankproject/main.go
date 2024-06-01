package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mattellis91/go-gamedev-sandbox/blankproject/blankproject"
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

	ebiten.SetWindowSize(blankproject.ScreenWidth*2, blankproject.ScreenHeight*2)
	ebiten.SetWindowTitle("Blank Project (Ebitengine Demo)")
	if err := ebiten.RunGame(&blankproject.Game{}); err != nil {
		log.Fatal(err)
	}
}
