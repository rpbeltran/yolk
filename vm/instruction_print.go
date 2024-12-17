package vm

import (
	"fmt"
)

type Instruction_PRINT struct{}

func (instruction *Instruction_PRINT) Parse(args *string) error {
	if len(*args) != 0 {
		return fmt.Errorf("PRINT instruction expected no arguments, received '%s'", *args)
	}
	return nil
}

func (instruction *Instruction_PRINT) String() string {
	return "PRINT"
}

func (instruction *Instruction_PRINT) Perform(vm *VirtualMachine) error {
	if value, err := vm.stack.Pop(); err != nil {
		return err
	} else {
		vm.output_buffer.Push(value.Display())
	}
	return nil
}
