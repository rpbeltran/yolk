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

func TestNumType(t *testing.T) {
	expected := "num"
	if actual := makeNumOrFail("100", t).Type(); actual != expected {
		t.Fatalf("Expected (100).type() == %q, instead got %q", expected, actual)
	}
}

func TestNumNonArithmetic(t *testing.T) {
	if actual, err := makeNumOrFail("100", t).Concatenate(makeNumOrFail("100", t)); err == nil {
		t.Fatalf("100 ++ 100 should have errored but instead succeeded and returned %s", actual.Display())
	}
	if actual, err := makeNumOrFail("100", t).Concatenate(MakeString("foo")); err == nil {
		t.Fatalf("100 ++ foo should have errored but instead succeeded and returned %s", actual.Display())
	}
	if err := makeNumOrFail("100", t).ConcatenateInPlace(makeNumOrFail("100", t)); err == nil {
		t.Fatalf("100 ++= 100 should have errored but instead succeeded")
	}
	if err := makeNumOrFail("100", t).ConcatenateInPlace(MakeString("foo")); err == nil {
		t.Fatalf("100 ++= foo should have errored but instead succeeded")
	}

	if actual, err := makeNumOrFail("100", t).And(MakeBool(true)); err == nil {
		t.Fatalf("100 && true should have errored but instead succeeded and returned %s", actual.Display())
	}
	if actual, err := makeNumOrFail("100", t).Or(MakeBool(true)); err == nil {
		t.Fatalf("100 || true should have errored but instead succeeded and returned %s", actual.Display())
	}
	if err := makeNumOrFail("100", t).AndInplace(MakeBool(true)); err == nil {
		t.Fatalf("100 &&= true should have errored but instead succeeded")
	}
	if err := makeNumOrFail("100", t).OrInplace(MakeBool(true)); err == nil {
		t.Fatalf("100 ||= true should have errored but instead succeeded")
	}

	if _, err := makeNumOrFail("100", t).Not(); err == nil {
		t.Fatal("!(100) should have errored but instead succeeded")
	}
}

func TestNumNegate(t *testing.T) {
	pos := makeNumOrFail("10.5", t)
	neg := makeNumOrFail("-10.5", t)
	if actual, err := pos.Negate(); err != nil {
		t.Fatalf("(10.5).Negate() returned the error %v but should have succeeded", err)
	} else if !actual.Equal(neg) {
		t.Fatalf("(10.5).Negate() returned %s, expected %s", actual.Display(), neg.Display())
	}
	if actual, err := neg.Negate(); err != nil {
		t.Fatalf("(-10.5).Negate() returned the error %v but should have succeeded", err)
	} else if !actual.Equal(pos) {
		t.Fatalf("(-10.5).Negate() returned %s, expected %s", actual.Display(), pos.Display())
	}
}

type binop_num_test_case struct {
	a            *PrimitiveNum
	b            Primitive
	c            Primitive
	should_error bool
}

// Add

func TestNumAdd(t *testing.T) {
	test_cases := []binop_num_test_case{
		{makeNumOrFail("123", t), makeNumOrFail("456", t), makeNumOrFail("579", t), false},
		{makeNumOrFail("12.03", t), makeNumOrFail("45.06", t), makeNumOrFail("57.09", t), false},
		{makeNumOrFail("12", t), makeNumOrFail("45.06", t), makeNumOrFail("57.06", t), false},
		{makeNumOrFail("123", t), MakeInt(450), makeNumOrFail("573", t), false},
		{makeNumOrFail("123", t), MakeBool(true), nil, true},
		{makeNumOrFail("123", t), MakeString("hello"), nil, true},
	}
	for _, test := range test_cases {
		if test.should_error {
			if _, err := test.a.Add(test.b); err == nil {
				t.Fatalf("Got %v + %v succeeded, expected an error", test.a.Display(), test.b.Display())
			}
		} else if val, err := test.a.Add(test.b); err != nil {
			t.Fatalf("Got %v + %v returned error %v, expected %v", test.a.Display(), test.b.Display(), err, test.c.Display())
		} else if !val.Equal(test.c) {
			t.Fatalf("Got %v + %v == %v, expected %v", test.a.Display(), test.b.Display(), val.Display(), test.c.Display())
		}

		if test.should_error {
			if err := test.a.AddInplace(test.b); err == nil {
				t.Fatalf("Got %v += %v succeeded, expected an error", test.a.Display(), test.b.Display())
			}
		} else if err := test.a.AddInplace(test.b); err != nil {
			t.Fatalf("Got %v += %v returned error %v, expected %v", test.a.Display(), test.b.Display(), err, test.c.Display())
		} else if !test.a.Equal(test.c) {
			t.Fatalf("Got %v += %v => %v, expected %v", test.a.Display(), test.b.Display(), test.a.Display(), test.c.Display())
		}
	}
}

