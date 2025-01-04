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

func TestBinopAdd(t *testing.T) {
	program := []string{
		"PUSH_NUM 5",
		"PUSH_NUM 8",
		"PUSH_NUM 12",
		"BINOP add",
		"BINOP add",
	}
	expected := "25"

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
			t.Fatalf("Output is not a number: %v", err)
		} else if actual := num.Display(); actual != expected {
			t.Fatalf("Calculating (12 + 8) + 5 gave %s, expected %s", actual, expected)
		}
	}
}

func TestBinopSubtract(t *testing.T) {
	program := []string{
		"PUSH_NUM 5",
		"PUSH_NUM 8",
		"PUSH_NUM 12",
		"BINOP subtract",
		"BINOP subtract",
	}
	expected := "-1"

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
			t.Fatalf("Output is not a number: %v", err)
		} else if actual := num.Display(); actual != expected {
			t.Fatalf("Calculating (12 - 8) - 5 gave %s, expected %s", actual, expected)
		}
	}
}

func TestBinopMuiltiply(t *testing.T) {
	program := []string{
		"PUSH_NUM 5",
		"PUSH_NUM 8",
		"PUSH_NUM 12",
		"BINOP multiply",
		"BINOP multiply",
	}
	expected := "480"

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
			t.Fatalf("Output is not a number: %v", err)
		} else if actual := num.Display(); actual != expected {
			t.Fatalf("Calculating (12 * 8) * 5 gave %s, expected %s", actual, expected)
		}
	}
}

func TestBinopDivide(t *testing.T) {
	program := []string{
		"PUSH_NUM 5",
		"PUSH_NUM 8",
		"PUSH_NUM 12",
		"BINOP divide",
		"BINOP divide",
	}
	expected := "0.3"

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
			t.Fatalf("Output is not a number: %v", err)
		} else if actual := num.Display(); actual != expected {
			t.Fatalf("Calculating (12 / 8) / 5 gave %s, expected %s", actual, expected)
		}
	}
}

func TestBinopIntDivide(t *testing.T) {
	program := []string{
		"PUSH_NUM 2",
		"PUSH_NUM 10",
		"PUSH_NUM 99",
		"BINOP int_divide",
		"BINOP int_divide",
	}
	expected := "4"

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
			t.Fatalf("Output is not a number: %v", err)
		} else if actual := num.Display(); actual != expected {
			t.Fatalf("Calculating (99 // 10) // 2 gave %s, expected %s", actual, expected)
		}
	}
}
func TestBinopPower(t *testing.T) {
	program := []string{
		"PUSH_NUM 2",
		"PUSH_NUM 3",
		"PUSH_NUM 4",
		"BINOP power",
		"BINOP power",
	}
	expected := "4096"

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
			t.Fatalf("Output is not a number: %v", err)
		} else if actual := num.Display(); actual != expected {
			t.Fatalf("Calculating (4 ** 3) ** 2 gave %s, expected %s", actual, expected)
		}
	}
}

func TestBinopMod(t *testing.T) {
	program := []string{
		"PUSH_NUM 12",
		"PUSH_NUM 30",
		"PUSH_NUM 44",
		"BINOP modulus",
		"BINOP modulus",
	}
	expected := "2"

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
			t.Fatalf("Output is not a number: %v", err)
		} else if actual := num.Display(); actual != expected {
			t.Fatalf("Calculating (44 mod 30) mod 12 gave %s, expected %s", actual, expected)
		}
	}
}

func TestBinopConcat(t *testing.T) {
	program := []string{
		`PUSH_STR "bar"`,
		"PUSH_NUM 0",
		`PUSH_STR "fo"`,
		"BINOP concat",
		"BINOP concat",
	}
	expected := "fo0bar"

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
		if str, err := value.RequireStr(); err != nil {
			t.Fatalf("Output is not a string: %v", err)
		} else if actual := str.Display(); actual != expected {
			t.Fatalf("Calculating (44 mod 30) mod 12 gave %s, expected %s", actual, expected)
		}
	}
}

func TestBinopAnd(t *testing.T) {
	program := []string{
		`PUSH_BOOL true`,
		"PUSH_BOOL true",
		`PUSH_BOOL false`,
		"BINOP and",
		"BINOP and",
	}
	expected := "false"

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
		if as_bool, err := value.RequireBool(); err != nil {
			t.Fatalf("Output is not a string: %v", err)
		} else if actual := as_bool.Display(); actual != expected {
			t.Fatalf("Calculating (true && true) && false gave %s, expected %s", actual, expected)
		}
	}
}

func TestBinopOr(t *testing.T) {
	program := []string{
		`PUSH_BOOL true`,
		"PUSH_BOOL true",
		`PUSH_BOOL false`,
		"BINOP or",
		"BINOP or",
	}
	expected := "true"

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
		if as_bool, err := value.RequireBool(); err != nil {
			t.Fatalf("Output is not a string: %v", err)
		} else if actual := as_bool.Display(); actual != expected {
			t.Fatalf("Calculating (true || true) || false gave %s, expected %s", actual, expected)
		}
	}
}
