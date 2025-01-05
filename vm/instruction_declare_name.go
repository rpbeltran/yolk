package vm

import (
	"fmt"
)

type Instruction_DECLARE struct {
	name string
}

func (instruction *Instruction_DECLARE) Parse(args *string) error {
	if len(*args) == 0 {
		return fmt.Errorf("DECLARE instruction needs a name")
	}
	instruction.name = *args
	return nil
}

func (instruction *Instruction_DECLARE) String() string {
	return fmt.Sprintf("DECLARE %s", instruction.name)
}

func (instruction *Instruction_DECLARE) Perform(vm *VirtualMachine) error {
	value, err := vm.stack.Pop()
	if err != nil {
		return fmt.Errorf("unexpected error popping value for assignment: %w", err)
	}

	if err := vm.StoreNewVariable(instruction.name, value); err != nil {
		return err
	}
	return nil
}
