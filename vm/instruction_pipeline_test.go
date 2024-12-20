package vm

import (
	"strings"
	"testing"
)

func TestPipelineParsing(t *testing.T) {
	expected_type := "*vm.Instruction_PIPELINE"

	ExpectParse(t, "PIPELINE begin", expected_type, "PIPELINE begin")
	ExpectParse(t, "PIPELINE next", expected_type, "PIPELINE next")
	ExpectParse(t, "PIPELINE end", expected_type, "PIPELINE end")
	ExpectParseFailure(t, "PIPELINE", "needs operation")
	ExpectParseFailure(t, "PIPELINE foo", `unexpected operation "foo"`)
}

func TestPipelinePerform(t *testing.T) {
	//  "a" | ("b" | "c") | "d"
	vm := VirtualMachine{}
	// PIPELINE begin
	if err := RequireParse(t, "PIPELINE begin").Perform(&vm); err != nil {
		t.Fatalf(`Unexpected error from "PIPELINE begin": %v`, err)
	}
	if pipe_count := vm.pipeline_states.Size(); pipe_count != 1 {
		t.Fatalf(`Pipeline states had %d entrues, ecpected 1`, pipe_count)
	}
	if value, err := vm.pipeline_states.Peek(); err != nil {
		t.Fatalf("Unexpected error peeking at pipeline state: %v", err)
	} else if *value != nil {
		t.Fatalf("Unexpected value for pipeline state %v expected nil", *value)
	}
	// PUSH_STR "a"
	if err := RequireParse(t, `PUSH_STR "a"`).Perform(&vm); err != nil {
		t.Fatalf(`Unexpected error from 'PUSH_STR "a"': %v`, err)
	}
	// PIPELINE next
	if err := RequireParse(t, "PIPELINE next").Perform(&vm); err != nil {
		t.Fatalf(`Unexpected error from "PIPELINE next": %v`, err)
	}
	if pipe_count := vm.pipeline_states.Size(); pipe_count != 1 {
		t.Fatalf(`Pipeline states had %d entrues, ecpected 1`, pipe_count)
	}
	if value, err := vm.pipeline_states.Peek(); err != nil {
		t.Fatalf("Unexpected error peeking at pipeline state: %v", err)
	} else if *value == nil {
		t.Fatal(`Unexpected nil value for pipeline state, expected "a"`)
	} else if actual := (**value).Display(); actual != "a" {
		t.Fatalf(`Unexpected value %q for pipeline state, expected "a"`, actual)
	}
	// PIPELINE begin
	if err := RequireParse(t, "PIPELINE begin").Perform(&vm); err != nil {
		t.Fatalf(`Unexpected error from "PIPELINE begin": %v`, err)
	}
	if pipe_count := vm.pipeline_states.Size(); pipe_count != 2 {
		t.Fatalf(`Pipeline states had %d entrues, ecpected 2`, pipe_count)
	}
	if value, err := vm.pipeline_states.Peek(); err != nil {
		t.Fatalf("Unexpected error peeking at pipeline state: %v", err)
	} else if *value != nil {
		t.Fatalf("Unexpected value for pipeline state %v expected nil", *value)
	}
	// PUSH_STR "b"
	if err := RequireParse(t, `PUSH_STR "b"`).Perform(&vm); err != nil {
		t.Fatalf(`Unexpected error from 'PUSH_STR "b"': %v`, err)
	}
	// PIPELINE next
	if err := RequireParse(t, "PIPELINE next").Perform(&vm); err != nil {
		t.Fatalf(`Unexpected error from "PIPELINE next": %v`, err)
	}
	if pipe_count := vm.pipeline_states.Size(); pipe_count != 2 {
		t.Fatalf(`pPpeline states had %d entrues, ecpected 1`, pipe_count)
	}
	if value, err := vm.pipeline_states.Peek(); err != nil {
		t.Fatalf("Unexpected error peeking at pipeline state: %v", err)
	} else if value == nil {
		t.Fatal(`Unexpected nil value for pipeline state, expected "b"`)
	} else if actual := (**value).Display(); actual != "b" {
		t.Fatalf(`Unexpected value %q for pipeline state, expected "b"`, actual)
	}
	// PUSH_STR "c"
	if err := RequireParse(t, `PUSH_STR "c"`).Perform(&vm); err != nil {
		t.Fatalf(`Unexpected error from 'PUSH_STR "c"': %v`, err)
	}
	// PIPELINE end
	if err := RequireParse(t, "PIPELINE end").Perform(&vm); err != nil {
		t.Fatalf(`Unexpected error from "PIPELINE next": %v`, err)
	}
	if pipe_count := vm.pipeline_states.Size(); pipe_count != 1 {
		t.Fatalf(`Pipeline states had %d entries, expected 1`, pipe_count)
	}
	if value, err := vm.pipeline_states.Peek(); err != nil {
		t.Fatalf("Unexpected error peeking at pipeline state: %v", err)
	} else if *value == nil {
		t.Fatal(`Unexpected nil value for pipeline state, expected "a"`)
	} else if actual := (**value).Display(); actual != "a" {
		t.Fatalf(`Unexpected value %q for pipeline state, expected "a"`, actual)
	}
	// PIPELINE next
	if err := RequireParse(t, "PIPELINE next").Perform(&vm); err != nil {
		t.Fatalf(`Unexpected error from "PIPELINE next": %v`, err)
	}
	if pipe_count := vm.pipeline_states.Size(); pipe_count != 1 {
		t.Fatalf(`Pipeline states had %d entries, expected 1`, pipe_count)
	}
	if value, err := vm.pipeline_states.Peek(); err != nil {
		t.Fatalf("Unexpected error peeking at pipeline state: %v", err)
	} else if *value == nil {
		t.Fatal(`Unexpected nil value for pipeline state, expected "c"`)
	} else if actual := (**value).Display(); actual != "c" {
		t.Fatalf(`Unexpected value %q for pipeline state, expected "c"`, actual)
	}
	// PUSH_STR "d"
	if err := RequireParse(t, `PUSH_STR "d"`).Perform(&vm); err != nil {
		t.Fatalf(`Unexpected error from 'PUSH_STR "d"': %v`, err)
	}
	// PIPELINE end
	if err := RequireParse(t, "PIPELINE end").Perform(&vm); err != nil {
		t.Fatalf(`Unexpected error from "PIPELINE next": %v`, err)
	}
	if pipe_count := vm.pipeline_states.Size(); pipe_count != 0 {
		t.Fatalf(`Pipeline states had %d entries, expected 0`, pipe_count)
	}

	if actual := vm.stack.Size(); actual != 1 {
		t.Fatalf("Stack had %d items after operations, expected 1", actual)
	}
	if value, err := vm.stack.Pop(); err != nil {
		t.Fatalf("Unexpected error popping stack: %v", err)
	} else if actual := value.Display(); actual != "d" {
		t.Fatalf(`Unexpected value %q for pipeline state, expected "d"`, actual)
	}

}
func TestPipelinePerformFailureNextBeforeBegin(t *testing.T) {
	vm := VirtualMachine{}
	if err := RequireParse(t, "PIPELINE begin").Perform(&vm); err != nil {
		t.Fatalf(`Unexpected error from "PIPELINE begin": %v`, err)
	}
	expected_error := "could not pop stack"
	if err := RequireParse(t, "PIPELINE next").Perform(&vm); err == nil {
		t.Fatal(`Expected error from "PIPELINE end" but got none`)
	} else if !strings.Contains(err.Error(), (expected_error)) {
		t.Fatalf(`Expected error from "PIPELINE end" to contain %q but got: %v`, expected_error, err)
	}
}

func TestPipelinePerformFailureNextWithoutValue(t *testing.T) {
	vm := VirtualMachine{}
	expected_error := "could not pop pipeline state"
	if err := RequireParse(t, "PIPELINE next").Perform(&vm); err == nil {
		t.Fatal(`Expected error from "PIPELINE next" but got none`)
	} else if !strings.Contains(err.Error(), (expected_error)) {
		t.Fatalf(`Expected error from "PIPELINE next" to contain %q but got: %v`, expected_error, err)
	}
}

func TestPipelinePerformFailureEndBeforeBegin(t *testing.T) {
	vm := VirtualMachine{}
	expected_error := "could not pop pipeline state"
	if err := RequireParse(t, "PIPELINE end").Perform(&vm); err == nil {
		t.Fatal(`Expected error from "PIPELINE end" but got none`)
	} else if !strings.Contains(err.Error(), (expected_error)) {
		t.Fatalf(`Expected error from "PIPELINE next" to contain %q but got: %v`, expected_error, err)
	}
}
