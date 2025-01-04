package vm

import (
	"fmt"
	"testing"
	"yolk/types"
)

func TestBinopParsing(t *testing.T) {
	expected_type := "*vm.Instruction_BINOP"

	ExpectParseSame(t, "BINOP add", expected_type)
	ExpectParseSame(t, "BINOP subtract", expected_type)
	ExpectParseSame(t, "BINOP multiply", expected_type)
	ExpectParseSame(t, "BINOP divide", expected_type)
	ExpectParseSame(t, "BINOP int_divide", expected_type)
	ExpectParseSame(t, "BINOP power", expected_type)
	ExpectParseSame(t, "BINOP modulus", expected_type)
	ExpectParseSame(t, "BINOP concat", expected_type)
	ExpectParseSame(t, "BINOP and", expected_type)
	ExpectParseSame(t, "BINOP or", expected_type)
	ExpectParse(t, "BINOP   add", expected_type, "BINOP add")
	ExpectParse(t, "BINOP   add  ", expected_type, "BINOP add")
	ExpectParse(t, "  BINOP   add  ", expected_type, "BINOP add")
	ExpectParse(t, "\tBINOP   add  ", expected_type, "BINOP add")
	ExpectParseFailure(t, "BINOP", "needs operator")
	ExpectParseFailure(t, "BINOP foo", `unexpected operator "foo"`)
}

type BinOpTestCase struct {
	operation string
	lhs       types.Primitive
	rhs       types.Primitive
	result    types.Primitive
}

func TestBinop(t *testing.T) {
	tests := []BinOpTestCase{
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

	for _, test := range tests {
		vm := NewVM()
		vm.stack.Push(test.rhs)
		vm.stack.Push(test.lhs)
		instruction := fmt.Sprintf("BINOP %s", test.operation)
		if err := RequireParse(t, instruction).Perform(&vm); err != nil {
			t.Fatalf("Unexpected error performing %q: %v", instruction, err)
		} else if vm.stack.Size() != 1 {
			t.Fatalf("Expected stack size after BINOP to be %d, got %d", 1, vm.stack.Size())
		} else if value, err := vm.stack.Pop(); err != nil {
			t.Fatalf("Unexpected error performing popping stack after BINOP: %v", err)
		} else if !value.Equal(test.result) {
			t.Fatalf("Expected %q with values %q and %q to give %q, instead got %q",
				instruction, test.lhs.Display(), test.rhs.Display(), test.result.Display(), value.Display())
		}
	}
}
