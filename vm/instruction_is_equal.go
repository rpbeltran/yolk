package vm

import (
	"fmt"
	"yolk/types"
)

type Instruction_IS_EQUAL struct{}

func (instruction *Instruction_IS_EQUAL) Parse(args *string) error {
	if len(*args) != 0 {
		return fmt.Errorf("IS_EQUAL instruction expected no arguments, received %q", *args)
	}
	return nil
}

func (instruction *Instruction_IS_EQUAL) String() string {
	return "IS_EQUAL"
}

func (instruction *Instruction_IS_EQUAL) Perform(vm *VirtualMachine) error {
	if left, err := vm.stack.Pop(); err != nil {
		return fmt.Errorf("popping lhs for IS_EQUAL: %v", err)
	} else if right, err := vm.stack.Pop(); err != nil {
		return fmt.Errorf("popping rhs for IS_EQUAL: %v", err)
	} else {
		equal := types.MakeBool(left.Equal(right))
		vm.stack.Push(equal)
	}
	return nil
}
