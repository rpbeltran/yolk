package vm

import (
	"fmt"
)

type Instruction_DUPLICATE struct{}

func (instruction *Instruction_DUPLICATE) Parse(args *string) error {
	if len(*args) != 0 {
		return fmt.Errorf("DUPLIOCATE instruction expected no arguments, received '%s'", *args)
	}
	return nil
}

func (instruction *Instruction_DUPLICATE) String() string {
	return "DUPLICATE"
}

func (instruction *Instruction_DUPLICATE) Perform(vm *VirtualMachine) error {
	if value, err := vm.stack.Peek(); err != nil {
		return fmt.Errorf("popping value for DUPLICATE: %v", err)
	} else {
		vm.stack.Push(*value)
	}
	return nil
}
