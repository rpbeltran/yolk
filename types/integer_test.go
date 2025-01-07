package types

import (
	"testing"
)

func TestIntType(t *testing.T) {
	expected := "int"
	if actual := MakeInt(100).Type(); actual != expected {
		t.Fatalf("Expected (100).type() == %q, instead got %q", expected, actual)
	}
}

func TestIntNonArithmetic(t *testing.T) {
	if actual, err := MakeInt(100).Concatenate(MakeInt(100)); err == nil {
		t.Fatalf("100 ++ 100 should have errored but instead succeeded and returned %s", actual.Display())
	}
	if actual, err := MakeInt(100).Concatenate(MakeString("foo")); err == nil {
		t.Fatalf("100 ++ foo should have errored but instead succeeded and returned %s", actual.Display())
	}
	if err := MakeInt(100).ConcatenateInPlace(MakeInt(100)); err == nil {
		t.Fatalf("100 ++= 100 should have errored but instead succeeded")
	}
	if err := MakeInt(100).ConcatenateInPlace(MakeString("foo")); err == nil {
		t.Fatalf("100 ++= foo should have errored but instead succeeded")
	}

	if actual, err := MakeInt(100).And(MakeBool(true)); err == nil {
		t.Fatalf("100 && true should have errored but instead succeeded and returned %s", actual.Display())
	}
	if actual, err := MakeInt(100).Or(MakeBool(true)); err == nil {
		t.Fatalf("100 || true should have errored but instead succeeded and returned %s", actual.Display())
	}
	if err := MakeInt(100).AndInplace(MakeBool(true)); err == nil {
		t.Fatalf("100 &&= true should have errored but instead succeeded")
	}
	if err := MakeInt(100).OrInplace(MakeBool(true)); err == nil {
		t.Fatalf("100 ||= true should have errored but instead succeeded")
	}

	if _, err := MakeInt(100).Not(); err == nil {
		t.Fatal("!(100) should have errored but instead succeeded")
	}
}

func TestIntNegate(t *testing.T) {
	pos := MakeInt(10)
	neg := MakeInt(-10)
	if actual, err := pos.Negate(); err != nil {
		t.Fatalf("int(10).Negate() returned the error %v but should have succeeded", err)
	} else if !actual.Equal(neg) {
		t.Fatalf("int(10).Negate() returned %s, expected %s", actual.Display(), neg.Display())
	}
	if actual, err := neg.Negate(); err != nil {
		t.Fatalf("int(-10).Negate() returned the error %v but should have succeeded", err)
	} else if !actual.Equal(pos) {
		t.Fatalf("int(-10).Negate() returned %s, expected %s", actual.Display(), pos.Display())
	}
}

type binop_int_test_case struct {
	a                    *PrimitiveInt
	b                    Primitive
	c                    Primitive
	should_error         bool
	should_error_inplace bool
}

// Add

func TestIntAdd(t *testing.T) {
	test_cases := []binop_int_test_case{
		{MakeInt(100), MakeInt(50), MakeInt(150), false, false},
		{MakeInt(100), MakeInt(-50), MakeInt(50), false, false},
		{MakeInt(100), MakeInt(-250), MakeInt(-150), false, false},
		{MakeInt(100), makeNumOrFail("50", t), MakeInt(150), false, false},
		{MakeInt(100), makeNumOrFail("-50", t), MakeInt(50), false, false},
		{MakeInt(100), makeNumOrFail("-250", t), MakeInt(-150), false, false},
		{MakeInt(100), makeNumOrFail("50.5", t), makeNumOrFail("150.5", t), false, true},
		{MakeInt(100), makeNumOrFail("-50.5", t), makeNumOrFail("49.5", t), false, true},
		{MakeInt(100), makeNumOrFail("-250.5", t), makeNumOrFail("-150.5", t), false, true},
		{MakeInt(1), MakeString("1"), nil, true, true},
		{MakeInt(1), MakeBool(true), nil, true, true},
	}
	for _, test := range test_cases {
		if test.should_error {
			if _, err := test.a.Add(test.b); err == nil {
				t.Fatalf("Got %v + %v succeeded, expected an error", test.a, test.b)
			}
		} else if val, err := test.a.Add(test.b); err != nil {
			t.Fatalf("Got %v + %v returned error %v, expected %v", test.a, test.b, err, test.c)
		} else if !val.Equal(test.c) {
			t.Fatalf("Got %v + %v == %v, expected %v", test.a, test.b, val, test.c)
		}

		if test.should_error_inplace {
			if err := test.a.AddInplace(test.b); err == nil {
				t.Fatalf("Got %v += %v succeeded, expected an error", test.a, test.b)
			}
		} else if err := test.a.AddInplace(test.b); err != nil {
			t.Fatalf("Got %v += %v returned error %v, expected %v", test.a, test.b, err, test.c)
		} else if !test.a.Equal(test.c) {
			t.Fatalf("Got %v += %v => %v, expected %v", test.a, test.b, test.a, test.c)
		}
	}
}

