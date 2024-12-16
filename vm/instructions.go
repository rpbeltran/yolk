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
	case "EXEC":
		var instruction Instruction_EXEC
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
	case "PUSH_STR":
		var instruction Instruction_PUSH_STR
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
	case "PRINT":
		var instruction Instruction_PRINT
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
	}

	return nil, fmt.Errorf("Unknown operator: %s", operator)

}
