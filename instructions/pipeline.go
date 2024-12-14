package instructions

import (
	"fmt"
	"yolk/vm"
)

type operation uint8

const (
	pipeline_begin operation = iota
	pipeline_next  operation = iota
	pipeline_end   operation = iota
)

type Instruction_PIPELINE struct {
	operation operation
}

func (instruction *Instruction_PIPELINE) Parse(args *string) error {
	switch *args {
	case "begin":
		instruction.operation = pipeline_begin
	case "next":
		instruction.operation = pipeline_next
	case "end":
		instruction.operation = pipeline_end
	case "":
		return fmt.Errorf("PIPELINE instruction needs operation")
	default:
		return fmt.Errorf("PIPELINE instruction has unexpected operation %q", *args)
	}
	return nil
}

func (instruction *Instruction_PIPELINE) String() string {
	switch instruction.operation {
	case pipeline_begin:
		return "PIPELINE begin"
	case pipeline_next:
		return "PIPELINE next"
	case pipeline_end:
		return "PIPELINE end"
	}
	panic(fmt.Sprintf("PIPELINE instruction deserialized with unexpected mode %d", instruction.operation))
}

func (instruction *Instruction_PIPELINE) Perform(machine *vm.VirtualMachine) error {
	switch instruction.operation {
	case pipeline_begin:
		fmt.Println("//TODO: implement `PIPELINE begin`")
	case pipeline_next:
		fmt.Println("//TODO: implement `PIPELINE next`")
	case pipeline_end:
		fmt.Println("//TODO: implement `PIPELINE end`")
	}
	return nil
}
