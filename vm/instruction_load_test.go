package vm

import (
	"fmt"
	"testing"
	"yolk/types"
)

func TestLoadNameParsing(t *testing.T) {
	expected_type := "*vm.Instruction_LOAD"

	ExpectParseSame(t, `LOAD <foo>`, expected_type)
	ExpectParseSame(t, `LOAD <foo_bar>`, expected_type)
	ExpectParseWrappedFailure(t, "LOAD", ErrParsingLOAD)
	ExpectParseWrappedFailure(t, "LOAD foo", ErrParsingLOAD)
	ExpectParseWrappedFailure(t, `LOAD "foo"`, ErrParsingLOAD)
	ExpectParseWrappedFailure(t, `LOAD "<foo> hello"`, ErrParsingLOAD)
}

func TestLoadNamePerform(t *testing.T) {
	message := "hello world"
	name := "foo"

	vm := NewVM()

	message_id := vm.memory.StorePrimitive(types.MakeString(message))
	if err := vm.memory.BindNewVariable(name, message_id); err != nil {
		t.Fatalf("Unexpected error storing variable: %v", err)
	}

	if instruction, err := ParseInstruction(fmt.Sprintf("LOAD <%s>", name)); err != nil {
		t.Fatalf("Error parsing instruction %q: %v", instruction, err)
	} else if err := instruction.Perform(&vm); err != nil {
		t.Fatalf("Error executing instruction %v: %v", instruction, err)
	} else if value, err := vm.stack.Pop(); err != nil {
		t.Fatalf("Error popping stack after performing LOAD: %v", err)
	} else if display := value.Display(); display != message {
		t.Fatalf("Expected top of stack to be %q, instead got %q", message, display)
	} else {
		if err := instruction.Perform(&vm); err != nil {
			t.Fatalf("Error executing instruction %q: %v", instruction, err)
		} else if value, err := vm.stack.Pop(); err != nil {
			t.Fatalf("Error popping stack after performing LOAD: %v", err)
		} else if display := value.Display(); display != message {
			t.Fatalf("Expected top of stack to be %q, instead got %q", message, display)
		}
	}
	if instruction, err := ParseInstruction("LOAD <fake>"); err != nil {
		t.Fatalf("Error parsing instruction %q: %v", instruction, err)
	} else if err := instruction.Perform(&vm); err == nil {
		t.Fatalf("expecte error executing instruction %v, got none", instruction)
	}
}
