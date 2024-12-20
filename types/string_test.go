package types

import (
	"testing"
)

func TestStrDisplay(t *testing.T) {
	test_cases := []string{"hello", "", "123", " @#()*", `""`}
	for _, tc := range test_cases {
		if actual := MakeString(tc).Display(); actual != tc {
			t.Fatalf("makePrimitiveNumber(%q).Display() returned %q, expected %q", tc, actual, tc)
		}
	}
}

// Add

func TestStrAdd(t *testing.T) {
	if _, err := MakeString("foo").Add(makeNumOrFail("100", t)); err == nil {
		t.Fatal(`"foo" + 100 succeeded but should have failed`)
	}
	if _, err := MakeString("foo").Add(MakeString("bar")); err == nil {
		t.Fatal(`"foo" + "bar" succeeded but should have failed`)
	}
}

// Casting

func TestStrRequireNum(t *testing.T) {
	s := "100"
	if _, err := MakeString(s).RequireNum(); err == nil {
		t.Fatalf("PrimitiveStr().RequireNum() succeeded but should have failed")
	}
}

func TestStrRequireStr(t *testing.T) {
	s := "foo"
	if val, err := MakeString(s).RequireStr(); err != nil {
		t.Fatalf("PrimitiveStr().RequireStr() returned the error %v but should have succeeded", err)
	} else if actual := val.Display(); actual != s {
		t.Fatalf("PrimitiveStr().RequireStr() returned %s, expected %s", actual, s)
	}
}

func TestStrCastNum(t *testing.T) {
	s := "100"
	if value, err := MakeString(s).CastNum(); err != nil {
		t.Fatalf("PrimitiveStr(%q).CastNum() returned the error %v but should have succeeded", s, err)
	} else if actual := value.Display(); actual != "100" {
		t.Fatalf("PrimitiveStr(%q).CastNum() returned %s, expected %s", s, actual, s)
	}
}

func TestStrCastNumFailure(t *testing.T) {
	if _, err := MakeString("foo").CastNum(); err == nil {
		t.Fatalf(`PrimitiveStr("foo").RequireStr() succeeded but should have failed`)
	}
}

func TestStrCastStr(t *testing.T) {
	s := "foo"
	if val, err := MakeString(s).CastStr(); err != nil {
		t.Fatalf("PrimitiveStr().CastStr() returned the error %v but should have succeeded", err)
	} else if actual := val.Display(); actual != s {
		t.Fatalf("PrimitiveStr().CastStr() returned %s, expected %s", actual, s)
	}
}
