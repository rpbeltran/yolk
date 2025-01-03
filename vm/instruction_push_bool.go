package vm

import (
	"fmt"
	"yolk/types"
)

type Instruction_PUSH_BOOL struct {
	value *types.PrimitiveBool
}

func (instruction *Instruction_PUSH_BOOL) Parse(args *string) error {
	switch *args {
	case "true":
		instruction.value = types.MakeBool(true)
		return nil
	case "false":
		instruction.value = types.MakeBool(false)
		return nil
	case "":
		return fmt.Errorf("PUSH_BOOL instruction needs a value")
	default:
		return fmt.Errorf("PUSH_BOOL instruction has invalid value %q", *args)
	}
}

func (instruction *Instruction_PUSH_BOOL) String() string {
	return fmt.Sprintf("PUSH_BOOL %s", instruction.value.Display())
}

func (instruction *Instruction_PUSH_BOOL) Perform(vm *VirtualMachine) error {
	vm.stack.Push(instruction.value)
	return nil
}
