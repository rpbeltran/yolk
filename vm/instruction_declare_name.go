package vm

import (
	"fmt"
)

type Instruction_DECLARE_NAME struct {
	name     *string
	id       uint64
	executed bool
}

func (instruction *Instruction_DECLARE_NAME) Parse(args *string) error {
	if len(*args) == 0 {
		return fmt.Errorf("DECLARE_NAME instruction needs a name")
	}
	instruction.name = args
	instruction.executed = false
	return nil
}

func (instruction *Instruction_DECLARE_NAME) String() string {
	return fmt.Sprintf("DECLARE_NAME %s", *instruction.name)
}

func (instruction *Instruction_DECLARE_NAME) Perform(vm *VirtualMachine) error {
	value, err := vm.stack.Pop()
	if err != nil {
		return fmt.Errorf("unexpected error popping value for assignment: %w", err)
	}

	if instruction.executed {
		if err := vm.StoreNewVariableWithID(*instruction.name, instruction.id, value); err != nil {
			return fmt.Errorf("unexpected error updating variable %q with a cached id: %v", *instruction.name, instruction.id)
		}
	} else if id, err := vm.StoreNewVariable(*instruction.name, value); err != nil {
		return err
	} else {
		instruction.id = id
		instruction.executed = true
	}
	return nil
}
