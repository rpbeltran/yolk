package vm

import (
	"fmt"
	"testing"
	"yolk/types"
)

func TestAssignNameParsing(t *testing.T) {
	expected_type := "*vm.Instruction_ASSIGN_NAME"

	ExpectParse(t, `ASSIGN_NAME foo`, expected_type, `ASSIGN_NAME foo`)
	ExpectParse(t, `ASSIGN_NAME foo_bar`, expected_type, `ASSIGN_NAME foo_bar`)
	ExpectParseFailure(t, "ASSIGN_NAME", "needs a name")
}

func TestAssignNamePerform(t *testing.T) {
	message := "Hello world!!!"
	name := "foo"

	vm := NewVM()
	vm.StoreNewVariable(name, types.MakeString(""))
	vm.stack.Push(types.MakeString(message))

	if instruction, err := ParseInstruction(fmt.Sprintf("ASSIGN_NAME %s", name)); err != nil {
		t.Fatalf("Error parsing instruction %q: %v", instruction, err)
	} else if err := instruction.Perform(&vm); err != nil {
		t.Fatalf("Error executing instruction %q: %v", instruction, err)
	} else if value, _, err := vm.FetchVariable(name); err != nil {
		t.Fatalf("Error popping stack after performing ASSIGN_NAME: %v", err)
	} else if display := value.Display(); display != message {
		t.Fatalf("Expected top of stack to be %q, instead got %q", message, display)
	} else {
		message = "Goodbye world!!"
		vm.stack.Push(types.MakeString(message))
		if err := instruction.Perform(&vm); err != nil {
			t.Fatalf("Error executing instruction %q: %v", instruction, err)
		} else if value, _, err := vm.FetchVariable(name); err != nil {
			t.Fatalf("Error popping stack after performing ASSIGN_NAME: %v", err)
		} else if display := value.Display(); display != message {
			t.Fatalf("Expected top of stack to be %q, instead got %q", message, display)
		}
	}
}
