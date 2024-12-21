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

func (instruction *Instruction_BINOP) Perform(vm *VirtualMachine) error {
	left, err := vm.stack.Pop()
	if err != nil {
		return fmt.Errorf("popping lhs for binop: %v", err)
	}
	right, err := vm.stack.Pop()
	if err != nil {
		return fmt.Errorf("popping rhs for binop: %v", err)
	}

	switch instruction.binop {
	case binop_add:
		if sum, err := left.Add(right); err != nil {
			return err
		} else {
			vm.stack.Push(sum)
		}
	default:
		return fmt.Errorf("BINOP instruction has binop code '%d'", instruction.binop)
	}
	return nil
}
