package types

import (
	"fmt"
	"math/big"
	"strconv"
	"yolk/utils"
)

type PrimitiveInt struct {
	value int64
}

func MakeInt(value int64) *PrimitiveInt {
	return &PrimitiveInt{value}
}

func (integer *PrimitiveInt) Type() string {
	return "int"
}

// Comparisons

func (integer *PrimitiveInt) Equal(other Primitive) bool {
	if as_int, err := other.CastImplicitInt(); err == nil {
		return integer.value == as_int.value
	}
	return false
}

func (integer *PrimitiveInt) LessThan(other Primitive) (bool, error) {
	if as_int, err := other.CastImplicitInt(); err == nil {
		return integer.value < as_int.value, nil
	}

	if as_num, err := other.CastImplicitNum(); err != nil {
		return false, fmt.Errorf("cannot compare int to %s", other.Type())
	} else {
		return big.NewRat(integer.value, 1).Cmp(&as_num.value) == -1, nil
	}
}

// Operators

func (integer *PrimitiveInt) Negate() (Primitive, error) {
	return &PrimitiveInt{-integer.value}, nil
}

func (integer *PrimitiveInt) Add(other Primitive) (Primitive, error) {
	if as_int, err := other.CastImplicitInt(); err == nil {
		return &PrimitiveInt{integer.value + as_int.value}, nil
	}
	if as_num, err := other.CastImplicitNum(); err != nil {
		return nil, fmt.Errorf("cannot add int and %v", other.Display())
	} else {
		var sum big.Rat
		sum.Add(big.NewRat(integer.value, 1), &as_num.value)
		return &PrimitiveNum{sum}, nil
	}
}

func (integer *PrimitiveInt) AddInplace(other Primitive) error {
	if other_int, err := other.CastImplicitInt(); err != nil {
		return fmt.Errorf("cannot add inplace int and %v", other.Display())
	} else {
		integer.value += other_int.value
		return nil
	}
}

func (integer *PrimitiveInt) Subtract(other Primitive) (Primitive, error) {
	if as_int, err := other.CastImplicitInt(); err == nil {
		return &PrimitiveInt{integer.value - as_int.value}, nil
	}
	if as_num, err := other.CastImplicitNum(); err != nil {
		return nil, fmt.Errorf("cannot subtract int and %s", other.Display())
	} else {
		var difference big.Rat
		difference.Sub(big.NewRat(integer.value, 1), &as_num.value)
		return &PrimitiveNum{difference}, nil
	}
}

func (integer *PrimitiveInt) SubtractInplace(other Primitive) error {
	if other_int, err := other.CastImplicitInt(); err != nil {
		return fmt.Errorf("cannot subtract inplace int and %s", other.Display())
	} else {
		integer.value -= other_int.value
		return nil
	}
}

func (integer *PrimitiveInt) Multiply(other Primitive) (Primitive, error) {
	if as_int, err := other.CastImplicitInt(); err == nil {
		return &PrimitiveInt{integer.value * as_int.value}, nil
	}
	if as_num, err := other.CastImplicitNum(); err != nil {
		return nil, fmt.Errorf("cannot multiply int and %s", other.Display())
	} else {
		var product big.Rat
		product.Mul(big.NewRat(integer.value, 1), &as_num.value)
		return &PrimitiveNum{product}, nil
	}
}

func (integer *PrimitiveInt) MultiplyInplace(other Primitive) error {
	if other_int, err := other.CastImplicitInt(); err != nil {
		return fmt.Errorf("cannot multiply inplace int and %s", other.Display())
	} else {
		integer.value *= other_int.value
		return nil
	}
}

func (integer *PrimitiveInt) Divide(other Primitive) (Primitive, error) {
	if as_int, err := other.CastImplicitInt(); err == nil {
		if as_int.value == 0 {
			return nil, fmt.Errorf("cannot compute division by zero")
		} else if integer.value%as_int.value == 0 {
			return &PrimitiveInt{integer.value / as_int.value}, nil
		}
	}
	if as_num, err := other.CastImplicitNum(); err != nil {
		return nil, fmt.Errorf("cannot divide int by %v", other.Display())
	} else {
		var quo big.Rat
		quo.Quo(big.NewRat(integer.value, 1), &as_num.value)
		return &PrimitiveNum{quo}, nil
	}
}

func (integer *PrimitiveInt) DivideInplace(other Primitive) error {
	if other_int, err := other.CastImplicitInt(); err != nil {
		return err
	} else if other_int.value == 0 {
		return fmt.Errorf("cannot compute division by zero")
	} else if integer.value != 0 && (integer.value%other_int.value) != 0 {
		return fmt.Errorf("cannot assign result of fractional division to integer")
	} else {
		integer.value /= other_int.value
		return nil
	}
}

func (integer *PrimitiveInt) IntDivide(other Primitive) (Primitive, error) {
	if as_int, err := other.CastImplicitInt(); err == nil {
		if as_int.value == 0 {
			return nil, fmt.Errorf("cannot compute division by zero")
		} else {
			return &PrimitiveInt{integer.value / as_int.value}, nil
		}
	}
	if as_num, err := other.CastImplicitNum(); err != nil {
		return nil, fmt.Errorf("cannot integer-divide int by %s", other.Display())
	} else {
		var quo big.Rat
		quo.Quo(big.NewRat(integer.value, 1), &as_num.value)
		utils.TruncateInPlace(&quo)
		return &PrimitiveNum{quo}, nil
	}
}

