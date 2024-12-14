package instructions

import (
	"testing"
)

func TestPipelineParsing(t *testing.T) {
	expected_type := "*instructions.Instruction_PIPELINE"

	ExpectParse(t, "PIPELINE begin", expected_type, "PIPELINE begin")
	ExpectParse(t, "PIPELINE next", expected_type, "PIPELINE next")
	ExpectParse(t, "PIPELINE end", expected_type, "PIPELINE end")
	ExpectParseFailure(t, "PIPELINE", "needs operation")
	ExpectParseFailure(t, "PIPELINE foo", `unexpected operation "foo"`)
}
