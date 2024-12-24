package types

import (
	"strings"
	"testing"
)

func makeNumOrFail(value string, t *testing.T) *PrimitiveNum {
	num, err := MakeNumber(value)
	if err != nil {
		t.Fatal(err)
	}
	return num
}

func TestNumDisplay(t *testing.T) {
	basic_test_cases := []string{"123", "0", "-123", "0.12345", "-0.102"}
	for _, tc := range basic_test_cases {
		if actual := makeNumOrFail(tc, t).Display(); actual != tc {
			t.Fatalf("makePrimitiveNumber(%q).Display() returned %q, expected %q", tc, actual, tc)
		}
	}

	big_input := "1.23e+123"
	equivalent := "123" + strings.Repeat("0", 121)
	if actual := makeNumOrFail(big_input, t).Display(); actual != equivalent {
		t.Fatalf("makePrimitiveNumber(%q).Display() returned %q, expected %q", big_input, actual, equivalent)
	}
}

func TestNumNonArithmetic(t *testing.T) {
	if actual, err := makeNumOrFail("100", t).Concatenate(makeNumOrFail("100", t)); err == nil {
		t.Fatalf("100 + 100 should have errored but instead succeeded and returned %s", actual.Display())
	}
	if actual, err := makeNumOrFail("100", t).Concatenate(MakeString("foo")); err == nil {
		t.Fatalf("100 + foo should have errored but instead succeeded and returned %s", actual.Display())
	}
	if err := makeNumOrFail("100", t).ConcatenateInPLace(makeNumOrFail("100", t)); err == nil {
		t.Fatalf("100 + 100 should have errored but instead succeeded")
	}
	if err := makeNumOrFail("100", t).ConcatenateInPLace(MakeString("foo")); err == nil {
		t.Fatalf("100 + foo should have errored but instead succeeded")
	}
}

type binop_test_case struct {
	a            string
	b            string
	c            string
	should_error bool
}

// Add

func TestNumAdd(t *testing.T) {
	test_cases := []binop_test_case{
		{"123", "456", "579", false},
		{"12.03", "45.06", "57.09", false},
		{"12", "45.06", "57.06", false},
	}
	for _, test := range test_cases {
		a_num := makeNumOrFail(test.a, t)
		b_num := makeNumOrFail(test.b, t)
		if test.should_error {
			if actual, err := a_num.Add(b_num); err == nil {
				t.Fatalf("%s + %s should have errored but instead succeeded and returned %s", test.a, test.b, actual.Display())
			}
		} else if actual, err := a_num.Add(b_num); err != nil {
			t.Fatalf("Got %s + %s returned error %v, expected %s", test.a, test.b, err, test.c)
		} else if actual_str := actual.Display(); actual_str != test.c {
			t.Fatalf("Got %s + %s = %s, expected %s", test.a, test.b, actual_str, test.c)
		}

		if test.should_error {
			if err := a_num.AddInplace(b_num); err == nil {
				t.Fatalf("%s + %s should have errored but instead succeeded and returned %s", test.a, test.b, &a_num.value)
			}
		} else if err := a_num.AddInplace(b_num); err != nil {
			t.Fatalf("Got %s + %s returned error %v, expected %s", test.a, test.b, err, test.c)
		} else if actual_str := a_num.Display(); actual_str != test.c {
			t.Fatalf("Got %s + %s = %s, expected %s", test.a, test.b, actual_str, test.c)
		}
	}

	if _, err := makeNumOrFail("100", t).Add(&PrimitiveStr{"foo"}); err == nil {
		t.Fatal(`Got 100 + "foo" succeeded where it should have failed`)
	}
	if err := makeNumOrFail("100", t).AddInplace(&PrimitiveStr{"foo"}); err == nil {
		t.Fatal(`Got 100 + "foo" succeeded inplace where it should have failed`)
	}
}

