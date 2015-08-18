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
	aa: AA,
	alpha: Alpha,
	animated: Animated,
	antialiased: Antialiased,
	axis: Axis,
	b: B,
	backgroundcolor: BackgroundColor,
	barsabove: BarsAbove,
	basex: BaseX,
	basey: BaseY,
	bboxToAnchor: BboxToAnchor,
	borderaxespad: BorderAxesPad,
	borderpad: BorderPad,
	c: C,
	capsize: CapSize,
	capthick: CapThick,
	clipOn: ClipOn,
	color: Color,
	columnspacing: ColumnSpacing,
	dashCapstyle: DashCapstyle,
	dashJoinstyle: DashJoinstyle,
	dashes: Dashes,
	dpi: DPI,
	ecolor: EColor,
	edgecolor: EdgeColor,
	elinewidth: ELineWidth,
	errorevery: ErrorEvery,
	facecolor: FaceColor,
	family: Family,
	fancybox: FancyBox,
	figsize: FigSize,
	fmtFlag: Fmt,
	fontname: FontName,
	fontsize: FontSize,
	fontstyle: FontStyle,
	fontweight: FontWeight,
	framealpha: FrameAlpha,
	frameon: FrameOn,
	ha: HA,
	handlelength: HandleLength,
	handletextpad: HandleTextPad,
	horizontalalignment: HorizontalAlignment,
	label: Label,
	labelspacing: LabelSpacing,
	linespacing: LineSpacing,
	linestyle: LineStyle,
	linewidth: LineWidth,
	linscalex: LinScaleX,
	linscaley: LinScaleY,
	linthreshx: LinThreshX,
	linthreshy: LinThreshY,
	ls: LS,
	lw: LW,
	loc: Loc,
	lod: LoD,
	lolims: LoLims,
	marker: Marker,
	markeredgecolor: MarkerEdgeColor,
	markeredgewidth: MarkerEdgeWidth,
	markerfacecolor: MarkerFaceColor,
	markerscale: MarkerScale,
	markersize: MarkerSize,
	markevery: MarkEvery,
	mec: MEC,
	mew: MEW,
	mfc: MFC,
	mode: Mode,
	ms: MS,
	multialignment: MultiAlignment,
	name: Name,
	ncol: NCol,
	num: Num,
	numpoints: NumPoints,
	nonposx: NonPosX,
	nonposy: NonPosY,
	pickradius: PickRadius,
	position: Position,
	rotation: Rotation,
	rotationMode: RotationMode,
	scatterpoints: ScatterPoints,
	scatteryoffsets: ScatterYOffsets,
	shadow: Shadow,
	size: Size,
	solidCapstyle: SolidCapstyle,
	solidJoinstyle: SolidJoinstyle,
	subsx: SubsX,
	subsy: SubsY,
	style: Style,
	text: Text,
	title: Title,
	uplims: UpLims,
	va: VA,
	variant: Variant,
	verticalalignment: VerticalAlignment,
	visible: Visible,
	weight: Weight,
	which: Which,
	x: X,
	xerr: XErr,
	xdata: XData,
	xlolims: XLoLims,
	xmax: XMax,
	xmin: XMin,
	xuplims: XUpLims,
	y: Y,
	yerr: YErr,
	ymax: YMax,
	ymin: YMin,
	ydata: YData,
	zorder: ZOrder,
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

