package vm

import (
	"fmt"
	"math/big"
	"yolk/types"
	"yolk/utils"
)

type Instruction_PUSH_NUM struct {
	value big.Rat
}

func (instruction *Instruction_PUSH_NUM) Parse(args *string) error {
	if _, success := instruction.value.SetString(*args); !success {
		if len(*args) == 0 {
			return fmt.Errorf("PUSH_NUM instruction needs a value")
		}
		return fmt.Errorf("PUSH_NUM instruction has invalid value %q", *args)
	}
	return nil
}

func (instruction *Instruction_PUSH_NUM) String() string {
	return fmt.Sprintf("PUSH_NUM %v", utils.EncodeNum(&instruction.value))
}

func (instruction *Instruction_PUSH_NUM) Perform(vm *VirtualMachine) error {
	vm.stack.Push(types.AsPrimitiveNumber(&instruction.value))
	return nil
}
