package vm

import (
	"fmt"
	"yolk/types"
)

// FIXME: Currently only global variables are supported, support local variables soon

func (vm *VirtualMachine) FetchVariable(name string) (types.Primitive, error) {
	if object, ok := vm.globals[name]; !ok {
		return nil, fmt.Errorf("could not find a variable with the name %q", name)
	} else {
		return object, nil
	}
}

func (vm *VirtualMachine) StoreNewVariable(name string, value types.Primitive) error {
	if _, ok := vm.globals[name]; ok {
		return fmt.Errorf("variable %q cannot be redeclared", name)
	}
	vm.globals[name] = value
	return nil
}

func (vm *VirtualMachine) UpdateVariable(name string, value types.Primitive) error {
	if _, ok := vm.globals[name]; !ok {
		return fmt.Errorf("no variable with name %q to be updated", name)
	}
	vm.globals[name] = value
	return nil
}
