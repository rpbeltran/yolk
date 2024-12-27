package types

import (
	"testing"
)

// String Operations

func TestStrConcatenate(t *testing.T) {
	if actual, err := MakeString("hello ").Concatenate(makeNumOrFail("100", t)); err != nil {
		t.Fatalf("\"hello \" ++ 100 should succeeded but instead gave the error: %v", err)
	} else if actual.Display() != "hello 100" {
		t.Fatalf("\"hello \" ++ 100 gave %q, expected %q", actual.Display(), "hello 100")
	}
	if actual, err := MakeString("hello ").Concatenate(MakeString("world")); err != nil {
		t.Fatalf("\"hello \" ++ \"world\" should succeeded but instead gave the error: %v", err)
	} else if actual.Display() != "hello world" {
		t.Fatalf("\"hello \" ++ \"world\" gave %q, expected %q", actual.Display(), "hello world")
	}

	s := MakeString("foo ")
	if err := s.ConcatenateInPlace(makeNumOrFail("100", t)); err != nil {
		t.Fatalf("\"hello \" ++=  100 should succeeded but instead gave the error: %v", err)
	} else if s.Display() != "foo 100" {
		t.Fatalf("\"hello \" ++=  100 gave %q, expected %q", s.Display(), "foo 100")
	}

	s2 := MakeString("foo ")
	if err := s2.ConcatenateInPlace(MakeString("world")); err != nil {
		t.Fatalf("\"hello \" ++=  \"world\" should succeeded but instead gave the error: %v", err)
	} else if s2.Display() != "foo world" {
		t.Fatalf("\"hello \" ++=  \"world\" gave %q, expected %q", s2.Display(), "foo world")
	}
}

// Math

func TestStrMath(t *testing.T) {
	if _, err := MakeString("foo").Add(makeNumOrFail("100", t)); err == nil {
		t.Fatal(`"foo" + 100 succeeded but should have failed`)
	}
	if _, err := MakeString("foo").Add(MakeString("bar")); err == nil {
		t.Fatal(`"foo" + "bar" succeeded but should have failed`)
	}

	if err := MakeString("foo").AddInplace(makeNumOrFail("100", t)); err == nil {
		t.Fatal(`"foo" + 100 succeeded but should have failed`)
	}
	if err := MakeString("foo").AddInplace(MakeString("bar")); err == nil {
		t.Fatal(`"foo" + "bar" succeeded but should have failed`)
	}

	if _, err := MakeString("foo").Subtract(makeNumOrFail("100", t)); err == nil {
		t.Fatal(`"foo" + 100 succeeded but should have failed`)
	}
	if _, err := MakeString("foo").Subtract(MakeString("bar")); err == nil {
		t.Fatal(`"foo" + "bar" succeeded but should have failed`)
	}

	if err := MakeString("foo").SubtractInplace(makeNumOrFail("100", t)); err == nil {
		t.Fatal(`"foo" + 100 succeeded but should have failed`)
	}
	if err := MakeString("foo").SubtractInplace(MakeString("bar")); err == nil {
		t.Fatal(`"foo" + "bar" succeeded but should have failed`)
	}

	if _, err := MakeString("foo").Multiply(makeNumOrFail("100", t)); err == nil {
		t.Fatal(`"foo" + 100 succeeded but should have failed`)
	}
	if _, err := MakeString("foo").Multiply(MakeString("bar")); err == nil {
		t.Fatal(`"foo" + "bar" succeeded but should have failed`)
	}

	if err := MakeString("foo").MultiplyInplace(makeNumOrFail("100", t)); err == nil {
		t.Fatal(`"foo" + 100 succeeded but should have failed`)
	}
	if err := MakeString("foo").MultiplyInplace(MakeString("bar")); err == nil {
		t.Fatal(`"foo" + "bar" succeeded but should have failed`)
	}

	if _, err := MakeString("foo").Divide(makeNumOrFail("100", t)); err == nil {
		t.Fatal(`"foo" + 100 succeeded but should have failed`)
	}
	if _, err := MakeString("foo").Divide(MakeString("bar")); err == nil {
		t.Fatal(`"foo" + "bar" succeeded but should have failed`)
	}

	if err := MakeString("foo").DivideInplace(makeNumOrFail("100", t)); err == nil {
		t.Fatal(`"foo" + 100 succeeded but should have failed`)
	}
	if err := MakeString("foo").DivideInplace(MakeString("bar")); err == nil {
		t.Fatal(`"foo" + "bar" succeeded but should have failed`)
	}

	if _, err := MakeString("foo").IntDivide(makeNumOrFail("100", t)); err == nil {
		t.Fatal(`"foo" + 100 succeeded but should have failed`)
	}
	if _, err := MakeString("foo").IntDivide(MakeString("bar")); err == nil {
		t.Fatal(`"foo" + "bar" succeeded but should have failed`)
	}

	if err := MakeString("foo").IntDivideInplace(makeNumOrFail("100", t)); err == nil {
		t.Fatal(`"foo" + 100 succeeded but should have failed`)
	}
	if err := MakeString("foo").IntDivideInplace(MakeString("bar")); err == nil {
		t.Fatal(`"foo" + "bar" succeeded but should have failed`)
	}

	if _, err := MakeString("foo").Modulo(makeNumOrFail("100", t)); err == nil {
		t.Fatal(`"foo" + 100 succeeded but should have failed`)
	}
	if _, err := MakeString("foo").Modulo(MakeString("bar")); err == nil {
		t.Fatal(`"foo" + "bar" succeeded but should have failed`)
	}

	if err := MakeString("foo").ModuloInplace(makeNumOrFail("100", t)); err == nil {
		t.Fatal(`"foo" + 100 succeeded but should have failed`)
	}
	if err := MakeString("foo").ModuloInplace(MakeString("bar")); err == nil {
		t.Fatal(`"foo" + "bar" succeeded but should have failed`)
	}

	if _, err := MakeString("foo").RaisePower(makeNumOrFail("100", t)); err == nil {
		t.Fatal(`"foo" + 100 succeeded but should have failed`)
	}
	if _, err := MakeString("foo").RaisePower(MakeString("bar")); err == nil {
		t.Fatal(`"foo" + "bar" succeeded but should have failed`)
	}

	if err := MakeString("foo").RaisePowerInplace(makeNumOrFail("100", t)); err == nil {
		t.Fatal(`"foo" + 100 succeeded but should have failed`)
	}
	if err := MakeString("foo").RaisePowerInplace(MakeString("bar")); err == nil {
		t.Fatal(`"foo" + "bar" succeeded but should have failed`)
	}
}

