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
	new_value_id := vm.memory.StorePrimitive(value)
	if err := vm.memory.RebindVariable(instruction.name, new_value_id); err != nil {
		return err
	}
	return nil
}
