package vm

import (
	"fmt"
)

type binop uint8

const (
	binop_add binop = iota
)

type Instruction_BINOP struct {
	binop binop
}

func (instruction *Instruction_BINOP) Parse(args *string) error {
	switch *args {
	case "add":
		instruction.binop = binop_add
	default:
		if len(*args) == 0 {
			return fmt.Errorf("BINOP instruction needs operator, none provided")
		}
		return fmt.Errorf("BINOP instruction specifies unexpected operator %q", *args)
	}
	return nil
}

func (instruction *Instruction_BINOP) String() string {
	switch instruction.binop {
	case binop_add:
		return "BINOP add"
	default:
		panic(fmt.Sprintf("Unimplemented BINOP serialization for mode %d", instruction.binop))
	}
}

func (instruction *Instruction_BINOP) Perform(machine *VirtualMachine) error {
	switch instruction.binop {
	case binop_add:
		fmt.Println("//TODO: implement `BINOP add`")
	default:
		return fmt.Errorf("BINOP instruction has binop code '%d'", instruction.binop)
	}
	return nil
}
