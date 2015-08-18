package pyplot

import (
	"fmt"
	"strings"
)

var (
	errorbarOptions = make(map[optionFlag]bool)
	figureOptions = make(map[optionFlag]bool)
	gridOptions = make(map[optionFlag]bool)
	labelOptions = make(map[optionFlag]bool)
	legendOptions = make(map[optionFlag]bool)
	plotOptions = make(map[optionFlag]bool)
	rcOptions = make(map[optionFlag]bool)
	titleOptions = make(map[optionFlag]bool)

	xlimOptions = make(map[optionFlag]bool)
	// xscale options
	xLinearScaleOptions = make(map[optionFlag]bool)
	xLogScaleOptions = make(map[optionFlag]bool)
	xSymlogScaleOptions = make(map[optionFlag]bool)

	ylimOptions = make(map[optionFlag]bool)
	// yscale options
	yLinearScaleOptions = make(map[optionFlag]bool)
	yLogScaleOptions = make(map[optionFlag]bool)
	ySymlogScaleOptions = make(map[optionFlag]bool)

)

func ErrorBar(xs, ys []float64, opts ...Option) {
	args := make([]string, 2)
	args[0], _ = convertArray(xs)
	args[1], _ = convertArray(ys)

	for _, opt := range opts {
		str, ok := opt(errorbarOptions)
		if !ok { panic("Invalid Option argument to Errorbar.") }
		args = append(args, str)
	}

	line := fmt.Sprintf("plt.errorbar(%s)", strings.Join(args, ","))
	lines = append(lines, line)
}

func Figure(opts ...Option) {
	args := make([]string, 0)
	
	for _, opt := range opts {
		str, ok := opt(figureOptions)
		if !ok { panic("Invalid Option argument for Figure.") }
		args = append(args, str)
	}

	line := fmt.Sprintf("plt.figure(%s)", strings.Join(args, ","))
	lines = append(lines, line)
}

func Grid(opts ...Option) {
	args := make([]string, 0)

	for _, opt := range opts {
		str, ok := opt(gridOptions)
		if !ok { panic("Invalid Option argument to Errorbar.") }
		args = append(args, str)
	}

	line := fmt.Sprintf("plt.grid(%s)", strings.Join(args, ","))
	lines = append(lines, line)
}

func Legend(opts ...Option) {
	args := make([]string, 0)
	for _, opt := range opts {
		str, ok := opt(legendOptions)
		if !ok { panic("Invalid Option argument for Legend.") }
		args = append(args, str)
	}

	line := fmt.Sprintf("plt.legend(%s)", strings.Join(args, ","))
	lines = append(lines, line)
}

func Plot(args ...interface{}) {
	plotArgs := []string{}
	optStart := optionStart(args)

	// TODO: not actually consistent with all valid pyplot vargs sets.
	for i := 0; i < optStart; i += 3 {
		plotArg, ok := convertType(args[i], Array)
		// TODO: Better errors messages!!!
		if !ok { panic(fmt.Sprintf("Invalid type for argument %d.", i+1)) }
		plotArgs = append(plotArgs, plotArg)
		if i + 1 < optStart {
			plotArg, ok = convertType(args[i+1], Array)
			if !ok { panic(fmt.Sprintf("Invalid type for argument %d.", i+2)) }
			plotArgs = append(plotArgs, plotArg)
		}
		if i + 2 < optStart {
			plotArg, ok = convertType(args[i+2], String)
			if !ok { panic(fmt.Sprintf("Invalid type for argument %d.", i+3)) }
			plotArgs = append(plotArgs, plotArg)
		}
	}

	for i := optStart; i < len(args); i++ {
		opt, ok := args[i].(Option)
		if !ok { panic("Non-Option argument given after Option argument.") }
		str, ok := opt(plotOptions)
		if !ok { panic("Invalid Option argument for Plot function.") }
		plotArgs = append(plotArgs, str)
	}

	line := fmt.Sprintf("plt.plot(%s)", strings.Join(plotArgs, ","))
	lines = append(lines, line)
}

func RC(group string, opts ...Option) {
	panic("NYI")
}

func Title(s string, opts ...Option) {
	args := make([]string, 1)
	args[0], _ = convertString(s)

	for _, opt := range opts {
		str, ok := opt(titleOptions)
		if !ok { panic("Invalid Option argument for Title function.") }
		args = append(args, str)
	}

	line := fmt.Sprintf("plt.title(%s)", strings.Join(args, ","))
	lines = append(lines, line)
}

func XLabel(s string, opts ...Option) {
	args := make([]string, 1)
	args[0], _ = convertString(s)

	for _, opt := range opts {
		str, ok := opt(labelOptions)
		if !ok { panic("Invalid Option argument for Xlabel function.") }
		args = append(args, str)
	}

	line := fmt.Sprintf("plt.xlabel(%s)", strings.Join(args, ","))
	lines = append(lines, line)
}

func XLim(args ...interface{}) {
	limArgs := make([]string, 0)

	switch len(args) {
	case 1:
		opt, ok := args[0].(Option)
		if !ok { panic("Single arguments to Xlim must be Options.") }
		str, ok := opt(xlimOptions)
		if !ok { panic("Invalid Option argument for Xlim function.") }
		limArgs = append(limArgs, str)
	case 2:
		strX, ok := convertType(args[0], NoneNumber)
		if !ok { panic("Double arguments to Xlim must be a number or nil.") }
		limArgs = append(limArgs, strX)
		strY, ok := convertType(args[1], NoneNumber)
		if !ok { panic("Double arguments to Xlim must be a number or nil.") }
		limArgs = append(limArgs, strY)
	default:
		panic("Invalid number of args to Xlim.")
	}

	line := fmt.Sprintf("plt.xlim(%s)", strings.Join(limArgs, ","))
	lines = append(lines, line)
}

