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

func (instruction *Instruction_PRINT) Perform(machine *VirtualMachine) error {
	fmt.Println("//TODO: Implement PRINT instruction")
	return nil
}
