package vm

import (
	"fmt"
)

type Instruction_LOAD_NAME struct {
	name     *string
	id       uint64
	executed bool
}

func (instruction *Instruction_LOAD_NAME) Parse(args *string) error {
	if len(*args) == 0 {
		return fmt.Errorf("LOAD_NAME instruction needs a name")
	}
	instruction.name = args
	instruction.executed = false
	return nil
}

func (instruction *Instruction_LOAD_NAME) String() string {
	return fmt.Sprintf("LOAD_NAME %s", *instruction.name)
}

func (instruction *Instruction_LOAD_NAME) Perform(vm *VirtualMachine) error {
	if instruction.executed {
		if value, err := vm.FetchVariableById(instruction.id); err != nil {
			return fmt.Errorf("unexpected error fetching variable %q with a cached id: %v", *instruction.name, instruction.id)
		} else {
			vm.stack.Push(value)
		}
	} else if value, id, err := vm.FetchVariable(*instruction.name); err != nil {
		return err
	} else {
		instruction.id = id
		instruction.executed = true
		vm.stack.Push(value)
	}
	return nil
}
