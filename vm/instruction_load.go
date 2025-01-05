package vm

import (
	"errors"
	"fmt"
	"yolk/utils"
)

type Instruction_LOAD struct {
	name string
}

var ErrParsingLOAD = errors.New("failed to parse LOAD")

func (instruction *Instruction_LOAD) Parse(args *string) error {
	if name, err := utils.DeserializeName(*args); err != nil {
		return fmt.Errorf("%w: bad name arg: %w", ErrParsingLOAD, err)
	} else {
		instruction.name = name
		return nil
	}
}

func (instruction *Instruction_LOAD) String() string {
	return fmt.Sprintf("LOAD %s", utils.SerializeName(instruction.name))
}

func (instruction *Instruction_LOAD) Perform(vm *VirtualMachine) error {
	if value, err := vm.FetchVariable(instruction.name); err != nil {
		return err
	} else {
		vm.stack.Push(value)
	}
	return nil
}
