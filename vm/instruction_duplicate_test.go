package vm

import (
	"testing"
	"yolk/types"
)

func TestDuplicateParsing(t *testing.T) {
	expected_type := "*vm.Instruction_DUPLICATE"

	ExpectParseSame(t, `DUPLICATE`, expected_type)
	ExpectParseFailure(t, "DUPLICATE foo", "expected no arguments")
	ExpectParseFailure(t, "DUPLICATE true", "expected no arguments")
	ExpectParseFailure(t, "DUPLICATE 1", "expected no arguments")
}

func TestDuplicate(t *testing.T) {
	vm := NewVM()

	instruction := RequireParse(t, "DUPLICATE")

	if err := instruction.Perform(&vm); err == nil {
		t.Fatalf("Expected an error performing DUPLICATE without arguments: %v", err)
	}

	item := types.MakeBool(true)

	vm.stack.Push(item)

	if err := instruction.Perform(&vm); err != nil {
		t.Fatalf("Unexpected error performing DUPLICATE: %v", err)
	} else if vm.stack.Size() != 2 {
		t.Fatalf("Expected stack size to be 2 after performing DUPLICATE, got %d", vm.stack.Size())
	} else if value, err := vm.stack.Pop(); err != nil {
		t.Fatalf("Unexpected error popping stack after DUPLICATE: %v", err)
	} else if !value.Equal(item) {
		t.Fatalf("Expected stack to contain %q after performing DUPLICATE with %q, got %q", item.Display(), value.Display(), item.Display())
	} else if value, err := vm.stack.Pop(); err != nil {
		t.Fatalf("Unexpected error popping stack a second time after DUPLICATE: %v", err)
	} else if !value.Equal(item) {
		t.Fatalf("Expected stack to contain %q after performing DUPLICATE with %q, got %q", item.Display(), value.Display(), item.Display())
	}

}
