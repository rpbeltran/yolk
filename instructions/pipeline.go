package instructions

import (
	"fmt"
	"yolk/vm"
)

type Instruction_PIPELINE struct {
	is_begin bool
	is_next  bool
	is_end   bool
}

func (instruction *Instruction_PIPELINE) Parse(args *string) error {
	switch *args {
	case "begin":
		instruction.is_begin = true
	case "next":
		instruction.is_next = true
	case "end":
		instruction.is_end = true
	default:
		return fmt.Errorf("PIPELINE instruction has unexpected mode '%s'", *args)
	}
	return nil
}

func (instruction *Instruction_PIPELINE) Perform(machine *vm.VirtualMachine) error {
	if instruction.is_begin {
		fmt.Println("//TODO: implement `PIPELINE begin`")
	} else if instruction.is_next {
		fmt.Println("//TODO: implement `PIPELINE next`")
	} else if instruction.is_end {
		fmt.Println("//TODO: implement `PIPELINE end`")
	}

	return nil
}
