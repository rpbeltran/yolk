package memory

import (
	"errors"
	"fmt"
	"yolk/types"
	"yolk/utils"
)

// FIXME: Make sure primitive values are not being passed as references (i.e {a = 1; b = a; b += 1; should not make a = 2})
// FIXME: Currently only global variables are supported, support local variables soon
// FIXME: Currently types are checked by name alone, but implicitly castable types should be permitted

type memID uint64

var ErrVariableDoesNotExist = errors.New("could not find a variable with given name")
var ErrVariableDoesNotHaveValidID = errors.New("variable has invalid id")
var ErrVariableRedeclaration = errors.New("cannot redeclare a variable with given name")
var ErrVariableTypeError = errors.New("type annotated variables cannot be declared with a value of an incompatible type")

type Memory struct {
	globals    map[string]*Variable
	objects    map[memID]types.Primitive
	ref_counts map[memID]uint
}

func NewVMMemory() Memory {
	return Memory{
		globals:    make(map[string]*Variable, 0),
		objects:    make(map[memID]types.Primitive, 0),
		ref_counts: make(map[memID]uint, 0),
	}
}

type Variable struct {
	bound_id        memID
	type_annotation string
	has_type        bool
}

func (vm *Memory) StorePrimitive(object types.Primitive) memID {
	next_id := memID(len(vm.objects))
	vm.objects[memID(next_id)] = object
	vm.ref_counts[memID(next_id)] = 0
	return memID(next_id)
}

func (vm *Memory) FetchVariableByName(name string) (types.Primitive, error) {
	if variable, ok := vm.globals[name]; !ok {
		return nil, fmt.Errorf("call to FetchVariable(%q): %w", name, ErrVariableDoesNotExist)
	} else if value, ok := vm.objects[variable.bound_id]; !ok {
		return nil, fmt.Errorf("call to FetchVariable(%q): %w", name, ErrVariableDoesNotHaveValidID)
	} else {
		return value, nil
	}
}

func (vm *Memory) BindNewVariable(name string, id memID) error {
	if _, ok := vm.globals[name]; ok {
		return fmt.Errorf("failed to make new variable %s: %w", utils.SerializeName(name), ErrVariableRedeclaration)
	}
	if _, ok := vm.objects[id]; !ok {
		return fmt.Errorf("call to BindNewVariable(%q, %d): %w", name, id, ErrVariableDoesNotHaveValidID)
	}
	vm.globals[name] = &Variable{bound_id: id}
	vm.ref_counts[id] += 1
	return nil
}

func (vm *Memory) BindNewVariableWithType(name string, type_annotation string, id memID) error {
	if _, ok := vm.globals[name]; ok {
		return fmt.Errorf("failed to make new variable %s: %w", utils.SerializeName(name), ErrVariableRedeclaration)
	} else if value, ok := vm.objects[id]; !ok {
		return fmt.Errorf("call to BindNewVariable(%q, %d): %w", name, id, ErrVariableDoesNotHaveValidID)
	} else if value_type := value.Type(); value_type != type_annotation {
		return fmt.Errorf("%w: cannot bind data of type %s to %s, which expects %s", ErrVariableTypeError,
			utils.SerializeName(value.Type()), utils.SerializeName(name), utils.SerializeName(type_annotation))
	}
	vm.globals[name] = &Variable{
		bound_id:        id,
		type_annotation: type_annotation,
		has_type:        true,
	}
	vm.ref_counts[id] += 1
	return nil
}

func (vm *Memory) RebindVariable(name string, id memID) error {
	variable, ok := vm.globals[name]
	if !ok {
		return fmt.Errorf("call to RebindVariable(%q): %w", utils.SerializeName(name), ErrVariableDoesNotExist)
	}
	if value, ok := vm.objects[id]; !ok {
		return fmt.Errorf("call to RebindVariable(%q, %d): %w", name, id, ErrVariableDoesNotHaveValidID)
	} else if variable.has_type && value.Type() != variable.type_annotation {
		return fmt.Errorf("%w: cannot bind data of type %s to %s, which expects %s", ErrVariableTypeError,
			utils.SerializeName(value.Type()), utils.SerializeName(name), utils.SerializeName(variable.type_annotation))
	} else {
		vm.DecrementRefcount(variable.bound_id)
		variable.bound_id = id
		vm.ref_counts[id] += 1
	}
	return nil
}

func (vm *Memory) DecrementRefcount(id memID) {
	vm.ref_counts[id] -= 1
	if vm.ref_counts[id] == 0 {
		delete(vm.objects, id)
	}
}
