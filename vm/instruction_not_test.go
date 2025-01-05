package vm

import (
	"testing"
	"yolk/types"
)

func TestNotParsing(t *testing.T) {
	expected_type := "*vm.Instruction_NOT"

	ExpectParseSame(t, `NOT`, expected_type)
	ExpectParseFailure(t, "NOT foo", "expected no arguments")
	ExpectParseFailure(t, "NOT true", "expected no arguments")
	ExpectParseFailure(t, "NOT 1", "expected no arguments")
}

func TestNot(t *testing.T) {
	vm := NewVM()

	instruction := RequireParse(t, "NOT")

	if err := instruction.Perform(&vm); err == nil {
		t.Fatalf("Expected an error performing NOT without arguments: %v", err)
	}

	vm.stack.Push(types.MakeBool(true))

	if err := instruction.Perform(&vm); err != nil {
		t.Fatalf("Unexpected error performing NOT: %v", err)
	} else if vm.stack.Size() != 1 {
		t.Fatalf("Expected stack size to be 1 after performing NOT, got %d", vm.stack.Size())
	} else if value, err := vm.stack.Pop(); err != nil {
		t.Fatalf("Unexpected error popping stack after NOT: %v", err)
	} else if !value.Equal(types.MakeBool(false)) {
		t.Fatalf("Expected stack to contain 'false' after performing NOT with 'true', got %q", value.Display())
	}

	vm.stack.Push(types.MakeBool(false))

	if err := instruction.Perform(&vm); err != nil {
		t.Fatalf("Unexpected error performing NOT: %v", err)
	} else if vm.stack.Size() != 1 {
		t.Fatalf("Expected stack size to be 1 after performing NOT, got %d", vm.stack.Size())
	} else if value, err := vm.stack.Pop(); err != nil {
		t.Fatalf("Unexpected error popping stack after NOT: %v", err)
	} else if !value.Equal(types.MakeBool(true)) {
		t.Fatalf("Expected stack to contain 'true' after performing NOT with 'false', got %q", value.Display())
	}

	vm.stack.Push(types.MakeString("foo"))

	if err := instruction.Perform(&vm); err == nil {
		t.Fatal(`Expected error performing NOT with "foo", got none`)
	}

	vm.stack.Push(RequireNum(t, "123"))

	if err := instruction.Perform(&vm); err == nil {
		t.Fatal(`Expected error performing NOT with 123, got none`)
	}
}
