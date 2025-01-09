package vm

import (
	"testing"
)

func TestParseEmpty(t *testing.T) {
	if instruction, err := ParseInstruction(""); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	} else if instruction != nil {
		t.Fatalf(`ParseInstruction("") gave %v, expected nil`, instruction)
	}
}

func TestParseComments(t *testing.T) {
	if instruction, err := ParseInstruction("# hello"); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	} else if instruction != nil {
		t.Fatalf(`ParseInstruction("") gave %v, expected nil`, instruction)
	}
}

func TestParseUnknownCommand(t *testing.T) {
	expected_error := "unknown operator: FAKE_OP_CODE"
	if instruction, err := ParseInstruction("FAKE_OP_CODE"); err == nil {
		t.Fatal(`ParseInstruction("FAKE_OP_CODE") succeeded, expected an error`)
	} else if instruction != nil {
		t.Fatalf(`ParseInstruction("FAKE_OP_CODE") gave %v, expected nil`, instruction)
	} else if err.Error() != expected_error {
		t.Fatalf(`Expected error from ParseInstruction("FAKE_OP_CODE") to contain %q but got: %v`, expected_error, err)
	}
}
