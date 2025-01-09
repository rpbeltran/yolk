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
	if len(yolk_line) == 0 || yolk_line[0] == '#' {
		return nil, nil
	}

	operator, args, _ := strings.Cut(yolk_line, " ")
	args = strings.TrimSpace(args)

	switch operator {
	case "ASSIGN":
		var instruction Instruction_ASSIGN
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
	case "BINOP_INPLACE":
		var instruction Instruction_BINOP_INPLACE
		if err := instruction.Parse(&args); err != nil {
			return nil, err
		}
		return &instruction, nil
	case "COMPARE":
		var instruction Instruction_COMPARE
		if err := instruction.Parse(&args); err != nil {
			return nil, err
		}
		return &instruction, nil
	case "DECLARE":
		var instruction Instruction_DECLARE
		if err := instruction.Parse(&args); err != nil {
			return nil, err
		}
		return &instruction, nil
	case ".DEFINE":
		var instruction Instruction_DEFINE
		if err := instruction.Parse(&args); err != nil {
			return nil, err
		}
		return &instruction, nil
	case ".DEFINE_END":
		var instruction Instruction_DEFINE_END
		if err := instruction.Parse(&args); err != nil {
			return nil, err
		}
		return &instruction, nil
	case "DUPLICATE":
		var instruction Instruction_DUPLICATE
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
	case "LOAD":
		var instruction Instruction_LOAD
		if err := instruction.Parse(&args); err != nil {
			return nil, err
		}
		return &instruction, nil
	case "NEGATE":
		var instruction Instruction_NEGATE
		if err := instruction.Parse(&args); err != nil {
			return nil, err
		}
		return &instruction, nil
	case "NOT":
		var instruction Instruction_NOT
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
	case "PUSH_INT":
		var instruction Instruction_PUSH_INT
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
	case "PUSH_STR":
		var instruction Instruction_PUSH_STR
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
