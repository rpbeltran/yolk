package vm

import (
	"testing"
)

func TestJumpParsing(t *testing.T) {
	expected_type := "*vm.Instruction_JUMP"

	ExpectParse(t, `JUMP 0`, expected_type, `JUMP 0`)
	ExpectParse(t, `JUMP 123`, expected_type, `JUMP 123`)
	ExpectParseFailure(t, "JUMP", "needs a destination")
	ExpectParseFailure(t, "JUMP hello", `invalid destination "hello"`)
	ExpectParseFailure(t, "JUMP -10", `invalid destination "-10"`)
}

func TestJumpPerform(t *testing.T) {
	vm := NewVM()
	vm.labels[100] = 50
	vm.instruction_pointer = 0
	if err := RequireParse(t, "JUMP 100").Perform(&vm); err != nil {
		t.Fatalf("Unexpected error from %q: %v", "JUMP 100", err)
	} else if vm.instruction_pointer != 50 {
		t.Fatalf("Expected ipointer=50, got %d", vm.instruction_pointer)
	}

	if err := RequireParse(t, "JUMP 0").Perform(&vm); err == nil {
		t.Fatalf("Expected error from %q, got none", "JUMP 0")
	}
}
