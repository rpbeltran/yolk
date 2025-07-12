package vm

import (
	"bytes"
	"testing"
	"yolk/types"
)

func TestFicParsing(t *testing.T) {
	expected_type := "*vm.Instruction_FINISHED_INTERACTIVE_COMMAND"
	expected_string := "FINISHED_INTERACTIVE_COMMAND"

	ExpectParse(t, "FINISHED_INTERACTIVE_COMMAND", expected_type, expected_string)
	ExpectParseFailure(t, "PRINT 2", "expected no arguments")
}

func TestFicPerformPrintsFICCommand(t *testing.T) {
	vm := VirtualMachine{}
	var output bytes.Buffer
	vm.stdout = &output

	fic_instruction := RequireParse(t, "FINISHED_INTERACTIVE_COMMAND")

	if err := fic_instruction.Perform(&vm); err != nil {
		t.Fatalf("Unexpected error executing FINISHED_INTERACTIVE_COMMAND: %v", err)
	}

	expected_output := "$$$###@@@::FINISHED_INTERACTIVE_COMMAND::1bc82ncv3yur::@@@###$$$\n"
	if actual := output.String(); actual != expected_output {
		t.Fatalf("FINISHED_INTERACTIVE_COMMAND outputed %q, expected %q", actual, expected_output)
	}

}

func TestFicPerformFailsIfStackNonempty(t *testing.T) {
	vm := VirtualMachine{}
	vm.stack.Push(types.MakeString("foo"))

	fic_instruction := RequireParse(t, "FINISHED_INTERACTIVE_COMMAND")
	if err := fic_instruction.Perform(&vm); err == nil {
		t.Fatalf("Expected error executing FINISHED_INTERACTIVE_COMMAND, got none")
	}
}
