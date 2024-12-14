package instructions

import (
	"testing"
)

func TestPushIntParsing(t *testing.T) {
	expected_type := "*instructions.Instruction_PUSH_INT"

	ExpectParse(t, "PUSH_INT 123", expected_type, "PUSH_INT 123")
	ExpectParse(t, "PUSH_INT -123", expected_type, "PUSH_INT -123")
	ExpectParse(t, "PUSH_INT   1", expected_type, "PUSH_INT 1")
	ExpectParse(t, "PUSH_INT   0  ", expected_type, "PUSH_INT 0")
	ExpectParse(t, "  PUSH_INT   123456789  ", expected_type, "PUSH_INT 123456789")
	ExpectParse(t, "PUSH_INT   314  ", expected_type, "PUSH_INT 314")
	ExpectParseFailure(t, "PUSH_INT", "needs a value")
	ExpectParseFailure(t, "PUSH_INT foo", `invalid value "foo"`)
}
