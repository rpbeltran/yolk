package vm

import (
	"bytes"
	"strings"
	"testing"
	"yolk/types"
)

func TestPrintParsing(t *testing.T) {
	expected_type := "*vm.Instruction_PRINT"
	expected_string := "PRINT"

	ExpectParse(t, "PRINT", expected_type, expected_string)
	ExpectParseFailure(t, "PRINT 2", "expected no arguments")
}

func TestPrintPerform(t *testing.T) {
	vm := VirtualMachine{}
	var output bytes.Buffer
	vm.stdout = &output

	print_instruction := RequireParse(t, "PRINT")

	if err := print_instruction.Perform(&vm); err == nil {
		t.Fatalf("Epected error executing PRINT: %v", err)
	}

	phrases := []string{"Hello", "", "foo!!", "12345", "''", `""`}

	for i := range phrases {
		vm.stack.Push(types.MakeString(phrases[len(phrases)-1-i]))
	}

	for range phrases {
		if err := print_instruction.Perform(&vm); err != nil {
			t.Fatalf("Unexpected error executing PRINT: %v", err)
		}
	}

	expected_output := strings.Join(phrases, "\n") + "\n"
	if actual := output.String(); actual != expected_output {
		t.Fatalf("PRINTS outputed %q, expected %q", actual, expected_output)
	}

	if err := print_instruction.Perform(&vm); err == nil {
		t.Fatalf("Epected error executing PRINT: %v", err)
	}

}
