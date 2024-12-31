package vm

import (
	"fmt"
	"strconv"
)

type Instruction_LABEL struct {
	address uint64
}

func (instruction *Instruction_LABEL) Parse(args *string) error {
	if address, err := strconv.ParseUint(*args, 10, 64); err == nil {
		instruction.address = address
	} else {
		if len(*args) == 0 {
			return fmt.Errorf(".LABEL instruction needs an address")
		}
		return fmt.Errorf(".LABEL instruction has invalid address %q", *args)
	}
	return nil
}

func (instruction *Instruction_LABEL) String() string {
	return fmt.Sprintf(".LABEL %d", instruction.address)
}

func (instruction *Instruction_LABEL) Perform(vm *VirtualMachine) error {
	return nil
}

func (instruction *Instruction_LABEL) AddTo(vm *VirtualMachine) error {
	vm.labels[instruction.address] = len(vm.program)
	return nil
}
