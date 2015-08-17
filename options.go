package pyplot

import (
	"fmt"
)

type optionFlag int
type funcOptions map[optionFlag]bool
type internalOption func(funcOptions) (string, bool)
type Option internalOption

const (
	optionFlagStart optionFlag = iota
	aa
	alpha
	animated
	antialiased
	axis
	b
	backgroundcolor
	barsabove
	basex
	basey
	bboxToAnchor
	borderaxespad
	borderpad
	c
	capsize
	capthick
	clipOn
	color
	columnspacing
	dashCapstyle
	dashJoinstyle
	dashes
	dpi
	ecolor
	edgecolor
	elinewidth
	errorevery
	facecolor
	family
	fancybox
	figsize
	fmtFlag
	fontname
	fontsize
	fontstyle
	fontweight
	framealpha
	frameon
	ha
	handlelength
	handletextpad
	horizontalalignment
	label
	labelspacing
	linestyle
	linewidth
	linespacing
	linscalex
	linscaley
	linthreshx
	linthreshy
	lolims
	ls
	lw
	loc
	lod
	marker
	markeredgecolor
	markeredgewidth
	markerfacecolor
	markerscale
	markersize
	markevery
	mec
	mew
	mfc
	mode
	ms
	multialignment
	name
	num
	ncol
	nonposx
	nonposy
	numpoints
	pickradius
	position
	rotation
	rotationMode
	scatterpoints
	scatteryoffsets
	shadow
	size
	solidCapstyle
	solidJoinstyle
	style
	subsx
	subsy
	text
	title
	uplims
	va
	variant
	verticalalignment
	visible
	weight
	which
	x
	xerr
	xmax
	xmin
	xdata
	xlolims
	xuplims
	y
	yerr
	ymax
	ymin
	ydata
	zorder
	optionFlagEnd
)

var optionNames = map[optionFlag]string {
	aa: "aa",
	alpha: "alpha",
	animated: "animated",
	antialiased: "antialiased",
	axis: "axis",
	b: "b",
	backgroundcolor: "backgroundcolor",
	basex: "basex",
	basey: "basey",
	barsabove: "barsabove",
	bboxToAnchor: "bbox_to_anchor",
	borderaxespad: "borderaxespad",
	borderpad: "borderpad",
	c: "c",
	capsize: "capsize",
	capthick: "capthick",
	clipOn: "clip_on",
	color: "color",
	columnspacing: "columnspacing",
	dashCapstyle: "dash_capstyle",
	dashJoinstyle: "dash_joinstyle",
	dashes: "dashes",
	dpi: "dpi",
	ecolor: "ecolor",
	edgecolor: "edgecolor",
	elinewidth: "elinewidth",
	errorevery: "errorevery",
	facecolor: "facecolor",
	family: "family",
	fancybox: "fancybox",
	figsize: "figsize",
	fmtFlag: "fmt",
	framealpha: "framealpha",
	frameon: "frameon",
	fontname: "fontname",
	fontsize: "fontsize",
	fontstyle: "fontstyle",
	fontweight: "fontweight",
	ha: "ha",
	handlelength: "handlelength",
	handletextpad: "handletextpad",
	horizontalalignment: "horizontalalignment",
	label: "label",
	labelspacing: "labelspacing",
	linespacing: "linespacing",
	linestyle: "linestyle",
	linewidth: "linewidth",
	linscalex: "linscalex",
	linscaley: "linscaley",
	linthreshx: "linthreshx",
	linthreshy: "linthreshy",
	ls: "ls",
	lw: "lw",
	loc: "loc",
	lod: "lod",
	lolims: "lolims",
	marker: "marker",
	markeredgecolor: "markeredgecolor",
	markeredgewidth: "markeredgewidth",
	markerfacecolor: "markerfacecolor",
	markerscale: "markerscale",
	markersize: "markersize",
	markevery: "markevery",
	mec: "mec",
	mew: "mew",
	mfc: "mfc",
	mode: "mode",
	ms: "ms",
	multialignment: "multialignment",
	name: "name",
	ncol: "ncol",
	nonposx: "nonposx",
	nonposy: "nonposy",
	num: "num",
	numpoints: "numpoints",
	pickradius: "pickradius",
	position: "position",
	rotation: "rotation",
	rotationMode: "rotation_mode",
	scatterpoints: "scatterpoints",
	scatteryoffsets: "scatteryoffsets",
	shadow: "shadow",
	size: "size",
	solidCapstyle: "solid_capstyle",
	solidJoinstyle: "solid_joinstyle",
	style: "style",
	subsx: "subsx",
	subsy: "subsy",
	text: "text",
	title: "title",
	uplims: "uplims",
	va: "va",
	variant: "variant",
	verticalalignment: "verticalalignment",
	visible: "visible",
	weight: "weight",
	which: "which",
	x: "x",
	xerr: "xerr",
	xmax: "xmax",
	xmin: "xmin",
	xdata: "xdata",
	xlolims: "xlolims",
	xuplims: "xuplims",
	y: "y",
	yerr: "yerr",
	ymax: "ymax",
	ymin: "ymin",
	ydata: "ydata",
	zorder: "zorder",
}