func TestNumSubtract(t *testing.T) {
	test_cases := []binop_test_case{
		{"456", "123", "333", false},
		{"123", "456", "-333", false},
		{"10", "4.5", "5.5", false},
		{"12.1", "1.2", "10.9", false},
	}
	for _, test := range test_cases {
		a_num := makeNumOrFail(test.a, t)
		b_num := makeNumOrFail(test.b, t)
		if test.should_error {
			if actual, err := a_num.Subtract(b_num); err == nil {
				t.Fatalf("%s - %s should have errored but instead succeeded and returned %s", test.a, test.b, actual.Display())
			}
		} else if actual, err := a_num.Subtract(b_num); err != nil {
			t.Fatalf("Got %s - %s returned error %v, expected %s", test.a, test.b, err, test.c)
		} else if actual_str := actual.Display(); actual_str != test.c {
			t.Fatalf("Got %s - %s = %s, expected %s", test.a, test.b, actual_str, test.c)
		}

		if test.should_error {
			if err := a_num.SubtractInplace(b_num); err == nil {
				t.Fatalf("%s - %s should have errored but instead succeeded and returned %s", test.a, test.b, &a_num.value)
			}
		} else if err := a_num.SubtractInplace(b_num); err != nil {
			t.Fatalf("Got %s - %s returned error %v, expected %s", test.a, test.b, err, test.c)
		} else if actual_str := a_num.Display(); actual_str != test.c {
			t.Fatalf("Got %s - %s = %s, expected %s", test.a, test.b, actual_str, test.c)
		}
	}

	if _, err := makeNumOrFail("100", t).Subtract(&PrimitiveStr{"foo"}); err == nil {
		t.Fatal(`Got 100 - "foo" succeeded where it should have failed`)
	}
	if err := makeNumOrFail("100", t).SubtractInplace(&PrimitiveStr{"foo"}); err == nil {
		t.Fatal(`Got 100 - "foo" succeeded inplace where it should have failed`)
	}
}

func TestNumMultiply(t *testing.T) {
	test_cases := []binop_test_case{
		{"456", "123", "56088", false},
		{"10", "4.5", "45", false},
		{"12.1", "1.2", "14.52", false},
		{"12.1", "0", "0", false},
	}
	for _, test := range test_cases {
		a_num := makeNumOrFail(test.a, t)
		b_num := makeNumOrFail(test.b, t)
		if test.should_error {
			if actual, err := a_num.Multiply(b_num); err == nil {
				t.Fatalf("%s * %s should have errored but instead succeeded and returned %s", test.a, test.b, actual.Display())
			}
		} else if actual, err := a_num.Multiply(b_num); err != nil {
			t.Fatalf("Got %s * %s returned error %v, expected %s", test.a, test.b, err, test.c)
		} else if actual_str := actual.Display(); actual_str != test.c {
			t.Fatalf("Got %s * %s = %s, expected %s", test.a, test.b, actual_str, test.c)
		}

		if test.should_error {
			if err := a_num.MultiplyInplace(b_num); err == nil {
				t.Fatalf("%s * %s should have errored but instead succeeded and returned %s", test.a, test.b, &a_num.value)
			}
		} else if err := a_num.MultiplyInplace(b_num); err != nil {
			t.Fatalf("Got %s * %s returned error %v, expected %s", test.a, test.b, err, test.c)
		} else if actual_str := a_num.Display(); actual_str != test.c {
			t.Fatalf("Got %s * %s = %s, expected %s", test.a, test.b, actual_str, test.c)
		}
	}

	if _, err := makeNumOrFail("100", t).Multiply(&PrimitiveStr{"foo"}); err == nil {
		t.Fatal(`Got 100 * "foo" succeeded where it should have failed`)
	}
	if err := makeNumOrFail("100", t).MultiplyInplace(&PrimitiveStr{"foo"}); err == nil {
		t.Fatal(`Got 100 * "foo" succeeded inplace where it should have failed`)
	}
}

func TestNumDivide(t *testing.T) {
	test_cases := []binop_test_case{
		{"1", "100", "0.01", false},
		{"75", "3", "25", false},
		{"75.3", "3", "25.1", false},
		{"100", "2.5", "40", false},
		{"100", "0", "", true},
		{"0", "0", "", true},
	}
	for _, test := range test_cases {
		a_num := makeNumOrFail(test.a, t)
		b_num := makeNumOrFail(test.b, t)
		if test.should_error {
			if actual, err := a_num.Divide(b_num); err == nil {
				t.Fatalf("%s / %s should have errored but instead succeeded and returned %s", test.a, test.b, actual.Display())
			}
		} else if actual, err := a_num.Divide(b_num); err != nil {
			t.Fatalf("Got %s / %s returned error %v, expected %s", test.a, test.b, err, test.c)
		} else if actual_str := actual.Display(); actual_str != test.c {
			t.Fatalf("Got %s / %s = %s, expected %s", test.a, test.b, actual_str, test.c)
		}

		if test.should_error {
			if err := a_num.DivideInplace(b_num); err == nil {
				t.Fatalf("%s / %s should have errored but instead succeeded and returned %s", test.a, test.b, &a_num.value)
			}
		} else if err := a_num.DivideInplace(b_num); err != nil {
			t.Fatalf("Got %s / %s returned error %v, expected %s", test.a, test.b, err, test.c)
		} else if actual_str := a_num.Display(); actual_str != test.c {
			t.Fatalf("Got %s / %s = %s, expected %s", test.a, test.b, actual_str, test.c)
		}
	}

	if _, err := makeNumOrFail("100", t).Divide(&PrimitiveStr{"foo"}); err == nil {
		t.Fatal(`Got 100 / "foo" succeeded where it should have failed`)
	}
	if err := makeNumOrFail("100", t).DivideInplace(&PrimitiveStr{"foo"}); err == nil {
		t.Fatal(`Got 100 / "foo" succeeded inplace where it should have failed`)
	}
}

