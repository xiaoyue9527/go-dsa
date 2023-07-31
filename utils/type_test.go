package utils

import (
	"testing"
)

func TestCheckType(t *testing.T) {
	tests := []struct {
		input  interface{}
		output string
	}{
		{42, "int"},
		{int8(42), "int8"},
		{int16(42), "int16"},
		{int32(42), "int32"},
		{int64(42), "int64"},
		{uint(42), "uint"},
		{uint8(42), "uint8 (byte)"},
		{uint16(42), "uint16"},
		{uint32(42), "uint32"},
		{uint64(42), "uint64"},
		{float32(42.0), "float32"},
		{float64(42.0), "float64"},
		{complex64(1 + 2i), "complex64"},
		{complex128(1 + 2i), "complex128"},
		{true, "bool"},
		{"hello", "string"},
		{struct{}{}, "struct {}"},
	}

	for _, test := range tests {
		result := CheckType(test.input)
		if result != test.output {
			t.Errorf("CheckType(%v) = %s, want %s", test.input, result, test.output)
		}
	}
}