func TestIntSubtract(t *testing.T) {
	test_cases := []binop_int_test_case{
		{MakeInt(100), MakeInt(50), MakeInt(50), false, false},
		{MakeInt(100), MakeInt(-50), MakeInt(150), false, false},
		{MakeInt(100), MakeInt(-250), MakeInt(350), false, false},
		{MakeInt(100), makeNumOrFail("50", t), MakeInt(50), false, false},
		{MakeInt(100), makeNumOrFail("-50", t), MakeInt(150), false, false},
		{MakeInt(100), makeNumOrFail("-250", t), MakeInt(350), false, false},
		{MakeInt(100), makeNumOrFail("50.5", t), makeNumOrFail("49.5", t), false, true},
		{MakeInt(100), makeNumOrFail("-50.5", t), makeNumOrFail("150.5", t), false, true},
		{MakeInt(100), makeNumOrFail("-250.5", t), makeNumOrFail("350.5", t), false, true},
		{MakeInt(1), MakeString("1"), nil, true, true},
		{MakeInt(1), MakeBool(true), nil, true, true},
	}
	for _, test := range test_cases {
		if test.should_error {
			if _, err := test.a.Subtract(test.b); err == nil {
				t.Fatalf("Got %v - %v succeeded, expected an error", test.a.Display(), test.b.Display())
			}
		} else if val, err := test.a.Subtract(test.b); err != nil {
			t.Fatalf("Got %v - %v returned error %v, expected %v", test.a.Display(), test.b.Display(), err, test.c.Display())
		} else if !val.Equal(test.c) {
			t.Fatalf("Got %v - %v == %v, expected %v", test.a.Display(), test.b.Display(), val.Display(), test.c.Display())
		}

		if test.should_error_inplace {
			if err := test.a.SubtractInplace(test.b); err == nil {
				t.Fatalf("Got %v -= %v succeeded, expected an error", test.a.Display(), test.b.Display())
			}
		} else if err := test.a.SubtractInplace(test.b); err != nil {
			t.Fatalf("Got %v -= %v returned error %v, expected %v", test.a.Display(), test.b.Display(), err, test.c.Display())
		} else if !test.a.Equal(test.c) {
			t.Fatalf("Got %v -= %v => %v, expected %v", test.a.Display(), test.b.Display(), test.a.Display(), test.c.Display())
		}
	}
}

func TestIntMultiply(t *testing.T) {
	test_cases := []binop_int_test_case{
		{MakeInt(10), MakeInt(50), MakeInt(500), false, false},
		{MakeInt(10), MakeInt(-50), MakeInt(-500), false, false},
		{MakeInt(10), MakeInt(-250), MakeInt(-2500), false, false},
		{MakeInt(10), makeNumOrFail("50", t), MakeInt(500), false, false},
		{MakeInt(10), makeNumOrFail("-50", t), MakeInt(-500), false, false},
		{MakeInt(10), makeNumOrFail("-250", t), MakeInt(-2500), false, false},
		{MakeInt(10), makeNumOrFail("50.5", t), makeNumOrFail("505", t), false, true},
		{MakeInt(10), makeNumOrFail("-50.5", t), makeNumOrFail("-505", t), false, true},
		{MakeInt(10), makeNumOrFail("-250.5", t), makeNumOrFail("-2505", t), false, true},
		{MakeInt(1), MakeString("1"), nil, true, true},
		{MakeInt(1), MakeBool(true), nil, true, true},
	}
	for _, test := range test_cases {
		if test.should_error {
			if _, err := test.a.Multiply(test.b); err == nil {
				t.Fatalf("Got %v * %v succeeded, expected an error", test.a.Display(), test.b.Display())
			}
		} else if val, err := test.a.Multiply(test.b); err != nil {
			t.Fatalf("Got %v * %v returned error %v, expected %v", test.a.Display(), test.b.Display(), err, test.c.Display())
		} else if !val.Equal(test.c) {
			t.Fatalf("Got %v * %v == %v, expected %v", test.a.Display(), test.b.Display(), val.Display(), test.c.Display())
		}

		if test.should_error_inplace {
			if err := test.a.MultiplyInplace(test.b); err == nil {
				t.Fatalf("Got %v *= %v succeeded, expected an error", test.a.Display(), test.b.Display())
			}
		} else if err := test.a.MultiplyInplace(test.b); err != nil {
			t.Fatalf("Got %v *= %v returned error %v, expected %v", test.a.Display(), test.b.Display(), err, test.c.Display())
		} else if !test.a.Equal(test.c) {
			t.Fatalf("Got %v *= %v => %v, expected %v", test.a.Display(), test.b.Display(), test.a.Display(), test.c.Display())
		}
	}
}

