package vm

import (
	"fmt"
	"testing"
)

func TestPushIntParsing(t *testing.T) {
	expected_type := "*vm.Instruction_PUSH_NUM"

	ExpectParse(t, "PUSH_NUM 123", expected_type, "PUSH_NUM 123")
	ExpectParse(t, "PUSH_NUM -123", expected_type, "PUSH_NUM -123")
	ExpectParse(t, "PUSH_NUM   1", expected_type, "PUSH_NUM 1")
	ExpectParse(t, "PUSH_NUM   0  ", expected_type, "PUSH_NUM 0")
	ExpectParse(t, "  PUSH_NUM   123456789  ", expected_type, "PUSH_NUM 123456789")
	ExpectParse(t, "PUSH_NUM   314  ", expected_type, "PUSH_NUM 314")
	ExpectParseFailure(t, "PUSH_NUM", "needs a value")
	ExpectParseFailure(t, "PUSH_NUM foo", `invalid value "foo"`)
}

func TestPushIntPerform(t *testing.T) {
	vm := VirtualMachine{}

	for i := range 10 {
		instruction, err := ParseInstruction(fmt.Sprintf("PUSH_NUM %d", i))
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
		num, err := value.RequireNum()
		if err != nil {
			t.Fatalf("Unexpected error casting popped value to number: %v", err)
		}
		actual := num.Display()
		expected := fmt.Sprint(9 - i)
		if actual != expected {
			t.Fatalf("Unexpected popped value, got %s expected %s", actual, expected)
		}
	}

}
