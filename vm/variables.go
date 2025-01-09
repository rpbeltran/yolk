package vm

import (
	"errors"
	"fmt"
	"yolk/types"
	"yolk/utils"
)

// FIXME: Make sure primitive values are not being passed as references (i.e {a = 1; b = a; b += 1; should not make a = 2})
// FIXME: Currently only global variables are supported, support local variables soon
// FIXME: Currently types are checked by name alone, but implicitly castable types should be permitted

var ErrVariableDoesNotExist = errors.New("could not find a variable with given name")
var ErrVariableRedeclaration = errors.New("cannot redeclare a variable with given name")
var ErrVariableTypeError = errors.New("type annotated variables cannot be declared with a value of an incompatible type")

func (vm *VirtualMachine) FetchVariable(name string) (types.Primitive, error) {
	if object, ok := vm.globals[name]; !ok {
		return nil, fmt.Errorf("attempting to fetch variable %q: %w", utils.SerializeName(name), ErrVariableDoesNotExist)
	} else {
		return object, nil
	}
}

func (vm *VirtualMachine) StoreNewVariable(name string, value types.Primitive) error {
	if _, ok := vm.globals[name]; ok {
		return fmt.Errorf("failed to make new variable %s: %w", utils.SerializeName(name), ErrVariableDoesNotExist)
	}
	vm.globals[name] = value
	return nil
}

func (vm *VirtualMachine) StoreNewVariableWithType(name string, type_annotation string, value types.Primitive) error {
	if _, ok := vm.globals[name]; ok {
		return fmt.Errorf("failed to make new variable %s: %w", utils.SerializeName(name), ErrVariableDoesNotExist)
	}
	if value_type := value.Type(); value_type != type_annotation {
		return fmt.Errorf("%w: got `%s` of type %s, expected %s (the type of %s)", ErrVariableTypeError,
			value.Display(), utils.SerializeName(value.Type()), utils.SerializeName(type_annotation), utils.SerializeName(name))
	}
	vm.globals[name] = value
	vm.globals_types[name] = type_annotation
	return nil
}

func (vm *VirtualMachine) UpdateVariable(name string, value types.Primitive) error {
	if _, ok := vm.globals[name]; !ok {
		return fmt.Errorf("attempting to update %q: %w", utils.SerializeName(name), ErrVariableDoesNotExist)
	}
	if req_type, has_type := vm.globals_types[name]; has_type && value.Type() != req_type {
		return fmt.Errorf("%w: got `%s` of type %s, expected %s (the type of %s)", ErrVariableTypeError,
			value.Display(), utils.SerializeName(value.Type()), utils.SerializeName(req_type), utils.SerializeName(name))
	}
	vm.globals[name] = value
	return nil
}
