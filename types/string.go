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

// Comparisons

func (str *PrimitiveStr) Equal(other Primitive) bool {
	if as_str, err := other.RequireStr(); err == nil {
		return str.value == as_str.value
	}
	return false
}

// String Operators

func (str *PrimitiveStr) ConcatenateInPlace(other Primitive) error {
	str.value = str.value + other.Display()
	return nil
}

func (str *PrimitiveStr) Concatenate(other Primitive) (Primitive, error) {
	return &PrimitiveStr{str.value + other.Display()}, nil
}

// Arithmetic Operators

func (str *PrimitiveStr) Add(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("string type does not support numeral arithmetic")
}

func (str *PrimitiveStr) AddInplace(other Primitive) error {
	return fmt.Errorf("string type does not support numeral arithmetic")
}

func (str *PrimitiveStr) Subtract(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("string type does not support numeral arithmetic")
}

func (str *PrimitiveStr) SubtractInplace(other Primitive) error {
	return fmt.Errorf("string type does not support numeral arithmetic")
}

func (str *PrimitiveStr) Multiply(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("string type does not support numeral arithmetic")
}

func (str *PrimitiveStr) MultiplyInplace(other Primitive) error {
	return fmt.Errorf("string type does not support numeral arithmetic")
}

func (str *PrimitiveStr) Divide(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("string type does not support numeral arithmetic")
}

func (str *PrimitiveStr) DivideInplace(other Primitive) error {
	return fmt.Errorf("string type does not support numeral arithmetic")
}

func (str *PrimitiveStr) IntDivide(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("string type does not support numeral arithmetic")
}

func (str *PrimitiveStr) IntDivideInplace(other Primitive) error {
	return fmt.Errorf("string type does not support numeral arithmetic")
}

func (str *PrimitiveStr) Modulo(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("string type does not support numeral arithmetic")
}

func (str *PrimitiveStr) ModuloInplace(other Primitive) error {
	return fmt.Errorf("string type does not support numeral arithmetic")
}

func (str *PrimitiveStr) RaisePower(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("string type does not support numeral arithmetic")
}

func (str *PrimitiveStr) RaisePowerInplace(other Primitive) error {
	return fmt.Errorf("string type does not support numeral arithmetic")
}

// Logical Operators

func (str *PrimitiveStr) And(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("string type does not support boolean logic")
}

func (str *PrimitiveStr) Or(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("string type does not support boolean logic")
}

// Casting

func (str *PrimitiveStr) RequireNum() (*PrimitiveNum, error) {
	return nil, fmt.Errorf("string value %q used where number was required", str.value)
}

func (str *PrimitiveStr) RequireStr() (*PrimitiveStr, error) {
	return str, nil
}

func (str *PrimitiveStr) RequireBool() (*PrimitiveBool, error) {
	return nil, fmt.Errorf("string value %q used where boolean was required", str.value)
}

func (str *PrimitiveStr) CastImplicitNum() (*PrimitiveNum, error) {
	return nil, fmt.Errorf("string value %q used where implicit number was required", str.value)
}

func (str *PrimitiveStr) CastExplicitNum() (*PrimitiveNum, error) {
	var num big.Rat
	if _, success := num.SetString(str.value); success {
		return &PrimitiveNum{num}, nil
	}
	return nil, fmt.Errorf("cannot interpret the string %q as a number", str.value)
}

func (str *PrimitiveStr) Display() string {
	return str.value
}

func (str *PrimitiveStr) Truthy() bool {
	return len(str.value) != 0
}
