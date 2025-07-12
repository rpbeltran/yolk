package vm

import (
	"errors"
	"fmt"
)

type Instruction_FINISHED_INTERACTIVE_COMMAND struct{}

func (instruction *Instruction_FINISHED_INTERACTIVE_COMMAND) Parse(args *string) error {
	if len(*args) != 0 {
		return fmt.Errorf("FINISHED_INTERACTIVE_COMMAND instruction expected no arguments, received '%s'", *args)
	}
	return nil
}

func (instruction *Instruction_FINISHED_INTERACTIVE_COMMAND) String() string {
	return "FINISHED_INTERACTIVE_COMMAND"
}

func (instruction *Instruction_FINISHED_INTERACTIVE_COMMAND) Perform(vm *VirtualMachine) error {
	if !vm.stack.Empty() {
		return errors.New("stack should be empty after interactive command finished executing but was not")
	}
	fmt.Fprintln(vm.stdout, "$$$###@@@::FINISHED_INTERACTIVE_COMMAND::1bc82ncv3yur::@@@###$$$")
	return nil
}
