package vm

import (
	"testing"
	"yolk/types"
)

const name = "foo"
const type_name = "str"
const bad_name = "foobar"
const expected_str = "Hello World!!!"
const other_str = "duck duck goose."

func TestFetchVariable(t *testing.T) {
	vm := NewVM()
	vm.globals[name] = types.MakeString(expected_str)

	if obj, err := vm.FetchVariable(name); err != nil {
		t.Fatalf("Unexpected error calling vm.FetchVariable(%q): %v", name, err)
	} else if display := obj.Display(); display != expected_str {
		t.Fatalf("Expected vm.FetchVariable(%q) to give %q, instead got: %q", name, expected_str, display)
	}

	if obj, err := vm.FetchVariable(bad_name); err == nil {
		t.Fatalf("Expected error calling vm.FetchVariable(%q), instead returned: %q", bad_name, obj.Display())
	}
}

func TestStoreNewVariable(t *testing.T) {
	vm := NewVM()
	if err := vm.StoreNewVariable(name, types.MakeString(expected_str)); err != nil {
		t.Fatalf("Unexpected error calling vm.StoreNewVariable(%q, %q): %v", name, expected_str, err)
	}

	if obj, err := vm.FetchVariable(name); err != nil {
		t.Fatalf("Unexpected error calling vm.FetchVariable(%q): %v", name, err)
	} else if display := obj.Display(); display != expected_str {
		t.Fatalf("Expected FetchVariable(%q) to give %q, instead got: %q", name, expected_str, display)
	}

	if err := vm.StoreNewVariable(name, types.MakeString(expected_str)); err == nil {
		t.Fatalf("Expected an error calling vm.StoreNewVariable(%q, %q) a second time, got none", name, expected_str)
	}
}

func TestUpdateVariable(t *testing.T) {
	vm := NewVM()
	if err := vm.UpdateVariable(name, types.MakeString(expected_str)); err == nil {
		t.Fatalf("Expected an error calling vm.UpdateVariable(%q, %q) for undefined %q, got none", name, expected_str, name)
	}

	if err := vm.StoreNewVariable(name, types.MakeString(expected_str)); err != nil {
		t.Fatalf("Unexpected error calling vm.StoreNewVariable(%q, %q): %v", name, expected_str, err)
	}

	new_string := "walawoo"
	if err := vm.UpdateVariable(name, types.MakeString(new_string)); err != nil {
		t.Fatalf("Unexpected error calling vm.UpdateVariable(%q, %q): %v", name, new_string, err)
	} else if obj, err := vm.FetchVariable(name); err != nil {
		t.Fatalf("Unexpected error calling vm.FetchVariable(%q): %v", name, err)
	} else if display := obj.Display(); display != new_string {
		t.Fatalf("Expected vm.FetchVariable(%q) to give %q, instead got: %s", name, new_string, display)
	}
}

func TestVariablesWithType(t *testing.T) {
	vm := NewVM()

	if err := vm.StoreNewVariableWithType(name, type_name, types.MakeString(expected_str)); err != nil {
		t.Fatalf("Unexpected error calling vm.StoreNewVariableWithType(%q, %q, %q): %v", name, type_name, expected_str, err)
	}

	if obj, err := vm.FetchVariable(name); err != nil {
		t.Fatalf("Unexpected error calling vm.FetchVariable(%q): %v", name, err)
	} else if !obj.Equal(types.MakeString(expected_str)) {
		t.Fatalf("Expected FetchVariable(%q) to give %q, instead got: %q", name, expected_str, obj.Display())
	}

	if err := vm.StoreNewVariableWithType(name, type_name, types.MakeString(expected_str)); err == nil {
		t.Fatalf("Expected an error calling vm.StoreNewVariableWithType(%q, %q, %q) a second time, got none", name, type_name, expected_str)
	}

	if err := vm.UpdateVariable(name, types.MakeBool(true)); err == nil {
		t.Fatalf("Expected an error calling vm.UpdateVariable(%q, true) with the wrong type, got none", name)
	}

	if err := vm.UpdateVariable(name, types.MakeString(other_str)); err != nil {
		t.Fatalf("Unexpected error calling vm.UpdateVariable(%q, %q): %v", name, other_str, err)
	}

	if err := vm.StoreNewVariableWithType(bad_name, type_name, types.MakeBool(true)); err == nil {
		t.Fatalf("Expected an error from vm.StoreNewVariableWithType(%q, %q, true), got none", bad_name, type_name)
	}
}