func AA(val bool) Option {
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

func BackgroundColor(val string) Option {
	return singletonOption(val, String, backgroundcolor)
}

func BarsAbove(val bool) Option {
	return singletonOption(val, Bool, barsabove)
}

func BaseX(val float64) Option {
	return singletonOption(val, Number, basex)
}

func BaseY(val float64) Option {
	return singletonOption(val, Number, basey)
}

func BboxToAnchor(x, y float64) Option {
	return func(fo funcOptions) (string, bool) {
		if _, ok := fo[bboxToAnchor]; ok { return "", false }
		return fmt.Sprintf("(%g, %g)", x, y), true
	}
}

func BorderAxesPad(intr interface{}) Option {
	return singletonOption(intr, NoneNumber, borderaxespad)
}

func BorderPad(intr interface{}) Option {
	return singletonOption(intr, NoneNumber, borderpad)
}

func C(val string) Option {
	return singletonOption(val, String, c)
}

func CapSize(val float64) Option {
	return singletonOption(val, Number, capsize)
}

func CapThick(val float64) Option {
	return singletonOption(val, Number, capthick)
}

func ClipOn(val bool) Option {
	return singletonOption(val, Bool, clipOn)
}

func Color(val string) Option {
	return singletonOption(val, String, color)
}

func ColumnSpacing(intr interface{}) Option {
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

func DPI(intr interface{}) Option {
	return singletonOption(intr, NoneInt, dpi)
}

func EColor(intr interface{}) Option {
	return singletonOption(intr, NoneString, ecolor)
}

func EdgeColor(val string) Option {
	return singletonOption(val, String, edgecolor)
}

func ELineWidth(val float64) Option {
	return singletonOption(val, Number, elinewidth)
}

func ErrorEvery(val int) Option {
	return singletonOption(val, Int, errorevery)
}

func FaceColor(val string) Option {
	return singletonOption(val, String, facecolor)
}

func Family(val string) Option {
	if val != "serif" && val != "sans-serif" && val != "cursive" &&
		val != "fantasy" && val != "monospace" {
		panic("Invalid value for Family Option.")
	}
	return singletonOption(val, String, family)
}

func FancyBox(intr interface{}) Option {
	return singletonOption(intr, NoneBool, fancybox)
}

func FigSize(x, y int) Option {
	return func(fo funcOptions) (string, bool) {
		if _, ok := fo[figsize]; !ok { return "", false }
		return fmt.Sprintf("%s=(%d,%d)", optionNames[figsize], x, y), true
	}
}

func Fmt(intr interface{}) Option {
	return singletonOption(intr, NoneString, fmtFlag)
}

func FontName(val string) Option {
	return singletonOption(val, String, fontname)
}

func FontSize(val interface{}) Option {
	if _, ok := convertNumber(val); ok {
		return singletonOption(val, Number, fontsize)
	} else if _, ok := convertString(val); ok {
		return singletonOption(val, String, fontsize)
	}
	panic("Fonsize Option only accepts strings and numbers.")
}

func FontStyle(val string) Option {
	if val != "normal" && val != "italic" && val != "oblique" {
		panic("Invalid value for FontStyle Option.")
	}
	return singletonOption(val, String, fontstyle)
}

func FontWeight(val string) Option {
	if val != "normal" && val != "bold" && val != "heavy" &&
		val != "light" && val != "ultrabold" && val != "ultralight" {
		panic("Invalid value for FontWeight Option.")
	}

	return singletonOption(val, String, fontweight)
}

func FrameAlpha(intr interface{}) Option {
	return singletonOption(intr, NoneNumber, framealpha)
}

func FrameOn(intr interface{}) Option {
	return singletonOption(intr, NoneBool, frameon)
}

func HA(val string) Option {
	if val != "center" && val != "right" && val != "left" {
		panic("Invalid value for HA option.")
	}
	return singletonOption(val, String, ha)
}

func HandleLength(intr interface{}) Option {
	return singletonOption(intr, NoneNumber, handlelength)
}

func HandleTextPad(intr interface{}) Option {
	return singletonOption(intr, NoneNumber, handletextpad)
}

func HorizontalAlignment(val string) Option {
	if val != "center" && val != "right" && val != "left" {
		panic("Invalid value for HorizontalAlignment option.")
	}
	return singletonOption(val, String, horizontalalignment)
}

func Label(val string) Option {
	singletonOption(val, String, label)
	return singletonOption(val, String, label)
}

func LabelSpacing(intr interface{}) Option {
	return singletonOption(intr, NoneNumber, labelspacing)
}

func LineSpacing(val float64) Option {
	return singletonOption(val, Number, linespacing)
}

func LineStyle(val string) Option {
	return singletonOption(val, String, linestyle)
}

func LinScaleX(val float64) Option { 
	return singletonOption(val, Number, linscalex)
}

func LinScaleY(val float64) Option { 
	return singletonOption(val, Number, linscaley)
}

func LinThreshX(val float64) Option {
	return singletonOption(val, Number, linthreshx)
}

func LinThreshY(val float64) Option {
	return singletonOption(val, Number, linthreshy)
}

func LoLims(intr interface{}) Option {
	if _, ok := convertBool(intr); ok {
		return singletonOption(intr, Bool, lolims)
	} else if _, ok := convertBoolArray(intr); ok {
		return singletonOption(intr, BoolArray, lolims)
	}
	panic("Invalid type for arugment to LoLims Option.")
}

func LS(val string) Option {
	return singletonOption(val, String, ls)
}

func LineWidth(val string) Option {
	return singletonOption(val, String, linewidth)
}

func LW(val float64) Option {
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

func LoD(val bool) Option {
	return singletonOption(val, Bool, lod)
}

func Marker(val string) Option {
	if val != "+" && val != "," && val != "." &&
		val != "1" && val != "2" && val != "3" && val != "4" {
		panic("Invalid value for Marker Option.")
	}
	return singletonOption(val, String, marker)
}

func MarkerEdgeColor(val string) Option {
	return singletonOption(val, String, markeredgecolor)
}

func MarkerEdgeWidth(val float64) Option {
	return singletonOption(val, Number, markeredgewidth)
}

func MarkerFaceColor(val string) Option {
	return singletonOption(val, String, markerfacecolor)
}

func MarkerScale(intr interface{}) Option {
	return singletonOption(intr, NoneNumber, markerscale)
}

func MarkerSize(val float64) Option {
	return singletonOption(val, Number, markersize)
}

func MarkEvery(val float64) Option {
	panic("NYI")
}

func MEC(val string) Option {
	return singletonOption(val, String, mec)
}

func MEW(val float64) Option {
	return singletonOption(val, Number, mew)
}

func MFC(val string) Option {
	return singletonOption(val, String, mfc)
}

func Mode(intr interface{}) Option {
	return singletonOption(intr, NoneString, mode)
}

func MS(val float64) Option {
	return singletonOption(val, Number, markersize)
}

func MultiAlignment(val string) Option {
	if val != "left" && val != "center" && val != "right" {
		panic("Invalid value for MultiAlignment Option.")
	}
	return singletonOption(val, String, multialignment)
}

func Name(val string) Option {
	return singletonOption(val, String, name)
}

func NCol(val int) Option {
	return singletonOption(val, Int, ncol)
}

func NonPosX(val string) Option {
	if val != "mask" && val != "clip" {
		panic("Invalid value for NonPosX Option")
	}
	return singletonOption(val, String, nonposx)
}

func NonPosY(val string) Option {
	if val != "mask" && val != "clip" {
		panic("Invalid value for NonPosY Option")
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

func NumPoints(intr interface{}) Option {
	return singletonOption(intr, NoneInt, numpoints)
}

func PickRadius(val float64) Option {
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

func ScatterPoints(intr interface{}) Option {
	return singletonOption(intr, NoneInt, scatterpoints)
}

func ScatterYOffsets(vals []float64) Option {
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

func SubsX(vals []int) Option {
	return singletonOption(vals, Array, subsx)
}

func SubsY(vals []int) Option {
	return singletonOption(vals, Array, subsy)
}

func Text(val string) Option {
	return singletonOption(val, String, text)
}

func UpLims(intr interface{}) Option {
	if _, ok := convertBool(intr); ok {
		return singletonOption(intr, Bool, uplims)
	} else if _, ok := convertBoolArray(intr); ok {
		return singletonOption(intr, BoolArray, uplims)
	}
	panic("Invalid type for arugment to UpLims Option.")
}

func VA(val string) Option {
	if val != "center" && val != "top" &&
		val != "bottom" && val != "baseline" {
		panic("Invalid value for VA Option.")
	}

	return singletonOption(val, String, va)
}

func Variant(val string) Option {
	if val != "notmal" && val != "small-caps" {
		panic("Invalid value for Variant Option.")
	}
	return singletonOption(val, String, variant)
}

func VerticalAlignment(val string) Option {
	if val != "center" && val != "top" &&
		val != "bottom" && val != "baseline" {
		panic("Invalid value for VerticalAlignment Option.")
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

func XData(vals []float64) Option {
	return singletonOption(vals, Array, xdata)
}

func XErr(vals []float64) Option {
	return singletonOption(vals, Array, xerr)
}

func XLoLims(intr interface{}) Option {
	if _, ok := convertBool(intr); ok {
		return singletonOption(intr, Bool, xlolims)
	} else if _, ok := convertBoolArray(intr); ok {
		return singletonOption(intr, BoolArray, xlolims)
	}
	panic("Invalid type for arugment to XLoLims Option.")
}

func XMax(intr interface{}) Option {
	return singletonOption(intr, NoneNumber, xmax)
}

func XMin(intr interface{}) Option {
	return singletonOption(intr, NoneNumber, xmin)
}

func XUpLims(intr interface{}) Option {
	if _, ok := convertBool(intr); ok {
		return singletonOption(intr, Bool, xuplims)
	} else if _, ok := convertBoolArray(intr); ok {
		return singletonOption(intr, BoolArray, xuplims)
	}
	panic("Invalid type for arugment to XUpLims Option.")
}

func Y(val float64) Option {
	return singletonOption(val, Number, y)
}

func YData(vals []float64) Option {
	return singletonOption(vals, Array, ydata)
}

func YErr(vals []float64) Option {
	return singletonOption(vals, Array, yerr)
}

func YMax(intr interface{}) Option {
	return singletonOption(intr, NoneNumber, ymax)
}

func YMin(intr interface{}) Option {
	return singletonOption(intr, NoneNumber, ymin)
}

func ZOrder(val float64) Option {
	return singletonOption(val, Number, zorder)
}
