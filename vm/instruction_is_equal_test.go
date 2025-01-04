package vm

import (
	"testing"
	"yolk/types"
)

func TestIsEqualParsing(t *testing.T) {
	expected_type := "*vm.Instruction_IS_EQUAL"

	ExpectParse(t, "IS_EQUAL", expected_type, "IS_EQUAL")
	ExpectParseFailure(t, "IS_EQUAL 1", "expected no arguments")
	ExpectParseFailure(t, "IS_EQUAL foo", "expected no arguments")
	ExpectParseFailure(t, "IS_EQUAL true", "expected no arguments")
}

func TestIsEqualPerform(t *testing.T) {
	vm := NewVM()

	vm.stack.Push(types.MakeString("foo"))
	vm.stack.Push(types.MakeString("bar"))

	if instruction, err := ParseInstruction("IS_EQUAL"); err != nil {
		t.Fatalf("Error parsing instruction %q: %v", instruction, err)
	} else if err := instruction.Perform(&vm); err != nil {
		t.Fatalf("Unexpected error executing IS_EQUAL: %v", err)
	} else if stack_size := vm.stack.Size(); stack_size != 1 {
		t.Fatalf("Unexpected stack to have 1 element after IS_EQUAL, had: %d", stack_size)
	} else if top, err := vm.stack.Pop(); err != nil {
		t.Fatalf("Unexpected error popping stack after IS_EQUAL: %v", err)
	} else if top_bool, err := top.RequireBool(); err != nil {
		t.Fatalf("Unexpected error interpretting result of IS_EQUAL as a bool: %v", err)
	} else if top_bool.Truthy() {
		t.Fatal("Expected IS_EQUAL to push false, got true")
	}
}

func TestIsEqualFailure(t *testing.T) {
	vm := NewVM()

	if instruction, err := ParseInstruction("IS_EQUAL"); err != nil {
		t.Fatalf("Error parsing instruction %q: %v", instruction, err)
	} else if err := instruction.Perform(&vm); err == nil {
		t.Fatalf("Expected error executing IS_EQUAL, got success")
	}

	vm.stack.Push(types.MakeString("foo"))

	if instruction, err := ParseInstruction("IS_EQUAL"); err != nil {
		t.Fatalf("Error parsing instruction %q: %v", instruction, err)
	} else if err := instruction.Perform(&vm); err == nil {
		t.Fatalf("Expected error executing IS_EQUAL, got success")
	}
}