func TestNumIntDivide(t *testing.T) {
	test_cases := []binop_test_case{
		{"3", "2", "1", false},
		{"5", "6", "0", false},
		{"1", "100", "0", false},
		{"75", "3", "25", false},
		{"75.3", "3", "25", false},
		{"100", "0", "", true},
		{"0", "0", "", true},
	}
	for _, test := range test_cases {
		a_num := makeNumOrFail(test.a, t)
		b_num := makeNumOrFail(test.b, t)
		if test.should_error {
			if actual, err := a_num.IntDivide(b_num); err == nil {
				t.Fatalf("%s // %s should have errored but instead succeeded and returned %s", test.a, test.b, actual.Display())
			}
		} else if actual, err := a_num.IntDivide(b_num); err != nil {
			t.Fatalf("Got %s // %s returned error %v, expected %s", test.a, test.b, err, test.c)
		} else if actual_str := actual.Display(); actual_str != test.c {
			t.Fatalf("Got %s // %s = %s, expected %s", test.a, test.b, actual_str, test.c)
		}

		if test.should_error {
			if err := a_num.IntDivideInplace(b_num); err == nil {
				t.Fatalf("%s // %s should have errored but instead succeeded and returned %s", test.a, test.b, &a_num.value)
			}
		} else if err := a_num.IntDivideInplace(b_num); err != nil {
			t.Fatalf("Got %s // %s returned error %v, expected %s", test.a, test.b, err, test.c)
		} else if actual_str := a_num.Display(); actual_str != test.c {
			t.Fatalf("Got %s // %s = %s, expected %s", test.a, test.b, actual_str, test.c)
		}
	}

	if _, err := makeNumOrFail("100", t).IntDivide(&PrimitiveStr{"foo"}); err == nil {
		t.Fatal(`Got 100 // "foo" succeeded where it should have failed`)
	}
	if err := makeNumOrFail("100", t).IntDivideInplace(&PrimitiveStr{"foo"}); err == nil {
		t.Fatal(`Got 100 // "foo" succeeded inplace where it should have failed`)
	}
}

func TestNumModulo(t *testing.T) {
	test_cases := []binop_test_case{
		{"3", "2", "1", false},
		{"5", "6", "5", false},
		{"1", "100", "1", false},
		{"100", "1", "0", false},
		{"100.456", "1", "0.456", false},
		{"3", "2.5", "0.5", false},
		{"5.1", "1.5", "0.6", false},
		{"100", "0", "", true},
		{"0", "0", "", true},
	}
	for _, test := range test_cases {
		a_num := makeNumOrFail(test.a, t)
		b_num := makeNumOrFail(test.b, t)
		if test.should_error {
			if actual, err := a_num.Modulo(b_num); err == nil {
				t.Fatalf("%s mod %s should have errored but instead succeeded and returned %s", test.a, test.b, actual.Display())
			}
		} else if actual, err := a_num.Modulo(b_num); err != nil {
			t.Fatalf("Got %s mod %s returned error %v, expected %s", test.a, test.b, err, test.c)
		} else if actual_str := actual.Display(); actual_str != test.c {
			t.Fatalf("Got %s mod %s = %s, expected %s", test.a, test.b, actual_str, test.c)
		}

		if test.should_error {
			if err := a_num.ModuloInplace(b_num); err == nil {
				t.Fatalf("%s mod %s should have errored but instead succeeded and returned %s", test.a, test.b, &a_num.value)
			}
		} else if err := a_num.ModuloInplace(b_num); err != nil {
			t.Fatalf("Got %s mod %s returned error %v, expected %s", test.a, test.b, err, test.c)
		} else if actual_str := a_num.Display(); actual_str != test.c {
			t.Fatalf("Got %s mod %s = %s, expected %s", test.a, test.b, actual_str, test.c)
		}
	}

	if _, err := makeNumOrFail("100", t).Modulo(&PrimitiveStr{"foo"}); err == nil {
		t.Fatal(`Got 100 mod "foo" succeeded where it should have failed`)
	}
	if err := makeNumOrFail("100", t).ModuloInplace(&PrimitiveStr{"foo"}); err == nil {
		t.Fatal(`Got 100 mod "foo" succeeded inplace where it should have failed`)
	}
}

