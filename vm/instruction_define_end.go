package vm

import (
	"errors"
	"fmt"
)

var ErrDefineEndParsing = errors.New("failed to parse .DEFINE_END")
var ErrDefineEndPerform = errors.New(".DEFINE_END instruction is never supposed to be executed")

type Instruction_DEFINE_END struct{}

func (instruction *Instruction_DEFINE_END) Parse(args *string) error {
	if len(*args) != 0 {
		return fmt.Errorf("%w: expected no arguments, received %q", ErrDefineEndParsing, *args)
	}
	return nil
}

func (instruction *Instruction_DEFINE_END) String() string {
	return ".DEFINE_END"
}

func (instruction *Instruction_DEFINE_END) Perform(vm *VirtualMachine) error {
	return ErrDefineEndPerform
}
