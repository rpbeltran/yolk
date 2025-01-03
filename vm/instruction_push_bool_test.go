package vm

import (
	"testing"
)

func TestPusBoolParsing(t *testing.T) {
	expected_type := "*vm.Instruction_PUSH_BOOL"

	ExpectParse(t, "PUSH_BOOL true", expected_type, "PUSH_BOOL true")
	ExpectParse(t, "PUSH_BOOL false", expected_type, "PUSH_BOOL false")
	ExpectParseFailure(t, "PUSH_BOOL", "needs a value")
	ExpectParseFailure(t, "PUSH_BOOL foo", `invalid value "foo"`)
	ExpectParseFailure(t, "PUSH_BOOL 1", `invalid value "1"`)
}

func TestPusBoolPerform(t *testing.T) {
	vm := VirtualMachine{}

	if instruction, err := ParseInstruction("PUSH_BOOL true"); err != nil {
		t.Fatalf("Error parsing instruction %q: %v", "PUSH_BOOL true", err)
	} else if err := instruction.Perform(&vm); err != nil {
		t.Fatalf("Error executing instruction %q: %v", instruction, err)
	} else if value, err := vm.stack.Pop(); err != nil {
		t.Fatalf("Error popping stack after PUSH_BOOL: %v", err)
	} else if as_bool, err := value.RequireBool(); err != nil {
		t.Fatalf("Error interpretting result of PUSH_BOOL as a bool: %v", err)
	} else if !as_bool.Truthy() {
		t.Fatalf("Expected %q to result in true, but got false", "PUSH_BOOL true")
	}

	if instruction, err := ParseInstruction("PUSH_BOOL false"); err != nil {
		t.Fatalf("Error parsing instruction %q: %v", "PUSH_BOOL false", err)
	} else if err := instruction.Perform(&vm); err != nil {
		t.Fatalf("Error executing instruction %q: %v", instruction, err)
	} else if value, err := vm.stack.Pop(); err != nil {
		t.Fatalf("Error popping stack after PUSH_BOOL: %v", err)
	} else if as_bool, err := value.RequireBool(); err != nil {
		t.Fatalf("Error interpretting result of PUSH_BOOL as a bool: %v", err)
	} else if as_bool.Truthy() {
		t.Fatalf("Expected %q to result in false, but got true", "PUSH_BOOL false")
	}
}
