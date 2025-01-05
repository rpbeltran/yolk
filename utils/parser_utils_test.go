package utils

import (
	"errors"
	"testing"
)

func TestStripAngleQuotes(t *testing.T) {
	s := "<foo>"
	if stripped, err := StripAngleQuotes(s); err != nil {
		t.Fatalf("Unexpected error from StripAngleQuotes(%q): %v", s, err)
	} else if stripped != "foo" {
		t.Fatalf("Expected StripAngleQuotes(%q) == %q, got %q", s, "foo", stripped)
	}

	s = "<hello <> world>>"
	if stripped, err := StripAngleQuotes(s); err != nil {
		t.Fatalf("Unexpected error from StripAngleQuotes(%q): %v", s, err)
	} else if stripped != "hello <> world>" {
		t.Fatalf("Expected StripAngleQuotes(%q) == %q, got %q", s, "hello <> world>", stripped)
	}

	s = "<>"
	if stripped, err := StripAngleQuotes(s); err != nil {
		t.Fatalf("Unexpected error from StripAngleQuotes(%q): %v", s, err)
	} else if stripped != "" {
		t.Fatalf("Expected StripAngleQuotes(%q) == %q, got %q", s, "", stripped)
	}

	s = ""
	if result, err := StripAngleQuotes(s); err == nil {
		t.Fatalf("StripAngleQuotes(%q) was expected to fail with %v, instead succeeded and gave %q", s, ErrStripAngleQuotedGotEmptyString, result)
	} else if !errors.Is(err, ErrStripAngleQuotedGotEmptyString) {
		t.Fatalf("StripAngleQuotes(%q) was expected to fail with %v, instead failed with %v", s, ErrStripAngleQuotedGotEmptyString, err)
	}

	s = "foo"
	if result, err := StripAngleQuotes(s); err == nil {
		t.Fatalf("StripAngleQuotes(%q) was expected to fail with %v, instead succeeded and gave %q", s, ErrStripAngleQuotedGotInvalidString, result)
	} else if !errors.Is(err, ErrStripAngleQuotedGotInvalidString) {
		t.Fatalf("StripAngleQuotes(%q) was expected to fail with %v, instead failed with %v", s, ErrStripAngleQuotedGotInvalidString, err)
	}

	s = "foo>"
	if result, err := StripAngleQuotes(s); err == nil {
		t.Fatalf("StripAngleQuotes(%q) was expected to fail with %v, instead succeeded and gave %q", s, ErrStripAngleQuotedGotInvalidString, result)
	} else if !errors.Is(err, ErrStripAngleQuotedGotInvalidString) {
		t.Fatalf("StripAngleQuotes(%q) was expected to fail with %v, instead failed with %v", s, ErrStripAngleQuotedGotInvalidString, err)
	}

	s = "<foo"
	if result, err := StripAngleQuotes(s); err == nil {
		t.Fatalf("StripAngleQuotes(%q) was expected to fail with %v, instead succeeded and gave %q", s, ErrStripAngleQuotedGotInvalidString, result)
	} else if !errors.Is(err, ErrStripAngleQuotedGotInvalidString) {
		t.Fatalf("StripAngleQuotes(%q) was expected to fail with %v, instead failed with %v", s, ErrStripAngleQuotedGotInvalidString, err)
	}

	s = ">foo>"
	if result, err := StripAngleQuotes(s); err == nil {
		t.Fatalf("StripAngleQuotes(%q) was expected to fail with %v, instead succeeded and gave %q", s, ErrStripAngleQuotedGotInvalidString, result)
	} else if !errors.Is(err, ErrStripAngleQuotedGotInvalidString) {
		t.Fatalf("StripAngleQuotes(%q) was expected to fail with %v, instead failed with %v", s, ErrStripAngleQuotedGotInvalidString, err)
	}

	s = "<foo<"
	if result, err := StripAngleQuotes(s); err == nil {
		t.Fatalf("StripAngleQuotes(%q) was expected to fail with %v, instead succeeded and gave %q", s, ErrStripAngleQuotedGotInvalidString, result)
	} else if !errors.Is(err, ErrStripAngleQuotedGotInvalidString) {
		t.Fatalf("StripAngleQuotes(%q) was expected to fail with %v, instead failed with %v", s, ErrStripAngleQuotedGotInvalidString, err)
	}
}
