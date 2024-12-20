package vm

import "testing"

func TestVMGetPipeIn(t *testing.T) {
	vm := VirtualMachine{}
	if value, has_input := vm.GetPipeIn(); has_input {
		t.Fatalf("vm.GetPipeIn() gave %v, expected no value", *value)
	}

	if err := RequireParse(t, "PIPELINE begin").Perform(&vm); err != nil {
		t.Fatalf(`Unexpected error from "PIPELINE begin": %v`, err)
	}

	if value, has_input := vm.GetPipeIn(); !has_input {
		t.Fatal("vm.GetPipeIn() had no value, expected nil")
	} else if value != nil {
		t.Fatalf("vm.GetPipeIn() gave %v, expected no nil", *value)
	}

	if err := RequireParse(t, `PUSH_STR "a"`).Perform(&vm); err != nil {
		t.Fatalf(`Unexpected error from 'PUSH_STR "a"': %v`, err)
	}

	if value, has_input := vm.GetPipeIn(); !has_input {
		t.Fatal("vm.GetPipeIn() had no value, expected nil")
	} else if value != nil {
		t.Fatalf("vm.GetPipeIn() gave %v, expected no nil", *value)
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
