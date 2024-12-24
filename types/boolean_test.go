package types

import (
	"testing"
)

func TestBoolDisplay(t *testing.T) {
	if actual := MakeBool(true).Display(); actual != "true" {
		t.Fatalf("MakeBool(true).Display() returned %q, expected %q", actual, "true")
	}
	if actual := MakeBool(false).Display(); actual != "false" {
		t.Fatalf("MakeBool(false).Display() returned %q, expected %q", actual, "false")
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

func TestBoolCastNum(t *testing.T) {
	if value, err := MakeBool(true).CastNum(); err != nil {
		t.Fatalf("MakeBool(true).CastNum() returned the error %v but should have succeeded", err)
	} else if actual := value.Display(); actual != "1" {
		t.Fatalf("MakeBool(true).CastNum() returned %s, expected 1", actual)
	}

	if value, err := MakeBool(false).CastNum(); err != nil {
		t.Fatalf("MakeBool(false).CastNum() returned the error %v but should have succeeded", err)
	} else if actual := value.Display(); actual != "0" {
		t.Fatalf("MakeBool(false).CastNum() returned %s, expected 0", actual)
	}
}

func TestBoolCastStr(t *testing.T) {
	if value, err := MakeBool(true).CastStr(); err != nil {
		t.Fatalf("MakeBool(true).CastStr() returned the error %v but should have succeeded", err)
	} else if actual := value.Display(); actual != "true" {
		t.Fatalf("MakeBool(true).CastNum() returned %s, expected %q", actual, "true")
	}

	if value, err := MakeBool(false).CastStr(); err != nil {
		t.Fatalf("MakeBool(false).CastStr() returned the error %v but should have succeeded", err)
	} else if actual := value.Display(); actual != "false" {
		t.Fatalf("MakeBool(false).CastNum() returned %s, expected %q", actual, "false")
	}
}

func TestBoolCastBool(t *testing.T) {
	if actual, err := MakeBool(false).CastBool(); err != nil {
		t.Fatalf("MakeBool(false).CastBool() failed with error: %v", err)
	} else if actual.value {
		t.Fatalf("MakeBool(false).CastBool() gave true, expected false")
	}
}
