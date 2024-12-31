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

func TestLabelAddTo(t *testing.T) {
	vm := NewVM()
	label := Instruction_LABEL{100}
	if err := label.AddTo(&vm); err != nil {
		t.Fatalf("Unexpected error from %q: %v", ".LABEL 100", err)
	} else if dst, ok := vm.labels[100]; !ok {
		t.Fatal("Expected vm to have label for 100, had none")
	} else if dst != 0 {
		t.Fatalf("Expected label %d to point to %d, instead points to %d", 100, 0, dst)
	}
}
