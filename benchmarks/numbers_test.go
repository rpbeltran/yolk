package benchmarks

import (
	"math/big"
	"testing"
	"yolk/types"
	"yolk/utils"
)

func BenchmarkMakeNumber(b *testing.B) {
	value := "13"
	for n := 0; n < b.N; n++ {
		if _, err := types.MakeNumber(value); err != nil {
			b.Fatalf("Error making number 7: %v", err)
		}
	}
}

func BenchmarkNumbersAddition(b *testing.B) {
	// add 2**63-1 to an accumulator b.N times

	inc, err := types.MakeNumber("7")
	if err != nil {
		b.Fatalf("Error making number 7: %v", err)
	}

	var acc types.Primitive

	acc, err = types.MakeNumber("0")
	if err != nil {
		b.Fatalf("Error making number 7: %v", err)
	}

	for n := 0; n < b.N; n++ {
		acc, err = acc.Add(inc)
		if err != nil {
			b.Fatalf("Error perfoming addition: %v", err)
		}
	}
}

func BenchmarkNumbersAdditionInplace(b *testing.B) {
	// add 2**63-1 to an accumulator b.N times

	inc, err := types.MakeNumber("7")
	if err != nil {
		b.Fatalf("Error making number 7: %v", err)
	}

	var acc types.Primitive

	acc, err = types.MakeNumber("0")
	if err != nil {
		b.Fatalf("Error making number 7: %v", err)
	}

	for n := 0; n < b.N; n++ {
		err = acc.AddInplace(inc)
		if err != nil {
			b.Fatalf("Error perfoming addition: %v", err)
		}
	}
}

func BenchmarkNumbersTruncate(b *testing.B) {
	// truncate a number b.N times
	value := big.NewRat(7621354, 3)
	for n := 0; n < b.N; n++ {
		utils.Truncate(value)
	}
}

func BenchmarkNumbersTruncateInplace(b *testing.B) {
	// truncate a number b.N times
	value := big.NewRat(7621354, 3)
	for n := 0; n < b.N; n++ {
		utils.TruncateInPlace(value)
	}
}

// For reference
func BenchmarkGolangAddition(b *testing.B) {
	// add 7 to an accumulator b.N times (for reference)
	acc := 0
	for n := 0; n < b.N; n++ {
		acc += 7
	}
}

// For reference

func BenchmarkGolangBigIntAddition(b *testing.B) {
	// add 7 to an accumulator b.N times (for reference)
	var acc big.Int
	seven := big.NewInt(9223372036854775807)
	for n := 0; n < b.N; n++ {
		acc.Add(&acc, seven)
	}
}

func BenchmarkGolangBigFloatAddition(b *testing.B) {
	// add 7 to an accumulator b.N times (for reference)
	var acc big.Float
	seven := big.NewFloat(9223372036854775807.0)
	for n := 0; n < b.N; n++ {
		acc.Add(&acc, seven)
	}
}

func BenchmarkGolangRationalAddition(b *testing.B) {
	// add 7 to an accumulator b.N times (for reference)
	var acc big.Rat
	seven := big.NewRat(9223372036854775807, 1)
	for n := 0; n < b.N; n++ {
		acc.Add(&acc, seven)
	}
}
