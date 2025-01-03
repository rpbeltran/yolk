package vm

import (
	"fmt"
)

type Instruction_LOAD_NAME struct {
	name *string
}

func (instruction *Instruction_LOAD_NAME) Parse(args *string) error {
	if len(*args) == 0 {
		return fmt.Errorf("LOAD_NAME instruction needs a name")
	}
	instruction.name = args
	return nil
}

func (instruction *Instruction_LOAD_NAME) String() string {
	return fmt.Sprintf("LOAD_NAME %s", *instruction.name)
}

func (instruction *Instruction_LOAD_NAME) Perform(vm *VirtualMachine) error {
	if value, err := vm.FetchVariable(*instruction.name); err != nil {
		return err
	} else {
		vm.stack.Push(value)
	}
	return nil
}
