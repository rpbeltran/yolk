package vm

import (
	"fmt"
)

type Instruction_DECLARE_NAME struct {
	name string
}

func (instruction *Instruction_DECLARE_NAME) Parse(args *string) error {
	if len(*args) == 0 {
		return fmt.Errorf("DECLARE_NAME instruction needs a name")
	}
	instruction.name = *args
	return nil
}

func (instruction *Instruction_DECLARE_NAME) String() string {
	return fmt.Sprintf("DECLARE_NAME %s", instruction.name)
}

func (instruction *Instruction_DECLARE_NAME) Perform(vm *VirtualMachine) error {
	value, err := vm.stack.Pop()
	if err != nil {
		return fmt.Errorf("unexpected error popping value for assignment: %w", err)
	}

	if err := vm.StoreNewVariable(instruction.name, value); err != nil {
		return err
	}
	return nil
}
