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

func TruncateInPlace(num *big.Rat) {
	if num.Sign() == -1 {
		TruncateInPlace(num.Abs(num))
		num.Neg(num)
	} else {
		n, d := num.Num(), num.Denom()
		n.Div(n, d)
		d.Set(big.NewInt(1))
	}
}

func Truncate(num *big.Rat) *big.Rat {
	if num.Sign() == -1 {
		var result big.Rat
		result.Abs(num)

		n, d := result.Num(), result.Denom()
		result.SetInt(big.NewInt(0).Div(n, d))
		return result.Neg(&result)
	} else {
		n, d := num.Num(), num.Denom()
		var truncated big.Rat
		truncated.SetInt(big.NewInt(0).Div(n, d))
		return &truncated
	}
}
