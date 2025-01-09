package vm

import (
	"errors"
	"testing"
)

func TestDefineEndParsing(t *testing.T) {
	expected_type := "*vm.Instruction_DEFINE_END"

	ExpectParseSame(t, `.DEFINE_END`, expected_type)
	ExpectParseWrappedFailure(t, ".DEFINE_END foo", ErrDefineEndParsing)
	ExpectParseWrappedFailure(t, ".DEFINE_END <foo>", ErrDefineEndParsing)
	ExpectParseWrappedFailure(t, `.DEFINE_END "foo"`, ErrDefineEndParsing)
	ExpectParseWrappedFailure(t, `.DEFINE_END 1`, ErrDefineEndParsing)
	ExpectParseWrappedFailure(t, `.DEFINE_END "<foo> bar"`, ErrDefineEndParsing)
	ExpectParseWrappedFailure(t, `.DEFINE_END "<foo> "bar""`, ErrDefineEndParsing)
	ExpectParseWrappedFailure(t, `.DEFINE_END "<foo> 1"`, ErrDefineEndParsing)
}

func TestDefineEndPerform(t *testing.T) {
	vm := NewVM()
	if instruction, err := ParseInstruction(".DEFINE_END"); err != nil {
		t.Fatalf("Error parsing instruction %q: %v", instruction, err)
	} else if err := instruction.Perform(&vm); err == nil {
		t.Fatalf("Instruction %q was supposed to fail but instead succeeded", instruction)
	} else if !errors.Is(err, ErrDefineEndPerform) {
		t.Fatalf("Instruction %q was supposed to fail with %v instead but instead gave %v", instruction, ErrDefineEndPerform, err)
	}
}
