package vm

import (
	"fmt"
	"testing"
	"yolk/types"
)

func TestLoadNameParsing(t *testing.T) {
	expected_type := "*vm.Instruction_LOAD_NAME"

	ExpectParse(t, `LOAD_NAME foo`, expected_type, `LOAD_NAME foo`)
	ExpectParse(t, `LOAD_NAME foo_bar`, expected_type, `LOAD_NAME foo_bar`)
	ExpectParseFailure(t, "LOAD_NAME", "needs a name")
}

func TestLoadNamePerform(t *testing.T) {
	message := "hello world"
	name := "foo"

	vm := NewVM()

	if err := vm.StoreNewVariable(name, types.MakeString(message)); err != nil {
		t.Fatalf("Unexpected error storing variable: %v", err)
	}

	if instruction, err := ParseInstruction(fmt.Sprintf("LOAD_NAME %s", name)); err != nil {
		t.Fatalf("Error parsing instruction %q: %v", instruction, err)
	} else if err := instruction.Perform(&vm); err != nil {
		t.Fatalf("Error executing instruction %v: %v", instruction, err)
	} else if value, err := vm.stack.Pop(); err != nil {
		t.Fatalf("Error popping stack after performing LOAD_NAME: %v", err)
	} else if display := value.Display(); display != message {
		t.Fatalf("Expected top of stack to be %q, instead got %q", message, display)
	} else {
		if err := instruction.Perform(&vm); err != nil {
			t.Fatalf("Error executing instruction %q: %v", instruction, err)
		} else if value, err := vm.stack.Pop(); err != nil {
			t.Fatalf("Error popping stack after performing LOAD_NAME: %v", err)
		} else if display := value.Display(); display != message {
			t.Fatalf("Expected top of stack to be %q, instead got %q", message, display)
		}
	}
	if instruction, err := ParseInstruction("LOAD_NAME fake"); err != nil {
		t.Fatalf("Error parsing instruction %q: %v", instruction, err)
	} else if err := instruction.Perform(&vm); err == nil {
		t.Fatalf("expecte error executing instruction %v, got none", instruction)
	}
}
