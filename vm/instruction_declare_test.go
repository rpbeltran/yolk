package vm

import (
	"fmt"
	"testing"
	"yolk/types"
)

func TestDeclareNameParsing(t *testing.T) {
	expected_type := "*vm.Instruction_DECLARE"

	ExpectParse(t, `DECLARE <foo>`, expected_type, `DECLARE <foo>`)
	ExpectParse(t, `DECLARE <foo_bar>`, expected_type, `DECLARE <foo_bar>`)
	ExpectParseWrappedFailure(t, "DECLARE", ErrParsingDECLARE)
	ExpectParseWrappedFailure(t, "DECLARE foo", ErrParsingDECLARE)
	ExpectParseWrappedFailure(t, `DECLARE "foo"`, ErrParsingDECLARE)
	ExpectParseWrappedFailure(t, `DECLARE "<foo> hello"`, ErrParsingDECLARE)
}

func TestDeclareNamePerform(t *testing.T) {
	message := "Hello world!!!"
	name := "foo"

	vm := NewVM()
	vm.stack.Push(types.MakeString(message))

	if instruction, err := ParseInstruction(fmt.Sprintf("DECLARE <%s>", name)); err != nil {
		t.Fatalf("Error parsing instruction %q: %v", instruction, err)
	} else if err := instruction.Perform(&vm); err != nil {
		t.Fatalf("Error executing instruction %q: %v", instruction, err)
	} else if value, err := vm.FetchVariable(name); err != nil {
		t.Fatalf("Error popping stack after performing DECLARE: %v", err)
	} else if display := value.Display(); display != message {
		t.Fatalf("Expected top of stack to be %q, instead got %q", message, display)
	} else {
		vm.globals = make(map[string]types.Primitive)
		message = "Goodbye world!!"
		vm.stack.Push(types.MakeString(message))
		if err := instruction.Perform(&vm); err != nil {
			t.Fatalf("Error executing instruction %q: %v", instruction, err)
		} else if value, err := vm.FetchVariable(name); err != nil {
			t.Fatalf("Error popping stack after performing DECLARE: %v", err)
		} else if display := value.Display(); display != message {
			t.Fatalf("Expected top of stack to be %q, instead got %q", message, display)
		}
	}
}
