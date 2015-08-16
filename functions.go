package pyplot

import (
	"fmt"
	"strings"
)

var plotOptions = make(map[optionFlag]bool)
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

	for i := optStart; i < len(args); i += 2 {
		opt, ok := args[i].(Option)
		if !ok { panic("Non-Option argument given after Option argument.") }
		str, ok := opt(plotOptions)
		if !ok { panic("Invalid Option argument for Plot function.") }
		plotArgs = append(plotArgs, str)
	}

	line := strings.Join(
		[]string{"plt.plot(", strings.Join(plotArgs, ","), ")"}, "",
	)
	lines = append(lines, line)
}

func register(fo funcOptions, flags ...optionFlag) {
	for _, flag := range flags { fo[flag] = true }
}

func init() {
	register(plotOptions,
		alpha, animated, antialiased, aa, clipOn, color, c, dashCapstyle,
		dashJoinstyle, data, label, linestyle, ls, linewidth, lw, lod, marker,
		markeredgecolor, mec, markeredgewidth, mew, markerfacecolor, mfc,
		markersize, ms, markevery, pickradius, solidCapstyle, solidJoinstyle,
		visible, xdata, ydata, zorder,
	)
}
