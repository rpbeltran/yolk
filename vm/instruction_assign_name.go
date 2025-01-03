package vm

import (
	"fmt"
)

type Instruction_ASSIGN_NAME struct {
	name string
}

func (instruction *Instruction_ASSIGN_NAME) Parse(args *string) error {
	if len(*args) == 0 {
		return fmt.Errorf("ASSIGN_NAME instruction needs a name")
	}
	instruction.name = *args
	return nil
}

func (instruction *Instruction_ASSIGN_NAME) String() string {
	return fmt.Sprintf("ASSIGN_NAME %s", instruction.name)
}

func (instruction *Instruction_ASSIGN_NAME) Perform(vm *VirtualMachine) error {
	value, err := vm.stack.Pop()
	if err != nil {
		return fmt.Errorf("unexpected error popping value for assignment: %w", err)
	}

	if err := vm.UpdateVariable(instruction.name, value); err != nil {
		return err
	}
	return nil
}
