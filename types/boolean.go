package types

import (
	"fmt"
	"math/big"
)

type PrimitiveBool struct {
	value bool
}

func MakeBool(value bool) *PrimitiveBool {
	return &PrimitiveBool{value}
}

func (boolean *PrimitiveBool) Display() string {
	if boolean.value {
		return "true"
	} else {
		return "false"
	}
}

// String Operators

func (boolean *PrimitiveBool) ConcatenateInPlace(other Primitive) error {
	return fmt.Errorf("bool type does not support numeral arithmetic")
}

func (boolean *PrimitiveBool) Concatenate(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("bool type does not support numeral arithmetic")
}

// Arithmetic Operators

func (boolean *PrimitiveBool) Add(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("bool type does not support numeral arithmetic")
}

func (boolean *PrimitiveBool) AddInplace(other Primitive) error {
	return fmt.Errorf("bool type does not support numeral arithmetic")
}

func (boolean *PrimitiveBool) Subtract(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("bool type does not support numeral arithmetic")
}

func (boolean *PrimitiveBool) SubtractInplace(other Primitive) error {
	return fmt.Errorf("bool type does not support numeral arithmetic")
}

func (boolean *PrimitiveBool) Multiply(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("bool type does not support numeral arithmetic")
}

func (boolean *PrimitiveBool) MultiplyInplace(other Primitive) error {
	return fmt.Errorf("bool type does not support numeral arithmetic")
}

func (boolean *PrimitiveBool) Divide(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("bool type does not support numeral arithmetic")
}

func (boolean *PrimitiveBool) DivideInplace(other Primitive) error {
	return fmt.Errorf("bool type does not support numeral arithmetic")
}

func (boolean *PrimitiveBool) IntDivide(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("bool type does not support numeral arithmetic")
}

func (boolean *PrimitiveBool) IntDivideInplace(other Primitive) error {
	return fmt.Errorf("bool type does not support numeral arithmetic")
}

func (boolean *PrimitiveBool) Modulo(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("bool type does not support numeral arithmetic")
}

func (boolean *PrimitiveBool) ModuloInplace(other Primitive) error {
	return fmt.Errorf("bool type does not support numeral arithmetic")
}

func (boolean *PrimitiveBool) RaisePower(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("bool type does not support numeral arithmetic")
}

func (boolean *PrimitiveBool) RaisePowerInplace(other Primitive) error {
	return fmt.Errorf("bool type does not support numeral arithmetic")
}

// Casting

func (boolean *PrimitiveBool) RequireNum() (*PrimitiveNum, error) {
	return nil, fmt.Errorf("bool %q used where number was required", boolean.Display())
}

func (boolean *PrimitiveBool) RequireStr() (*PrimitiveStr, error) {
	return nil, fmt.Errorf("bool %q used where number was required", boolean.Display())
}

func (boolean *PrimitiveBool) RequireBool() (*PrimitiveBool, error) {
	return boolean, nil
}

func (boolean *PrimitiveBool) CastNum() (*PrimitiveNum, error) {
	if boolean.value {
		return &PrimitiveNum{*big.NewRat(1, 1)}, nil
	}
	return &PrimitiveNum{*big.NewRat(0, 1)}, nil
}

func (boolean *PrimitiveBool) CastStr() (*PrimitiveStr, error) {
	return &PrimitiveStr{boolean.Display()}, nil
}

func (boolean *PrimitiveBool) CastBool() (*PrimitiveBool, error) {
	return boolean, nil
}