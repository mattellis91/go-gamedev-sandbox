package tiles

import (
	_ "embed"
)

var (
	//go:embed tiles.png
	Tiles_png []byte
	//go:embed dungeon_tiles.png
	Dungeon_tiles_png []byte
)
