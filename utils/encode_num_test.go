package utils

import (
	"math/big"
	"strings"
	"testing"
)

func TestEncodeNum(t *testing.T) {
	test_cases := []string{
		"0",
		"123",
		"123456789",
		"0.123",
		"1.234",
		"-1",
		"-123",
		"-123.456789",
		strings.Repeat("123456789", 999),
		"123456789." + strings.Repeat("123456789", 999),
		"-123456789" + strings.Repeat("123456789", 999),
		"-123456789." + strings.Repeat("123456789", 999),
	}
	for _, tc := range test_cases {
		var num big.Rat
		if _, success := num.SetString(tc); !success {
			t.Fatalf("Failed to parse test string %q into rational number", tc)
		}
		if actual := EncodeNum(&num); actual != tc {
			t.Fatalf("Encode(%s) returned %q, expected %q", tc, actual, tc)
		}
	}
}