func XScale(scale string, opts ...Option) {
	args := make([]string, 1)
	var fo funcOptions
	switch scale {
	case "linear": fo = xLinearScaleOptions
	case "log": fo = xLogScaleOptions
	case "symlog": fo = xSymlogScaleOptions
	default:
		panic("Invalid value for scale parameter to Xscale.")
	}

	args[0], _ = convertString(scale)

	for _, opt := range opts {
		str, ok := opt(fo)
		if !ok { panic("Invalid Option argument for Xscale function") }
		args = append(args, str)
	}

	line := fmt.Sprintf("plt.xscale(%s)", strings.Join(args, ","))
	lines = append(lines, line)
}

func YLabel(s string, opts ...Option) {
	args := make([]string, 1)
	args[0], _ = convertString(s)

	for _, opt := range opts {
		str, ok := opt(labelOptions)
		if !ok { panic("Invalid Option argument for Ylabel function.") }
		args = append(args, str)
	}

	line := fmt.Sprintf("plt.ylabel(%s)", strings.Join(args, ","))
	lines = append(lines, line)
}

func YLim(args ...interface{}) {
	limArgs := make([]string, 0)

	switch len(args) {
	case 1:
		opt, ok := args[0].(Option)
		if !ok { panic("Single arguments to Ylim must be Options.") }
		str, ok := opt(ylimOptions)
		if !ok { panic("Invalid Option argument for Ylim function.") }
		limArgs = append(limArgs, str)
	case 2:
		strX, ok := convertType(args[0], NoneNumber)
		if !ok { panic("Double arguments to Ylim must be a number or nil.") }
		limArgs = append(limArgs, strX)
		strY, ok := convertType(args[1], NoneNumber)
		if !ok { panic("Double arguments to Ylim must be a number or nil.") }
		limArgs = append(limArgs, strY)
	default:
		panic("Invalid number of args to Ylim.")
	}

	line := fmt.Sprintf("plt.ylim(%s)", strings.Join(limArgs, ","))
	lines = append(lines, line)
}

func YScale(scale string, opts ...Option) {
	args := make([]string, 1)
	var fo funcOptions
	switch scale {
	case "linear": fo = yLinearScaleOptions
	case "log": fo = yLogScaleOptions
	case "symlog": fo = ySymlogScaleOptions
	default:
		panic("Invalid value for scale parameter to Xscale.")
	}

	args[0], _ = convertString(scale)

	for _, opt := range opts {
		str, ok := opt(fo)
		if !ok { panic("Invalid Option argument for Xscale function") }
		args = append(args, str)
	}

	line := fmt.Sprintf("plt.yscale(%s)", strings.Join(args, ","))
	lines = append(lines, line)
}

func register(fo funcOptions, flags ...optionFlag) {
	for _, flag := range flags { fo[flag] = true }
}

func init() {
	checkOptionNames()
	checkOptionFuncs()

	textOptions := []optionFlag{
		alpha, animated, backgroundcolor, clipOn, color, family,
		horizontalalignment, ha, label, linespacing, lod, multialignment,
		name, fontname, position, rotation, rotationMode, size, fontsize,
		style, fontstyle, text, variant, verticalalignment, va, visible,
		weight, fontweight, x, y, zorder,
	}

	line2DOptions := []optionFlag{
		alpha, animated, antialiased, aa, clipOn, color, c, dashCapstyle,
		dashJoinstyle, label, linestyle, ls, linewidth, lw, lod, marker,
		markeredgecolor, mec, markeredgewidth, mew, markerfacecolor, mfc,
		markersize, ms, markevery, pickradius, solidCapstyle, solidJoinstyle,
		visible, xdata, ydata, zorder,
	}

	register(errorbarOptions, append([]optionFlag{yerr, xerr, fmtFlag, ecolor,
		elinewidth, capsize, capthick, barsabove, lolims, uplims, xlolims,
		xuplims, errorevery}, line2DOptions...)...)
	register(figureOptions, num, figsize, dpi, facecolor, edgecolor)
	register(gridOptions, append([]optionFlag{b, which, axis},
		line2DOptions...)...)
	register(labelOptions, textOptions...)
	register(legendOptions, loc, bboxToAnchor, ncol, fontsize, numpoints,
		scatterpoints, scatteryoffsets, markerscale, frameon, fancybox, shadow,
		framealpha, mode, title, borderpad, labelspacing, handlelength,
		handletextpad, borderaxespad, columnspacing)
	register(plotOptions, line2DOptions...)
	register(titleOptions, append([]optionFlag{loc}, textOptions...)...)

	register(xlimOptions, xmin, xmax)
	register(xlimOptions, ymin, ymax)

	register(xLogScaleOptions, basex, nonposx, subsx)
	register(xSymlogScaleOptions, basex, nonposx, subsx, linscalex)
	register(yLogScaleOptions, basey, nonposy, subsy)
	register(ySymlogScaleOptions, basey, nonposy, subsy, linscaley)
}