var optionFuncs = map[optionFlag]interface{} {
	aa: Aa,
	alpha: Alpha,
	animated: Animated,
	antialiased: Antialiased,
	axis: Axis,
	b: B,
	backgroundcolor: Backgroundcolor,
	barsabove: Barsabove,
	basex: Basex,
	basey: Basey,
	bboxToAnchor: Bboxtoanchor,
	borderaxespad: Borderaxespad,
	borderpad: Borderpad,
	c: C,
	capsize: Capsize,
	capthick: Capthick,
	clipOn: ClipOn,
	color: Color,
	columnspacing: Columnspacing,
	dashCapstyle: DashCapstyle,
	dashJoinstyle: DashJoinstyle,
	dashes: Dashes,
	dpi: Dpi,
	ecolor: Ecolor,
	edgecolor: Edgecolor,
	elinewidth: Elinewidth,
	errorevery: Errorevery,
	facecolor: Facecolor,
	family: Family,
	fancybox: Fancybox,
	figsize: Figsize,
	fmtFlag: Fmt,
	fontname: Fontname,
	fontsize: Fontsize,
	fontstyle: Fontstyle,
	fontweight: Fontweight,
	framealpha: Framealpha,
	frameon: Frameon,
	ha: Ha,
	handlelength: Handlelength,
	handletextpad: Handletextpad,
	horizontalalignment: Horizontalalignment,
	label: Label,
	labelspacing: Labelspacing,
	linespacing: Linespacing,
	linestyle: Linestyle,
	linewidth: Linewidth,
	linscalex: Linscalex,
	linscaley: Linscaley,
	linthreshx: Linthreshx,
	linthreshy: Linthreshy,
	ls: Ls,
	lw: Lw,
	loc: Loc,
	lod: Lod,
	lolims: Lolims,
	marker: Marker,
	markeredgecolor: Markeredgecolor,
	markeredgewidth: Markeredgewidth,
	markerfacecolor: Markerfacecolor,
	markerscale: Markerscale,
	markersize: Markersize,
	markevery: Markevery,
	mec: Mec,
	mew: Mew,
	mfc: Mfc,
	mode: Mode,
	ms: Ms,
	multialignment: Multialignment,
	name: Name,
	ncol: Ncol,
	num: Num,
	numpoints: Numpoints,
	nonposx: Nonposx,
	nonposy: Nonposy,
	pickradius: Pickradius,
	position: Position,
	rotation: Rotation,
	rotationMode: RotationMode,
	scatterpoints: Scatterpoints,
	scatteryoffsets: Scatteryoffsets,
	shadow: Shadow,
	size: Size,
	solidCapstyle: SolidCapstyle,
	solidJoinstyle: SolidJoinstyle,
	subsx: Subsx,
	subsy: Subsy,
	style: Style,
	text: Text,
	title: Title,
	uplims: Uplims,
	va: Va,
	variant: Variant,
	verticalalignment: Verticalalignment,
	visible: Visible,
	weight: Weight,
	which: Which,
	x: X,
	xerr: Xerr,
	xdata: Xdata,
	xlolims: Xlolims,
	xmax: Xmax,
	xmin: Xmin,
	xuplims: Xuplims,
	y: Y,
	yerr: Yerr,
	ymax: Ymax,
	ymin: Ymin,
	ydata: Ydata,
	zorder: Zorder,
}

