package utils

import (
	"errors"
	"testing"
)

func TestDeserializeName(t *testing.T) {
	s := "<foo>"
	if stripped, err := DeserializeName(s); err != nil {
		t.Fatalf("Unexpected error from DeserializeName(%q): %v", s, err)
	} else if stripped != "foo" {
		t.Fatalf("Expected DeserializeName(%q) == %q, got %q", s, "foo", stripped)
	}

	s = "<hello <> world>>"
	if stripped, err := DeserializeName(s); err != nil {
		t.Fatalf("Unexpected error from DeserializeName(%q): %v", s, err)
	} else if stripped != "hello <> world>" {
		t.Fatalf("Expected DeserializeName(%q) == %q, got %q", s, "hello <> world>", stripped)
	}

	s = "<>"
	if result, err := DeserializeName(s); err == nil {
		t.Fatalf("DeserializeName(%q) was expected to fail with %v, instead succeeded and gave %q", s, ErrDeserializeNameGotEmptyName, result)
	} else if !errors.Is(err, ErrDeserializeNameGotEmptyName) {
		t.Fatalf("DeserializeName(%q) was expected to fail with %v, instead failed with %v", s, ErrDeserializeNameGotEmptyName, err)
	}

	s = ""
	if result, err := DeserializeName(s); err == nil {
		t.Fatalf("DeserializeName(%q) was expected to fail with %v, instead succeeded and gave %q", s, ErrDeserializeNameGotEmptyString, result)
	} else if !errors.Is(err, ErrDeserializeNameGotEmptyString) {
		t.Fatalf("DeserializeName(%q) was expected to fail with %v, instead failed with %v", s, ErrDeserializeNameGotEmptyString, err)
	}

	s = "foo"
	if result, err := DeserializeName(s); err == nil {
		t.Fatalf("DeserializeName(%q) was expected to fail with %v, instead succeeded and gave %q", s, ErrDeserializeNameGotInvalidString, result)
	} else if !errors.Is(err, ErrDeserializeNameGotInvalidString) {
		t.Fatalf("DeserializeName(%q) was expected to fail with %v, instead failed with %v", s, ErrDeserializeNameGotInvalidString, err)
	}

	s = "foo>"
	if result, err := DeserializeName(s); err == nil {
		t.Fatalf("DeserializeName(%q) was expected to fail with %v, instead succeeded and gave %q", s, ErrDeserializeNameGotInvalidString, result)
	} else if !errors.Is(err, ErrDeserializeNameGotInvalidString) {
		t.Fatalf("DeserializeName(%q) was expected to fail with %v, instead failed with %v", s, ErrDeserializeNameGotInvalidString, err)
	}

	s = "<foo"
	if result, err := DeserializeName(s); err == nil {
		t.Fatalf("DeserializeName(%q) was expected to fail with %v, instead succeeded and gave %q", s, ErrDeserializeNameGotInvalidString, result)
	} else if !errors.Is(err, ErrDeserializeNameGotInvalidString) {
		t.Fatalf("DeserializeName(%q) was expected to fail with %v, instead failed with %v", s, ErrDeserializeNameGotInvalidString, err)
	}

	s = ">foo>"
	if result, err := DeserializeName(s); err == nil {
		t.Fatalf("DeserializeName(%q) was expected to fail with %v, instead succeeded and gave %q", s, ErrDeserializeNameGotInvalidString, result)
	} else if !errors.Is(err, ErrDeserializeNameGotInvalidString) {
		t.Fatalf("DeserializeName(%q) was expected to fail with %v, instead failed with %v", s, ErrDeserializeNameGotInvalidString, err)
	}

	s = "<foo<"
	if result, err := DeserializeName(s); err == nil {
		t.Fatalf("DeserializeName(%q) was expected to fail with %v, instead succeeded and gave %q", s, ErrDeserializeNameGotInvalidString, result)
	} else if !errors.Is(err, ErrDeserializeNameGotInvalidString) {
		t.Fatalf("DeserializeName(%q) was expected to fail with %v, instead failed with %v", s, ErrDeserializeNameGotInvalidString, err)
	}
}

func TestSerializeName(t *testing.T) {
	s := "foo"
	expected := "<foo>"
	if quoted := SerializeName(s); quoted != expected {
		t.Fatalf("Expected SerializeName(%q) to return %q, instead got %q", s, expected, quoted)
	}
}
