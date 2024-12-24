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

// String Operators

func (str *PrimitiveStr) ConcatenateInPlace(other Primitive) error {
	if other_str, err := other.CastStr(); err != nil {
		return fmt.Errorf("attempting to concatenate: %w", err)
	} else {
		str.value = str.value + other_str.value
		return nil
	}
}

func (str *PrimitiveStr) Concatenate(other Primitive) (Primitive, error) {
	if other_str, err := other.CastStr(); err != nil {
		return nil, fmt.Errorf("attempting to concatenate: %w", err)
	} else {
		return &PrimitiveStr{str.value + other_str.value}, nil
	}
}

// Arithmetic Operators

func (str *PrimitiveStr) Add(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("string type does not support numeral arithmetic")
}

func (str *PrimitiveStr) AddInplace(other Primitive) error {
	return fmt.Errorf("string type does not support numeral arithmetic")
}

func (num *PrimitiveStr) Subtract(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("string type does not support numeral arithmetic")
}

func (num *PrimitiveStr) SubtractInplace(other Primitive) error {
	return fmt.Errorf("string type does not support numeral arithmetic")
}

func (num *PrimitiveStr) Multiply(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("string type does not support numeral arithmetic")
}

func (num *PrimitiveStr) MultiplyInplace(other Primitive) error {
	return fmt.Errorf("string type does not support numeral arithmetic")
}

func (num *PrimitiveStr) Divide(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("string type does not support numeral arithmetic")
}

func (num *PrimitiveStr) DivideInplace(other Primitive) error {
	return fmt.Errorf("string type does not support numeral arithmetic")
}

func (num *PrimitiveStr) IntDivide(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("string type does not support numeral arithmetic")
}

func (num *PrimitiveStr) IntDivideInplace(other Primitive) error {
	return fmt.Errorf("string type does not support numeral arithmetic")
}

func (num *PrimitiveStr) Modulo(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("string type does not support numeral arithmetic")
}

func (num *PrimitiveStr) ModuloInplace(other Primitive) error {
	return fmt.Errorf("string type does not support numeral arithmetic")
}

func (num *PrimitiveStr) RaisePower(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("string type does not support numeral arithmetic")
}

func (num *PrimitiveStr) RaisePowerInplace(other Primitive) error {
	return fmt.Errorf("string type does not support numeral arithmetic")
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
