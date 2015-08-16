package pyplot

import (
	"testing"
)

func TestPlot(t *testing.T) {
	xs := []float64{-2, -1, 0, 1, 2, 3, 4}
	ys1 := []float64{9, 4, 1, 0, 1, 4, 9}
	ys2 := []float64{0, 5, 8, 9, 8, 5, 0}
	Plot(xs, ys1, "r", xs, ys2, "g", Lw(3.0))
	Show()
}
