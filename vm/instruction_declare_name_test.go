package vm

import (
	"fmt"
	"testing"
	"yolk/types"
)

func TestDeclareNameParsing(t *testing.T) {
	expected_type := "*vm.Instruction_DECLARE_NAME"

	ExpectParse(t, `DECLARE_NAME foo`, expected_type, `DECLARE_NAME foo`)
	ExpectParse(t, `DECLARE_NAME foo_bar`, expected_type, `DECLARE_NAME foo_bar`)
	ExpectParseFailure(t, "DECLARE_NAME", "needs a name")
}

func TestDeclareNamePerform(t *testing.T) {
	message := "Hello world!!!"
	name := "foo"

	vm := NewVM()
	vm.stack.Push(types.MakeString(message))

	if instruction, err := ParseInstruction(fmt.Sprintf("DECLARE_NAME %s", name)); err != nil {
		t.Fatalf("Error parsing instruction %q: %v", instruction, err)
	} else if err := instruction.Perform(&vm); err != nil {
		t.Fatalf("Error executing instruction %q: %v", instruction, err)
	} else if value, err := vm.FetchVariable(name); err != nil {
		t.Fatalf("Error popping stack after performing DECLARE_NAME: %v", err)
	} else if display := value.Display(); display != message {
		t.Fatalf("Expected top of stack to be %q, instead got %q", message, display)
	} else {
		vm.globals = make(map[string]types.Primitive)
		message = "Goodbye world!!"
		vm.stack.Push(types.MakeString(message))
		if err := instruction.Perform(&vm); err != nil {
			t.Fatalf("Error executing instruction %q: %v", instruction, err)
		} else if value, err := vm.FetchVariable(name); err != nil {
			t.Fatalf("Error popping stack after performing DECLARE_NAME: %v", err)
		} else if display := value.Display(); display != message {
			t.Fatalf("Expected top of stack to be %q, instead got %q", message, display)
		}
	}
}
