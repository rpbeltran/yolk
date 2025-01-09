package vm

import (
	"testing"
)

func TestLabelParsing(t *testing.T) {
	expected_type := "*vm.Instruction_LABEL"

	ExpectParse(t, `.LABEL 0`, expected_type, `.LABEL 0`)
	ExpectParse(t, `.LABEL 123`, expected_type, `.LABEL 123`)
	ExpectParseFailure(t, ".LABEL", "needs an address")
	ExpectParseFailure(t, ".LABEL hello", `invalid address "hello"`)
	ExpectParseFailure(t, ".LABEL -10", `invalid address "-10"`)
}
