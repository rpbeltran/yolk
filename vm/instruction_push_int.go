package vm

import (
	"fmt"
	"strconv"
)

type Instruction_PUSH_INT struct {
	value int64
}

func (instruction *Instruction_PUSH_INT) Parse(args *string) error {
	if value, err := strconv.ParseInt(*args, 10, 64); err == nil {
		instruction.value = value
	} else {
		if len(*args) == 0 {
			return fmt.Errorf("PUSH_INT instruction needs a value")
		}
		return fmt.Errorf("PUSH_INT instruction has invalid value %q", *args)
	}
	return nil
}

func (instruction *Instruction_PUSH_INT) String() string {
	return fmt.Sprintf("PUSH_INT %d", instruction.value)
}

func (instruction *Instruction_PUSH_INT) Perform(machine *VirtualMachine) error {
	fmt.Printf("//TODO: Implement PUSH_INT (called with value %d)\n", instruction.value)
	return nil
}