func TestIntDivide(t *testing.T) {
	test_cases := []binop_int_test_case{
		{MakeInt(50), MakeInt(5), MakeInt(10), false, false},
		{MakeInt(0), MakeInt(10), MakeInt(0), false, false},
		{MakeInt(50), MakeInt(1000), makeNumOrFail(".05", t), false, true},
		{MakeInt(10), MakeInt(0), nil, true, true},
		{MakeInt(0), MakeInt(0), nil, true, true},
		{MakeInt(50), makeNumOrFail("5", t), MakeInt(10), false, false},
		{MakeInt(0), makeNumOrFail("10", t), MakeInt(0), false, false},
		{MakeInt(10), makeNumOrFail("0", t), nil, true, true},
		{MakeInt(0), makeNumOrFail("0", t), nil, true, true},
		{MakeInt(50), makeNumOrFail(".5", t), makeNumOrFail("100", t), false, true},
		{MakeInt(50), makeNumOrFail("2.5", t), makeNumOrFail("20", t), false, true},
		{MakeInt(50), makeNumOrFail("1000", t), makeNumOrFail(".05", t), false, true},
		{MakeInt(1), MakeString("1"), nil, true, true},
		{MakeInt(1), MakeBool(true), nil, true, true},
	}
	for _, test := range test_cases {
		if test.should_error {
			if _, err := test.a.Divide(test.b); err == nil {
				t.Fatalf("Got %v / %v succeeded, expected an error", test.a.Display(), test.b.Display())
			}
		} else if val, err := test.a.Divide(test.b); err != nil {
			t.Fatalf("Got %v / %v returned error %v, expected %v", test.a.Display(), test.b.Display(), err, test.c.Display())
		} else if !val.Equal(test.c) {
			t.Fatalf("Got %v / %v == %v, expected %v", test.a.Display(), test.b.Display(), val.Display(), test.c.Display())
		}

		if test.should_error_inplace {
			if err := test.a.DivideInplace(test.b); err == nil {
				t.Fatalf("Got %v /= %v succeeded, expected an error", test.a.Display(), test.b.Display())
			}
		} else if err := test.a.DivideInplace(test.b); err != nil {
			t.Fatalf("Got %v /= %v returned error %v, expected %v", test.a.Display(), test.b.Display(), err, test.c.Display())
		} else if !test.a.Equal(test.c) {
			t.Fatalf("Got %v /= %v => %v, expected %v", test.a.Display(), test.b.Display(), test.a.Display(), test.c.Display())
		}
	}
}

