package vm

import (
	"errors"
	"fmt"
	"strconv"
	"yolk/types"
)

type Instruction_EXEC struct {
	arg_count uint64
}

func (instruction *Instruction_EXEC) Parse(args *string) error {
	if arg_count, err := strconv.ParseUint(*args, 10, 64); err == nil {
		instruction.arg_count = arg_count
	} else {
		if len(*args) == 0 {
			return fmt.Errorf("EXEC instruction needs argument count")
		}
		return fmt.Errorf("EXEC instruction has invalid argument count %q", *args)
	}
	return nil
}

func (instruction *Instruction_EXEC) String() string {
	return fmt.Sprintf("EXEC %d", instruction.arg_count)
}

func (instruction *Instruction_EXEC) mockPerform(vm *VirtualMachine) error {
	cmd, err := vm.stack.Pop()
	if err != nil {
		return fmt.Errorf("popping command from stack: %v", err)
	}
	cmd_str, err := cmd.RequireStr()
	if err != nil {
		return fmt.Errorf("command must be a string: %v", err)
	}

	args := []string{}
	for i := range instruction.arg_count {
		if arg, err := vm.stack.Pop(); err != nil {
			return fmt.Errorf("popping arg %d from stack: %v", i, err)
		} else {
			args = append(args, arg.Display())
		}
	}

	var json_str string
	if stdin, has_stdin := vm.GetPipeIn(); !has_stdin {
		json_str = fmt.Sprintf(`{\n\t"command":%q,\n\t"args":%v\n}`, cmd_str.Display(), args)
	} else if stdin == nil {
		return errors.New("vm.GetPipeIn() gave a null ptr")
	} else {
		json_str = fmt.Sprintf(`{\n\t"command":%q,\n\t"args":%v,\n\t"stdin":%q\n}`, cmd_str.Display(), args, (*stdin).Display())
	}
	json := types.MakeString(json_str)
	vm.stack.Push(json)
	return nil
}

func (instruction *Instruction_EXEC) Perform(vm *VirtualMachine) error {
	if vm.MockExecutions {
		return instruction.mockPerform(vm)
	} else {
		return errors.New("nonmocked executions are currently unimplemented")
	}
}
