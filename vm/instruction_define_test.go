package vm

import (
	"testing"
)

func TestDefineParsing(t *testing.T) {
	expected_type := "*vm.Instruction_DEFINE"

	ExpectParseSame(t, `.DEFINE <foo>`, expected_type)
	ExpectParseSame(t, `.DEFINE <foo> <bar>`, expected_type)
	ExpectParseWrappedFailure(t, ".DEFINE", ErrDefineParsingName)
	ExpectParseWrappedFailure(t, ".DEFINE foo", ErrDefineParsingName)
	ExpectParseWrappedFailure(t, `.DEFINE "foo"`, ErrDefineParsingName)
	ExpectParseWrappedFailure(t, `.DEFINE 1`, ErrDefineParsingName)
	ExpectParseWrappedFailure(t, `.DEFINE "<foo> bar"`, ErrDefineParsingType)
	ExpectParseWrappedFailure(t, `.DEFINE "<foo> "bar""`, ErrDefineParsingType)
	ExpectParseWrappedFailure(t, `.DEFINE "<foo> 1"`, ErrDefineParsingType)
}

func TestDefinePerform(t *testing.T) {
	vm := NewVM()
	if instruction, err := ParseInstruction(".DEFINE <foo>"); err != nil {
		t.Fatalf("Error parsing instruction %q: %v", instruction, err)
	} else if err := instruction.Perform(&vm); err == nil {
		t.Fatalf("Error executing instruction %q: %v", instruction, err)
	}
}
