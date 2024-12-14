package instructions

import (
	"testing"
)

func TestPrintParsing(t *testing.T) {
	expected_type := "*instructions.Instruction_PRINT"
	expected_string := "PRINT"

	ExpectParse(t, "PRINT", expected_type, expected_string)
	ExpectParse(t, "PRINT  ", expected_type, expected_string)
	ExpectParse(t, "  PRINT  ", expected_type, expected_string)
	ExpectParseFailure(t, "PRINT 2", "expected no arguments")
}
