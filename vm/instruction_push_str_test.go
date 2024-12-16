package vm

import (
	"testing"
)

func TestPushStrParsing(t *testing.T) {
	expected_type := "*vm.Instruction_PUSH_STR"

	ExpectParse(t, `PUSH_STR "FOO"`, expected_type, `PUSH_STR "FOO"`)
	ExpectParse(t, `PUSH_STR "123"`, expected_type, `PUSH_STR "123"`)
	ExpectParse(t, `PUSH_STR ""`, expected_type, `PUSH_STR ""`)
	ExpectParseFailure(t, "PUSH_STR", "needs a value")
	ExpectParseFailure(t, "PUSH_STR hello", `invalid value "hello"`)
	ExpectParseFailure(t, "PUSH_STR 0", `invalid value "0"`)
}