func TestIntIntegerDivide(t *testing.T) {
	test_cases := []binop_int_test_case{
		{MakeInt(50), MakeInt(5), MakeInt(10), false, false},
		{MakeInt(0), MakeInt(10), MakeInt(0), false, false},
		{MakeInt(50), MakeInt(1000), MakeInt(0), false, false},
		{MakeInt(10), MakeInt(0), nil, true, true},
		{MakeInt(0), MakeInt(0), nil, true, true},
		{MakeInt(50), makeNumOrFail("5", t), MakeInt(10), false, false},
		{MakeInt(0), makeNumOrFail("10", t), MakeInt(0), false, false},
		{MakeInt(10), makeNumOrFail("0", t), nil, true, true},
		{MakeInt(0), makeNumOrFail("0", t), nil, true, true},
		{MakeInt(50), makeNumOrFail(".5", t), makeNumOrFail("100", t), false, true},
		{MakeInt(50), makeNumOrFail("2.5", t), makeNumOrFail("20", t), false, true},
		{MakeInt(50), makeNumOrFail("1000", t), MakeInt(0), false, false},
		{MakeInt(1), MakeString("1"), nil, true, true},
		{MakeInt(1), MakeBool(true), nil, true, true},
	}
	for _, test := range test_cases {
		if test.should_error {
			if _, err := test.a.IntDivide(test.b); err == nil {
				t.Fatalf("Got %v // %v succeeded, expected an error", test.a.Display(), test.b.Display())
			}
		} else if val, err := test.a.IntDivide(test.b); err != nil {
			t.Fatalf("Got %v // %v returned error %v, expected %v", test.a.Display(), test.b.Display(), err, test.c.Display())
		} else if !val.Equal(test.c) {
			t.Fatalf("Got %v // %v == %v, expected %v", test.a.Display(), test.b.Display(), val.Display(), test.c.Display())
		}

		if test.should_error_inplace {
			if err := test.a.IntDivideInplace(test.b); err == nil {
				t.Fatalf("Got %v //= %v succeeded, expected an error", test.a.Display(), test.b.Display())
			}
		} else if err := test.a.IntDivideInplace(test.b); err != nil {
			t.Fatalf("Got %v //= %v returned error %v, expected %v", test.a.Display(), test.b.Display(), err, test.c.Display())
		} else if !test.a.Equal(test.c) {
			t.Fatalf("Got %v //= %v => %v, expected %v", test.a.Display(), test.b.Display(), test.a.Display(), test.c.Display())
		}
	}
}

func TestIntModulo(t *testing.T) {
	test_cases := []binop_int_test_case{
		{MakeInt(50), MakeInt(5), MakeInt(0), false, false},
		{MakeInt(0), MakeInt(10), MakeInt(0), false, false},
		{MakeInt(12), MakeInt(5), MakeInt(2), false, false},
		{MakeInt(-10), MakeInt(4), MakeInt(2), false, false},
		{MakeInt(-10), MakeInt(3), MakeInt(2), false, false},
		{MakeInt(50), MakeInt(0), nil, true, true},

		{MakeInt(50), makeNumOrFail("5", t), MakeInt(0), false, false},
		{MakeInt(0), makeNumOrFail("10", t), MakeInt(0), false, false},
		{MakeInt(12), makeNumOrFail("5", t), MakeInt(2), false, false},
		{MakeInt(-10), makeNumOrFail("4", t), MakeInt(2), false, false},
		{MakeInt(-10), makeNumOrFail("3", t), MakeInt(2), false, false},
		{MakeInt(50), makeNumOrFail("0", t), nil, true, true},

		{MakeInt(10), makeNumOrFail("3.3", t), makeNumOrFail("0.1", t), false, true},
		{MakeInt(-10), makeNumOrFail("3.3", t), makeNumOrFail("3.2", t), false, true},
		{MakeInt(10), makeNumOrFail("-3.3", t), makeNumOrFail("-3.2", t), false, true},
		{MakeInt(-10), makeNumOrFail("-3.3", t), makeNumOrFail("-0.1", t), false, true},

		{MakeInt(1), MakeString("1"), nil, true, true},
		{MakeInt(1), MakeBool(true), nil, true, true},
	}
	for _, test := range test_cases {
		if test.should_error {
			if _, err := test.a.Modulo(test.b); err == nil {
				t.Fatalf("Got %v mod %v succeeded, expected an error", test.a.Display(), test.b.Display())
			}
		} else if val, err := test.a.Modulo(test.b); err != nil {
			t.Fatalf("Got %v mod %v returned error %v, expected %v", test.a.Display(), test.b.Display(), err, test.c.Display())
		} else if !val.Equal(test.c) {
			t.Fatalf("Got %v mod %v == %v, expected %v", test.a.Display(), test.b.Display(), val.Display(), test.c.Display())
		}

		if test.should_error_inplace {
			if err := test.a.ModuloInplace(test.b); err == nil {
				t.Fatalf("Got %v mod %v (inplace) succeeded, expected an error", test.a.Display(), test.b.Display())
			}
		} else if err := test.a.ModuloInplace(test.b); err != nil {
			t.Fatalf("Got %v mod %v (in place) returned error %v, expected %v", test.a.Display(), test.b.Display(), err, test.c.Display())
		} else if !test.a.Equal(test.c) {
			t.Fatalf("Got %v mod %v (inplace) => %v, expected %v", test.a.Display(), test.b.Display(), test.a.Display(), test.c.Display())
		}
	}
}

