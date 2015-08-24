package pyplot

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type pyplotType int

const (
	invalid pyplotType = iota
	Int
	Number
	Bool
	String
	Array
	BoolArray
	NoneInt
	NoneNumber
	NoneBool
	NoneString
	NoneArray
)

func convertInt(val interface{}) (string, bool) {
	switch val.(type) {
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8:
		return fmt.Sprintf("%d", val), true
	default: return "", false
	}
}

func convertNumber(val interface{}) (string, bool) {
	str, isInt := convertInt(val)
	if isInt { return str, true }
	switch t := val.(type) {
	case float32: return convertFloat64(float64(t)), true
	case float64: return convertFloat64(t), true
	default: return "", false
	}
}

func convertBool(val interface{}) (string, bool) {
	switch val := val.(type) {
	case bool:
		if val { return "True", true }
		return "False", true
	default: return "", false
	}
}

func convertString(val interface{}) (string, bool) {
	str, ok := val.(string)
	if !ok { return "", false }
	return strconv.Quote(str), true
}

func convertBoolArray(val interface{}) (string, bool) {
	bs, ok := val.([]bool)
	if !ok { return "", false }
	ss := make([]string, len(bs))
	for i := range bs {
		if bs[i] {
			ss[i] = "True"
		} else {
			ss[i] = "False"
		}
	}
	return fmt.Sprintf("[%s]", strings.Join(ss, ",")), true
}

func convertFloat64(val float64) string {
	if math.IsNaN(val) { return "np.nan" }
	if math.IsInf(val, +1)  { return "+np.inf" }
	if math.IsInf(val, -1)  { return "-np.inf" }
	return fmt.Sprintf("%g", val)
}

func convertArray(val interface{}) (string, bool) {
	switch xs := val.(type) {
	case []float64:
		items := make([]string, len(xs))
		for i, x := range xs { items[i] = convertFloat64(x) }
		return fmt.Sprintf("[%s]", strings.Join(items, ",")), true
	case []float32:
		items := make([]string, len(xs))
		for i, x := range xs { items[i] = convertFloat64(float64(x)) }
		return fmt.Sprintf("[%s]", strings.Join(items, ",")), true
	case []int:
		items := make([]string, len(xs))
		for i, x := range xs { items[i] = fmt.Sprintf("%d", x) }
		return fmt.Sprintf("[%s]", strings.Join(items, ",")), true
	case []int64:
		items := make([]string, len(xs))
		for i, x := range xs { items[i] = fmt.Sprintf("%d", x) }
		return fmt.Sprintf("[%s]", strings.Join(items, ",")), true
	case []int32:
		items := make([]string, len(xs))
		for i, x := range xs { items[i] = fmt.Sprintf("%d", x) }
		return fmt.Sprintf("[%s]", strings.Join(items, ",")), true
	case []int16:
		items := make([]string, len(xs))
		for i, x := range xs { items[i] = fmt.Sprintf("%d", x) }
		return fmt.Sprintf("[%s]", strings.Join(items, ",")), true
	case []int8:
		items := make([]string, len(xs))
		for i, x := range xs { items[i] = fmt.Sprintf("%d", x) }
		return fmt.Sprintf("[%s]", strings.Join(items, ",")), true
	case []uint:
		items := make([]string, len(xs))
		for i, x := range xs { items[i] = fmt.Sprintf("%d", x) }
		return fmt.Sprintf("[%s]", strings.Join(items, ",")), true
	case []uint64:
		items := make([]string, len(xs))
		for i, x := range xs { items[i] = fmt.Sprintf("%d", x) }
		return fmt.Sprintf("[%s]", strings.Join(items, ",")), true
	case []uint32:
		items := make([]string, len(xs))
		for i, x := range xs { items[i] = fmt.Sprintf("%d", x) }
		return fmt.Sprintf("[%s]", strings.Join(items, ",")), true
	case []uint16:
		items := make([]string, len(xs))
		for i, x := range xs { items[i] = fmt.Sprintf("%d", x) }
		return fmt.Sprintf("[%s]", strings.Join(items, ",")), true
	case []uint8:
		items := make([]string, len(xs))
		for i, x := range xs { items[i] = fmt.Sprintf("%d", x) }
		return fmt.Sprintf("[%s]", strings.Join(items, ",")), true
	default:
		return "", false
	}
}

func convertNone(val interface{}) (string, bool) {
	if val == nil { return "None", true }
	return "", false
}

func convertType(val interface{}, t pyplotType) (str string, ok bool) {
	switch t {
	case Int: return convertInt(val)
	case Number: return convertNumber(val)
	case Bool: return convertBool(val)
	case String: return convertString(val)
	case Array: return convertArray(val)
	case BoolArray: return convertBoolArray(val)

	case NoneInt:
		if str, ok := convertNone(val); ok { return str, ok }
		return convertInt(val)
	case NoneNumber:
		if str, ok := convertNone(val); ok { return str, ok }
		return convertNumber(val)
	case NoneBool:
		if str, ok := convertNone(val); ok { return str, ok }
		return convertBool(val)
	case NoneString:
		if str, ok := convertNone(val); ok { return str, ok }
		return convertString(val)
	case NoneArray:
		if str, ok := convertNone(val); ok { return str, ok }
		return convertArray(val)
	default: return "", false
	}
}
