package instructions

import (
	"fmt"
	"yolk/vm"
)

type Instruction_PRINT struct{}

func (instruction *Instruction_PRINT) Parse(args *string) error {
	if len(*args) != 0 {
		return fmt.Errorf("PRINT instruction expected no arguments, received '%s'", *args)
	}
	return nil
}

func (instruction *Instruction_PRINT) Perform(machine *vm.VirtualMachine) error {
	fmt.Println("//TODO: Implement PRINT instruction")
	return nil
}
