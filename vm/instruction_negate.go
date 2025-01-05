package vm

import (
	"fmt"
)

type Instruction_NEGATE struct{}

func (instruction *Instruction_NEGATE) Parse(args *string) error {
	if len(*args) != 0 {
		return fmt.Errorf("NEGATE instruction expected no arguments, received '%s'", *args)
	}
	return nil
}

func (instruction *Instruction_NEGATE) String() string {
	return "NEGATE"
}

func (instruction *Instruction_NEGATE) Perform(vm *VirtualMachine) error {
	if value, err := vm.stack.Pop(); err != nil {
		return fmt.Errorf("popping value for NEGATE: %v", err)
	} else if negated_value, err := value.Negate(); err != nil {
		return fmt.Errorf("error computing NEGATE: %v", err)
	} else {
		vm.stack.Push(negated_value)
	}
	return nil
}
