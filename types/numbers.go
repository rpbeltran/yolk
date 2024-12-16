package types

import (
	"fmt"
	"math/big"
	"yolk/utils"
)

type PrimitiveNum struct {
	value big.Rat
}

func MakeNumber(value string) (*PrimitiveNum, error) {
	var num big.Rat
	if _, success := num.SetString(value); !success {
		return nil, fmt.Errorf("cannot parse %q into a rational number", value)
	}
	return &PrimitiveNum{num}, nil
}

func AsPrimitiveNumber(value *big.Rat) *PrimitiveNum {
	return &PrimitiveNum{value: *value}
}

func (num *PrimitiveNum) Display() string {
	return utils.EncodeNum(&num.value)
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
