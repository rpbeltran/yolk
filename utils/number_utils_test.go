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

func TestTruncate(t *testing.T) {
	test_cases := map[string]string{
		"0":       "0",
		"1":       "1",
		"-1":      "-1",
		"3":       "3",
		"-3":      "-3",
		"3.14":    "3",
		"-3.14":   "-3",
		"1.5":     "1",
		"-1.5":    "-1",
		"1.75":    "1",
		"-1.75":   "-1",
		"1.9999":  "1",
		"-1.9999": "-1",
	}
	for test, expected := range test_cases {
		var num big.Rat
		if _, success := num.SetString(test); !success {
			t.Fatalf("Failed to parse test string %q into rational number", test)
		}

		// Test not in place
		if actual := EncodeNum(Truncate(&num)); actual != expected {
			t.Fatalf("Truncate(%q) gave %s, expected %s", test, actual, expected)
		}

		// To ensure that num was not modified
		if actual := EncodeNum(&num); actual != test {
			t.Fatalf("Truncate(%q) turned %s into %s, no modification expected", test, test, actual)
		}

		// Test in place
		TruncateInPlace(&num)
		if actual := EncodeNum(&num); actual != expected {
			t.Fatalf("TruncateInPlace(%q) gave %s, expected %s", test, actual, expected)
		}
	}
}
