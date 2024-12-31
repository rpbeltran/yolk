package vm

import (
	"testing"
	"yolk/types"
)

func TestJumpIfTrueParsing(t *testing.T) {
	expected_type := "*vm.Instruction_JUMP_IF_TRUE"

	ExpectParse(t, `JUMP_IF_TRUE 0`, expected_type, `JUMP_IF_TRUE 0`)
	ExpectParse(t, `JUMP_IF_TRUE 123`, expected_type, `JUMP_IF_TRUE 123`)
	ExpectParseFailure(t, "JUMP_IF_TRUE", "needs a destination")
	ExpectParseFailure(t, "JUMP_IF_TRUE hello", `invalid destination "hello"`)
	ExpectParseFailure(t, "JUMP_IF_TRUE -10", `invalid destination "-10"`)
}

func TestJumpIfTruePerform(t *testing.T) {
	vm := NewVM()
	vm.labels[100] = 50

	vm.instruction_pointer = 0
	vm.stack.Push(types.MakeBool(true))

	if err := RequireParse(t, "JUMP_IF_TRUE 100").Perform(&vm); err != nil {
		t.Fatalf("Unexpected error from %q: %v", "JUMP_IF_TRUE 100", err)
	} else if vm.instruction_pointer != 50 {
		t.Fatalf("Expected ipointer=50, got %d", vm.instruction_pointer)
	}

	vm.instruction_pointer = 0
	vm.stack.Push(types.MakeBool(false))

	if err := RequireParse(t, "JUMP_IF_TRUE 100").Perform(&vm); err != nil {
		t.Fatalf("Unexpected error from %q: %v", "JUMP_IF_TRUE 100", err)
	} else if vm.instruction_pointer != 0 {
		t.Fatalf("Expected ipointer=0, got %d", vm.instruction_pointer)
	}

	vm.instruction_pointer = 0
	vm.stack.Push(types.MakeString("hello"))

	if err := RequireParse(t, "JUMP_IF_TRUE 100").Perform(&vm); err == nil {
		t.Fatalf("Expected error from %q, got none", "JUMP_IF_TRUE 0")
	}

	vm.instruction_pointer = 0
	vm.stack.Push(types.MakeBool(true))

	if err := RequireParse(t, "JUMP_IF_TRUE 0").Perform(&vm); err == nil {
		t.Fatalf("Expected error from %q, got none", "JUMP_IF_TRUE 0")
	}
}
