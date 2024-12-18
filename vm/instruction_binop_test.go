package vm

import (
	"testing"
)

func TestBinopParsing(t *testing.T) {
	expected_type := "*vm.Instruction_BINOP"

	ExpectParse(t, "BINOP add", expected_type, "BINOP add")
	ExpectParse(t, "BINOP   add", expected_type, "BINOP add")
	ExpectParse(t, "BINOP   add  ", expected_type, "BINOP add")
	ExpectParse(t, "  BINOP   add  ", expected_type, "BINOP add")
	ExpectParse(t, "\tBINOP   add  ", expected_type, "BINOP add")
	ExpectParseFailure(t, "BINOP", "needs operator")
	ExpectParseFailure(t, "BINOP foo", `unexpected operator "foo"`)
}

func TestBinopPerform(t *testing.T) {
	program := []string{
		"PUSH_NUM 5",
		"PUSH_NUM 8",
		"PUSH_NUM 12",
		"BINOP add",
		"BINOP add",
	}

	vm := VirtualMachine{}

	for _, line := range program {
		line_instruction, err := ParseInstruction(line)
		if err != nil {
			t.Fatalf("Error parsing instruction %q: %v", line_instruction, err)
		}
		if err := line_instruction.Perform(&vm); err != nil {
			t.Fatalf("Unexpected error executing %q: %v", line, err)
		}
	}
	if actual := vm.stack.Size(); actual != 1 {
		t.Fatalf("Stack had %d items after operations, expected 1", actual)
	}
	if value, err := vm.stack.Pop(); err != nil {
		t.Fatalf("Unexpected error popping stack: %v", err)
	} else {
		if num, err := value.RequireNum(); err != nil {
			t.Fatalf("Output of addition is not a number: %v", err)
		} else if actual := num.Display(); actual != "25" {
			t.Fatalf("Calculating 8 + 12 + 5 gave %s, expecte 25 ", actual)
		}
	}
}