func checkOptionNames() {
	for i := optionFlagStart + 1; i < optionFlagEnd; i++ {
		if _, ok := optionNames[i]; !ok {
			panic(fmt.Sprintf("option %d does not have a name", i))
		}
	}
}

func checkOptionFuncs() {
	for i := optionFlagStart + 1; i < optionFlagEnd; i++ {
		if _, ok := optionFuncs[i]; !ok {
			panic(fmt.Sprintf("option %d does not have a function", i))
		}
	}
}

func optionStart(args []interface{}) int {
	for i, intr := range args {
		switch intr.(type) {
		case Option: return i
		}
	}
	return len(args)
}

func singletonOption(val interface{}, t pyplotType, f optionFlag) Option {
	return func(fo funcOptions) (string, bool) {
		if _, ok := fo[f]; ok {
			str, ok := convertType(val, t)
			// TODO: better error message
			if !ok {
				panic(fmt.Sprintf("Value %v has invalid type for Option %s.",
					val, optionNames[f]))
			}
			return fmt.Sprintf("%s=%s", optionNames[f], str), true
		} else {
			return "", false
		}
	}
}

func Aa(val bool) Option {
	return singletonOption(val, Bool, aa)
}

func Alpha(val float64) Option {
	return singletonOption(val, Number, alpha)
}

func Animated(val bool) Option {
	return singletonOption(val, Bool, animated)
}

func Antialiased(val bool) Option {
	return singletonOption(val, Bool, antialiased)
}

func Axis(val string) Option {
	if val != "both" && val != "x" && val != "y" {
		panic("Invalid value for Axis Option.")
	}

	return singletonOption(val, String, axis)
}

func B(intr interface{}) Option {
	return singletonOption(intr, NoneBool, b)
}

func Backgroundcolor(val string) Option {
	return singletonOption(val, String, backgroundcolor)
}

func Barsabove(val bool) Option {
	return singletonOption(val, Bool, barsabove)
}

func Basex(val float64) Option {
	return singletonOption(val, Number, basex)
}

func Basey(val float64) Option {
	return singletonOption(val, Number, basey)
}

func Bboxtoanchor(x, y float64) Option {
	return func(fo funcOptions) (string, bool) {
		if _, ok := fo[bboxToAnchor]; ok { return "", false }
		return fmt.Sprintf("(%g, %g)", x, y), true
	}
}

func Borderaxespad(intr interface{}) Option {
	return singletonOption(intr, NoneNumber, borderaxespad)
}

func Borderpad(intr interface{}) Option {
	return singletonOption(intr, NoneNumber, borderpad)
}

func C(val string) Option {
	return singletonOption(val, String, c)
}

func Capsize(val float64) Option {
	return singletonOption(val, Number, capsize)
}

func Capthick(val float64) Option {
	return singletonOption(val, Number, capthick)
}

func ClipOn(val bool) Option {
	return singletonOption(val, Bool, clipOn)
}

func Color(val string) Option {
	return singletonOption(val, String, color)
}

func Columnspacing(intr interface{}) Option {
	return singletonOption(intr, NoneNumber, columnspacing)
}

func DashCapstyle(val string) Option {
	if val != "butt" && val != "round" && val != "projecting" {
		// TODO: better error message.
		panic("Invalid value for DashCapstyle Option.")
	}
	return singletonOption(val, String, dashCapstyle)
}

func DashJoinstyle(val string) Option {
	if val != "miter" && val != "round" && val != "bevel" {
		// TODO: better error message.
		panic("Invalid value for DashJoinstyle Option.")
	}
	return singletonOption(val, String, dashJoinstyle)
}

func Dashes(vals []float64) Option {
	return singletonOption(vals, Array, dashes)
}

func Dpi(intr interface{}) Option {
	return singletonOption(intr, NoneInt, dpi)
}

func Ecolor(intr interface{}) Option {
	return singletonOption(intr, NoneString, ecolor)
}

