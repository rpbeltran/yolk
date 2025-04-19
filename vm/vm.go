package vm

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"yolk/memory"
	"yolk/types"
	"yolk/utils"
)

type VirtualMachine struct {
	// Instructions
	program              []Instruction
	functions            map[string][]Instruction
	function_definitions map[string]Instruction_DEFINE

	labels              map[uint64]int
	instruction_pointer int

	// Runtime data
	stack           utils.Stack[types.Primitive]
	pipeline_states utils.Stack[*types.Primitive]
	memory          memory.Memory

	// Configuration
	stdout         io.Writer
	MockExecutions bool
}

func NewVM() VirtualMachine {
	return VirtualMachine{
		// Instructions
		program:              make([]Instruction, 0),
		functions:            make(map[string][]Instruction),
		function_definitions: make(map[string]Instruction_DEFINE),
		labels:               make(map[uint64]int),
		instruction_pointer:  0,

		// Runtime Data
		stack:           *utils.CreateStack[types.Primitive](),
		pipeline_states: *utils.CreateStack[*types.Primitive](),
		memory:          memory.NewVMMemory(),

		// Configuration
		stdout: os.Stdout,
	}
}

func NewDebugVM(stdout io.Writer) VirtualMachine {
	return VirtualMachine{
		// Instructions
		program:              make([]Instruction, 0),
		functions:            make(map[string][]Instruction),
		function_definitions: make(map[string]Instruction_DEFINE),
		labels:               make(map[uint64]int),
		instruction_pointer:  0,

		// Runtime Data
		stack:           *utils.CreateStack[types.Primitive](),
		pipeline_states: *utils.CreateStack[*types.Primitive](),
		memory:          memory.NewVMMemory(),

		// Configuration
		stdout:         stdout,
		MockExecutions: true,
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

func (vm *VirtualMachine) PutProgramInVM(scanner *bufio.Scanner) error {

	var function_name string
	var function_buffer []Instruction
	has_function_definition := false

	line_num := 0
	for scanner.Scan() {
		line_num += 1
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, ".LABEL") {
			_, args, _ := strings.Cut(line, " ")
			args = strings.TrimSpace(args)
			var instruction Instruction_LABEL
			if err := instruction.Parse(&args); err != nil {
				return err
			}
			vm.labels[instruction.address] = len(vm.program)
			vm.program = append(vm.program, &instruction)
		} else if strings.HasPrefix(line, ".DEFINE_END") {
			has_function_definition = false
			vm.functions[function_name] = function_buffer
			function_buffer = []Instruction{}
		} else if strings.HasPrefix(line, ".DEFINE") {
			if has_function_definition {
				return fmt.Errorf("parsing line %d %q: cannot nest function definitions", line_num, line)
			}
			_, args, _ := strings.Cut(line, " ")
			args = strings.TrimSpace(args)
			var function_definition Instruction_DEFINE
			if err := function_definition.Parse(&args); err != nil {
				return err
			}
			has_function_definition = true
			function_name = function_definition.name
			vm.function_definitions[function_name] = function_definition
		} else if instruction, err := ParseInstruction(line); err != nil {
			return fmt.Errorf("parsing line %d %q: %w", line_num, line, err)
		} else if instruction != nil {
			if has_function_definition {
				function_buffer = append(function_buffer, instruction)
			} else {
				vm.program = append(vm.program, instruction)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanning yolk: %v", err)
	}
	return nil
}
