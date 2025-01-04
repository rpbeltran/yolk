package vm

import (
	"fmt"
	"testing"
	"yolk/types"
)

func TestIsEqualParsing(t *testing.T) {
	expected_type := "*vm.Instruction_COMPARE"

	ExpectParse(t, "COMPARE equal", expected_type, "COMPARE equal")
	ExpectParse(t, "COMPARE unequal", expected_type, "COMPARE unequal")
	ExpectParse(t, "COMPARE less", expected_type, "COMPARE less")
	ExpectParse(t, "COMPARE lte", expected_type, "COMPARE lte")
	ExpectParse(t, "COMPARE greater", expected_type, "COMPARE greater")
	ExpectParse(t, "COMPARE gte", expected_type, "COMPARE gte")
	ExpectParseFailure(t, "COMPARE", "needs a test mode")
	ExpectParseFailure(t, "COMPARE 1", "unexpected test mode")
	ExpectParseFailure(t, "COMPARE foo", "unexpected test mode")
	ExpectParseFailure(t, "COMPARE true", "unexpected test mode")
}

type CompareTestCase struct {
	left       types.Primitive
	right      types.Primitive
	comparison string
	expected   bool
}

func TestComparePerform(t *testing.T) {
	tests := []CompareTestCase{
		{types.MakeString("foo"), types.MakeString("bar"), "equal", false},
		{types.MakeString("foo"), types.MakeString("foo"), "equal", true},
		{types.MakeString("foo"), types.MakeString("bar"), "unequal", true},
		{types.MakeString("foo"), types.MakeString("foo"), "unequal", false},
		{types.MakeString("a"), types.MakeString("b"), "less", true},
		{types.MakeString("a"), types.MakeString("a"), "less", false},
		{types.MakeString("b"), types.MakeString("a"), "less", false},
		{types.MakeString("a"), types.MakeString("b"), "lte", true},
		{types.MakeString("a"), types.MakeString("a"), "lte", true},
		{types.MakeString("b"), types.MakeString("a"), "lte", false},
		{types.MakeString("a"), types.MakeString("b"), "greater", false},
		{types.MakeString("a"), types.MakeString("a"), "greater", false},
		{types.MakeString("b"), types.MakeString("a"), "greater", true},
		{types.MakeString("a"), types.MakeString("b"), "gte", false},
		{types.MakeString("a"), types.MakeString("a"), "gte", true},
		{types.MakeString("b"), types.MakeString("a"), "gte", true},
	}

	for _, test := range tests {
		vm := NewVM()

		vm.stack.Push(test.right)
		vm.stack.Push(test.left)

		test_instruction := fmt.Sprintf("COMPARE %s", test.comparison)

		if instruction, err := ParseInstruction(test_instruction); err != nil {
			t.Fatalf("Error parsing instruction %q: %v", test_instruction, err)
		} else if err := instruction.Perform(&vm); err != nil {
			t.Fatalf("Unexpected error executing %q: %v", test_instruction, err)
		} else if stack_size := vm.stack.Size(); stack_size != 1 {
			t.Fatalf("Unexpected stack to have 1 element after %q, had: %d", test_instruction, stack_size)
		} else if top, err := vm.stack.Pop(); err != nil {
			t.Fatalf("Unexpected error popping stack after %q: %v", test_instruction, err)
		} else if top_bool, err := top.RequireBool(); err != nil {
			t.Fatalf("Unexpected error interpretting result of %q as a bool: %v", test_instruction, err)
		} else if top_bool.Truthy() != test.expected {
			t.Fatalf("Expected %q to push %t, got %t", test_instruction, test.expected, !test.expected)
		}
	}
}

func TestCompareArgFailures(t *testing.T) {
	vm := NewVM()

	if instruction, err := ParseInstruction("COMPARE equal"); err != nil {
		t.Fatalf("Error parsing instruction %q: %v", instruction, err)
	} else if err := instruction.Perform(&vm); err == nil {
		t.Fatalf("Expected error executing 'COMPARE equal', got success")
	}

	vm.stack.Push(types.MakeString("foo"))

	if instruction, err := ParseInstruction("COMPARE equal"); err != nil {
		t.Fatalf("Error parsing instruction %q: %v", instruction, err)
	} else if err := instruction.Perform(&vm); err == nil {
		t.Fatalf("Expected error executing 'COMPARE equal', got success")
	}
}
