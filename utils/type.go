package utils

import (
	"reflect"
)

func CheckType(data interface{}) string {
	switch data.(type) {
	case int:
		return "int"
	case int8:
		return "int8"
	case int16:
		return "int16"
	case int32:
		return "int32"
	case int64:
		return "int64"
	case uint:
		return "uint"
	case uint8:
		return "uint8 (byte)"
	case uint16:
		return "uint16"
	case uint32:
		return "uint32"
	case uint64:
		return "uint64"
	case float32:
		return "float32"
	case float64:
		return "float64"
	case complex64:
		return "complex64"
	case complex128:
		return "complex128"
	case bool:
		return "bool"
	case string:
		return "string"
	default:
		return reflect.TypeOf(data).String() // 返回未知类型的名称
	}
}
