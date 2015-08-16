package pyplot

import (
	"fmt"
)

type optionFlag int
type funcOptions map[optionFlag]bool
type internalOption func(funcOptions) (string, bool)
type Option internalOption

const (
	alpha optionFlag = iota
	animated
	antialiased
	aa
	clipOn
	color
	c
	dashCapstyle
	dashJoinstyle
	dashes
	data
	label
	linestyle
	ls
	linewidth
	lw
	lod
	marker
	markeredgecolor
	mec
	markeredgewidth
	mew
	markerfacecolor
	mfc
	markersize
	ms
	markevery
	pickradius
	solidCapstyle
	solidJoinstyle
	visible
	xdata
	ydata
	zorder
)

func optionStart(args []interface{}) int {
	for i, intr := range args {
		switch intr.(type) {
		case Option: return i
		}
	}
	return len(args)
}

func Alpha(val float64) Option {
	return func(fo funcOptions) (string, bool) {
		if _, ok := fo[alpha]; ok {
			return fmt.Sprintf("lw=%g", val), true
		} else {
			return "", false
		}
	}
}

func Animated(val bool) Option {
	return func(fo funcOptions) (string, bool) {
		if _, ok := fo[alpha]; ok {
			b, _ := convertBool(val)
			return fmt.Sprintf("lw=%s", b), true
		} else {
			return "", false
		}
	}
}

func Antialiased(val bool) Option {
	return func(fo funcOptions) (string, bool) {
		if _, ok := fo[antialiased]; ok {
			b, _ := convertBool(val)
			return fmt.Sprintf("lw=%s", b), true
		} else {
			return "", false
		}
	}
}

func Aa(val bool) Option {
	return func(fo funcOptions) (string, bool) {
		if _, ok := fo[aa]; ok {
			b, _ := convertBool(val)
			return fmt.Sprintf("lw=%s", b), true
		} else {
			return "", false
		}
	}
}

func Lw(val float64) Option {
	return func(fo funcOptions) (string, bool) {
		if _, ok := fo[lw]; ok {
			return fmt.Sprintf("lw=%g", val), true
		} else {
			return "", false
		}
	}
}