func (integer *PrimitiveInt) IntDivideInplace(other Primitive) error {
	if other_int, err := other.CastImplicitInt(); err != nil {
		return err
	} else if other_int.value == 0 {
		return fmt.Errorf("cannot compute division by zero")
	} else {
		integer.value /= other_int.value
		return nil
	}
}

func (integer *PrimitiveInt) Modulo(other Primitive) (Primitive, error) {
	if other_int, err := other.CastImplicitInt(); err == nil {
		if other_int.value == 0 {
			return nil, fmt.Errorf("cannot compute modulo by zero")
		}
		// Golang implements Euclidean mod where -A % +B is negative
		// the formula below corrects for that
		mod := (integer.value%other_int.value + other_int.value) % other_int.value
		return &PrimitiveInt{mod}, nil
	}
	as_rat := big.NewRat(integer.value, 1)
	if other_num, err := other.CastImplicitNum(); err != nil {
		return nil, fmt.Errorf("cannot compute int mod %s", other.Display())
	} else if mod, err := utils.ModNumber(as_rat, &other_num.value); err != nil {
		return nil, fmt.Errorf("attempting to perform modulus: %w", err)
	} else {
		return &PrimitiveNum{mod}, nil
	}

}

func (integer *PrimitiveInt) ModuloInplace(other Primitive) error {
	if other_int, err := other.CastImplicitInt(); err != nil {
		return err
	} else if other_int.value == 0 {
		return fmt.Errorf("cannot compute modulo by zero")
	} else {
		integer.value = (integer.value%other_int.value + other_int.value) % other_int.value
		return nil
	}
}

func (integer *PrimitiveInt) RaisePower(other Primitive) (Primitive, error) {
	if other_int, err := other.CastImplicitInt(); err == nil && other_int.value >= 0 {
		if other_int.value == 0 {
			return &PrimitiveInt{1}, nil
		}
		result := integer.value
		for range other_int.value - 1 {
			result *= integer.value
		}
		return &PrimitiveInt{result}, nil

	}
	as_rat := big.NewRat(integer.value, 1)
	if other_num, err := other.CastImplicitNum(); err != nil {
		return nil, fmt.Errorf("attempting to raise power: %w", err)
	} else if pow, err := utils.RaisePower(as_rat, &other_num.value); err != nil {
		return nil, fmt.Errorf("attempting to raise power: %w", err)
	} else {
		return &PrimitiveNum{pow}, nil
	}
}

func (integer *PrimitiveInt) RaisePowerInplace(other Primitive) error {
	if other_int, err := other.CastImplicitInt(); err != nil {
		return err
	} else if other_int.value == 0 {
		integer.value = 1
		return nil
	} else if other_int.value < 0 {
		return fmt.Errorf("cannot raise integer to negative power inplace")
	} else {
		multiplier := integer.value
		for range other_int.value - 1 {
			integer.value *= multiplier
		}
		return nil
	}
}

// String operations

func (integer *PrimitiveInt) Concatenate(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("integer type does not support concatenation")
}

func (integer *PrimitiveInt) ConcatenateInPlace(other Primitive) error {
	return fmt.Errorf("integer type does not support concatenation")
}

// Logical Operators

func (integer *PrimitiveInt) And(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("integer type does not support boolean logic")
}

func (integer *PrimitiveInt) AndInplace(other Primitive) error {
	return fmt.Errorf("integer type does not support boolean logic")
}

func (integer *PrimitiveInt) Or(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("integer type does not support boolean logic")
}

func (integer *PrimitiveInt) OrInplace(other Primitive) error {
	return fmt.Errorf("integer type does not support boolean logic")
}

func (integer *PrimitiveInt) Not() (Primitive, error) {
	return nil, fmt.Errorf("integer type does not support boolean logic")
}

// Casting

func (integer *PrimitiveInt) RequireNum() (*PrimitiveNum, error) {
	return nil, fmt.Errorf("integer %s used where a num was required", integer.Display())
}

func (integer *PrimitiveInt) RequireInt() (*PrimitiveInt, error) {
	return integer, nil
}

func (integer *PrimitiveInt) RequireStr() (*PrimitiveStr, error) {
	return nil, fmt.Errorf("integer %s used where a string was required", integer.Display())
}

func (integer *PrimitiveInt) RequireBool() (*PrimitiveBool, error) {
	return nil, fmt.Errorf("integer %s used where a bool was required", integer.Display())
}

func (integer *PrimitiveInt) CastImplicitNum() (*PrimitiveNum, error) {
	return &PrimitiveNum{*big.NewRat(integer.value, 1)}, nil
}

func (integer *PrimitiveInt) CastExplicitNum() (*PrimitiveNum, error) {
	return &PrimitiveNum{*big.NewRat(integer.value, 1)}, nil
}

func (integer *PrimitiveInt) CastImplicitInt() (*PrimitiveInt, error) {
	return integer, nil
}

func (integer *PrimitiveInt) CastExplicitInt() (*PrimitiveInt, error) {
	return integer, nil
}

func (integer *PrimitiveInt) Display() string {
	return strconv.FormatInt(integer.value, 10)
}

func (integer *PrimitiveInt) Truthy() bool {
	return integer.value != 0
}
