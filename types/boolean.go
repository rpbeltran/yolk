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

// Logical Operators

func (boolean *PrimitiveBool) And(other Primitive) (Primitive, error) {
	if other_bool, err := other.CastImplicitBool(); err != nil {
		return nil, fmt.Errorf("attempting to And: %w", err)
	} else {
		return &PrimitiveBool{boolean.value && other_bool.value}, nil
	}
}

func (boolean *PrimitiveBool) Or(other Primitive) (Primitive, error) {
	if other_bool, err := other.CastImplicitBool(); err != nil {
		return nil, fmt.Errorf("attempting to Or: %w", err)
	} else {
		return &PrimitiveBool{boolean.value || other_bool.value}, nil
	}
}

// Non logical Operators

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

func (boolean *PrimitiveBool) ConcatenateInPlace(other Primitive) error {
	return fmt.Errorf("bool type does not support numeral arithmetic")
}

func (boolean *PrimitiveBool) Concatenate(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("bool type does not support numeral arithmetic")
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

func (boolean *PrimitiveBool) CastImplicitNum() (*PrimitiveNum, error) {
	return nil, fmt.Errorf("bool %q used where number-like required", boolean.Display())
}

func (boolean *PrimitiveBool) CastImplicitStr() (*PrimitiveStr, error) {
	return nil, fmt.Errorf("bool %q used where string-like required", boolean.Display())
}

func (boolean *PrimitiveBool) CastImplicitBool() (*PrimitiveBool, error) {
	return boolean, nil
}

func (boolean *PrimitiveBool) CastExplicitNum() (*PrimitiveNum, error) {
	if boolean.value {
		return &PrimitiveNum{*big.NewRat(1, 1)}, nil
	}
	return &PrimitiveNum{*big.NewRat(0, 1)}, nil
}

func (boolean *PrimitiveBool) CastExplicitStr() (*PrimitiveStr, error) {
	return &PrimitiveStr{boolean.Display()}, nil
}

func (boolean *PrimitiveBool) CastExplicitBool() (*PrimitiveBool, error) {
	return boolean, nil
}
