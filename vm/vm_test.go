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

func TestPutInVm(t *testing.T) {
	input := strings.Join([]string{
		"PRINT",
		".LABEL 100",
		"JUMP 300",
		".LABEL 300",
	}, "\n")

	vm := NewVM()
	if err := vm.PutProgramInVM(bufio.NewScanner(strings.NewReader(input))); err != nil {
		t.Fatalf("Unexpected error in call to vm.PutProgramInVM: %v", err)
	}

	if label_count := len(vm.labels); label_count != 2 {
		t.Fatalf("Expected VM to have two labels, instead has: %d", label_count)
	} else if dst, ok := vm.labels[100]; !ok {
		t.Fatal("Expected VM to have the label, 100 -> 1 but found no label from 100")
	} else if dst != 1 {
		t.Fatalf("Expected VM to have the label, 100 -> 1, but instead found 100 -> %d", dst)
	} else if dst, ok := vm.labels[300]; !ok {
		t.Fatal("Expected VM to have the label, 300 -> 3 but found no label from 300")
	} else if dst != 3 {
		t.Fatalf("Expected VM to have the label, 300 -> 3, but instead found 300 -> %d", dst)
	} else if instruction_count := len(vm.program); instruction_count != 4 {
		t.Fatalf("Expected VM to have four instructions, instead has: %d (%v)", instruction_count, vm.program)
	}
}