func TestIntPower(t *testing.T) {
	test_cases := []binop_int_test_case{
		{MakeInt(2), MakeInt(3), MakeInt(8), false, false},
		{MakeInt(101), MakeInt(0), MakeInt(1), false, false},
		{MakeInt(100), MakeInt(2), MakeInt(10000), false, false},
		{MakeInt(2), makeNumOrFail("3", t), MakeInt(8), false, false},
		{MakeInt(101), makeNumOrFail("0", t), MakeInt(1), false, false},
		{MakeInt(100), makeNumOrFail("2", t), MakeInt(10000), false, false},
		{MakeInt(100), makeNumOrFail("-2", t), makeNumOrFail(".0001", t), false, true},
		{MakeInt(100), makeNumOrFail("0.5", t), makeNumOrFail("10", t), false, true},
		{MakeInt(-100), makeNumOrFail("0.5", t), nil, true, true},
		{MakeInt(-100), makeNumOrFail("2.5", t), nil, true, true},
		{MakeInt(-100), makeNumOrFail("-2.5", t), nil, true, true},
		{MakeInt(1), MakeString("1"), nil, true, true},
		{MakeInt(1), MakeBool(true), nil, true, true},
	}
	for _, test := range test_cases {
		if test.should_error {
			if _, err := test.a.RaisePower(test.b); err == nil {
				t.Fatalf("Got %v + %v succeeded, expected an error", test.a.Display(), test.b.Display())
			}
		} else if val, err := test.a.RaisePower(test.b); err != nil {
			t.Fatalf("Got %v ** %v returned error %v, expected %v", test.a.Display(), test.b.Display(), err, test.c.Display())
		} else if !val.Equal(test.c) {
			t.Fatalf("Got %v ** %v == %v, expected %v", test.a.Display(), test.b.Display(), val.Display(), test.c.Display())
		}

		if test.should_error_inplace {
			if err := test.a.RaisePowerInplace(test.b); err == nil {
				t.Fatalf("Got %v **= %v succeeded, expected an error", test.a, test.b)
			}
		} else if err := test.a.RaisePowerInplace(test.b); err != nil {
			t.Fatalf("Got %v **= %v returned error %v, expected %v", test.a, test.b, err, test.c)
		} else if !test.a.Equal(test.c) {
			t.Fatalf("Got %v **= %v => %v, expected %v", test.a, test.b, test.a, test.c)
		}
	}
}

// Casting

func TestIntRequireNum(t *testing.T) {
	if _, err := MakeInt(10).RequireNum(); err == nil {
		t.Fatalf("PrimitiveNum().RequireNum() succeeded but should have failed")
	}
}

func TestIntRequireInt(t *testing.T) {
	if actual, err := MakeInt(10).RequireInt(); err != nil {
		t.Fatalf("PrimitiveNum().RequireNum() failed with %v", err)
	} else if actual.value != 10 {
		t.Fatalf("PrimitiveNum().RequireNum() gave %d, expected %d", actual.value, 10)
	}
}

func TestIntCastExplicitInt(t *testing.T) {
	if actual, err := MakeInt(10).CastExplicitInt(); err != nil {
		t.Fatalf("PrimitiveNum().CastExplicitInt() failed with %v", err)
	} else if actual.value != 10 {
		t.Fatalf("PrimitiveNum().CastExplicitInt() gave %d, expected %d", actual.value, 10)
	}
}

func TestIntCastImplicitInt(t *testing.T) {
	if actual, err := MakeInt(10).CastImplicitInt(); err != nil {
		t.Fatalf("PrimitiveNum().CastImplicitInt() failed with %v", err)
	} else if actual.value != 10 {
		t.Fatalf("PrimitiveNum().CastImplicitInt() gave %d, expected %d", actual.value, 10)
	}
}

func TestIntCastImplicitNum(t *testing.T) {
	integer := MakeInt(10)
	if as_num, err := integer.CastImplicitNum(); err != nil {
		t.Fatalf("PrimitiveNum().CastNum() returned the error %v but should have succeeded", err)
	} else if !as_num.Equal(integer) {
		t.Fatalf("PrimitiveNum().CastNum() returned %s, expected %s", as_num.Display(), integer.Display())
	}
}