// Casting

func TestStrRequireNum(t *testing.T) {
	if _, err := MakeString("100").RequireNum(); err == nil {
		t.Fatalf("PrimitiveStr().RequireNum() succeeded but should have failed")
	}
}

func TestStrRequireStr(t *testing.T) {
	if val, err := MakeString("100").RequireStr(); err != nil {
		t.Fatalf("PrimitiveStr().RequireStr() returned the error %v but should have succeeded", err)
	} else if actual := val.Display(); actual != "100" {
		t.Fatalf("PrimitiveStr().RequireStr() returned %s, expected %s", actual, "100")
	}
}

func TestStrRequireBool(t *testing.T) {
	if _, err := MakeString("true").RequireBool(); err == nil {
		t.Fatalf("PrimitiveStr().RequireNum() succeeded but should have failed")
	}
}

func TestStrCastNumImplicit(t *testing.T) {
	if value, err := MakeString("100").CastImplicitNum(); err == nil {
		t.Fatalf("PrimitiveStr(%q).CastImplicitNum() gave %v, expected to fail with an error", "100", value)
	}
}

func TestStrCastNumExplicit(t *testing.T) {
	s := "100"
	if value, err := MakeString(s).CastExplicitNum(); err != nil {
		t.Fatalf("PrimitiveStr(%q).CastNum() returned the error %v but should have succeeded", s, err)
	} else if actual := value.Display(); actual != "100" {
		t.Fatalf("PrimitiveStr(%q).CastNum() returned %s, expected %s", s, actual, s)
	}

	if _, err := MakeString("foo").CastExplicitNum(); err == nil {
		t.Fatalf(`PrimitiveStr("foo").RequireStr() succeeded but should have failed`)
	}
}

func TestStrDisplay(t *testing.T) {
	test_cases := []string{"hello", "", "123", " @#()*", `""`}
	for _, tc := range test_cases {
		if actual := MakeString(tc).Display(); actual != tc {
			t.Fatalf("makePrimitiveNumber(%q).Display() returned %q, expected %q", tc, actual, tc)
		}
	}
}

func TestStrTruthy(t *testing.T) {
	if !MakeString("spam").Truthy() {
		t.Fatal(`PrimitiveStr("spam").Truthy() returned false, expected true`)
	}

	if MakeString("").Truthy() {
		t.Fatal(`PrimitiveStr("").Truthy() returned true, expected false`)
	}
}
