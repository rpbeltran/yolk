package vm

import (
	"errors"
	"fmt"
	"yolk/utils"
)

type Instruction_ASSIGN struct {
	name string
}

var ErrParsingASSIGN = errors.New("failed to parse ASSIGN")

func (instruction *Instruction_ASSIGN) Parse(args *string) error {
	if name, err := utils.DeserializeName(*args); err != nil {
		return fmt.Errorf("%w: bad name arg: %w", ErrParsingASSIGN, err)
	} else {
		instruction.name = name
		return nil
	}
}

func (instruction *Instruction_ASSIGN) String() string {
	return fmt.Sprintf("ASSIGN %s", utils.SerializeName(instruction.name))
}

func (instruction *Instruction_ASSIGN) Perform(vm *VirtualMachine) error {
	value, err := vm.stack.Pop()
	if err != nil {
		return fmt.Errorf("unexpected error popping value for assignment: %w", err)
	}

	if err := vm.UpdateVariable(instruction.name, value); err != nil {
		return err
	}
	return nil
}