func TestIntCastExplicitNum(t *testing.T) {
	integer := MakeInt(10)
	if as_num, err := integer.CastExplicitNum(); err != nil {
		t.Fatalf("PrimitiveNum().CastNum() returned the error %v but should have succeeded", err)
	} else if !as_num.Equal(integer) {
		t.Fatalf("PrimitiveNum().CastNum() returned %s, expected %s", as_num.Display(), integer.Display())
	}
}

func TestIntRequireStr(t *testing.T) {
	if _, err := MakeInt(10).RequireStr(); err == nil {
		t.Fatalf("PrimitiveNum().RequireStr() succeeded but should have failed")
	}
}

func TestIntRequireBool(t *testing.T) {
	if _, err := MakeInt(10).RequireBool(); err == nil {
		t.Fatalf("PrimitiveNum().RequireBool() succeeded but should have failed")
	}
}

func TestIntDisplay(t *testing.T) {
	test_cases := map[*PrimitiveInt]string{
		MakeInt(10):  "10",
		MakeInt(20):  "20",
		MakeInt(30):  "30",
		MakeInt(0):   "0",
		MakeInt(-30): "-30",
		MakeInt(-20): "-20",
		MakeInt(-10): "-10",
	}
	for integer, expected := range test_cases {
		if actual := integer.Display(); actual != expected {
			t.Fatalf("int(%d).Display() returned %q, expected %q", integer.value, actual, expected)
		}
	}
}

func TestIntTruthy(t *testing.T) {
	if !MakeInt(1).Truthy() {
		t.Fatalf("int(1).Truthy() returned false, expected true")
	}

	if !MakeInt(-1).Truthy() {
		t.Fatalf("int(-1).Truthy() returned false, expected true")
	}

	if MakeInt(0).Truthy() {
		t.Fatalf("int(0).Truthy() returned true, expected false")
	}
}

func TestIntEquality(t *testing.T) {
	if !MakeInt(1).Equal(MakeInt(1)) {
		t.Fatal("Equal(1, 1) gave false, expected true")
	}
	if MakeInt(1).Equal(MakeInt(2)) {
		t.Fatal("Equal(1, 1) gave true, expected false")
	}
	if !MakeInt(1).Equal(makeNumOrFail("1", t)) {
		t.Fatal("Equal(1, 1) gave false, expected true")
	}
	if MakeInt(1).Equal(makeNumOrFail("2", t)) {
		t.Fatal("Equal(1, 1) gave true, expected false")
	}
	if MakeInt(1).Equal(MakeBool(true)) {
		t.Fatal("Equal(1, true) gave true, expected false")
	}
	if MakeInt(1).Equal(MakeString("1")) {
		t.Fatal(`Equal("1", 1) gave true, expected false`)
	}
}

func TestIntLessThan(t *testing.T) {
	true_cases := map[*PrimitiveInt]Primitive{
		MakeInt(5): MakeInt(10),
		MakeInt(5): makeNumOrFail("10", t),
		MakeInt(5): makeNumOrFail("5.1", t),
	}

	false_cases := map[*PrimitiveInt]Primitive{
		MakeInt(5): MakeInt(4),
		MakeInt(5): makeNumOrFail("4", t),
		MakeInt(5): makeNumOrFail("4.9", t),
	}

	failure_cases := map[*PrimitiveInt]Primitive{
		MakeInt(5): MakeBool(true),
		MakeInt(5): MakeString("5"),
	}

	for a, b := range true_cases {
		if lt, err := a.LessThan(b); err != nil {
			t.Fatalf("Unexpected error testing %s < %s", a.Display(), b.Display())
		} else if !lt {
			t.Fatalf("Equal(%s, %s) gave true, expected false", a.Display(), b.Display())
		}
	}

	for a, b := range false_cases {
		if lt, err := a.LessThan(b); err != nil {
			t.Fatalf("Unexpected error testing %s < %s", a.Display(), b.Display())
		} else if lt {
			t.Fatalf("Equal(%s, %s) gave true, expected false", a.Display(), b.Display())
		}
	}

	for a, b := range failure_cases {
		if _, err := a.LessThan(b); err == nil {
			t.Fatalf("Eexpected error testing %s < %v, but got none", a.Display(), b)
		}
	}
}
