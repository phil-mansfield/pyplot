package pyplot

import (
	"testing"
)

func TestPlot(t *testing.T) {
	xs := []float64{-2, -1, 0, 1, 2, 3, 4}
	ys := []float64{9, 4, 1, 0, 1, 4, 9}
	Plot(xs, ys, "--r")
}
