package vm

import (
	"fmt"
	"testing"
)

func TestPushStrParsing(t *testing.T) {
	expected_type := "*vm.Instruction_PUSH_STR"

	ExpectParse(t, `PUSH_STR "FOO"`, expected_type, `PUSH_STR "FOO"`)
	ExpectParse(t, `PUSH_STR "123"`, expected_type, `PUSH_STR "123"`)
	ExpectParse(t, `PUSH_STR ""`, expected_type, `PUSH_STR ""`)
	ExpectParseFailure(t, "PUSH_STR", "needs a value")
	ExpectParseFailure(t, "PUSH_STR hello", `invalid value "hello"`)
	ExpectParseFailure(t, "PUSH_STR 0", `invalid value "0"`)
}

func TestPushStrPerform(t *testing.T) {
	vm := VirtualMachine{}

	phrases := []string{
		"Hello world",
		"foo",
		"",
		"123",
	}

	for _, phrase := range phrases {
		instruction, err := ParseInstruction(fmt.Sprintf("PUSH_STR %q", phrase))
		if err != nil {
			t.Fatalf("Error parsing instruction %q: %v", instruction, err)
		}
		if err := instruction.Perform(&vm); err != nil {
			t.Fatalf("Error executing instruction %q: %v", instruction, err)
		}
	}

	for i := range phrases {
		value, err := vm.stack.Pop()
		if err != nil {
			t.Fatalf("Unexpected error popping stack: %v", err)
		}
		str, err := value.RequireStr()
		if err != nil {
			t.Fatalf("Unexpected error casting popped value to primitive string: %v", err)
		}
		expected := phrases[len(phrases)-1-i]
		if actual := str.Display(); actual != expected {
			t.Fatalf("Unexpected popped value, got %q expected %q", actual, expected)
		}
	}
}
