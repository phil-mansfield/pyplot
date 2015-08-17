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

	Figure(Num(13), Facecolor("m"))
	Plot(xs, ys1, "r", xs, ys2, "g", Lw(3.0), Label(`$\frac{1}{2}$`))
	Errorbar(xs, []float64{3, 3, 3, 3, 3, 3, 3}, Xerr(errs), Yerr(errs),
		C("c"), Lolims(lims), Label("My Errors"))
	Title("My Test Plot", Fontsize(16), Loc("right"))
	Xlabel("$t$ [seconds]", Name("Sans"))
	Ylabel(`$2 \times\ \pi / {\rm meow}$`)
	Ylim(nil, 11)
	Xscale("symlog", Subsx([]int{2, 3, 4, 5, 6, 7, 8, 9}))
	Grid(Which("minor"), C("b"))
	Legend(Fancybox(true))
	Show()
}
