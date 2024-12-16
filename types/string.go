package types

import (
	"fmt"
	"math/big"
)

type PrimitiveStr struct {
	value string
}

func MakeString(value string) *PrimitiveStr {
	return &PrimitiveStr{value}
}

func (str *PrimitiveStr) Display() string {
	return str.value
}

// Operators

func (str *PrimitiveStr) Add(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("string type does not support numeric addition")
}

// Casting

func (str *PrimitiveStr) RequireNum() (*PrimitiveNum, error) {
	return nil, fmt.Errorf("string value %q used where number was required", str.value)
}

func (str *PrimitiveStr) RequireStr() (*PrimitiveStr, error) {
	return &PrimitiveStr{str.Display()}, nil
}

func (str *PrimitiveStr) CastNum() (*PrimitiveNum, error) {
	var num big.Rat
	if _, success := num.SetString(str.value); success {
		return &PrimitiveNum{num}, nil
	}
	return nil, fmt.Errorf("cannot interpret the string %q as a number", str.value)
}

func (str *PrimitiveStr) CastStr() (*PrimitiveStr, error) {
	return &PrimitiveStr{str.Display()}, nil
}
