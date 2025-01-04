package vm

import (
	"fmt"
	"yolk/types"
)

const (
	comparison_equal   uint8 = iota
	comparison_unequal uint8 = iota
	comparison_less    uint8 = iota
	comparison_lte     uint8 = iota
	comparison_greater uint8 = iota
	comparison_gte     uint8 = iota
)

type Instruction_COMPARE struct {
	mode uint8
}

func (instruction *Instruction_COMPARE) Parse(args *string) error {
	switch *args {
	case "equal":
		instruction.mode = comparison_equal
	case "unequal":
		instruction.mode = comparison_unequal
	case "less":
		instruction.mode = comparison_less
	case "lte":
		instruction.mode = comparison_lte
	case "greater":
		instruction.mode = comparison_greater
	case "gte":
		instruction.mode = comparison_gte
	default:
		if len(*args) == 0 {
			return fmt.Errorf("COMPARE instruction needs a test mode, none provided")
		}
		return fmt.Errorf("COMPARE instruction specifies unexpected test mode %q", *args)
	}
	return nil
}

func (instruction *Instruction_COMPARE) String() string {
	switch instruction.mode {
	case comparison_equal:
		return "COMPARE equal"
	case comparison_unequal:
		return "COMPARE unequal"
	case comparison_less:
		return "COMPARE less"
	case comparison_lte:
		return "COMPARE lte"
	case comparison_greater:
		return "COMPARE greater"
	case comparison_gte:
		return "COMPARE gte"
	default:
		panic(fmt.Sprintf("Unimplemented COMPARE serialization for test mode %d", instruction.mode))
	}
}

func (instruction *Instruction_COMPARE) Perform(vm *VirtualMachine) error {
	left, err := vm.stack.Pop()
	if err != nil {
		return fmt.Errorf("popping lhs for COMPARE: %v", err)
	}

	right, err := vm.stack.Pop()
	if err != nil {
		return fmt.Errorf("popping rhs for COMPARE: %v", err)
	}

	switch instruction.mode {
	case comparison_equal:
		vm.stack.Push(types.MakeBool(left.Equal(right)))
	case comparison_unequal:
		vm.stack.Push(types.MakeBool(!left.Equal(right)))
	case comparison_less:
		if lt, err := left.LessThan(right); err != nil {
			return fmt.Errorf("computing \"less than\": %v", err)
		} else {
			vm.stack.Push(types.MakeBool(lt))
		}
	case comparison_lte:
		if left.Equal(right) {
			vm.stack.Push(types.MakeBool(true))
		} else if lt, err := left.LessThan(right); err != nil {
			return fmt.Errorf("computing \"less than or equal to\": %v", err)
		} else {
			vm.stack.Push(types.MakeBool(lt))
		}
	case comparison_greater:
		if left.Equal(right) {
			vm.stack.Push(types.MakeBool(false))
		} else if lt, err := left.LessThan(right); err != nil {
			return fmt.Errorf("computing \"greater than\": %v", err)
		} else {
			vm.stack.Push(types.MakeBool(!lt))
		}
	case comparison_gte:
		if left.Equal(right) {
			vm.stack.Push(types.MakeBool(true))
		} else if lt, err := left.LessThan(right); err != nil {
			return fmt.Errorf("computing \"greater than\": %v", err)
		} else {
			vm.stack.Push(types.MakeBool(!lt))
		}
	default:
		panic(fmt.Sprintf("Unimplemented COMPARE for test mode %d", instruction.mode))
	}

	return nil
}
