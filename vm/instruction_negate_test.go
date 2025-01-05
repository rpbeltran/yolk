package vm

import (
	"testing"
	"yolk/types"
)

func TestNegateParsing(t *testing.T) {
	expected_type := "*vm.Instruction_NEGATE"

	ExpectParseSame(t, `NEGATE`, expected_type)
	ExpectParseFailure(t, "NEGATE foo", "expected no arguments")
	ExpectParseFailure(t, "NEGATE true", "expected no arguments")
	ExpectParseFailure(t, "NEGATE 1", "expected no arguments")
}

func TestNegate(t *testing.T) {
	vm := NewVM()

	instruction := RequireParse(t, "NEGATE")

	if err := instruction.Perform(&vm); err == nil {
		t.Fatalf("Expected an error performing NEGATE without arguments: %v", err)
	}

	vm.stack.Push(RequireNum(t, "123"))

	if err := instruction.Perform(&vm); err != nil {
		t.Fatalf("Unexpected error performing NEGATE: %v", err)
	} else if vm.stack.Size() != 1 {
		t.Fatalf("Expected stack size to be 1 after performing NEGATE, got %d", vm.stack.Size())
	} else if value, err := vm.stack.Pop(); err != nil {
		t.Fatalf("Unexpected error popping stack after NEGATE: %v", err)
	} else if !value.Equal(RequireNum(t, "-123")) {
		t.Fatalf("Expected stack to contain 'false' after performing NEGATE with 'true', got %q", value.Display())
	}

	vm.stack.Push(RequireNum(t, "-0.123"))

	if err := instruction.Perform(&vm); err != nil {
		t.Fatalf("Unexpected error performing NEGATE: %v", err)
	} else if vm.stack.Size() != 1 {
		t.Fatalf("Expected stack size to be 1 after performing NEGATE, got %d", vm.stack.Size())
	} else if value, err := vm.stack.Pop(); err != nil {
		t.Fatalf("Unexpected error popping stack after NEGATE: %v", err)
	} else if !value.Equal(RequireNum(t, "0.123")) {
		t.Fatalf("Expected stack to contain 'true' after performing NEGATE with 'false', got %q", value.Display())
	}

	vm.stack.Push(types.MakeString("foo"))

	if err := instruction.Perform(&vm); err == nil {
		t.Fatal(`Expected error performing NEGATE with "foo", got none`)
	}

	vm.stack.Push(types.MakeBool(true))

	if err := instruction.Perform(&vm); err == nil {
		t.Fatal(`Expected error performing NEGATE with true, got none`)
	}
}
