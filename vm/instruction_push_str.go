package vm

import (
	"fmt"
	"strconv"
	"yolk/types"
)

type Instruction_PUSH_STR struct {
	value *types.PrimitiveStr
}

func (instruction *Instruction_PUSH_STR) Parse(args *string) error {
	if value, err := strconv.Unquote(*args); err == nil {
		instruction.value = types.MakeString(value)
	} else {
		if len(*args) == 0 {
			return fmt.Errorf("PUSH_STR instruction needs a value")
		}
		return fmt.Errorf("PUSH_STR instruction has invalid value %q", *args)
	}
	return nil
}

func (instruction *Instruction_PUSH_STR) String() string {
	return fmt.Sprintf("PUSH_STR %q", instruction.value.Display())
}

func (instruction *Instruction_PUSH_STR) Perform(vm *VirtualMachine) error {
	vm.stack.Push(instruction.value)
	return nil
}
