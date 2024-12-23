package utils

import (
	"fmt"
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

func ModNumber(a *big.Rat, b *big.Rat) (*big.Rat, error) {
	if b.Num().BitLen() == 0 {
		return nil, fmt.Errorf("cannot compute mod zero")
	}
	if a.IsInt() && b.IsInt() {
		var mod big.Int
		mod.Mod(a.Num(), b.Num())
		if sign := mod.Sign(); sign != 0 && sign != b.Sign() {
			mod.Add(&mod, b.Num())
		}
		var mod_rat big.Rat
		mod_rat.SetInt(&mod)
		return &mod_rat, nil
	}
	var result big.Rat
	result.Mul(a, result.Inv(b))
	TruncateInPlace(&result)
	result.Mul(&result, b)
	result.Sub(a, &result)
	if sign := result.Sign(); sign != 0 && sign != b.Sign() {
		result.Add(&result, b)
	}
	return &result, nil
}
