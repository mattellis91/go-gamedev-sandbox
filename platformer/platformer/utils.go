package platformer

import (
	"fmt"
	"math"
)

// Various utilities functions

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func sign(x float64) float64 {
	if math.Signbit(x) {
		return -1.0
	} else {
		return 1.0
	}
}

func moveToward(start, stop, step float64) float64 {
	switch {
	case start < stop:
		if start+step >= stop {
			return stop
		}
		return start + step

	case start > stop:
		if start-step <= stop {
			return stop
		}
		return start - step

	default:
		return stop
	}
}

type Vector struct {
	x, y float64
}

func add(a, b Vector) Vector {
	return Vector{a.x + b.x, a.y + b.y}
}

func sub(a, b Vector) Vector {
	return Vector{a.x - b.x, a.y - b.y}
}

func dot(a, b Vector) float64 {
	return a.x*b.x + a.y*b.y
}
