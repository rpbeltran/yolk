package vm

import (
	"fmt"
	"testing"
	"yolk/types"
)

func TestAssignNameParsing(t *testing.T) {
	expected_type := "*vm.Instruction_ASSIGN"

	ExpectParse(t, `ASSIGN <foo>`, expected_type, `ASSIGN <foo>`)
	ExpectParse(t, `ASSIGN <foo_bar>`, expected_type, `ASSIGN <foo_bar>`)
	ExpectParseWrappedFailure(t, "ASSIGN", ErrParsingASSIGN)
	ExpectParseWrappedFailure(t, "ASSIGN foo", ErrParsingASSIGN)
	ExpectParseWrappedFailure(t, `ASSIGN "foo"`, ErrParsingASSIGN)
	ExpectParseWrappedFailure(t, `ASSIGN "<foo> hello"`, ErrParsingASSIGN)
}

func TestAssignNamePerform(t *testing.T) {
	message := "Hello world!!!"
	name := "foo"

	vm := NewVM()

	if err := vm.StoreNewVariable(name, types.MakeString("")); err != nil {
		t.Fatalf("Unexpected error storing variable: %v", err)
	}
	vm.stack.Push(types.MakeString(message))

	if instruction, err := ParseInstruction(fmt.Sprintf("ASSIGN <%s>", name)); err != nil {
		t.Fatalf("Error parsing instruction %q: %v", instruction, err)
	} else if err := instruction.Perform(&vm); err != nil {
		t.Fatalf("Error executing instruction %q: %v", instruction, err)
	} else if value, err := vm.FetchVariable(name); err != nil {
		t.Fatalf("Error popping stack after performing ASSIGN: %v", err)
	} else if display := value.Display(); display != message {
		t.Fatalf("Expected top of stack to be %q, instead got %q", message, display)
	} else {
		message = "Goodbye world!!"
		vm.stack.Push(types.MakeString(message))
		if err := instruction.Perform(&vm); err != nil {
			t.Fatalf("Error executing instruction %q: %v", instruction, err)
		} else if value, err := vm.FetchVariable(name); err != nil {
			t.Fatalf("Error popping stack after performing ASSIGN: %v", err)
		} else if display := value.Display(); display != message {
			t.Fatalf("Expected top of stack to be %q, instead got %q", message, display)
		}
	}
}
