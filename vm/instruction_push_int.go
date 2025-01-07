package vm

import (
	"fmt"
	"strconv"
	"yolk/types"
)

type Instruction_PUSH_INT struct {
	value *types.PrimitiveInt
}

func (instruction *Instruction_PUSH_INT) Parse(args *string) error {
	if value, err := strconv.ParseInt(*args, 10, 64); err == nil {
		instruction.value = types.MakeInt(value)
	} else {
		if len(*args) == 0 {
			return fmt.Errorf("PUSH_INT instruction needs a value")
		}
		return fmt.Errorf("PUSH_INT instruction has invalid value %q", *args)
	}
	return nil
}

func (instruction *Instruction_PUSH_INT) String() string {
	return fmt.Sprintf("PUSH_INT %s", (*instruction.value).Display())
}

func (instruction *Instruction_PUSH_INT) Perform(vm *VirtualMachine) error {
	vm.stack.Push(instruction.value)
	return nil
}
