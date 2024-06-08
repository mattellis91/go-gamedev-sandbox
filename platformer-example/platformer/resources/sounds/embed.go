package sounds

import (
	_ "embed"
)

var (
	//go:embed jump.wav
	Jump_wav []byte

	//go:embed hurt.wav
	Hurt_wav []byte
) 
