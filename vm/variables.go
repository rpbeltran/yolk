package vm

import (
	"fmt"
	"yolk/types"
)

// FIXME: Currently only global variables are supported, support local variables soon

func (vm *VirtualMachine) GetVariableID(name string) (uint64, error) {
	if id, ok := vm.global_names[name]; !ok {
		return 0, fmt.Errorf("could not find an associated id for the name: %q", name)
	} else {
		return id, nil
	}
}

func (vm *VirtualMachine) FetchVariableById(id uint64) (types.Primitive, error) {
	if object, ok := vm.globals[id]; !ok {
		return nil, fmt.Errorf("could not find an object for the id: %d", id)
	} else {
		return object, nil
	}
}

func (vm *VirtualMachine) FetchVariable(name string) (types.Primitive, uint64, error) {
	if id, err := vm.GetVariableID(name); err != nil {
		return nil, 0, err
	} else {
		value, err := vm.FetchVariableById(id)
		return value, id, err
	}
}

func (vm *VirtualMachine) StoreNewVariable(name string, value types.Primitive) (uint64, error) {
	if _, ok := vm.global_names[name]; ok {
		return 0, fmt.Errorf("variable %q cannot be redeclared", name)
	}
	new_id := vm.variable_id_counter
	vm.variable_id_counter += 1

	vm.global_names[name] = new_id
	vm.globals[new_id] = value
	return new_id, nil
}

func (vm *VirtualMachine) StoreNewVariableWithID(name string, id uint64, value types.Primitive) error {
	if _, ok := vm.global_names[name]; ok {
		return fmt.Errorf("variable %q cannot be redeclared", name)
	}
	if _, ok := vm.globals[id]; ok {
		return fmt.Errorf("variable id %d cannot be reused", id)
	}
	vm.global_names[name] = id
	vm.globals[id] = value
	return nil
}

func (vm *VirtualMachine) UpdateVariable(name string, value types.Primitive) (uint64, error) {
	if id, err := vm.GetVariableID(name); err != nil {
		return 0, fmt.Errorf("could not update the variable %q due to the error: %v (perhaps it does not exist)", name, err)
	} else {
		vm.globals[id] = value
		return id, nil
	}
}

func (vm *VirtualMachine) UpdateVariableByID(id uint64, value types.Primitive) error {
	if _, ok := vm.globals[id]; !ok {
		return fmt.Errorf("no variable with id %d to be updated", id)
	}
	vm.globals[id] = value
	return nil
}
