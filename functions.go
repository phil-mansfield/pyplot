package pyplot

import (
	"fmt"
	"strings"
)

var (
	plotOptions = make(map[optionFlag]bool)
	labelOptions = make(map[optionFlag]bool)
	titleOptions = make(map[optionFlag]bool)
	linearScaleOptions = make(map[optionFlag]bool)
	logScaleOptions = make(map[optionFlag]bool)
	symlogScaleOptions = make(map[optionFlag]bool)
	xlimOptions = make(map[optionFlag]bool)
	ylimOptions = make(map[optionFlag]bool)
)

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

func Xlabel(s string, opts ...Option) {
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

func Xlim(args ...interface{}) {
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

func Ylabel(s string, opts ...Option) {
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

func Ylim(args ...interface{}) {
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

	register(plotOptions, line2DOptions...)
	register(labelOptions, textOptions...)
	register(titleOptions, append([]optionFlag{loc}, textOptions...)...)

	register(xlimOptions, xmin, xmax)
	register(xlimOptions, ymin, ymax)

	register(logScaleOptions, basex, basey, nonposx, nonposy, subsx, subsy)
	register(symlogScaleOptions, basex, basey, nonposx, nonposy,
		subsx, subsy, linscalex, linscaley)

}
