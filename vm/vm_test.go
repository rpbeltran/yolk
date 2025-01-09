package vm

import (
	"bufio"
	"strings"
	"testing"
)

func TestVMGetPipeIn(t *testing.T) {
	vm := VirtualMachine{}
	if value, has_input := vm.GetPipeIn(); has_input {
		t.Fatalf("vm.GetPipeIn() gave %v, expected no value", *value)
	}

	if err := RequireParse(t, "PIPELINE begin").Perform(&vm); err != nil {
		t.Fatalf(`Unexpected error from "PIPELINE begin": %v`, err)
	}

	if value, has_input := vm.GetPipeIn(); has_input {
		t.Fatalf("vm.GetPipeIn() gave %v, expected no value", *value)
	}

	if err := RequireParse(t, `PUSH_STR "a"`).Perform(&vm); err != nil {
		t.Fatalf(`Unexpected error from 'PUSH_STR "a"': %v`, err)
	}

	if value, has_input := vm.GetPipeIn(); has_input {
		t.Fatalf("vm.GetPipeIn() gave %v, expected no value", *value)
	}

	if err := RequireParse(t, "PIPELINE next").Perform(&vm); err != nil {
		t.Fatalf(`Unexpected error from "PIPELINE next": %v`, err)
	}

	if value, has_input := vm.GetPipeIn(); !has_input {
		t.Fatal(`vm.GetPipeIn() had no value, expected pointer to "a"`)
	} else if value == nil {
		t.Fatal(`vm.GetPipeIn() gave nil, expected pointer to "a"`)
	} else if actual := (*value).Display(); actual != "a" {
		t.Fatalf(`vm.GetPipeIn() gave %q, expected pointer to "a"`, actual)
	}

	if err := RequireParse(t, `PUSH_STR "b"`).Perform(&vm); err != nil {
		t.Fatalf(`Unexpected error from 'PUSH_STR "a"': %v`, err)
	}

	if value, has_input := vm.GetPipeIn(); !has_input {
		t.Fatal(`vm.GetPipeIn() had no value, expected pointer to "a"`)
	} else if value == nil {
		t.Fatal(`vm.GetPipeIn() gave nil, expected pointer to "a"`)
	} else if actual := (*value).Display(); actual != "a" {
		t.Fatalf(`vm.GetPipeIn() gave %q, expected pointer to "a"`, actual)
	}

	if err := RequireParse(t, "PIPELINE end").Perform(&vm); err != nil {
		t.Fatalf(`Unexpected error from "PIPELINE next": %v`, err)
	}

	if value, has_input := vm.GetPipeIn(); has_input {
		t.Fatalf("vm.GetPipeIn() gave %v, expected no value", *value)
	}
}

func TestPutProgramInVM(t *testing.T) {
	input := strings.Join([]string{
		// main
		"PRINT",
		".LABEL 100",
		"JUMP 300",
		".LABEL 300",

		// foo
		".DEFINE <foo>",
		"LOAD <name>",
		".DEFINE_END",

		// spam return_type
		".DEFINE <spam> <return_type>",
		"PUSH_INT 1",
		"PRINT",
		".DEFINE_END",

		// main
		"PRINT",
	}, "\n")

	vm := NewVM()
	if err := vm.PutProgramInVM(bufio.NewScanner(strings.NewReader(input))); err != nil {
		t.Fatalf("Unexpected error in call to vm.PutProgramInVM: %v", err)
	}

	// Count program instructions
	if instruction_count := len(vm.program); instruction_count != 5 {
		t.Fatalf("Expected VM to have four instructions, instead has: %d (%v)", instruction_count, vm.program)
	}

	// Test labels
	if label_count := len(vm.labels); label_count != 2 {
		t.Fatalf("Expected VM to have two labels, instead has: %d", label_count)
	}
	if dst, ok := vm.labels[100]; !ok {
		t.Fatal("Expected VM to have the label, 100 -> 1 but found no label from 100")
	} else if dst != 1 {
		t.Fatalf("Expected VM to have the label, 100 -> 1, but instead found 100 -> %d", dst)
	}

	if dst, ok := vm.labels[300]; !ok {
		t.Fatal("Expected VM to have the label, 300 -> 3 but found no label from 300")
	} else if dst != 3 {
		t.Fatalf("Expected VM to have the label, 300 -> 3, but instead found 300 -> %d", dst)
	}

	// Test functions
	if function_count := len(vm.functions); function_count != 2 {
		t.Fatalf("Expected VM to have two functions, instead has: %d", function_count)
	}
	if function_definition_count := len(vm.function_definitions); function_definition_count != 2 {
		t.Fatalf("Expected VM to have two function definitions, instead has: %d", function_definition_count)
	}

	if f_instructions, has_f := vm.functions["foo"]; !has_f {
		t.Fatal("Expected VM to have a function 'foo' but it does not")
	} else if f_instruct_count := len(f_instructions); f_instruct_count != 1 {
		t.Fatalf("Expected VM function 'foo' to have length 1, instead had length %d", f_instruct_count)
	} else if f_instructions[0].String() != "LOAD <name>" {
		t.Fatalf("Expected VM function 'foo' to have LOAD instruction, instead had %q", f_instructions[0].String())
	} else if f_definition, has_f := vm.function_definitions["foo"]; !has_f {
		t.Fatal("Expected VM to have a function 'foo' but it does not")
	} else if f_definition.name != "foo" {
		t.Fatalf("Expected VM function 'foo' to have name 'foo', instead got %q", f_definition.name)
	} else if f_definition.has_type_annotation {
		t.Fatalf("Expected VM function 'foo' to have no type, instead got %q", f_definition.type_annotation)
	}

	if f_instructions, has_f := vm.functions["spam"]; !has_f {
		t.Fatal("Expected VM to have a function 'spam' but it does not")
	} else if f_instruct_count := len(f_instructions); f_instruct_count != 2 {
		t.Fatalf("Expected VM function 'spam' to have length 2, instead had length %d (%v)", f_instruct_count, f_instructions)
	} else if f_definition, has_f := vm.function_definitions["spam"]; !has_f {
		t.Fatal("Expected VM to have a function 'spam' but it does not")
	} else if f_definition.name != "spam" {
		t.Fatalf("Expected VM function 'spam' to have name 'spam', instead got %q", f_definition.name)
	} else if !f_definition.has_type_annotation {
		t.Fatal("Expected VM function 'spam' to have return type 'return_type', instead had none")
	} else if f_definition.type_annotation != "return_type" {
		t.Fatalf("Expected VM function 'spam' to have return type 'return_type', instead had %q", f_definition.type_annotation)
	}
}
