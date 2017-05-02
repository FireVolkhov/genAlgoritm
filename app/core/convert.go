package core

import (
	"strconv"
	"encoding/binary"
)

func StringToInt (value string) int {
	result, _ := strconv.ParseInt(value, 10, 64)
	return int(result);
}

func StringToFloat32 (value string) float32 {
	result, _ := strconv.ParseFloat(value, 32)
	return float32(result);
}

func Float64ToFloat32 (value float64) float32 {
	return float32(value);
}

func IntToFloat32 (value int) float32 {
	return float32(value)
}

func ToFloat32 (value interface{}) float32 {
	switch value.(type) {
	case int:
		return IntToFloat32(value.(int))
	case float32:
		return value.(float32)
	case float64:
		return Float64ToFloat32(value.(float64))
	case string:
		return StringToFloat32(value.(string))
	default:
		panic("Not have variant")
	}
}

func IntToFloat64 (value int) float64 {
	return float64(value)
}

func Float32ToFloat64 (value float32) float64 {
	return float64(value)
}

func StringToFloat64 (value string) float64 {
	result, _ := strconv.ParseFloat(value, 64)
	return result;
}

func ToFloat64 (value interface{}) float64 {
	switch value.(type) {
	case int:
		return IntToFloat64(value.(int))
	case float32:
		return Float32ToFloat64(value.(float32))
	case float64:
		return value.(float64)
	case string:
		return StringToFloat64(value.(string))
	default:
		panic("Not have variant")
	}
}

func Float64ToString(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}

func IntToString(value int) string {
	return strconv.FormatInt(int64(value), 10)
}

func Uint16ToBytes(value uint16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, uint16(700))
	return b
}
