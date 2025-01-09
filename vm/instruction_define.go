package vm

import (
	"errors"
	"fmt"
	"strings"
	"yolk/utils"
)

type Instruction_DEFINE struct {
	name                string
	type_annotation     string
	has_type_annotation bool
}

var ErrDefineParsing = errors.New("failed to parse .DEFINE")
var ErrDefineParsingName = fmt.Errorf("%w: invalid name arg", ErrDefineParsing)
var ErrDefineParsingType = fmt.Errorf("%w: invalid type arg", ErrDefineParsing)

var ErrDefinePerform = errors.New(".DEFINE instruction should not have been executed")

func (instruction *Instruction_DEFINE) Parse(args *string) error {
	name, type_annotation, has_type := strings.Cut(*args, " ")
	if has_type {
		if unquoted_type, err := utils.DeserializeName(type_annotation); err != nil {
			return fmt.Errorf("%w: bad type arg: %w", ErrDefineParsingType, err)
		} else {
			instruction.type_annotation = unquoted_type
			instruction.has_type_annotation = true
		}
	}
	if unquoted_name, err := utils.DeserializeName(name); err != nil {
		return fmt.Errorf("%w: bad name arg: %w", ErrDefineParsingName, err)
	} else {
		instruction.name = unquoted_name
	}
	return nil
}

func (instruction *Instruction_DEFINE) String() string {
	if instruction.has_type_annotation {
		return fmt.Sprintf(".DEFINE %s %s", utils.SerializeName(instruction.name), utils.SerializeName(instruction.type_annotation))
	}
	return fmt.Sprintf(".DEFINE %s", utils.SerializeName(instruction.name))
}

func (instruction *Instruction_DEFINE) Perform(vm *VirtualMachine) error {
	return ErrDeclarePerform
}
