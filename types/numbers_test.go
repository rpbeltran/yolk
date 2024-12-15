package types

import (
	"math/big"
	"testing"
)

func makeNumber(value string) *PrimitiveNum {
	var num big.Rat
	num.SetString(value)
	return &PrimitiveNum{num}
}

func TestNumDisplay(t *testing.T) {
	test_cases := []string{"123", "0", "-123", "0.12345", "-0.102", "1.23e+123"}
	for _, tc := range test_cases {
		if actual := makeNumber(tc).Display(); actual != tc {
			t.Fatalf("makePrimitiveNumber(%q).Display() returned %q, expected %q", tc, actual, tc)
		}
	}
}

// Add

func TestNumAddInt(t *testing.T) {
	const (
		a = "123"
		b = "456"
		c = "579"
	)
	if actual, err := makeNumber(a).Add(makeNumber(b)); err != nil {
		t.Fatalf("Got %s + %s returned error %v, expected %s", a, b, err, c)
	} else if actual_str := actual.Display(); actual_str != c {
		t.Fatalf("Got %s + %s = %s, expected %s", a, b, actual_str, c)
	}
}

func TestNumAddFloat(t *testing.T) {
	const (
		a = "12.03"
		b = "45.06"
		c = "57.09"
	)
	if actual, err := makeNumber(a).Add(makeNumber(b)); err != nil {
		t.Fatalf("Got %s + %s returned error %v, expected %s", a, b, err, c)
	} else if actual_str := actual.Display(); actual_str != c {
		t.Fatalf("Got %s + %s = %s, expected %s", a, b, actual_str, c)
	}
}

func TestNumAddMixed(t *testing.T) {
	const (
		a = "12"
		b = "45.06"
		c = "57.06"
	)
	if actual, err := makeNumber(a).Add(makeNumber(b)); err != nil {
		t.Fatalf("Got %s + %s returned error %v, expected %s", a, b, err, c)
	} else if actual_str := actual.Display(); actual_str != c {
		t.Fatalf("Got %s + %s = %s, expected %s", a, b, actual_str, c)
	}
}

func TestNumAddStr(t *testing.T) {
	if _, err := makeNumber("100").Add(&PrimitiveStr{"foo"}); err == nil {
		t.Fatal(`Got 100 + "foo" succeeded where it should have failed`)
	}
}

// Casting

func TestRequireNum(t *testing.T) {
	n := "100"
	if val, err := makeNumber(n).RequireNum(); err != nil {
		t.Fatalf("PrimitiveNum().RequireNum() returned the error %v but should have succeeded", err)
	} else if actual := val.Display(); actual != n {
		t.Fatalf("PrimitiveNum().RequireNum() returned %s, expected %s", actual, n)
	}
}

func TestRequireStr(t *testing.T) {
	n := "100"
	if _, err := makeNumber(n).RequireStr(); err == nil {
		t.Fatalf("PrimitiveNum().RequireStr() succeeded but should have failed")
	}
}

func TestCastNum(t *testing.T) {
	n := "100"
	if val, err := makeNumber(n).CastNum(); err != nil {
		t.Fatalf("PrimitiveNum().CastNum() returned the error %v but should have succeeded", err)
	} else if actual := val.Display(); actual != n {
		t.Fatalf("PrimitiveNum().CastNum() returned %s, expected %s", actual, n)
	}
}

func TestCastStr(t *testing.T) {
	n := "100"
	if val, err := makeNumber(n).CastStr(); err != nil {
		t.Fatalf("PrimitiveNum().CastStr() returned the error %v but should have succeeded", err)
	} else if actual := val.Display(); actual != n {
		t.Fatalf("PrimitiveNum().CastStr() returned %s, expected %s", actual, n)
	}
}
