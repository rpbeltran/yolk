package vm

import (
	"fmt"
	"yolk/types"
)

type Instruction_PUSH_NUM struct {
	value *types.PrimitiveNum
}

func (instruction *Instruction_PUSH_NUM) Parse(args *string) error {
	if num, err := types.MakeNumber(*args); err != nil {
		if len(*args) == 0 {
			return fmt.Errorf("PUSH_NUM instruction needs a value")
		}
		return fmt.Errorf("PUSH_NUM instruction has invalid value %q", *args)
	} else {
		instruction.value = num
	}
	return nil
}

func (instruction *Instruction_PUSH_NUM) String() string {
	return fmt.Sprintf("PUSH_NUM %s", (*instruction.value).Display())
}

func (instruction *Instruction_PUSH_NUM) Perform(vm *VirtualMachine) error {
	vm.stack.Push(instruction.value)
	return nil
}
