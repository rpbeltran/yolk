package vm

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type inplace_binop uint8

const (
	binop_add_inplace        inplace_binop = iota
	binop_subtract_inplace   inplace_binop = iota
	binop_multiply_inplace   inplace_binop = iota
	binop_divide_inplace     inplace_binop = iota
	binop_int_divide_inplace inplace_binop = iota
	binop_power_inplace      inplace_binop = iota
	binop_modulus_inplace    inplace_binop = iota
	binop_concat_inplace     inplace_binop = iota
	binop_and_inplace        inplace_binop = iota
	binop_or_inplace         inplace_binop = iota
)

type Instruction_BINOP_INPLACE struct {
	operation inplace_binop
	name      string
}

func (instruction *Instruction_BINOP_INPLACE) Parse(args *string) error {
	if len(*args) == 0 {
		return errors.New("BINOP_INPLACE instruction needs operator and name, neither provided")
	}

	operator, name, found := strings.Cut(*args, " ")
	if !found {
		return errors.New("BINOP_INPLACE instruction needs operator and name, only got one arg")
	}

	switch operator {
	case "add":
		instruction.operation = binop_add_inplace
	case "subtract":
		instruction.operation = binop_subtract_inplace
	case "multiply":
		instruction.operation = binop_multiply_inplace
	case "divide":
		instruction.operation = binop_divide_inplace
	case "int_divide":
		instruction.operation = binop_int_divide_inplace
	case "power":
		instruction.operation = binop_power_inplace
	case "modulus":
		instruction.operation = binop_modulus_inplace
	case "concat":
		instruction.operation = binop_concat_inplace
	case "and":
		instruction.operation = binop_and_inplace
	case "or":
		instruction.operation = binop_or_inplace
	default:
		return fmt.Errorf("BINOP_INPLACE instruction specifies unexpected operator %q", operator)
	}

	if name_unquoted, err := strconv.Unquote(name); err != nil {
		return fmt.Errorf("BINOP_INPLACE instruction has invalid name %q (needs quotes)", name)
	} else if len(name_unquoted) == 0 {
		return fmt.Errorf("BINOP_INPLACE instruction has invalid name %q", name)
	} else {
		instruction.name = name_unquoted
	}
	return nil
}

func (instruction *Instruction_BINOP_INPLACE) String() string {
	switch instruction.operation {
	case binop_add_inplace:
		return fmt.Sprintf("BINOP_INPLACE add %q", instruction.name)
	case binop_subtract_inplace:
		return fmt.Sprintf("BINOP_INPLACE subtract %q", instruction.name)
	case binop_multiply_inplace:
		return fmt.Sprintf("BINOP_INPLACE multiply %q", instruction.name)
	case binop_divide_inplace:
		return fmt.Sprintf("BINOP_INPLACE divide %q", instruction.name)
	case binop_int_divide_inplace:
		return fmt.Sprintf("BINOP_INPLACE int_divide %q", instruction.name)
	case binop_power_inplace:
		return fmt.Sprintf("BINOP_INPLACE power %q", instruction.name)
	case binop_modulus_inplace:
		return fmt.Sprintf("BINOP_INPLACE modulus %q", instruction.name)
	case binop_concat_inplace:
		return fmt.Sprintf("BINOP_INPLACE concat %q", instruction.name)
	case binop_and_inplace:
		return fmt.Sprintf("BINOP_INPLACE and %q", instruction.name)
	case binop_or_inplace:
		return fmt.Sprintf("BINOP_INPLACE or %q", instruction.name)
	default:
		panic(fmt.Sprintf("Unimplemented BINOP_INPLACE serialization for mode %d", instruction.operation))
	}
}

func (instruction *Instruction_BINOP_INPLACE) Perform(vm *VirtualMachine) error {

	lhs, err := vm.FetchVariable(instruction.name)
	if err != nil {
		return fmt.Errorf("getting variable for lhs: %v", err)
	}

	right, err := vm.stack.Pop()
	if err != nil {
		return fmt.Errorf("popping rhs for binop: %v", err)
	}

	switch instruction.operation {
	case binop_add_inplace:
		if err := lhs.AddInplace(right); err != nil {
			return err
		}
	case binop_subtract_inplace:
		if err := lhs.SubtractInplace(right); err != nil {
			return err
		}
	case binop_multiply_inplace:
		if err := lhs.MultiplyInplace(right); err != nil {
			return err
		}
	case binop_divide_inplace:
		if err := lhs.DivideInplace(right); err != nil {
			return err
		}
	case binop_int_divide_inplace:
		if err := lhs.IntDivideInplace(right); err != nil {
			return err
		}
	case binop_power_inplace:
		if err := lhs.RaisePowerInplace(right); err != nil {
			return err
		}
	case binop_modulus_inplace:
		if err := lhs.ModuloInplace(right); err != nil {
			return err
		}
	case binop_concat_inplace:
		if err := lhs.ConcatenateInPlace(right); err != nil {
			return err
		}
	case binop_and_inplace:
		if err := lhs.AndInplace(right); err != nil {
			return err
		}
	case binop_or_inplace:
		if err := lhs.OrInplace(right); err != nil {
			return err
		}
	default:
		return fmt.Errorf("BINO_INPLACE instruction has invalid operation code '%d'", instruction.operation)
	}
	return nil
}
