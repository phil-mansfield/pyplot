package pyplot

import (
	"testing"
)

func TestPlot(t *testing.T) {
	xs := []float64{-2, -1, 0, 1, 2, 3, 4}
	ys1 := []float64{9, 4, 1, 0, 1, 4, 9}
	ys2 := []float64{0, 5, 8, 9, 8, 5, 0}
	errs := []float64{1, 1, 1, 1, 1, 1, 1}
	lims := []bool{true, false, true, false, true, false, false}

	Figure(Num(13), FaceColor("m"))
	Plot(xs, ys1, "r", xs, ys2, "g", LW(3.0), Label(`$\frac{1}{2}$`))
	ErrorBar(xs, []float64{3, 3, 3, 3, 3, 3, 3}, XErr(errs), YErr(errs),
		C("c"), LoLims(lims), Label("My Errors"))
	Title("My Test Plot", FontSize(16), Loc("right"))
	XLabel("$t$ [seconds]", Name("Sans"))
	YLabel(`$2 \times\ \pi / {\rm meow}$`)
	YLim(nil, 11)
	XScale("symlog", SubsX([]int{2, 3, 4, 5, 6, 7, 8, 9}))
	Grid(Which("minor"), C("b"))
	Legend(FancyBox(true))
	Show()
}
