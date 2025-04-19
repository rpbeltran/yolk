package vm

import (
	"errors"
	"fmt"
	"strings"
	"yolk/utils"
)

type Instruction_DECLARE struct {
	name                string
	type_annotation     string
	has_type_annotation bool
}

var ErrDeclareParsing = errors.New("failed to parse DECLARE")
var ErrDeclareParsingName = fmt.Errorf("%w: invalid name arg", ErrDeclareParsing)
var ErrDeclareParsingType = fmt.Errorf("%w: invalid type arg", ErrDeclareParsing)

var ErrDeclarePerform = errors.New("failed to perform DECLARE")

func (instruction *Instruction_DECLARE) Parse(args *string) error {
	name, type_annotation, has_type := strings.Cut(*args, " ")
	if has_type {
		if unquoted_type, err := utils.DeserializeName(type_annotation); err != nil {
			return fmt.Errorf("%w: bad type arg: %w", ErrDeclareParsingType, err)
		} else {
			instruction.type_annotation = unquoted_type
			instruction.has_type_annotation = true
		}
	}
	if unquoted_name, err := utils.DeserializeName(name); err != nil {
		return fmt.Errorf("%w: bad name arg: %w", ErrDeclareParsingName, err)
	} else {
		instruction.name = unquoted_name
	}
	return nil
}

func (instruction *Instruction_DECLARE) String() string {
	if instruction.has_type_annotation {
		return fmt.Sprintf("DECLARE %s %s", utils.SerializeName(instruction.name), utils.SerializeName(instruction.type_annotation))
	} else {
		return fmt.Sprintf("DECLARE %s", utils.SerializeName(instruction.name))
	}
}

func (instruction *Instruction_DECLARE) Perform(vm *VirtualMachine) error {
	value, err := vm.stack.Pop()
	if err != nil {
		return fmt.Errorf("%q: %w", ErrDeclarePerform, err)
	}
	id := vm.memory.StorePrimitive(value)
	if instruction.has_type_annotation {
		if err := vm.memory.BindNewVariableWithType(instruction.name, instruction.type_annotation, id); err != nil {
			return fmt.Errorf("%q: %w", ErrDeclarePerform, err)
		}
	} else if err := vm.memory.BindNewVariable(instruction.name, id); err != nil {
		return fmt.Errorf("%q: %w", ErrDeclarePerform, err)
	}
	return nil
}
