package vm

import (
	"fmt"
)

type Instruction_NOT struct{}

func (instruction *Instruction_NOT) Parse(args *string) error {
	if len(*args) != 0 {
		return fmt.Errorf("NOT instruction expected no arguments, received '%s'", *args)
	}
	return nil
}

func (instruction *Instruction_NOT) String() string {
	return "NOT"
}

func (instruction *Instruction_NOT) Perform(vm *VirtualMachine) error {
	if value, err := vm.stack.Pop(); err != nil {
		return fmt.Errorf("popping value for NOT: %v", err)
	} else if not_value, err := value.Not(); err != nil {
		return fmt.Errorf("error computing NOT: %v", err)
	} else {
		vm.stack.Push(not_value)
	}
	return nil
}
