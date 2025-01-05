package vm

import (
	"errors"
	"fmt"
	"yolk/utils"
)

type Instruction_DECLARE struct {
	name string
}

var ErrParsingDECLARE = errors.New("failed to parse DECLARE")

func (instruction *Instruction_DECLARE) Parse(args *string) error {
	if name, err := utils.DeserializeName(*args); err != nil {
		return fmt.Errorf("%w: bad name arg: %w", ErrParsingDECLARE, err)
	} else {
		instruction.name = name
		return nil
	}
}

func (instruction *Instruction_DECLARE) String() string {
	return fmt.Sprintf("DECLARE %s", utils.SerializeName(instruction.name))
}

func (instruction *Instruction_DECLARE) Perform(vm *VirtualMachine) error {
	value, err := vm.stack.Pop()
	if err != nil {
		return fmt.Errorf("unexpected error popping value for assignment: %w", err)
	}

	if err := vm.StoreNewVariable(instruction.name, value); err != nil {
		return err
	}
	return nil
}
