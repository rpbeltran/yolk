package vm

import (
	"testing"
	"yolk/types"
)

func TestPrintParsing(t *testing.T) {
	expected_type := "*vm.Instruction_PRINT"
	expected_string := "PRINT"

	ExpectParse(t, "PRINT", expected_type, expected_string)
	ExpectParse(t, "PRINT  ", expected_type, expected_string)
	ExpectParse(t, "  PRINT  ", expected_type, expected_string)
	ExpectParseFailure(t, "PRINT 2", "expected no arguments")
}

func TestPrintPerform(t *testing.T) {
	vm := VirtualMachine{}

	phrases := []string{"Hello", "", "foo!!", "12345", "''", `""`}

	for _, phrase := range phrases {
		vm.stack.Push(types.MakeString(phrase))
	}

	print_instruction, err := ParseInstruction("PRINT")
	if err != nil {
		t.Fatalf("Error parsing instruction %q: %v", print_instruction, err)
	}

	for i := range phrases {
		if err := print_instruction.Perform(&vm); err != nil {
			t.Fatalf("Unexpected error executing PRINT: %v", err)
		}

		if actual := vm.output_buffer.Size(); actual != i+1 {
			t.Fatalf("Queue had size %d, expected %d", actual, i+1)
		}
	}

	for i := range phrases {
		expected_phrase := phrases[len(phrases)-1-i]
		if actual, err := vm.output_buffer.Pop(); err != nil {
			t.Fatalf("Unexpected error popping output buffer: %v", err)
		} else if actual != expected_phrase {
			t.Fatalf("output_buffer[%d] had message %q, expected %q", i, actual, expected_phrase)
		}
	}

}
