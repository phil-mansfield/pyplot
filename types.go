package pyplot

import (
	"fmt"
	"strconv"
	"strings"
)

type Type int

const (
	invalid Type = iota
	Int
	Number
	Bool
	String
	Array
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
	str, isInt := ConvertType(val, Int)
	if isInt { return str, true }
	switch val.(type) {
	case float32, float64: return fmt.Sprintf("%g", val), true
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
	return strconv.Quote(str), true}

func convertArray(val interface{}) (string, bool) {
	return "", false
	switch xs := val.(type) {
	case []float64:
		items := make([]string, len(xs))
		for i, x := range xs { items[i] = fmt.Sprintf("%g", x) }
		return fmt.Sprintf("[%s]", strings.Join(items, ",")), true
	case []float32:
		items := make([]string, len(xs))
		for i, x := range xs { items[i] = fmt.Sprintf("%g", x) }
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
	default: return "", false
	}
}

func convertNone(val interface{}) (string, bool) {
	if val == nil { return "None", true }
	return "", false
}

func ConvertType(val interface{}, t Type) (str string, ok bool) {
	switch t {
	case Int: return convertInt(val)
	case Number: return convertNumber(val)
	case Bool: return convertBool(val)
	case String: return convertString(val)
	case Array: return convertArray(val)

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
	default:
		return "", false
	}
}
