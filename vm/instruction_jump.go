package vm

import (
	"fmt"
	"strconv"
)

type Instruction_JUMP struct {
	destination uint64
}

func (instruction *Instruction_JUMP) Parse(args *string) error {
	if destination, err := strconv.ParseUint(*args, 10, 64); err == nil {
		instruction.destination = destination
	} else {
		if len(*args) == 0 {
			return fmt.Errorf("JUMP instruction needs a destination")
		}
		return fmt.Errorf("JUMP instruction has invalid destination %q", *args)
	}
	return nil
}

func (instruction *Instruction_JUMP) String() string {
	return fmt.Sprintf("JUMP %d", instruction.destination)
}

func (instruction *Instruction_JUMP) Perform(vm *VirtualMachine) error {
	if new_ipointer, ok := vm.labels[instruction.destination]; ok {
		vm.instruction_pointer = new_ipointer
		return nil
	}
	return fmt.Errorf("JUMP to unknown label: %d", instruction.destination)
}
