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

func (boolean *PrimitiveBool) Type() string {
	return "bool"
}

// Comparisons

func (boolean *PrimitiveBool) Equal(other Primitive) bool {
	if as_bool, err := other.RequireBool(); err == nil {
		return boolean.value == as_bool.value
	}
	return false
}

func (boolean *PrimitiveBool) LessThan(other Primitive) (bool, error) {
	if as_bool, err := other.RequireBool(); err != nil {
		return false, fmt.Errorf("cannot compare boolean %s to %v", boolean.Display(), other.Display())
	} else {
		return !boolean.value && as_bool.value, nil
	}
}

// Logical Operators

func (boolean *PrimitiveBool) Not() (Primitive, error) {
	return &PrimitiveBool{!boolean.value}, nil
}

func (boolean *PrimitiveBool) And(other Primitive) (Primitive, error) {
	return &PrimitiveBool{boolean.value && other.Truthy()}, nil
}

func (boolean *PrimitiveBool) AndInplace(other Primitive) error {
	boolean.value = boolean.value && other.Truthy()
	return nil
}

func (boolean *PrimitiveBool) Or(other Primitive) (Primitive, error) {
	return &PrimitiveBool{boolean.value || other.Truthy()}, nil
}

func (boolean *PrimitiveBool) OrInplace(other Primitive) error {
	boolean.value = boolean.value || other.Truthy()
	return nil
}

// Non logical Operators

func (boolean *PrimitiveBool) Negate() (Primitive, error) {
	return nil, fmt.Errorf("bool type does not support numeral arithmetic")
}

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

func (boolean *PrimitiveBool) CastExplicitNum() (*PrimitiveNum, error) {
	if boolean.value {
		return &PrimitiveNum{*big.NewRat(1, 1)}, nil
	}
	return &PrimitiveNum{*big.NewRat(0, 1)}, nil
}

func (boolean *PrimitiveBool) Display() string {
	if boolean.value {
		return "true"
	} else {
		return "false"
	}
}

func (boolean *PrimitiveBool) Truthy() bool {
	return boolean.value
}
