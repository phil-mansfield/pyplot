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
	backgroundcolor
	basex
	basey
	c
	clipOn
	color
	dashCapstyle
	dashJoinstyle
	dashes
	family
	fontname
	fontsize
	fontstyle
	fontweight
	ha
	horizontalalignment
	label
	linestyle
	linewidth
	linespacing
	linscalex
	linscaley
	linthreshx
	linthreshy
	ls
	lw
	loc
	lod
	marker
	markeredgecolor
	markeredgewidth
	markerfacecolor
	markersize
	markevery
	mec
	mew
	mfc
	ms
	multialignment
	name
	nonposx
	nonposy
	pickradius
	position
	rotation
	rotationMode
	size
	solidCapstyle
	solidJoinstyle
	style
	subsx
	subsy
	text
	va
	variant
	verticalalignment
	visible
	weight
	x
	xmax
	xmin
	xdata
	y
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
	backgroundcolor: "backgroundcolor",
	basex: "basex",
	basey: "basey",
	c: "c",
	clipOn: "clip_on",
	color: "color",
	dashCapstyle: "dash_capstyle",
	dashJoinstyle: "dash_joinstyle",
	dashes: "dashes",
	family: "family",
	fontname: "fontname",
	fontsize: "fontsize",
	fontstyle: "fontstyle",
	fontweight: "fontweight",
	ha: "ha",
	horizontalalignment: "horizontalalignment",
	label: "label",
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
	marker: "marker",
	markeredgecolor: "markeredgecolor",
	markeredgewidth: "markeredgewidth",
	markerfacecolor: "markerfacecolor",
	markersize: "markersize",
	markevery: "markevery",
	mec: "mec",
	mew: "mew",
	mfc: "mfc",
	ms: "ms",
	multialignment: "multialignment",
	name: "name",
	nonposx: "nonposx",
	nonposy: "nonposy",
	pickradius: "pickradius",
	position: "position",
	rotation: "rotation",
	rotationMode: "rotation_mode",
	size: "size",
	solidCapstyle: "solid_capstyle",
	solidJoinstyle: "solid_joinstyle",
	style: "style",
	subsx: "subsx",
	subsy: "subsy",
	text: "text",
	va: "va",
	variant: "variant",
	verticalalignment: "verticalalignment",
	visible: "visible",
	weight: "weight",
	x: "x",
	xmax: "xmax",
	xmin: "xmin",
	xdata: "xdata",
	y: "y",
	ymax: "ymax",
	ymin: "ymin",
	ydata: "ydata",
	zorder: "zorder",
}

var optionFuncs = map[optionFlag]interface{} {
	alpha: Alpha,
	animated: Animated,
	antialiased: Antialiased,
	aa: Aa,
	backgroundcolor: Backgroundcolor,
	basex: Basex,
	basey: Basey,
	c: C,
	clipOn: ClipOn,
	color: Color,
	dashCapstyle: DashCapstyle,
	dashJoinstyle: DashJoinstyle,
	dashes: Dashes,
	family: Family,
	fontname: Fontname,
	fontsize: Fontsize,
	fontstyle: Fontstyle,
	fontweight: Fontweight,
	ha: Ha,
	horizontalalignment: Horizontalalignment,
	label: Label,
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
	marker: Marker,
	markeredgecolor: Markeredgecolor,
	markeredgewidth: Markeredgewidth,
	markerfacecolor: Markerfacecolor,
	markersize: Markersize,
	markevery: Markevery,
	mec: Mec,
	mew: Mew,
	mfc: Mfc,
	ms: Ms,
	multialignment: Multialignment,
	name: Name,
	nonposx: Nonposx,
	nonposy: Nonposy,
	pickradius: Pickradius,
	position: Position,
	rotation: Rotation,
	rotationMode: RotationMode,
	size: Size,
	solidCapstyle: SolidCapstyle,
	solidJoinstyle: SolidJoinstyle,
	subsx: Subsx,
	subsy: Subsy,
	style: Style,
	text: Text,
	va: Va,
	variant: Variant,
	verticalalignment: Verticalalignment,
	visible: Visible,
	weight: Weight,
	x: X,
	xmax: Xmax,
	xmin: Xmin,
	xdata: Xdata,
	y: Y,
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
			if !ok { panic("Invalid type for Option.") }
			return fmt.Sprintf("%s=%s", optionNames[f], str), true
		} else {
			return "", false
		}
	}
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

func Aa(val bool) Option {
	return singletonOption(val, Bool, aa)
}

func Backgroundcolor(val string) Option {
	return singletonOption(val, String, backgroundcolor)
}

func Basex(val float64) Option {
	return singletonOption(val, Number, basex)
}

func Basey(val float64) Option {
	return singletonOption(val, Number, basey)
}

func C(val string) Option {
	return singletonOption(val, String, c)
}

func ClipOn(val bool) Option {
	return singletonOption(val, Bool, clipOn)
}

func Color(val string) Option {
	return singletonOption(val, String, color)
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

func Dashes(intr interface{}) Option {
	panic("NYI")
}

func Family(val string) Option {
	if val != "serif" && val != "sans-serif" && val != "cursive" &&
		val != "fantasy" && val != "monospace" {
		panic("Invalid value for Family Option.")
	}
	return singletonOption(val, String, family)
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

func Ha(val string) Option {
	if val != "center" && val != "right" && val != "left" {
		panic("Invalid value for Ha option.")
	}
	return singletonOption(val, String, ha)
}

func Horizontalalignment(val string) Option {
	if val != "center" && val != "right" && val != "left" {
		panic("Invalid value for Horizontalalignment option.")
	}
	return singletonOption(val, String, horizontalalignment)
}

func Label(val string) Option {
	return singletonOption(val, String, label)
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

func X(val float64) Option {
	return singletonOption(val, Number, x)
}

func Xmax(intr interface{}) Option {
	return singletonOption(intr, NoneNumber, xmax)
}

func Xmin(intr interface{}) Option {
	return singletonOption(intr, NoneNumber, xmin)
}

func Xdata(vals []float64) Option {
	return singletonOption(vals, Array, xdata)
}

func Y(val float64) Option {
	return singletonOption(val, Number, y)
}

func Ymax(intr interface{}) Option {
	return singletonOption(intr, NoneNumber, ymax)
}

func Ymin(intr interface{}) Option {
	return singletonOption(intr, NoneNumber, ymin)
}

func Ydata(vals []float64) Option {
	return singletonOption(vals, Array, ydata)
}

func Zorder(val float64) Option {
	return singletonOption(val, Number, zorder)
}