func TestNumRaisePower(t *testing.T) {
	test_cases := []binop_test_case{
		{"3", "2", "9", false},
		{"5", "6", "15625", false},
		{"9", "0.5", "3", false},
		{"4", "-0.5", "0.5", false},
		{"4", "-2", "0.0625", false},
		{"-4", "-2", "0.0625", false},
		{"-4", "2", "16", false},
		{"-4", "3", "-64", false},
		{"-4", "-2.5", "", true},
	}
	for _, test := range test_cases {
		a_num := makeNumOrFail(test.a, t)
		b_num := makeNumOrFail(test.b, t)
		if test.should_error {
			if actual, err := a_num.RaisePower(b_num); err == nil {
				t.Fatalf("%s ** %s should have errored but instead succeeded and returned %s", test.a, test.b, actual.Display())
			}
		} else if actual, err := a_num.RaisePower(b_num); err != nil {
			t.Fatalf("Got %s ** %s returned error %v, expected %s", test.a, test.b, err, test.c)
		} else if actual_str := actual.Display(); actual_str != test.c {
			t.Fatalf("Got %s ** %s = %s, expected %s", test.a, test.b, actual_str, test.c)
		}

		if test.should_error {
			if err := a_num.RaisePowerInplace(b_num); err == nil {
				t.Fatalf("%s ** %s should have errored but instead succeeded and returned %s", test.a, test.b, &a_num.value)
			}
		} else if err := a_num.RaisePowerInplace(b_num); err != nil {
			t.Fatalf("Got %s ** %s returned error %v, expected %s", test.a, test.b, err, test.c)
		} else if actual_str := a_num.Display(); actual_str != test.c {
			t.Fatalf("Got %s ** %s = %s, expected %s", test.a, test.b, actual_str, test.c)
		}
	}

	if _, err := makeNumOrFail("100", t).RaisePower(&PrimitiveStr{"foo"}); err == nil {
		t.Fatal(`Got 100 ** "foo" succeeded where it should have failed`)
	}
	if err := makeNumOrFail("100", t).RaisePowerInplace(&PrimitiveStr{"foo"}); err == nil {
		t.Fatal(`Got 100 ** "foo" succeeded inplace where it should have failed`)
	}
}

// Casting

func TestNumRequireNum(t *testing.T) {
	n := "100"
	if val, err := makeNumOrFail(n, t).RequireNum(); err != nil {
		t.Fatalf("PrimitiveNum().RequireNum() returned the error %v but should have succeeded", err)
	} else if actual := val.Display(); actual != n {
		t.Fatalf("PrimitiveNum().RequireNum() returned %s, expected %s", actual, n)
	}
}

func TestNumRequireStr(t *testing.T) {
	n := "100"
	if _, err := makeNumOrFail(n, t).RequireStr(); err == nil {
		t.Fatalf("PrimitiveNum().RequireStr() succeeded but should have failed")
	}
}

func TestNumRequireBool(t *testing.T) {
	if _, err := makeNumOrFail("100", t).RequireBool(); err == nil {
		t.Fatalf("PrimitiveNum().RequireBool() succeeded but should have failed")
	}
}

func TestNumCastNum(t *testing.T) {
	n := "100"
	if val, err := makeNumOrFail(n, t).CastNum(); err != nil {
		t.Fatalf("PrimitiveNum().CastNum() returned the error %v but should have succeeded", err)
	} else if actual := val.Display(); actual != n {
		t.Fatalf("PrimitiveNum().CastNum() returned %s, expected %s", actual, n)
	}
}

func TestNumCastStr(t *testing.T) {
	n := "100"
	if val, err := makeNumOrFail(n, t).CastStr(); err != nil {
		t.Fatalf("PrimitiveNum().CastStr() returned the error %v but should have succeeded", err)
	} else if actual := val.Display(); actual != n {
		t.Fatalf("PrimitiveNum().CastStr() returned %s, expected %s", actual, n)
	}
}

func TestNumCastBool(t *testing.T) {
	if val, err := makeNumOrFail("100", t).CastBool(); err != nil {
		t.Fatalf("PrimitiveNum().CastBool() returned the error %v but should have succeeded", err)
	} else if actual := val.Display(); actual != "true" {
		t.Fatalf("PrimitiveNum().CastBool() returned %s, expected %s", actual, "true")
	}

	if val, err := makeNumOrFail("0", t).CastBool(); err != nil {
		t.Fatalf("PrimitiveNum().CastBool() returned the error %v but should have succeeded", err)
	} else if actual := val.Display(); actual != "false" {
		t.Fatalf("PrimitiveNum().CastBool() returned %s, expected %s", actual, "false")
	}
}
