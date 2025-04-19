package vm

/*
var ErrVariableDoesNotExist = errors.New("could not find a variable with given name")
var ErrVariableDoesNotHaveValidID = errors.New("variable has invalid id")
var ErrVariableRedeclaration = errors.New("cannot redeclare a variable with given name")
var ErrVariableTypeError = errors.New("type annotated variables cannot be declared with a value of an incompatible type")

type MemoryManager struct {
	globals map[string]Variable
	memory  map[memID]types.Primitive
}

type Variable struct {
	bound_memory   memID
	constraint     string
	has_constraint bool
}

func (vm *VirtualMachine) FetchVariable(name string) (types.Primitive, error) {
	if variable, ok := vm.globals[name]; !ok {
		return nil, fmt.Errorf("call to FetchVariable(%q): %w", name, ErrVariableDoesNotExist)
	} else if value, ok := vm.memory[variable.bound_memory]; !ok {
		return nil, fmt.Errorf("call to FetchVariable(%q): %w", name, ErrVariableDoesNotHaveValidID)
	} else {
		return value, nil
	}
}

func (vm *VirtualMachine) BindNewVariable(name string, id memID) error {
	if _, ok := vm.globals[name]; ok {
		return fmt.Errorf("failed to make new variable %s: %w", utils.SerializeName(name), ErrVariableRedeclaration)
	}
	if _, ok := vm.memory[id]; !ok {
		return fmt.Errorf("call to BindNewVariable(%q, %d): %w", name, id, ErrVariableDoesNotHaveValidID)
	}
	vm.globals[name] = Variable{bound_memory: id}
	return nil
}

func (vm *VirtualMachine) BindNewVariableWithType(name string, type_annotation string, id memID) error {
	if _, ok := vm.globals[name]; ok {
		return fmt.Errorf("failed to make new variable %s: %w", utils.SerializeName(name), ErrVariableRedeclaration)
	} else if value, ok := vm.memory[id]; !ok {
		return fmt.Errorf("call to BindNewVariable(%q, %d): %w", name, id, ErrVariableDoesNotHaveValidID)
	} else if value_type := value.Type(); value_type != type_annotation {
		return fmt.Errorf("%w: cannot bind data of type %s to %s, which expects %s", ErrVariableTypeError,
			utils.SerializeName(value.Type()), utils.SerializeName(name), utils.SerializeName(type_annotation))
	}
	vm.globals[name] = Variable{
		bound_memory:   id,
		constraint:     type_annotation,
		has_constraint: true,
	}
	return nil
}

func (vm *VirtualMachine) RebindVariable(name string, id memID) error {
	variable, ok := vm.globals[name]
	if !ok {
		return fmt.Errorf("call to RebindVariable(%q): %w", utils.SerializeName(name), ErrVariableDoesNotExist)
	}
	if value, ok := vm.memory[id]; !ok {
		return fmt.Errorf("call to RebindVariable(%q, %d): %w", name, id, ErrVariableDoesNotHaveValidID)
	} else if variable.has_constraint && value.Type() != variable.constraint {
		return fmt.Errorf("%w: got `%s` of type %s, expected %s (the type of %s)", ErrVariableTypeError,
			value.Display(), utils.SerializeName(value.Type()), utils.SerializeName(req_type), utils.SerializeName(name))
	} else {
		variable.bound_memory = id
	}
	return nil
}
*/
