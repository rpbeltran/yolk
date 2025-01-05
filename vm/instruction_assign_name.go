package vm

import (
	"fmt"
)

type Instruction_ASSIGN struct {
	name string
}

func (instruction *Instruction_ASSIGN) Parse(args *string) error {
	if len(*args) == 0 {
		return fmt.Errorf("ASSIGN instruction needs a name")
	}
	instruction.name = *args
	return nil
}

func (instruction *Instruction_ASSIGN) String() string {
	return fmt.Sprintf("ASSIGN %s", instruction.name)
}

func (instruction *Instruction_ASSIGN) Perform(vm *VirtualMachine) error {
	value, err := vm.stack.Pop()
	if err != nil {
		return fmt.Errorf("unexpected error popping value for assignment: %w", err)
	}

	if err := vm.UpdateVariable(instruction.name, value); err != nil {
		return err
	}
	return nil
}
