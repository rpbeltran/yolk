package vm

import (
	"fmt"
	"io"
	"os"
	"yolk/types"
	"yolk/utils"
)

type VirtualMachine struct {
	program             []Instruction
	instruction_pointer int
	stack               utils.Stack[types.Primitive] //todo: benchmark this being a pointer
	pipeline_states     utils.Stack[*types.Primitive]
	MockExecutions      bool
	stdout              io.Writer
}

func NewVM() VirtualMachine {
	return VirtualMachine{
		stdout:              os.Stdout,
		stack:               *utils.CreateStack[types.Primitive](),
		pipeline_states:     *utils.CreateStack[*types.Primitive](),
		instruction_pointer: 0,
	}
}

func NewDebugVM(stdout io.Writer) VirtualMachine {
	return VirtualMachine{
		stdout:              stdout,
		stack:               *utils.CreateStack[types.Primitive](),
		pipeline_states:     *utils.CreateStack[*types.Primitive](),
		instruction_pointer: 0,
		MockExecutions:      true,
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

func (vm *VirtualMachine) AddProgramInstruction(instruction Instruction) {
	vm.program = append(vm.program, instruction)
}

func (vm *VirtualMachine) ClearProgram() {
	vm.program = make([]Instruction, 0)
}

func (vm *VirtualMachine) RunProgram(verbose_debug bool) error {
	vm.instruction_pointer = 0
	for vm.instruction_pointer < len(vm.program) {
		instruction := vm.program[vm.instruction_pointer]
		if verbose_debug {
			fmt.Printf("-- executing %d: %v\n", vm.instruction_pointer, instruction)
		}
		if err := instruction.Perform(vm); err != nil {
			return fmt.Errorf("executing instruction %d: %w", vm.instruction_pointer, err)
		}
		if verbose_debug {
			dbg_state := vm.GetDebugState()
			fmt.Printf("   -- stack size: %d\n", dbg_state.StackSize)
			if dbg_state.StackSize != 0 {
				fmt.Printf("   -- top of stack: %q\n", dbg_state.TopOfStack)
			}
		}
		vm.instruction_pointer += 1
	}
	return nil
}
