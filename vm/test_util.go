package vm

import (
	"fmt"
	"strings"
	"testing"
)

func RequireParse(t *testing.T, input string) Instruction {
	if instruction, err := ParseInstruction(input); err != nil {
		t.Fatalf("ParseInstruction(%q) has unexpected error: %v", input, err)
		return nil
	} else {
		return instruction
	}
}

func ExpectParse(t *testing.T, input string, expected_type string, expected_string string) {
	if instruction, err := ParseInstruction(input); err != nil {
		t.Fatalf("ParseInstruction(%q) has unexpected error: %v", input, err)
	} else if instruction_type := fmt.Sprintf("%T", instruction); instruction_type != expected_type {
		t.Fatalf("ParseInstruction(%q) expected type %q, actual type %q", input, expected_type, instruction_type)
	} else if instruction.String() != expected_string {
		t.Fatalf("ParseInstruction(%q).String() expected %q, actual %q", input, expected_string, instruction)
	}
}

func ExpectParseFailure(t *testing.T, input string, expect_contains string) {
	if instruction, err := ParseInstruction(input); err == nil {
		t.Fatalf("ParseInstruction(%q) expected error containing %q but received no error", input, expect_contains)
	} else if !strings.Contains(err.Error(), expect_contains) {
		t.Fatalf("ParseInstruction(%q) expected error containing %q but received %v", input, expect_contains, err)
	} else if instruction != nil {
		t.Fatalf("ParseInstruction(%q) expected nil instruction but received %v", input, instruction)
	}
}