func Edgecolor(val string) Option {
	return singletonOption(val, String, edgecolor)
}

func Elinewidth(val float64) Option {
	return singletonOption(val, Number, elinewidth)
}

func Errorevery(val int) Option {
	return singletonOption(val, Int, errorevery)
}

func Facecolor(val string) Option {
	return singletonOption(val, String, facecolor)
}

func Family(val string) Option {
	if val != "serif" && val != "sans-serif" && val != "cursive" &&
		val != "fantasy" && val != "monospace" {
		panic("Invalid value for Family Option.")
	}
	return singletonOption(val, String, family)
}

func Fancybox(intr interface{}) Option {
	return singletonOption(intr, NoneBool, fancybox)
}

func Figsize(x, y int) Option {
	return func(fo funcOptions) (string, bool) {
		if _, ok := fo[figsize]; !ok { return "", false }
		return fmt.Sprintf("%s=(%d,%d)", optionNames[figsize], x, y), true
	}
}

func Fmt(intr interface{}) Option {
	return singletonOption(intr, NoneString, fmtFlag)
}

func Fontname(val string) Option {
	return singletonOption(val, String, fontname)
}

func Fontsize(val interface{}) Option {
	if _, ok := convertNumber(val); ok {
		return singletonOption(val, Number, fontsize)
	} else if _, ok := convertString(val); ok {
		return singletonOption(val, String, fontsize)
	}
	panic("Fonsize Option only accepts strings and numbers.")
}

func Fontstyle(val string) Option {
	if val != "normal" && val != "italic" && val != "oblique" {
		panic("Invalid value for Fontstyle Option.")
	}
	return singletonOption(val, String, fontstyle)
}

func Fontweight(val string) Option {
	if val != "normal" && val != "bold" && val != "heavy" &&
		val != "light" && val != "ultrabold" && val != "ultralight" {
		panic("Invalid value for Fontweight Option.")
	}

	return singletonOption(val, String, fontweight)
}

func Framealpha(intr interface{}) Option {
	return singletonOption(intr, NoneNumber, framealpha)
}

func Frameon(intr interface{}) Option {
	return singletonOption(intr, NoneBool, frameon)
}

func Ha(val string) Option {
	if val != "center" && val != "right" && val != "left" {
		panic("Invalid value for Ha option.")
	}
	return singletonOption(val, String, ha)
}

func Handlelength(intr interface{}) Option {
	return singletonOption(intr, NoneNumber, handlelength)
}

func Handletextpad(intr interface{}) Option {
	return singletonOption(intr, NoneNumber, handletextpad)
}

func Horizontalalignment(val string) Option {
	if val != "center" && val != "right" && val != "left" {
		panic("Invalid value for Horizontalalignment option.")
	}
	return singletonOption(val, String, horizontalalignment)
}

func Label(val string) Option {
	singletonOption(val, String, label)
	return singletonOption(val, String, label)
}

func Labelspacing(intr interface{}) Option {
	return singletonOption(intr, NoneNumber, labelspacing)
}

func Linespacing(val float64) Option {
	return singletonOption(val, Number, linespacing)
}

func Linestyle(val string) Option {
	return singletonOption(val, String, linestyle)
}

func Linscalex(val float64) Option { 
	return singletonOption(val, Number, linscalex)
}

func Linscaley(val float64) Option { 
	return singletonOption(val, Number, linscaley)
}

func Linthreshx(val float64) Option {
	return singletonOption(val, Number, linthreshx)
}

func Linthreshy(val float64) Option {
	return singletonOption(val, Number, linthreshy)
}

func Lolims(intr interface{}) Option {
	if _, ok := convertBool(intr); ok {
		return singletonOption(intr, Bool, lolims)
	} else if _, ok := convertBoolArray(intr); ok {
		return singletonOption(intr, BoolArray, lolims)
	}
	panic("Invalid type for arugment to Lolims Option.")
}

func Ls(val string) Option {
	return singletonOption(val, String, ls)
}

func Linewidth(val string) Option {
	return singletonOption(val, String, linewidth)
}

func Lw(val float64) Option {
	return singletonOption(val, Number, lw)
}

