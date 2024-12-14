package instructions

import (
	"fmt"
	"strconv"
	"yolk/vm"
)

type Instruction_PUSH_STR struct {
	value string
}

func (instruction *Instruction_PUSH_STR) Parse(args *string) error {
	if value, err := strconv.Unquote(*args); err == nil {
		instruction.value = value
	} else {
		if len(*args) == 0 {
			return fmt.Errorf("PUSH_STR instruction needs a value")
		}
		return fmt.Errorf("PUSH_STR instruction has invalid value %q", *args)
	}
	return nil
}

func (instruction *Instruction_PUSH_STR) String() string {
	return fmt.Sprintf("PUSH_STR %q", instruction.value)
}

func (instruction *Instruction_PUSH_STR) Perform(machine *vm.VirtualMachine) error {
	fmt.Printf("//TODO: Implement PUSH_STR (called with value %q)\n", instruction.value)
	return nil
}
