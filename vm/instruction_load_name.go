package vm

import (
	"fmt"
)

type Instruction_LOAD struct {
	name *string
}

func (instruction *Instruction_LOAD) Parse(args *string) error {
	if len(*args) == 0 {
		return fmt.Errorf("LOAD instruction needs a name")
	}
	instruction.name = args
	return nil
}

func (instruction *Instruction_LOAD) String() string {
	return fmt.Sprintf("LOAD %s", *instruction.name)
}

func (instruction *Instruction_LOAD) Perform(vm *VirtualMachine) error {
	if value, err := vm.FetchVariable(*instruction.name); err != nil {
		return err
	} else {
		vm.stack.Push(value)
	}
	return nil
}
