package instructions

import (
	"fmt"
	"yolk/vm"
)

type Binop uint8

const (
	BINOP_ADD Binop = iota
)

type Instruction_BIN_OP struct {
	binop Binop
}

func (instruction *Instruction_BIN_OP) Parse(args *string) error {
	switch *args {
	case "add":
		instruction.binop = BINOP_ADD
	default:
		return fmt.Errorf("BINOP instruction specifies unexpected operator '%s'", *args)
	}
	return nil
}

func (instruction *Instruction_BIN_OP) Perform(machine *vm.VirtualMachine) error {
	switch instruction.binop {
	case BINOP_ADD:
		fmt.Println("//TODO: implement `BINOP add`")
	default:
		return fmt.Errorf("BINOP instruction has binop code '%d'", instruction.binop)
	}
	return nil
}
