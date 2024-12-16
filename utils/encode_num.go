package utils

import (
	"math/big"
)

func EncodeNum(value *big.Rat) string {
	if prec, exact := value.FloatPrec(); exact {
		return value.FloatString(prec)
	} else {
		return value.FloatString(prec + 6)
	}
}