func Loc(intr interface{}) Option {
	if _, ok := convertInt(intr); ok {
		return singletonOption(intr, Int, loc)
	} else if _, ok := convertString(intr); ok {
		return singletonOption(intr, String, loc)
	}
	panic("Option Loc only accepts ints and strings.")
}

func Lod(val bool) Option {
	return singletonOption(val, Bool, lod)
}

func Marker(val string) Option {
	if val != "+" && val != "," && val != "." &&
		val != "1" && val != "2" && val != "3" && val != "4" {
		panic("Invalid value for Marker Option.")
	}
	return singletonOption(val, String, marker)
}

func Markeredgecolor(val string) Option {
	return singletonOption(val, String, markeredgecolor)
}

func Markeredgewidth(val float64) Option {
	return singletonOption(val, Number, markeredgewidth)
}

func Markerfacecolor(val string) Option {
	return singletonOption(val, String, markerfacecolor)
}

func Markerscale(intr interface{}) Option {
	return singletonOption(intr, NoneNumber, markerscale)
}

func Markersize(val float64) Option {
	return singletonOption(val, Number, markersize)
}

func Markevery(val float64) Option {
	panic("NYI")
}

func Mec(val string) Option {
	return singletonOption(val, String, mec)
}

func Mew(val float64) Option {
	return singletonOption(val, Number, mew)
}

func Mfc(val string) Option {
	return singletonOption(val, String, mfc)
}

func Mode(intr interface{}) Option {
	return singletonOption(intr, NoneString, mode)
}

func Ms(val float64) Option {
	return singletonOption(val, Number, markersize)
}

func Multialignment(val string) Option {
	if val != "left" && val != "center" && val != "right" {
		panic("Invalid value for Multialignment Option.")
	}
	return singletonOption(val, String, multialignment)
}

func Name(val string) Option {
	return singletonOption(val, String, name)
}

func Ncol(val int) Option {
	return singletonOption(val, Int, ncol)
}

func Nonposx(val string) Option {
	if val != "mask" && val != "clip" {
		panic("Invalid value for Nonposx Option")
	}
	return singletonOption(val, String, nonposx)
}

func Nonposy(val string) Option {
	if val != "mask" && val != "clip" {
		panic("Invalid value for Nonposy Option")
	}
	return singletonOption(val, String, nonposy)
}

func Num(intr interface{}) Option {
	if _, ok := convertNumber(intr); ok {
		return singletonOption(intr, Int, num)
	} else if _, ok := convertString(intr); ok {
		return singletonOption(intr, String, num)
	}
	panic("Num Option only accepts numbers and strings.")
}

func Numpoints(intr interface{}) Option {
	return singletonOption(intr, NoneInt, numpoints)
}

func Pickradius(val float64) Option {
	return singletonOption(val, Number, pickradius)
}

func Position(x, y float64) Option {
	return func(fo funcOptions) (string, bool) {
		if _, ok := fo[position]; ok {
			return fmt.Sprintf("position=(%g, %g)", x, y), true
		} else {
			return "", false
		}
	}
}

func Rotation(intr interface{}) Option {
	if _, ok := convertNumber(intr); ok {
		return singletonOption(intr, Number, rotation)
	} else if _, ok := convertString(intr); ok {
		str, _ := intr.(string)
		if str != "vertical" && str != "horizontal" {
			panic("Invalid value for Rotation Option.")
		}
	}
	panic("Rotation Option only accepts numbers and strings.")
}

func RotationMode(intr interface{}) Option {
	return singletonOption(intr, NoneString, rotationMode)
}

func Scatterpoints(intr interface{}) Option {
	return singletonOption(intr, NoneInt, scatterpoints)
}

func Scatteryoffsets(vals []float64) Option {
	return singletonOption(vals, Array, scatteryoffsets)
}

func Shadow(intr interface{}) Option {
	return singletonOption(intr, NoneBool, shadow)
}

func Size(intr interface{}) Option {
	if _, ok := convertNumber(intr); ok {
		return singletonOption(intr, Number, size)
	} else if _, ok := convertString(intr); ok {
		return singletonOption(intr, String, size)
	}
	panic("Option Size only accepts number and strings.")
}

