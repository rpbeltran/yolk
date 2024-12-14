package instructions

import (
	"testing"
)

func TestExecParsing(t *testing.T) {
	expected_type := "*instructions.Instruction_EXEC"

	ExpectParse(t, "EXEC 123", expected_type, "EXEC 123")
	ExpectParse(t, "EXEC   1", expected_type, "EXEC 1")
	ExpectParse(t, "EXEC   0  ", expected_type, "EXEC 0")
	ExpectParse(t, "  EXEC   123456789  ", expected_type, "EXEC 123456789")
	ExpectParse(t, "EXEC   314  ", expected_type, "EXEC 314")
	ExpectParseFailure(t, "EXEC", "needs argument count")
	ExpectParseFailure(t, "EXEC foo", `invalid argument count "foo"`)
	ExpectParseFailure(t, "EXEC -1", `invalid argument count "-1"`)
}
