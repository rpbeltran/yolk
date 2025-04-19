package vm

import (
	"fmt"
)

type Instruction_COMPARE_CHAIN struct {
	mode uint8
}

func (instruction *Instruction_COMPARE_CHAIN) Parse(args *string) error {
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
			return fmt.Errorf("COMPARE_CHAIN instruction needs a test mode, none provided")
		}
		return fmt.Errorf("COMPARE_CHAIN instruction specifies unexpected test mode %q", *args)
	}
	return nil
}

func (instruction *Instruction_COMPARE_CHAIN) String() string {
	switch instruction.mode {
	case comparison_equal:
		return "COMPARE_CHAIN equal"
	case comparison_unequal:
		return "COMPARE_CHAIN unequal"
	case comparison_less:
		return "COMPARE_CHAIN less"
	case comparison_lte:
		return "COMPARE_CHAIN lte"
	case comparison_greater:
		return "COMPARE_CHAIN greater"
	case comparison_gte:
		return "COMPARE_CHAIN gte"
	default:
		panic(fmt.Sprintf("Unimplemented COMPARE_CHAIN serialization for test mode %d", instruction.mode))
	}
}

func (instruction *Instruction_COMPARE_CHAIN) Perform(vm *VirtualMachine) error {
	right, err := vm.stack.Pop()
	if err != nil {
		return fmt.Errorf("popping rhs for COMPARE_CHAIN: %v", err)
	}

	left, err := vm.stack.Pop()
	if err != nil {
		return fmt.Errorf("popping lhs for COMPARE_CHAIN: %v", err)
	}

	push_compare(vm, instruction.mode, left, right)
	vm.stack.Push(right)

	return nil
}
