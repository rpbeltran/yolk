package vm

import (
	"fmt"
)

type binop uint8

const (
	binop_add        binop = iota
	binop_subtract   binop = iota
	binop_multiply   binop = iota
	binop_divide     binop = iota
	binop_int_divide binop = iota
	binop_power      binop = iota
	binop_modulus    binop = iota
)

type Instruction_BINOP struct {
	binop binop
}

func (instruction *Instruction_BINOP) Parse(args *string) error {
	switch *args {
	case "add":
		instruction.binop = binop_add
	case "subtract":
		instruction.binop = binop_subtract
	case "multiply":
		instruction.binop = binop_multiply
	case "divide":
		instruction.binop = binop_divide
	case "int_divide":
		instruction.binop = binop_int_divide
	case "power":
		instruction.binop = binop_power
	case "modulus":
		instruction.binop = binop_modulus
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
	case binop_subtract:
		return "BINOP subtract"
	case binop_multiply:
		return "BINOP multiply"
	case binop_divide:
		return "BINOP divide"
	case binop_int_divide:
		return "BINOP int_divide"
	case binop_power:
		return "BINOP power"
	case binop_modulus:
		return "BINOP modulus"
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
	case binop_subtract:
		if difference, err := left.Subtract(right); err != nil {
			return err
		} else {
			vm.stack.Push(difference)
		}
	case binop_multiply:
		if product, err := left.Multiply(right); err != nil {
			return err
		} else {
			vm.stack.Push(product)
		}
	case binop_divide:
		if quotient, err := left.Divide(right); err != nil {
			return err
		} else {
			vm.stack.Push(quotient)
		}
	case binop_int_divide:
		if int_quotient, err := left.IntDivide(right); err != nil {
			return err
		} else {
			vm.stack.Push(int_quotient)
		}
	case binop_power:
		if power, err := left.RaisePower(right); err != nil {
			return err
		} else {
			vm.stack.Push(power)
		}
	case binop_modulus:
		if mod, err := left.Modulo(right); err != nil {
			return err
		} else {
			vm.stack.Push(mod)
		}
	default:
		return fmt.Errorf("BINOP instruction has binop code '%d'", instruction.binop)
	}
	return nil
}
