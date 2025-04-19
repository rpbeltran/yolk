package vm

import (
	"fmt"
	"testing"
	"yolk/types"
)

func TestDeclareNameParsing(t *testing.T) {
	expected_type := "*vm.Instruction_DECLARE"

	ExpectParseSame(t, `DECLARE <foo>`, expected_type)
	ExpectParseSame(t, `DECLARE <foo> <bar>`, expected_type)
	ExpectParseWrappedFailure(t, "DECLARE", ErrDeclareParsingName)
	ExpectParseWrappedFailure(t, "DECLARE foo", ErrDeclareParsingName)
	ExpectParseWrappedFailure(t, `DECLARE "foo"`, ErrDeclareParsingName)
	ExpectParseWrappedFailure(t, `DECLARE 1`, ErrDeclareParsingName)
	ExpectParseWrappedFailure(t, `DECLARE "<foo> bar"`, ErrDeclareParsingType)
	ExpectParseWrappedFailure(t, `DECLARE "<foo> "bar""`, ErrDeclareParsingType)
	ExpectParseWrappedFailure(t, `DECLARE "<foo> 1"`, ErrDeclareParsingType)
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
	} else if value, err := vm.memory.FetchVariableByName(name); err != nil {
		t.Fatalf("Error popping stack after performing DECLARE: %v", err)
	} else if display := value.Display(); display != message {
		t.Fatalf("Expected top of stack to be %q, instead got %q", message, display)
	}
}
