package vm

import (
	"fmt"
)

type Instruction_ASSIGN_NAME struct {
	name     *string
	id       uint64
	executed bool
}

func (instruction *Instruction_ASSIGN_NAME) Parse(args *string) error {
	if len(*args) == 0 {
		return fmt.Errorf("ASSIGN_NAME instruction needs a name")
	}
	instruction.name = args
	instruction.executed = false
	return nil
}

func (instruction *Instruction_ASSIGN_NAME) String() string {
	return fmt.Sprintf("ASSIGN_NAME %s", *instruction.name)
}

func (instruction *Instruction_ASSIGN_NAME) Perform(vm *VirtualMachine) error {
	value, err := vm.stack.Pop()
	if err != nil {
		return fmt.Errorf("unexpected error popping value for assignment: %w", err)
	}

	if instruction.executed {
		if err := vm.UpdateVariableByID(instruction.id, value); err != nil {
			return fmt.Errorf("unexpected error updating variable %q with a cached id: %v", *instruction.name, instruction.id)
		}
	} else if id, err := vm.UpdateVariable(*instruction.name, value); err != nil {
		return err
	} else {
		instruction.id = id
		instruction.executed = true
	}
	return nil
}
