package vm

import (
	"testing"
)

func TestBinopParsing(t *testing.T) {
	expected_type := "*vm.Instruction_BINOP"

	ExpectParse(t, "BINOP add", expected_type, "BINOP add")
	ExpectParse(t, "BINOP   add", expected_type, "BINOP add")
	ExpectParse(t, "BINOP   add  ", expected_type, "BINOP add")
	ExpectParse(t, "  BINOP   add  ", expected_type, "BINOP add")
	ExpectParse(t, "\tBINOP   add  ", expected_type, "BINOP add")
	ExpectParseFailure(t, "BINOP", "needs operator")
	ExpectParseFailure(t, "BINOP foo", `unexpected operator "foo"`)
}
