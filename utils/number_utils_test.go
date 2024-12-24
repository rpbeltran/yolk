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

func TestModulo(t *testing.T) {
	type testcase struct {
		a string
		b string
		c string
	}

	test_cases := []testcase{
		{"0", "1", "0"},
		{"1", "1", "0"},
		{"13", "10", "3"},
		{"190124873", "10", "3"},
		{"-10", "3", "2"},
		{"10.5", "2", "0.5"},
		{"-10.5", "3", "1.5"},
		{"-10.5", "3.3", "2.7"},
		{"10.5", "-3.3", "-2.7"},
		{"10", "-2", "0"},
		{"10", "-2.5", "0"},
		{"10", "-9", "-8"},
	}
	for _, test := range test_cases {
		var a big.Rat
		if _, success := a.SetString(test.a); !success {
			t.Fatalf("Failed to parse test string %q into rational number", test.a)
		}
		var b big.Rat
		if _, success := b.SetString(test.b); !success {
			t.Fatalf("Failed to parse test string %q into rational number", test.b)
		}
		var expected big.Rat
		if _, success := expected.SetString(test.c); !success {
			t.Fatalf("Failed to parse test string %q into rational number", test.c)
		}

		if actual, err := ModNumber(&a, &b); err != nil {
			t.Fatalf("unexpected error computing ModNumber(%s, %s): %v", test.a, test.b, err)
		} else if actual.Cmp(&expected) != 0 {
			t.Fatalf("ModNumber(%s, %s) gave %s, expected %s", test.a, test.b, EncodeNum(&actual), test.c)
		}

		// Ensure that inputs weren't modified
		if actual := EncodeNum(&a); actual != test.a {
			t.Fatalf("ModNumber(%s, %s) turned %s into %s", test.a, test.b, test.a, actual)
		}

		if actual := EncodeNum(&b); actual != test.b {
			t.Fatalf("ModNumber(%s, %s) turned %s into %s", test.a, test.b, test.b, actual)
		}
	}
}

func TestRaisePower(t *testing.T) {
	type testcase struct {
		a            string
		b            string
		c            string
		should_error bool
	}

	test_cases := []testcase{
		{"0", "1", "0", false},
		{"1", "1", "1", false},
		{"13", "10", "137858491849", false},
		{"10", "-5", "0.00001", false},
		{"-10", "3", "-1000", false},
		{"10.5", "2", "110.25", false},
		{"-10.5", "3", "-1157.625", false},
		{"10", "3.5", "3162.2776601683799526654183864593505859375", false},
		{"10.00000001", "3.5", "3162.27767123635248935897834599018096923828125", false},
		{"10.5", "3.3", "2343.822326711475398042239248752593994140625", false},
		{"-10.5", "-3.3", "", true},
		{"-10.5", "-3.3", "", true},
		{"10.5", "-3", "8/9261", false},
		{"99", "999", "", false},
		{"99", "9999", "", false},
		{"99", "999.1", "", true},
		{"1000000", "0.5", "1000", false},
		{"9", "0.5", "3", false},
	}

	for _, test := range test_cases {
		var a big.Rat
		if _, success := a.SetString(test.a); !success {
			t.Fatalf("Failed to parse test string %q into rational number", test.a)
		}
		var b big.Rat
		if _, success := b.SetString(test.b); !success {
			t.Fatalf("Failed to parse test string %q into rational number", test.b)
		}

		if test.should_error {
			if actual, err := RaisePower(&a, &b); err == nil {
				t.Fatalf("expected error from RaisePower(%s, %s), instead successfully returned %s", test.a, test.b, &actual)
			}
		} else {
			var expected big.Rat

			if actual, err := RaisePower(&a, &b); err != nil {
				t.Fatalf("unexpected error computing RaisePower(%s, %s): %v", test.a, test.b, err)
			} else if test.c == "" {
				// Skip checking value
			} else if _, success := expected.SetString(test.c); !success {
				t.Fatalf("Failed to parse test string %q into rational number", test.c)
			} else if !RatsAlmostEqual(&actual, &expected) {
				t.Fatalf("RaisePower(%s), %s) gave %s, expected %s", test.a, test.b, &actual, &expected)
			}
		}

		// Ensure that inputs weren't modified
		if actual := EncodeNum(&a); actual != test.a {
			t.Fatalf("RaisePower(%s, %s) turned %s into %s", test.a, test.b, test.a, actual)
		}
		if actual := EncodeNum(&b); actual != test.b {
			t.Fatalf("RaisePower(%s, %s) turned %s into %s", test.a, test.b, test.b, actual)
		}
	}
}

func TestRatsAlmostEqual(t *testing.T) {
	equals := map[string]string{
		"1": "1.0000000000000000000001",
		"2": "2",
		"3": "2.9999999999999999999999",
	}
	not_equals := map[string]string{
		"1": "1.00001",
		"2": "2.5",
		"3": "2",
	}

	for a, b := range equals {
		var a_rat, b_rat big.Rat
		a_rat.SetString(a)
		b_rat.SetString(b)
		if !RatsAlmostEqual(&a_rat, &b_rat) {
			t.Fatalf("RatsAlmostEqual(%s, %s) gave false, expected true, diff", a, b)
		}
	}
	for a, b := range not_equals {
		var a_rat, b_rat big.Rat
		a_rat.SetString(a)
		b_rat.SetString(b)
		if RatsAlmostEqual(&a_rat, &b_rat) {
			t.Fatalf("RatsAlmostEqual(%s, %s) gave true, expected false", a, b)
		}
	}
}
