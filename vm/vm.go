package vm

import (
	"fmt"
	"io"
	"os"
	"yolk/types"
	"yolk/utils"
)

type VirtualMachine struct {
	instructions        []Instruction
	instruction_pointer int
	stack               utils.Stack[types.Primitive] //todo: benchmark this being a pointer
	pipeline_states     utils.Stack[*types.Primitive]
	MockExecutions      bool
	stdout              io.Writer
}

func NewVM() VirtualMachine {
	return VirtualMachine{
		stdout: os.Stdout,
		stack:  *utils.CreateStack[types.Primitive](),
	}
}

func NewDebugVM(stdout io.Writer) VirtualMachine {
	return VirtualMachine{
		stdout: stdout,
		stack:  *utils.CreateStack[types.Primitive](),
	}
}

func (vm *VirtualMachine) GetPipeIn() (*types.Primitive, bool) {
	if vm.pipeline_states.Empty() {
		return nil, false
	}
	if value, err := vm.pipeline_states.Peek(); err != nil {
		panic(fmt.Sprintf("Failed to get pipeline state: %q", err))
	} else {
		if stdin := *value; stdin == nil {
			return nil, false
		} else {
			return *value, true
		}
	}
}

type VirtualMachineDebugState struct {
	StackSize  uint
	TopOfStack string
}

func (vm *VirtualMachine) GetDebugState() VirtualMachineDebugState {
	top_of_stack, err := vm.stack.Peek()
	if err != nil {
		return VirtualMachineDebugState{
			StackSize:  uint(vm.stack.Size()),
			TopOfStack: "",
		}
	}
	return VirtualMachineDebugState{
		StackSize:  uint(vm.stack.Size()),
		TopOfStack: (*top_of_stack).Display(),
	}
}
