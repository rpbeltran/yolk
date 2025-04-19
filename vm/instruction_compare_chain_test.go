package vm

import (
	"fmt"
	"testing"
	"yolk/types"
)

func TestCompareChainParsing(t *testing.T) {
	expected_type := "*vm.Instruction_COMPARE_CHAIN"

	ExpectParse(t, "COMPARE_CHAIN equal", expected_type, "COMPARE_CHAIN equal")
	ExpectParse(t, "COMPARE_CHAIN unequal", expected_type, "COMPARE_CHAIN unequal")
	ExpectParse(t, "COMPARE_CHAIN less", expected_type, "COMPARE_CHAIN less")
	ExpectParse(t, "COMPARE_CHAIN lte", expected_type, "COMPARE_CHAIN lte")
	ExpectParse(t, "COMPARE_CHAIN greater", expected_type, "COMPARE_CHAIN greater")
	ExpectParse(t, "COMPARE_CHAIN gte", expected_type, "COMPARE_CHAIN gte")
	ExpectParseFailure(t, "COMPARE_CHAIN", "needs a test mode")
	ExpectParseFailure(t, "COMPARE_CHAIN 1", "unexpected test mode")
	ExpectParseFailure(t, "COMPARE_CHAIN foo", "unexpected test mode")
	ExpectParseFailure(t, "COMPARE_CHAIN true", "unexpected test mode")
}

func TestCompareChainPerform(t *testing.T) {
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

		vm.stack.Push(test.left)
		vm.stack.Push(test.right)

		test_instruction := fmt.Sprintf("COMPARE_CHAIN %s", test.comparison)

		if instruction, err := ParseInstruction(test_instruction); err != nil {
			t.Fatalf("Error parsing instruction %q: %v", test_instruction, err)
		} else if err := instruction.Perform(&vm); err != nil {
			t.Fatalf("Unexpected error executing %q: %v", test_instruction, err)
		} else if stack_size := vm.stack.Size(); stack_size != 2 {
			t.Fatalf("Unexpected stack to have 1 element after %q, had: %d", test_instruction, stack_size)
		} else if top, err := vm.stack.Pop(); err != nil {
			t.Fatalf("Unexpected error popping stack after %q: %v", test_instruction, err)
		} else if !top.Equal(test.right) {
			t.Fatalf("Expected rhs (%v) on top, got (%v)", test.right, top)
		} else if result, err := vm.stack.Pop(); err != nil {
			t.Fatalf("Unexpected error popping stack after %q: %v", test_instruction, err)
		} else if result_bool, err := result.RequireBool(); err != nil {
			t.Fatalf("Unexpected error interpretting result of %q as a bool: %v", test_instruction, err)
		} else if result_bool.Truthy() != test.expected {
			t.Fatalf("Expected %q to push %t, got %t", test_instruction, test.expected, !test.expected)
		}
	}
}

func TestCompareChainArgFailures(t *testing.T) {
	vm := NewVM()

	if instruction, err := ParseInstruction("COMPARE_CHAIN equal"); err != nil {
		t.Fatalf("Error parsing instruction %q: %v", instruction, err)
	} else if err := instruction.Perform(&vm); err == nil {
		t.Fatalf("Expected error executing 'COMPARE_CHAIN equal', got success")
	}

	vm.stack.Push(types.MakeString("foo"))

	if instruction, err := ParseInstruction("COMPARE_CHAIN equal"); err != nil {
		t.Fatalf("Error parsing instruction %q: %v", instruction, err)
	} else if err := instruction.Perform(&vm); err == nil {
		t.Fatalf("Expected error executing 'COMPARE_CHAIN equal', got success")
	}
}
