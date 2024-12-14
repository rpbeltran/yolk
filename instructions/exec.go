package instructions

import (
	"fmt"
	"strconv"
	"yolk/vm"
)

type Instruction_EXEC struct {
	arg_count uint64
}

func (instruction *Instruction_EXEC) Parse(args *string) error {
	if arg_count, err := strconv.ParseUint(*args, 10, 8); err == nil {
		instruction.arg_count = arg_count
	} else {
		return err
	}
	return nil
}

func (instruction *Instruction_EXEC) Perform(machine *vm.VirtualMachine) error {
	fmt.Printf("//TODO: Implement EXEC (called with %d inputs)\n", instruction.arg_count)
	return nil
}
