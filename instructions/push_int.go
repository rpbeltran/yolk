package instructions

import (
	"fmt"
	"strconv"
	"yolk/vm"
)

type Instruction_PUSH_INT struct {
	value int64
}

func (instruction *Instruction_PUSH_INT) Parse(args *string) error {
	if value, err := strconv.ParseInt(*args, 10, 64); err == nil {
		instruction.value = value
	} else {
		return err
	}
	return nil
}

func (instruction *Instruction_PUSH_INT) Perform(machine *vm.VirtualMachine) error {
	fmt.Printf("//TODO: Implement PUSH_INT (called with value %d)\n", instruction.value)
	return nil
}
