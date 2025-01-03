package vm

import (
	"fmt"
	"strings"
)

type Instruction interface {
	Parse(args *string) error
	String() string
	Perform(machine *VirtualMachine) error
}

func ParseInstruction(yolk_line string) (Instruction, error) {

	yolk_line = strings.TrimSpace(yolk_line)
	if len(yolk_line) == 0 || yolk_line[0] == '#' {
		return nil, nil
	}

	operator, args, _ := strings.Cut(yolk_line, " ")
	args = strings.TrimSpace(args)

	switch operator {
	case "ASSIGN_NAME":
		var instruction Instruction_ASSIGN_NAME
		if err := instruction.Parse(&args); err != nil {
			return nil, err
		}
		return &instruction, nil
	case "BINOP":
		var instruction Instruction_BINOP
		if err := instruction.Parse(&args); err != nil {
			return nil, err
		}
		return &instruction, nil
	case "DECLARE_NAME":
		var instruction Instruction_DECLARE_NAME
		if err := instruction.Parse(&args); err != nil {
			return nil, err
		}
		return &instruction, nil
	case "EXEC":
		var instruction Instruction_EXEC
		if err := instruction.Parse(&args); err != nil {
			return nil, err
		}
		return &instruction, nil
	case "JUMP":
		var instruction Instruction_JUMP
		if err := instruction.Parse(&args); err != nil {
			return nil, err
		}
		return &instruction, nil
	case "JUMP_IF_FALSE":
		var instruction Instruction_JUMP_IF_FALSE
		if err := instruction.Parse(&args); err != nil {
			return nil, err
		}
		return &instruction, nil
	case "JUMP_IF_TRUE":
		var instruction Instruction_JUMP_IF_TRUE
		if err := instruction.Parse(&args); err != nil {
			return nil, err
		}
		return &instruction, nil
	case ".LABEL":
		var instruction Instruction_LABEL
		if err := instruction.Parse(&args); err != nil {
			return nil, err
		}
		return &instruction, nil
	case "LOAD_NAME":
		var instruction Instruction_LOAD_NAME
		if err := instruction.Parse(&args); err != nil {
			return nil, err
		}
		return &instruction, nil
	case "PIPELINE":
		var instruction Instruction_PIPELINE
		if err := instruction.Parse(&args); err != nil {
			return nil, err
		}
		return &instruction, nil
	case "PUSH_BOOL":
		var instruction Instruction_PUSH_BOOL
		if err := instruction.Parse(&args); err != nil {
			return nil, err
		}
		return &instruction, nil
	case "PUSH_STR":
		var instruction Instruction_PUSH_STR
		if err := instruction.Parse(&args); err != nil {
			return nil, err
		}
		return &instruction, nil
	case "PUSH_NUM":
		var instruction Instruction_PUSH_NUM
		if err := instruction.Parse(&args); err != nil {
			return nil, err
		}
		return &instruction, nil
	case "PRINT":
		var instruction Instruction_PRINT
		if err := instruction.Parse(&args); err != nil {
			return nil, err
		}
		return &instruction, nil

	}

	return nil, fmt.Errorf("unknown operator: %s", operator)

}