func SolidCapstyle(val string) Option {
	if val != "butt" && val != "round" && val != "projecting" {
		panic("Invalid value for SolidCapstyle Option.")
	}
	return singletonOption(val, String, solidCapstyle)
}

func SolidJoinstyle(val string) Option {
	if val != "miter" && val != "round" && val != "bevel" {
		panic("Invalid value for SolidJoinstyle Option.")
	}
	return singletonOption(val, String, solidJoinstyle)
}

func Style(val string) Option {
	if val != "normal" && val != "italic" && val != "oblique" {
		panic("Invalid value for Style option.")
	}
	return singletonOption(val, String, style)
}

func Subsx(vals []int) Option {
	return singletonOption(vals, Array, subsx)
}

func Subsy(vals []int) Option {
	return singletonOption(vals, Array, subsy)
}

func Text(val string) Option {
	return singletonOption(val, String, text)
}

func Uplims(intr interface{}) Option {
	if _, ok := convertBool(intr); ok {
		return singletonOption(intr, Bool, uplims)
	} else if _, ok := convertBoolArray(intr); ok {
		return singletonOption(intr, BoolArray, uplims)
	}
	panic("Invalid type for arugment to Uplims Option.")
}

func Va(val string) Option {
	if val != "center" && val != "top" &&
		val != "bottom" && val != "baseline" {
		panic("Invalid value for Va Option.")
	}

	return singletonOption(val, String, va)
}

func Variant(val string) Option {
	if val != "notmal" && val != "small-caps" {
		panic("Invalid value for Variant Option.")
	}
	return singletonOption(val, String, variant)
}

func Verticalalignment(val string) Option {
	if val != "center" && val != "top" &&
		val != "bottom" && val != "baseline" {
		panic("Invalid value for Verticalalignment Option.")
	}

	return singletonOption(val, String, verticalalignment)
}

func Visible(val bool) Option {
	return singletonOption(val, Bool, visible)
}

func Weight(val string) Option {
	if val != "normal" && val != "bold" && val != "heavy" &&
		val != "light" && val != "ultrabold" && val != "ultralight" {
		panic("Invalid value for Weight Option.")
	}

	return singletonOption(val, String, weight)
}

func Which(val string) Option {
	if val != "major" && val != "minor" && val != "both" {
		panic("Invalid value for Which Option.")
	}
	return singletonOption(val, String, which)
}

func X(val float64) Option {
	return singletonOption(val, Number, x)
}

func Xdata(vals []float64) Option {
	return singletonOption(vals, Array, xdata)
}

func Xerr(vals []float64) Option {
	return singletonOption(vals, Array, xerr)
}

func Xlolims(intr interface{}) Option {
	if _, ok := convertBool(intr); ok {
		return singletonOption(intr, Bool, xlolims)
	} else if _, ok := convertBoolArray(intr); ok {
		return singletonOption(intr, BoolArray, xlolims)
	}
	panic("Invalid type for arugment to Xlolims Option.")
}

func Xmax(intr interface{}) Option {
	return singletonOption(intr, NoneNumber, xmax)
}

func Xmin(intr interface{}) Option {
	return singletonOption(intr, NoneNumber, xmin)
}

func Xuplims(intr interface{}) Option {
	if _, ok := convertBool(intr); ok {
		return singletonOption(intr, Bool, xuplims)
	} else if _, ok := convertBoolArray(intr); ok {
		return singletonOption(intr, BoolArray, xuplims)
	}
	panic("Invalid type for arugment to Xuplims Option.")
}

func Y(val float64) Option {
	return singletonOption(val, Number, y)
}

func Ydata(vals []float64) Option {
	return singletonOption(vals, Array, ydata)
}

func Yerr(vals []float64) Option {
	return singletonOption(vals, Array, yerr)
}

func Ymax(intr interface{}) Option {
	return singletonOption(intr, NoneNumber, ymax)
}

func Ymin(intr interface{}) Option {
	return singletonOption(intr, NoneNumber, ymin)
}

func Zorder(val float64) Option {
	return singletonOption(val, Number, zorder)
}