func TestNumSubtract(t *testing.T) {
	test_cases := []binop_num_test_case{
		{makeNumOrFail("123", t), makeNumOrFail("456", t), makeNumOrFail("-333", t), false},
		{makeNumOrFail("45.06", t), makeNumOrFail("12.03", t), makeNumOrFail("33.03", t), false},
		{makeNumOrFail("12", t), makeNumOrFail("45.06", t), makeNumOrFail("-33.06", t), false},
		{makeNumOrFail("123", t), MakeInt(450), makeNumOrFail("-327", t), false},
		{makeNumOrFail("123", t), MakeBool(true), nil, true},
		{makeNumOrFail("123", t), MakeString("hello"), nil, true},
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

		if test.should_error {
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

func TestNumMultiply(t *testing.T) {
	test_cases := []binop_num_test_case{
		{makeNumOrFail("10", t), MakeInt(50), makeNumOrFail("500", t), false},
		{makeNumOrFail(".1", t), MakeInt(50), makeNumOrFail("5", t), false},
		{makeNumOrFail("10", t), MakeInt(-50), makeNumOrFail("-500", t), false},
		{makeNumOrFail("10", t), MakeInt(-250), makeNumOrFail("-2500", t), false},
		{makeNumOrFail("10", t), makeNumOrFail("50", t), makeNumOrFail("500", t), false},
		{makeNumOrFail("10", t), makeNumOrFail("-50", t), makeNumOrFail("-500", t), false},
		{makeNumOrFail("10", t), makeNumOrFail("-250", t), makeNumOrFail("-2500", t), false},
		{makeNumOrFail("10", t), makeNumOrFail("50.5", t), makeNumOrFail("505", t), false},
		{makeNumOrFail(".001", t), makeNumOrFail("50.5", t), makeNumOrFail(".0505", t), false},
		{makeNumOrFail("10", t), makeNumOrFail("-50.5", t), makeNumOrFail("-505", t), false},
		{makeNumOrFail("10", t), makeNumOrFail("-250.5", t), makeNumOrFail("-2505", t), false},
		{makeNumOrFail("1", t), MakeString("1"), nil, true},
		{makeNumOrFail("1", t), MakeBool(true), nil, true},
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

		if test.should_error {
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

func TestNumDivide(t *testing.T) {
	test_cases := []binop_num_test_case{
		{makeNumOrFail("50.5", t), MakeInt(5), makeNumOrFail("10.1", t), false},
		{makeNumOrFail("50", t), MakeInt(5), makeNumOrFail("10", t), false},
		{makeNumOrFail("0", t), MakeInt(10), makeNumOrFail("0", t), false},
		{makeNumOrFail("50", t), MakeInt(1000), makeNumOrFail(".05", t), false},
		{makeNumOrFail("10", t), MakeInt(0), nil, true},
		{makeNumOrFail("10.5", t), MakeInt(0), nil, true},
		{makeNumOrFail("0", t), MakeInt(0), nil, true},
		{makeNumOrFail("50", t), makeNumOrFail("5", t), makeNumOrFail("10", t), false},
		{makeNumOrFail("0", t), makeNumOrFail("10", t), makeNumOrFail("0", t), false},
		{makeNumOrFail("10", t), makeNumOrFail("0", t), nil, true},
		{makeNumOrFail("0", t), makeNumOrFail("0", t), nil, true},
		{makeNumOrFail("50", t), makeNumOrFail(".5", t), makeNumOrFail("100", t), false},
		{makeNumOrFail("50", t), makeNumOrFail("2.5", t), makeNumOrFail("20", t), false},
		{makeNumOrFail("50", t), makeNumOrFail("1000", t), makeNumOrFail(".05", t), false},
		{makeNumOrFail("1", t), MakeString("1"), nil, true},
		{makeNumOrFail("1", t), MakeBool(true), nil, true},
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

		if test.should_error {
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

func TestNumIntegerDivide(t *testing.T) {
	test_cases := []binop_num_test_case{
		{makeNumOrFail("50.5132", t), MakeInt(1), makeNumOrFail("50", t), false},
		{makeNumOrFail("50", t), MakeInt(5), makeNumOrFail("10", t), false},
		{makeNumOrFail("0", t), MakeInt(10), makeNumOrFail("0", t), false},
		{makeNumOrFail("50", t), MakeInt(1000), makeNumOrFail("0", t), false},
		{makeNumOrFail("10", t), MakeInt(0), nil, true},
		{makeNumOrFail("0", t), MakeInt(0), nil, true},
		{makeNumOrFail("50", t), makeNumOrFail("5", t), makeNumOrFail("10", t), false},
		{makeNumOrFail("0", t), makeNumOrFail("10", t), makeNumOrFail("0", t), false},
		{makeNumOrFail("10", t), makeNumOrFail("0", t), nil, true},
		{makeNumOrFail("0", t), makeNumOrFail("0", t), nil, true},
		{makeNumOrFail("50", t), makeNumOrFail(".5", t), makeNumOrFail("100", t), false},
		{makeNumOrFail("50", t), makeNumOrFail("2.5", t), makeNumOrFail("20", t), false},
		{makeNumOrFail("50", t), makeNumOrFail("1000", t), makeNumOrFail("0", t), false},
		{makeNumOrFail("1", t), MakeString("1"), nil, true},
		{makeNumOrFail("1", t), MakeBool(true), nil, true},
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

		if test.should_error {
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

func TestNumModulo(t *testing.T) {
	test_cases := []binop_num_test_case{
		{makeNumOrFail("50", t), MakeInt(5), makeNumOrFail("0", t), false},
		{makeNumOrFail("0", t), MakeInt(10), makeNumOrFail("0", t), false},
		{makeNumOrFail("12", t), MakeInt(5), makeNumOrFail("2", t), false},
		{makeNumOrFail("-10", t), MakeInt(4), makeNumOrFail("2", t), false},
		{makeNumOrFail("-10", t), MakeInt(3), makeNumOrFail("2", t), false},
		{makeNumOrFail("50", t), MakeInt(0), nil, true},
		{makeNumOrFail("50", t), makeNumOrFail("5", t), makeNumOrFail("0", t), false},
		{makeNumOrFail("0", t), makeNumOrFail("10", t), makeNumOrFail("0", t), false},
		{makeNumOrFail("12", t), makeNumOrFail("5", t), makeNumOrFail("2", t), false},
		{makeNumOrFail("-10", t), makeNumOrFail("4", t), makeNumOrFail("2", t), false},
		{makeNumOrFail("-10", t), makeNumOrFail("3", t), makeNumOrFail("2", t), false},
		{makeNumOrFail("50", t), makeNumOrFail("0", t), nil, true},
		{makeNumOrFail("10", t), makeNumOrFail("3.3", t), makeNumOrFail("0.1", t), false},
		{makeNumOrFail("-10", t), makeNumOrFail("3.3", t), makeNumOrFail("3.2", t), false},
		{makeNumOrFail("10", t), makeNumOrFail("-3.3", t), makeNumOrFail("-3.2", t), false},
		{makeNumOrFail("-10", t), makeNumOrFail("-3.3", t), makeNumOrFail("-0.1", t), false},
		{makeNumOrFail("50.123", t), MakeInt(5), makeNumOrFail("0.123", t), false},
		{makeNumOrFail("-50.123", t), MakeInt(5), makeNumOrFail("4.877", t), false},
		{makeNumOrFail("1", t), MakeString("1"), nil, true},
		{makeNumOrFail("1", t), MakeBool(true), nil, true},
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

		if test.should_error {
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

func TestNumPower(t *testing.T) {
	test_cases := []binop_num_test_case{
		{makeNumOrFail("2", t), MakeInt(3), makeNumOrFail("8", t), false},
		{makeNumOrFail("2.5", t), MakeInt(3), makeNumOrFail("15.625", t), false},
		{makeNumOrFail("101", t), MakeInt(0), makeNumOrFail("1", t), false},
		{makeNumOrFail("100", t), MakeInt(2), makeNumOrFail("10000", t), false},
		{makeNumOrFail("2", t), makeNumOrFail("3", t), makeNumOrFail("8", t), false},
		{makeNumOrFail("101", t), makeNumOrFail("0", t), makeNumOrFail("1", t), false},
		{makeNumOrFail("100", t), makeNumOrFail("2", t), MakeInt(10000), false},
		{makeNumOrFail("100", t), makeNumOrFail("-2", t), makeNumOrFail(".0001", t), false},
		{makeNumOrFail("100", t), makeNumOrFail("0.5", t), makeNumOrFail("10", t), false},
		{makeNumOrFail("-100", t), makeNumOrFail("0.5", t), nil, true},
		{makeNumOrFail("-100", t), makeNumOrFail("2.5", t), nil, true},
		{makeNumOrFail("-100", t), makeNumOrFail("-2.5", t), nil, true},
		{makeNumOrFail("1", t), MakeString("1"), nil, true},
		{makeNumOrFail("1", t), MakeBool(true), nil, true},
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

		if test.should_error {
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

func TestNumCastImplicitNum(t *testing.T) {
	n := "100"
	if val, err := makeNumOrFail(n, t).CastImplicitNum(); err != nil {
		t.Fatalf("PrimitiveNum().CastNum() returned the error %v but should have succeeded", err)
	} else if actual := val.Display(); actual != n {
		t.Fatalf("PrimitiveNum().CastNum() returned %s, expected %s", actual, n)
	}
}

func TestNumCastExplicitNum(t *testing.T) {
	n := "100"
	if val, err := makeNumOrFail(n, t).CastExplicitNum(); err != nil {
		t.Fatalf("PrimitiveNum().CastNum() returned the error %v but should have succeeded", err)
	} else if actual := val.Display(); actual != n {
		t.Fatalf("PrimitiveNum().CastNum() returned %s, expected %s", actual, n)
	}
}

func TestNumRequireInt(t *testing.T) {
	n := "100"
	if _, err := makeNumOrFail(n, t).RequireInt(); err == nil {
		t.Fatalf("PrimitiveNum(%s).RequireInt() succeeded but should have failed", n)
	}
}

func TestNumCastImplicitInt(t *testing.T) {
	n := "100"
	if val, err := makeNumOrFail(n, t).CastImplicitInt(); err != nil {
		t.Fatalf("PrimitiveNum().CastNum() returned the error %v but should have succeeded", err)
	} else if actual := val.Display(); actual != n {
		t.Fatalf("PrimitiveNum().CastNum() returned %s, expected %s", actual, n)
	}

	n = "100.5"
	if _, err := makeNumOrFail(n, t).CastImplicitInt(); err == nil {
		t.Fatalf("PrimitiveNum(%s).RequireInt() succeeded but should have failed", n)
	}

	n = "999999999999999999999999999999999999999999999999999999999999999999999999"
	if _, err := makeNumOrFail(n, t).CastImplicitInt(); err == nil {
		t.Fatalf("PrimitiveNum(%s).RequireInt() succeeded but should have failed", n)
	}
}

func TestNumCastExplicitInt(t *testing.T) {

	test_cases := map[*PrimitiveNum]*PrimitiveInt{
		makeNumOrFail("100", t):    MakeInt(100),
		makeNumOrFail("100.2", t):  MakeInt(100),
		makeNumOrFail("99.8", t):   MakeInt(99),
		makeNumOrFail("-100", t):   MakeInt(-100),
		makeNumOrFail("-100.9", t): MakeInt(-100),
		makeNumOrFail("-100.1", t): MakeInt(-100),
	}
	for num, expected := range test_cases {
		if val, err := num.CastExplicitInt(); err != nil {
			t.Fatalf("PrimitiveNum().CastNum() returned the error %v but should have succeeded", err)
		} else if !val.Equal(expected) {
			t.Fatalf("PrimitiveNum().CastNum() returned %s, expected %s", val.Display(), expected.Display())
		}
	}

	n := "999999999999999999999999999999999999999999999999999999999999999999999999"
	if _, err := makeNumOrFail(n, t).CastExplicitInt(); err == nil {
		t.Fatalf("PrimitiveNum(%s).RequireInt() succeeded but should have failed", n)
	}
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

func TestNumTruthy(t *testing.T) {
	if !makeNumOrFail("100", t).Truthy() {
		t.Fatalf("PrimitiveNum(100).Truthy() returned false, expected true")
	}

	if makeNumOrFail("0", t).Truthy() {
		t.Fatalf("PrimitiveNum(0).Truthy() returned true, expected false")
	}
}

func TestNumEquality(t *testing.T) {
	if !makeNumOrFail("1", t).Equal(makeNumOrFail("1.0", t)) {
		t.Fatal("Equal(1, 1) gave false, expected true")
	}
	if makeNumOrFail("1", t).Equal(makeNumOrFail("2", t)) {
		t.Fatal("Equal(1, 2) gave true, expected false")
	}

	if makeNumOrFail("1", t).Equal(MakeBool(true)) {
		t.Fatal("Equal(1, true) gave true, expected false")
	}
	if makeNumOrFail("1", t).Equal(MakeString("1")) {
		t.Fatal(`Equal("1", 1) gave true, expected false`)
	}

}

func TestNumLessThan(t *testing.T) {
	if lt, err := makeNumOrFail("5", t).LessThan(makeNumOrFail("-15", t)); err != nil {
		t.Fatalf("Unexpected error testing 5 < -15")
	} else if lt {
		t.Fatal("Equal(5, -15) gave true, expected false")
	}
	if lt, err := makeNumOrFail("5", t).LessThan(makeNumOrFail("5", t)); err != nil {
		t.Fatalf("Unexpected error testing 5 < 5")
	} else if lt {
		t.Fatal("Equal(5, 5) gave true, expected false")
	}
	if lt, err := makeNumOrFail("5", t).LessThan(makeNumOrFail("10", t)); err != nil {
		t.Fatalf("Unexpected error testing 5 < 10")
	} else if !lt {
		t.Fatal("Equal(5, 10) gave false, expected true")
	}

	if lt, err := makeNumOrFail("5", t).LessThan(MakeString("4")); err == nil {
		t.Fatalf("Expected error testing 5 < \"4\", instead succeded and gave %t", lt)
	}
	if lt, err := makeNumOrFail("5", t).LessThan(MakeBool(true)); err == nil {
		t.Fatalf("Expected error testing 5 < true, instead succeded and gave %t", lt)
	}
}
