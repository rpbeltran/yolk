package types

import (
	"fmt"
	"math/big"
	"strconv"
	"yolk/utils"
)

type PrimitiveNum struct {
	value big.Rat
}

func MakeNumber(value string) (*PrimitiveNum, error) {
	if value, err := strconv.Atoi(value); err == nil {
		return &PrimitiveNum{*big.NewRat(int64(value), 1)}, nil
	}
	var num big.Rat
	if _, success := num.SetString(value); !success {
		return nil, fmt.Errorf("cannot parse %q into a rational number", value)
	}
	return &PrimitiveNum{num}, nil
}

// Operators

func (num *PrimitiveNum) Add(other Primitive) (Primitive, error) {
	other_num, err := other.CastImplicitNum()
	if err != nil {
		return &PrimitiveNum{}, fmt.Errorf("attempting to perform addition: %w", err)
	}
	var sum big.Rat
	sum.Add(&num.value, &other_num.value)
	return &PrimitiveNum{sum}, nil
}

func (num *PrimitiveNum) AddInplace(other Primitive) error {
	other_num, err := other.CastImplicitNum()
	if err != nil {
		return err
	}
	num.value.Add(&num.value, &other_num.value)
	return nil
}

func (num *PrimitiveNum) Subtract(other Primitive) (Primitive, error) {
	other_num, err := other.CastImplicitNum()
	if err != nil {
		return &PrimitiveNum{}, fmt.Errorf("attempting to perform subtraction: %w", err)
	}
	var difference big.Rat
	difference.Sub(&num.value, &other_num.value)
	return &PrimitiveNum{difference}, nil
}

func (num *PrimitiveNum) SubtractInplace(other Primitive) error {
	other_num, err := other.CastImplicitNum()
	if err != nil {
		return err
	}
	num.value.Sub(&num.value, &other_num.value)
	return nil
}

func (num *PrimitiveNum) Multiply(other Primitive) (Primitive, error) {
	other_num, err := other.CastImplicitNum()
	if err != nil {
		return &PrimitiveNum{}, fmt.Errorf("attempting to perform multiplication: %w", err)
	}
	var product big.Rat
	product.Mul(&num.value, &other_num.value)
	return &PrimitiveNum{product}, nil
}

func (num *PrimitiveNum) MultiplyInplace(other Primitive) error {
	other_num, err := other.CastImplicitNum()
	if err != nil {
		return err
	}
	num.value.Mul(&num.value, &other_num.value)
	return nil
}

func (num *PrimitiveNum) Divide(other Primitive) (Primitive, error) {
	other_num, err := other.CastImplicitNum()
	if err != nil {
		return &PrimitiveNum{}, fmt.Errorf("attempting to perform division: %w", err)
	}
	if other_num.value.Num().BitLen() == 0 {
		return nil, fmt.Errorf("cannot compute division by zero")
	}
	var quotient big.Rat
	quotient.Quo(&num.value, &other_num.value)
	return &PrimitiveNum{quotient}, nil
}

func (num *PrimitiveNum) DivideInplace(other Primitive) error {
	other_num, err := other.CastImplicitNum()
	if err != nil {
		return err
	}
	if other_num.value.Num().BitLen() == 0 {
		return fmt.Errorf("cannot compute division by zero")
	}
	num.value.Quo(&num.value, &other_num.value)
	return nil
}

func (num *PrimitiveNum) IntDivide(other Primitive) (Primitive, error) {
	other_num, err := other.CastImplicitNum()
	if err != nil {
		return &PrimitiveNum{}, fmt.Errorf("attempting to perform integer division: %w", err)
	}
	if other_num.value.Num().BitLen() == 0 {
		return nil, fmt.Errorf("cannot compute integer division by zero")
	}
	var quotient big.Rat
	quotient.Quo(&num.value, &other_num.value)
	utils.TruncateInPlace(&quotient)
	return &PrimitiveNum{quotient}, nil
}

func (num *PrimitiveNum) IntDivideInplace(other Primitive) error {
	other_num, err := other.CastImplicitNum()
	if err != nil {
		return err
	}
	if other_num.value.Num().BitLen() == 0 {
		return fmt.Errorf("cannot compute integer division by zero")
	}
	num.value.Quo(&num.value, &other_num.value)
	utils.TruncateInPlace(&num.value)
	return nil
}

func (num *PrimitiveNum) Modulo(other Primitive) (Primitive, error) {
	if other_num, err := other.CastImplicitNum(); err != nil {
		return nil, fmt.Errorf("attempting to perform modulus: %w", err)
	} else if mod, err := utils.ModNumber(&num.value, &other_num.value); err != nil {
		return nil, fmt.Errorf("attempting to perform modulus: %w", err)
	} else {
		return &PrimitiveNum{mod}, nil
	}
}

func (num *PrimitiveNum) ModuloInplace(other Primitive) error {
	if other_num, err := other.CastImplicitNum(); err != nil {
		return fmt.Errorf("attempting to perform modulus: %w", err)
	} else if mod, err := utils.ModNumber(&num.value, &other_num.value); err != nil {
		return fmt.Errorf("attempting to perform modulus: %w", err)
	} else {
		num.value.Set(&mod)
		return nil
	}
}

func (num *PrimitiveNum) RaisePower(other Primitive) (Primitive, error) {
	if other_num, err := other.CastImplicitNum(); err != nil {
		return nil, fmt.Errorf("attempting to raise power: %w", err)
	} else if pow, err := utils.RaisePower(&num.value, &other_num.value); err != nil {
		return nil, fmt.Errorf("attempting to raise power: %w", err)
	} else {
		return &PrimitiveNum{pow}, nil
	}
}

func (num *PrimitiveNum) RaisePowerInplace(other Primitive) error {
	if other_num, err := other.RequireNum(); err != nil {
		return fmt.Errorf("attempting to raise power: %w", err)
	} else if pow, err := utils.RaisePower(&num.value, &other_num.value); err != nil {
		return fmt.Errorf("attempting to raise power: %w", err)
	} else {
		num.value.Set(&pow)
		return nil
	}
}

// Nonnumeric operations

func (str *PrimitiveNum) Concatenate(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("num type does not support concatenation")
}

func (str *PrimitiveNum) ConcatenateInPlace(other Primitive) error {
	return fmt.Errorf("num type does not support concatenation")
}

// Logical Operators

func (num *PrimitiveNum) And(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("num type does not support boolean logic")
}

func (num *PrimitiveNum) Or(other Primitive) (Primitive, error) {
	return nil, fmt.Errorf("num type does not support boolean logic")
}

// Casting

func (num *PrimitiveNum) RequireNum() (*PrimitiveNum, error) {
	return num, nil
}

func (num *PrimitiveNum) RequireStr() (*PrimitiveStr, error) {
	return nil, fmt.Errorf("num %s used where a string was required", num.Display())
}

func (num *PrimitiveNum) RequireBool() (*PrimitiveBool, error) {
	return nil, fmt.Errorf("num %s used where a bool was required", num.Display())
}

func (num *PrimitiveNum) CastImplicitNum() (*PrimitiveNum, error) {
	return num, nil
}

func (num *PrimitiveNum) CastExplicitNum() (*PrimitiveNum, error) {
	return num, nil
}

func (num *PrimitiveNum) Display() string {
	return utils.EncodeNum(&num.value)
}

func (num *PrimitiveNum) Truthy() bool {
	return num.value.Num().BitLen() != 0
}
