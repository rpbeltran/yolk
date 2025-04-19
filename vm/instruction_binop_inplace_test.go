package vm

import (
	"fmt"
	"strings"
	"testing"
	"yolk/types"
)

func TestBinopInplaceParsing(t *testing.T) {
	expected_type := "*vm.Instruction_BINOP_INPLACE"

	ExpectParseSame(t, `BINOP_INPLACE add <sum>`, expected_type)
	ExpectParseSame(t, `BINOP_INPLACE add <foo>`, expected_type)
	ExpectParseSame(t, `BINOP_INPLACE subtract <foo>`, expected_type)
	ExpectParseSame(t, `BINOP_INPLACE multiply <foo>`, expected_type)
	ExpectParseSame(t, `BINOP_INPLACE divide <foo>`, expected_type)
	ExpectParseSame(t, `BINOP_INPLACE int_divide <foo>`, expected_type)
	ExpectParseSame(t, `BINOP_INPLACE power <foo>`, expected_type)
	ExpectParseSame(t, `BINOP_INPLACE modulus <foo>`, expected_type)
	ExpectParseSame(t, `BINOP_INPLACE concat <foo>`, expected_type)
	ExpectParseSame(t, `BINOP_INPLACE and <foo>`, expected_type)
	ExpectParseSame(t, `BINOP_INPLACE or <foo>`, expected_type)
	ExpectParseFailure(t, "BINOP_INPLACE", "instruction needs operator and name")
	ExpectParseFailure(t, "BINOP_INPLACE foo", "instruction needs operator and name")
	ExpectParseFailure(t, "BINOP_INPLACE foo bar", "unexpected operator")
	ExpectParseWrappedFailure(t, `BINOP_INPLACE add ""`, ErrParsingBinopInplaceName)
	ExpectParseWrappedFailure(t, `BINOP_INPLACE add <>`, ErrParsingBinopInplaceName)
}

type BinOpInplaceTestCase struct {
	operation string
	lhs       types.Primitive
	rhs       types.Primitive
	result    types.Primitive
}

func TestBinopInplace(t *testing.T) {
	name := "foo"
	tests := []BinOpInplaceTestCase{
		{"add", RequireNum(t, "10"), RequireNum(t, "5"), RequireNum(t, "15")},
		{"subtract", RequireNum(t, "10"), RequireNum(t, "5"), RequireNum(t, "5")},
		{"multiply", RequireNum(t, "11"), RequireNum(t, "5"), RequireNum(t, "55")},
		{"divide", RequireNum(t, "12"), RequireNum(t, "5"), RequireNum(t, "2.4")},
		{"int_divide", RequireNum(t, "12"), RequireNum(t, "5"), RequireNum(t, "2")},
		{"power", RequireNum(t, "12"), RequireNum(t, "5"), RequireNum(t, "248832")},
		{"modulus", RequireNum(t, "12"), RequireNum(t, "5"), RequireNum(t, "2")},
		{"concat", types.MakeString("foo"), types.MakeString("bar"), types.MakeString("foobar")},
		{"and", types.MakeBool(true), types.MakeBool(false), types.MakeBool(false)},
		{"or", types.MakeBool(true), types.MakeBool(false), types.MakeBool(true)},
	}

	for _, tc := range tests {
		vm := NewVM()

		initial := tc.lhs.Display()

		lhs_id := vm.memory.StorePrimitive(tc.lhs)
		if err := vm.memory.BindNewVariable(name, lhs_id); err != nil {
			t.Fatalf("Unexpected error storing variable: %v", err)
		}

		vm.stack.Push(tc.rhs)

		instruction := fmt.Sprintf("BINOP_INPLACE %s <%s>", tc.operation, name)

		if err := RequireParse(t, instruction).Perform(&vm); err != nil {
			t.Fatalf("Unexpected error performing %q: %v", instruction, err)
		}

		if actual, err := vm.memory.FetchVariableByName(name); err != nil {
			t.Fatalf("Unexpected error fetching variable %q: %v", name, err)
		} else if !actual.Equal(tc.result) {
			t.Fatalf("Expected %q with values %q and %q to give %q, instead got %q",
				instruction, initial, tc.rhs.Display(), tc.result.Display(), actual.Display())
		}
	}
}

func TestBinopInplaceFailure(t *testing.T) {
	name := "foo"
	tests := []BinOpInplaceTestCase{
		{"add", types.MakeString("10"), RequireNum(t, "5"), nil},
		{"subtract", types.MakeString("10"), RequireNum(t, "5"), nil},
		{"multiply", types.MakeString("11"), RequireNum(t, "5"), nil},
		{"divide", RequireNum(t, "12"), RequireNum(t, "0"), nil},
		{"int_divide", RequireNum(t, "12"), RequireNum(t, "0"), nil},
		{"power", RequireNum(t, "-12"), RequireNum(t, "1.2"), nil},
		{"modulus", RequireNum(t, "12"), RequireNum(t, "0"), nil},
		{"concat", RequireNum(t, "12"), types.MakeString("bar"), nil},
		{"and", RequireNum(t, "12"), RequireNum(t, "0"), nil},
		{"or", RequireNum(t, "12"), types.MakeString("bar"), nil},
	}

	for _, tc := range tests {
		vm := NewVM()

		initial := tc.lhs.Display()

		lhs_id := vm.memory.StorePrimitive(tc.lhs)
		if err := vm.memory.BindNewVariable(name, lhs_id); err != nil {
			t.Fatalf("Unexpected error storing variable: %v", err)
		}
		vm.stack.Push(tc.rhs)

		instruction := fmt.Sprintf("BINOP_INPLACE %s <%s>", tc.operation, name)

		if err := RequireParse(t, instruction).Perform(&vm); err == nil {
			t.Fatalf("Expected an error performing %q with %q and %q, got none",
				tc.operation, initial, tc.rhs.Display())
		}
	}
}

func TestBinopInplaceArgsFailure(t *testing.T) {
	name := "foo"

	vm := NewVM()

	lhs_id := vm.memory.StorePrimitive(types.MakeString("hello world!"))
	if err := vm.memory.BindNewVariable(name, lhs_id); err != nil {
		t.Fatalf("Unexpected error storing variable: %v", err)
	}
	instruction := fmt.Sprintf("BINOP_INPLACE concat <%s>", name)

	if err := RequireParse(t, instruction).Perform(&vm); err == nil {
		t.Fatal("Expected an error performing BINOP_INPLACE with empty stack, got none")
	} else if !strings.Contains(err.Error(), "rhs") {
		t.Fatalf("Expected an error containing 'rhs' got %v", err)
	}

	instruction = fmt.Sprintf("BINOP_INPLACE concat <%s>", "fake_name")
	vm.stack.Push(types.MakeString("hello"))
	if err := RequireParse(t, instruction).Perform(&vm); err == nil {
		t.Fatal("Expected an error performing BINOP_INPLACE with empty stack, got none")
	} else if !strings.Contains(err.Error(), "lhs") {
		t.Fatalf("Expected an error containing 'lhs' got %v", err)
	}

}
