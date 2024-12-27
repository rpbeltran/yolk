package types

import (
	"testing"
)

// Test logical operators

type binop_bool_test_case struct {
	a bool
	b bool
	c bool
}

func TestBooleanAnd(t *testing.T) {
	test_cases := []binop_bool_test_case{
		{false, false, false},
		{false, true, false},
		{true, false, false},
		{true, true, true},
	}
	for _, test := range test_cases {
		a_bool := MakeBool(test.a)
		b_bool := MakeBool(test.b)
		if actual, err := a_bool.And(b_bool); err != nil {
			t.Fatalf("Got %t and %t returned error %v, expected %t", test.a, test.b, err, test.c)
		} else if actual_bool, err := actual.RequireBool(); err != nil {
			t.Fatalf("%t and %t should have returned a bool, but RequireBool() failed with: %v", test.a, test.b, err)
		} else if actual_bool.value != test.c {
			t.Fatalf("Got %t and %t = %t, expected %t", test.a, test.b, actual_bool.value, test.c)
		}
	}
}

func TestBooleanOr(t *testing.T) {
	test_cases := []binop_bool_test_case{
		{false, false, false},
		{false, true, true},
		{true, false, true},
		{true, true, true},
	}
	for _, test := range test_cases {
		a_bool := MakeBool(test.a)
		b_bool := MakeBool(test.b)
		if actual, err := a_bool.Or(b_bool); err != nil {
			t.Fatalf("Got %t or %t returned error %v, expected %t", test.a, test.b, err, test.c)
		} else if actual_bool, err := actual.RequireBool(); err != nil {
			t.Fatalf("%t or %t should have returned a bool, but RequireBool() failed with: %v", test.a, test.b, err)
		} else if actual_bool.value != test.c {
			t.Fatalf("Got %t or %t = %t, expected %t", test.a, test.b, actual_bool.value, test.c)
		}
	}
}

// Non logical operators

func TestBoolNonlogicalOps(t *testing.T) {
	if _, err := MakeBool(true).Add(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}
	if _, err := MakeBool(true).Add(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}

	if err := MakeBool(true).AddInplace(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}
	if err := MakeBool(true).AddInplace(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}

	if _, err := MakeBool(true).Subtract(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}
	if _, err := MakeBool(true).Subtract(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}

	if err := MakeBool(true).SubtractInplace(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}
	if err := MakeBool(true).SubtractInplace(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}

	if _, err := MakeBool(true).Multiply(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}
	if _, err := MakeBool(true).Multiply(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}

	if err := MakeBool(true).MultiplyInplace(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}
	if err := MakeBool(true).MultiplyInplace(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}

	if _, err := MakeBool(true).Divide(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}
	if _, err := MakeBool(true).Divide(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}

	if err := MakeBool(true).DivideInplace(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}
	if err := MakeBool(true).DivideInplace(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}

	if _, err := MakeBool(true).IntDivide(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}
	if _, err := MakeBool(true).IntDivide(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}

	if err := MakeBool(true).IntDivideInplace(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}
	if err := MakeBool(true).IntDivideInplace(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}

	if _, err := MakeBool(true).Modulo(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}
	if _, err := MakeBool(true).Modulo(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}

	if err := MakeBool(true).ModuloInplace(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}
	if err := MakeBool(true).ModuloInplace(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}

	if _, err := MakeBool(true).RaisePower(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}
	if _, err := MakeBool(true).RaisePower(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}

	if err := MakeBool(true).RaisePowerInplace(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}
	if err := MakeBool(true).RaisePowerInplace(MakeBool(false)); err == nil {
		t.Fatal(`true + false succeeded but should have failed`)
	}

	if _, err := MakeBool(true).Concatenate(MakeBool(false)); err == nil {
		t.Fatal(`true ++ false succeeded but should have failed`)
	}
	if _, err := MakeBool(true).Concatenate(MakeBool(false)); err == nil {
		t.Fatal(`true ++ false succeeded but should have failed`)
	}

	if err := MakeBool(true).ConcatenateInPlace(MakeBool(false)); err == nil {
		t.Fatal(`true ++ false succeeded but should have failed`)
	}
	if err := MakeBool(true).ConcatenateInPlace(MakeBool(false)); err == nil {
		t.Fatal(`true ++ false succeeded but should have failed`)
	}
}

// Casting

func TestBoolRequireNum(t *testing.T) {
	if _, err := MakeBool(true).RequireNum(); err == nil {
		t.Fatalf("MakeBool(true).RequireNum() succeeded but should have failed")
	}
}
func TestBoolRequireStr(t *testing.T) {
	if _, err := MakeBool(true).RequireStr(); err == nil {
		t.Fatalf("MakeBool(true).RequireNum() succeeded but should have failed")
	}
}

func TestBoolRequireBool(t *testing.T) {
	if actual, err := MakeBool(false).RequireBool(); err != nil {
		t.Fatalf("MakeBool(false).RequireBool() failed with error: %v", err)
	} else if actual.value {
		t.Fatalf("MakeBool(false).RequireBool() gave true, expected false")
	}
}

func TestBoolCastImplicit(t *testing.T) {
	if value, err := MakeBool(true).CastImplicitNum(); err == nil {
		t.Fatalf("MakeBool(true).CastNum() gave %v, expected to fail with an error:", value)
	}
}

func TestBoolCastExplicit(t *testing.T) {
	if value, err := MakeBool(true).CastExplicitNum(); err != nil {
		t.Fatalf("MakeBool(true).CastNum() returned the error %v but should have succeeded", err)
	} else if actual := value.Display(); actual != "1" {
		t.Fatalf("MakeBool(true).CastNum() returned %s, expected 1", actual)
	}

	if value, err := MakeBool(false).CastExplicitNum(); err != nil {
		t.Fatalf("MakeBool(false).CastNum() returned the error %v but should have succeeded", err)
	} else if actual := value.Display(); actual != "0" {
		t.Fatalf("MakeBool(false).CastNum() returned %s, expected 0", actual)
	}
}

func TestBoolDisplay(t *testing.T) {
	if actual := MakeBool(true).Display(); actual != "true" {
		t.Fatalf("MakeBool(true).Display() returned %q, expected %q", actual, "true")
	}
	if actual := MakeBool(false).Display(); actual != "false" {
		t.Fatalf("MakeBool(false).Display() returned %q, expected %q", actual, "false")
	}
}

func TestTruthy(t *testing.T) {
	if MakeBool(false).Truthy() {
		t.Fatalf("MakeBool(false).Truthy() gave true, expected false")
	}

	if !MakeBool(true).Truthy() {
		t.Fatalf("MakeBool(true).Truthy() gave false, expected true")
	}
}
