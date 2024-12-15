package types

import (
	"fmt"
	"math/big"
)

type PrimitiveNum struct {
	value big.Rat
}

func (num *PrimitiveNum) Display() string {
	as_float, _ := num.value.Float64()
	return fmt.Sprintf("%v", as_float)
}

// Operators

func (num *PrimitiveNum) Add(other Primitive) (Primitive, error) {
	other_num, err := other.RequireNum()
	if err != nil {
		return &PrimitiveNum{}, err
	}
	var sum big.Rat
	sum.Add(&num.value, &other_num.value)
	return &PrimitiveNum{sum}, nil
}

// Casting

func (num *PrimitiveNum) RequireNum() (*PrimitiveNum, error) {
	return num, nil
}

func (num *PrimitiveNum) RequireStr() (*PrimitiveStr, error) {
	return nil, fmt.Errorf("%s used where a string was required", num.Display())
}

func (num *PrimitiveNum) CastNum() (*PrimitiveNum, error) {
	return num, nil
}

func (num *PrimitiveNum) CastStr() (*PrimitiveStr, error) {
	return &PrimitiveStr{num.Display()}, nil
}
