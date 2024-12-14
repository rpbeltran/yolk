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
	if arg_count, err := strconv.ParseUint(*args, 10, 64); err == nil {
		instruction.arg_count = arg_count
	} else {
		if len(*args) == 0 {
			return fmt.Errorf("EXEC instruction needs argument count")
		}
		return fmt.Errorf("EXEC instruction has invalid argument count %q", *args)
	}
	return nil
}

func (instruction *Instruction_EXEC) String() string {
	return fmt.Sprintf("EXEC %d", instruction.arg_count)
}

func (instruction *Instruction_EXEC) Perform(machine *vm.VirtualMachine) error {
	fmt.Printf("//TODO: Implement EXEC (called with %d inputs)\n", instruction.arg_count)
	return nil
}
