package vm

import (
	"fmt"
	"testing"
)

func TestPushIntParsing(t *testing.T) {
	expected_type := "*vm.Instruction_PUSH_INT"

	ExpectParseSame(t, "PUSH_INT 123", expected_type)
	ExpectParseSame(t, "PUSH_INT -123", expected_type)
	ExpectParseSame(t, "PUSH_INT 1", expected_type)
	ExpectParseSame(t, "PUSH_INT 0", expected_type)
	ExpectParseSame(t, "PUSH_INT 123456789", expected_type)
	ExpectParseSame(t, "PUSH_INT 314", expected_type)
	ExpectParseFailure(t, "PUSH_INT", "needs a value")
	ExpectParseFailure(t, "PUSH_INT foo", `invalid value "foo"`)
	ExpectParseFailure(t, "PUSH_INT 1.5", `invalid value "1.5"`)
}

func TestPushIntPerform(t *testing.T) {
	vm := VirtualMachine{}

	for i := range 10 {
		instruction, err := ParseInstruction(fmt.Sprintf("PUSH_INT %d", i))
		if err != nil {
			t.Fatalf("Error parsing instruction %q: %v", instruction, err)
		}
		if err := instruction.Perform(&vm); err != nil {
			t.Fatalf("Error executing instruction %q: %v", instruction, err)
		}
	}

	for i := range 10 {
		value, err := vm.stack.Pop()
		if err != nil {
			t.Fatalf("Unexpected error popping stack: %v", err)
		}
		integer, err := value.RequireInt()
		if err != nil {
			t.Fatalf("Unexpected error casting popped value to number: %v", err)
		}
		actual := integer.Display()
		expected := fmt.Sprint(9 - i)
		if actual != expected {
			t.Fatalf("Unexpected popped value, got %s expected %s", actual, expected)
		}
	}
}
