package vm

import (
	"testing"
)

func TestExecParsing(t *testing.T) {
	expected_type := "*vm.Instruction_EXEC"

	ExpectParse(t, "EXEC 123", expected_type, "EXEC 123")
	ExpectParse(t, "EXEC   1", expected_type, "EXEC 1")
	ExpectParse(t, "EXEC   0  ", expected_type, "EXEC 0")
	ExpectParse(t, "  EXEC   123456789  ", expected_type, "EXEC 123456789")
	ExpectParse(t, "EXEC   314  ", expected_type, "EXEC 314")
	ExpectParseFailure(t, "EXEC", "needs argument count")
	ExpectParseFailure(t, "EXEC foo", `invalid argument count "foo"`)
	ExpectParseFailure(t, "EXEC -1", `invalid argument count "-1"`)
}

func TestExecMockPerformNoArgs(t *testing.T) {
	vm := VirtualMachine{}
	vm.MockExecutions = true

	program := []string{
		`PUSH_STR "foo"`,
		"EXEC 0",
	}

	for _, line := range program {
		if err := RequireParse(t, line).Perform(&vm); err != nil {
			t.Fatalf(`Unexpected error from %q: %v`, line, err)
		}
	}

	expected_string := `{\n\t"command":"foo",\n\t"args":[]\n}`
	if actual_stack_size := vm.stack.Size(); actual_stack_size != 1 {
		t.Fatalf("Stack had %d items after operations, expected 1", actual_stack_size)
	}
	if value, err := vm.stack.Pop(); err != nil {
		t.Fatalf("Unexpected error popping stack: %v", err)
	} else if actual_string := value.Display(); actual_string != expected_string {
		t.Fatalf(`Executing "foo" gave %q, expected %q`, actual_string, expected_string)
	}
}

func TestExecMockPerformOneArg(t *testing.T) {
	vm := VirtualMachine{}
	vm.MockExecutions = true

	program := []string{
		`PUSH_STR "bar"`,
		`PUSH_STR "foo"`,
		"EXEC 1",
	}

	for _, line := range program {
		if err := RequireParse(t, line).Perform(&vm); err != nil {
			t.Fatalf(`Unexpected error from %q: %v`, line, err)
		}
	}

	expected_string := `{\n\t"command":"foo",\n\t"args":[bar]\n}`
	if actual_stack_size := vm.stack.Size(); actual_stack_size != 1 {
		t.Fatalf("Stack had %d items after operations, expected 1", actual_stack_size)
	}
	if value, err := vm.stack.Pop(); err != nil {
		t.Fatalf("Unexpected error popping stack: %v", err)
	} else if actual_string := value.Display(); actual_string != expected_string {
		t.Fatalf(`Executing "foo" gave %q, expected %q`, actual_string, expected_string)
	}
}

func TestExecMockPerformMultipleArgs(t *testing.T) {
	vm := VirtualMachine{}
	vm.MockExecutions = true

	program := []string{
		`PUSH_STR "c"`,
		`PUSH_STR "b"`,
		`PUSH_STR "a"`,
		`PUSH_STR "foo"`,
		"EXEC 3",
	}

	for _, line := range program {
		if err := RequireParse(t, line).Perform(&vm); err != nil {
			t.Fatalf(`Unexpected error from %q: %v`, line, err)
		}
	}

	expected_string := `{\n\t"command":"foo",\n\t"args":[a b c]\n}`
	if actual_stack_size := vm.stack.Size(); actual_stack_size != 1 {
		t.Fatalf("Stack had %d items after operations, expected 1", actual_stack_size)
	}
	if value, err := vm.stack.Pop(); err != nil {
		t.Fatalf("Unexpected error popping stack: %v", err)
	} else if actual_string := value.Display(); actual_string != expected_string {
		t.Fatalf(`Executing "foo" gave %q, expected %q`, actual_string, expected_string)
	}
}

func TestExecMockPerformPipein(t *testing.T) {
	vm := VirtualMachine{}
	vm.MockExecutions = true

	program := []string{
		"PIPELINE begin",
		`PUSH_STR "Hello World!"`,
		"PIPELINE next",
		`PUSH_STR "c"`,
		`PUSH_STR "b"`,
		`PUSH_STR "a"`,
		`PUSH_STR "foo"`,
		"EXEC 3",
	}

	for _, line := range program {
		if err := RequireParse(t, line).Perform(&vm); err != nil {
			t.Fatalf(`Unexpected error from %q: %v`, line, err)
		}
	}

	expected_string := `{\n\t"command":"foo",\n\t"args":[a b c],\n\t"stdin":"Hello World!"\n}`
	if actual_stack_size := vm.stack.Size(); actual_stack_size != 1 {
		t.Fatalf("Stack had %d items after operations, expected 1", actual_stack_size)
	}
	if value, err := vm.stack.Pop(); err != nil {
		t.Fatalf("Unexpected error popping stack: %v", err)
	} else if actual_string := value.Display(); actual_string != expected_string {
		t.Fatalf(`Executing "foo" gave %q, expected %q`, actual_string, expected_string)
	}
}
