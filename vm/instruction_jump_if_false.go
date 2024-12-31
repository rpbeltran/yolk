package vm

import (
	"fmt"
	"strconv"
)

type Instruction_JUMP_IF_FALSE struct {
	destination uint64
}

func (instruction *Instruction_JUMP_IF_FALSE) Parse(args *string) error {
	if destination, err := strconv.ParseUint(*args, 10, 64); err == nil {
		instruction.destination = destination
	} else {
		if len(*args) == 0 {
			return fmt.Errorf("JUMP_IF_FALSE instruction needs a destination")
		}
		return fmt.Errorf("JUMP_IF_FALSE instruction has invalid destination %q", *args)
	}
	return nil
}

func (instruction *Instruction_JUMP_IF_FALSE) String() string {
	return fmt.Sprintf("JUMP_IF_FALSE %d", instruction.destination)
}

func (instruction *Instruction_JUMP_IF_FALSE) Perform(vm *VirtualMachine) error {
	if value, err := vm.stack.Pop(); err != nil {
		return fmt.Errorf("error to getting popping condition for JUMP_IF_FALSE: %w", err)
	} else if as_bool, err := value.RequireBool(); err != nil {
		return fmt.Errorf("condition for JUMP_IF_FALSE must be a bool, but top of stack had: %v", value)
	} else if as_bool.Truthy() {
		return nil
	} else if new_ipointer, ok := vm.labels[instruction.destination]; ok {
		vm.instruction_pointer = new_ipointer
		return nil
	} else {
		return fmt.Errorf("JUMP_IF_FALSE to unknown label: %d", instruction.destination)
	}
}
