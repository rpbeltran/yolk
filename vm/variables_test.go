package vm

import (
	"testing"
	"yolk/types"
)

const name = "foo"
const bad_name = "foobar"
const expected_id = uint64(7)
const bad_id = uint64(0)
const expected_str = "Hello World!!!"

func TestGetVariableID(t *testing.T) {
	vm := NewVM()
	vm.global_names[name] = expected_id

	if id, err := vm.GetVariableID(name); err != nil {
		t.Fatalf("Unexpected error calling vm.GetVariableID(%q): %v", name, err)
	} else if id != expected_id {
		t.Fatalf("Expected vm.GetVariableID(%q) to give %d, instead got: %d", name, expected_id, id)
	}

	if id, err := vm.GetVariableID(bad_name); err == nil {
		t.Fatalf("Expected error calling vm.GetVariableID(%q), instead returned: %d", bad_name, id)
	}
}

func TestFetchVariableById(t *testing.T) {
	vm := NewVM()
	vm.globals[expected_id] = types.MakeString(expected_str)

	if obj, err := vm.FetchVariableById(expected_id); err != nil {
		t.Fatalf("Unexpected error calling vm.FetchVariableById(%d): %v", expected_id, err)
	} else if display := obj.Display(); display != expected_str {
		t.Fatalf("Expected vm.FetchVariableById(%d) to give %q, instead got: %s", expected_id, expected_str, display)
	}

	if obj, err := vm.FetchVariableById(bad_id); err == nil {
		t.Fatalf("Expected error calling vm.FetchVariableById(%d), instead returned: %q", bad_id, obj.Display())
	}
}

func TestFetchVariable(t *testing.T) {
	vm := NewVM()
	vm.global_names[name] = expected_id
	vm.globals[expected_id] = types.MakeString(expected_str)

	if obj, id, err := vm.FetchVariable(name); err != nil {
		t.Fatalf("Unexpected error calling vm.FetchVariable(%q): %v", name, err)
	} else if display := obj.Display(); display != expected_str {
		t.Fatalf("Expected vm.FetchVariable(%q) to give %d, instead got: %s", name, expected_id, display)
	} else if id != expected_id {
		t.Fatalf("Expected FetchVariable to give the id %d, but instead gave %d", expected_id, id)
	}

	if obj, _, err := vm.FetchVariable(bad_name); err == nil {
		t.Fatalf("Expected error calling vm.FetchVariable(%q), instead returned: %q", bad_name, obj.Display())
	}
}

func TestStoreNewVariable(t *testing.T) {
	vm := NewVM()
	vm.variable_id_counter = expected_id
	if id, err := vm.StoreNewVariable(name, types.MakeString(expected_str)); err != nil {
		t.Fatalf("Unexpected error calling vm.StoreNewVariable(%q, %q): %v", name, expected_str, err)
	} else if id != expected_id {
		t.Fatalf("Expected StoreNewVariable to give the id %d, but instead gave %d", expected_id, id)
	}

	if obj, id, err := vm.FetchVariable(name); err != nil {
		t.Fatalf("Unexpected error calling vm.FetchVariable(%q): %v", name, err)
	} else if display := obj.Display(); display != expected_str {
		t.Fatalf("Expected FetchVariable(%q) to give %d, instead got: %s", name, expected_id, display)
	} else if id != expected_id {
		t.Fatalf("Expected FetchVariable to give the id %d, but instead gave %d", expected_id, id)
	}

	if _, err := vm.StoreNewVariable(name, types.MakeString(expected_str)); err == nil {
		t.Fatalf("Expected an error calling vm.StoreNewVariable(%q, %q) a second time, got none", name, expected_str)
	}
}

func TestStoreNewVariableWithID(t *testing.T) {
	vm := NewVM()
	if err := vm.StoreNewVariableWithID(name, expected_id, types.MakeString(expected_str)); err != nil {
		t.Fatalf("Unexpected error calling vm.StoreNewVariable(%q, %q): %v", name, expected_str, err)
	}

	if obj, id, err := vm.FetchVariable(name); err != nil {
		t.Fatalf("Unexpected error calling vm.FetchVariable(%q): %v", name, err)
	} else if display := obj.Display(); display != expected_str {
		t.Fatalf("Expected vm.FetchVariable(%q) to give %d, instead got: %s", name, expected_id, display)
	} else if id != expected_id {
		t.Fatalf("Expected FetchVariable to give the id %d, but instead gave %d", expected_id, id)
	}

	if err := vm.StoreNewVariableWithID(name, bad_id, types.MakeString(expected_str)); err == nil {
		t.Fatal("Expected an error calling vm.StoreNewVariableWithID using name a second time, got none")
	}

	if err := vm.StoreNewVariableWithID(bad_name, expected_id, types.MakeString(expected_str)); err == nil {
		t.Fatal("Expected an error calling vm.StoreNewVariableWithID using id a second time, got none")
	}
}

func TestUpdateVariable(t *testing.T) {
	vm := NewVM()
	vm.variable_id_counter = expected_id

	if _, err := vm.UpdateVariable(name, types.MakeString(expected_str)); err == nil {
		t.Fatalf("Expected an error calling vm.UpdateVariable(%q, %q) for undefined %q, got none", name, expected_str, name)
	}

	if id, err := vm.StoreNewVariable(name, types.MakeString(expected_str)); err != nil {
		t.Fatalf("Unexpected error calling vm.StoreNewVariable(%q, %q): %v", name, expected_str, err)
	} else if id != expected_id {
		t.Fatalf("Expected vm.StoreNewVariable to give the id %d, but instead gave %d", expected_id, id)
	}

	new_string := "walawoo"
	if id, err := vm.UpdateVariable(name, types.MakeString(new_string)); err != nil {
		t.Fatalf("Unexpected error calling vm.UpdateVariable(%q, %q): %v", name, new_string, err)
	} else if obj, _, err := vm.FetchVariable(name); err != nil {
		t.Fatalf("Unexpected error calling vm.FetchVariable(%q): %v", name, err)
	} else if display := obj.Display(); display != new_string {
		t.Fatalf("Expected vm.FetchVariable(%q) to give %q, instead got: %s", name, new_string, display)
	} else if id != expected_id {
		t.Fatalf("Expected UpdateVariable to give the id %d, but instead gave %d", expected_id, id)
	}
}

func TestUpdateVariableByID(t *testing.T) {
	vm := NewVM()
	if err := vm.UpdateVariableByID(expected_id, types.MakeString(expected_str)); err == nil {
		t.Fatalf("Expected an error calling vm.UpdateVariableByID(%d, %q) for undefined %d, got none", expected_id, expected_str, expected_id)
	}

	if err := vm.StoreNewVariableWithID(name, expected_id, types.MakeString(expected_str)); err != nil {
		t.Fatalf("Unexpected error calling vm.StoreNewVariable(%q, %q): %v", name, expected_str, err)
	}

	new_string := "walawoo"
	if err := vm.UpdateVariableByID(expected_id, types.MakeString(new_string)); err != nil {
		t.Fatalf("Unexpected error calling vm.UpdateVariableByID(%d, %q): %v", expected_id, new_string, err)
	} else if obj, id, err := vm.FetchVariable(name); err != nil {
		t.Fatalf("Unexpected error calling vm.FetchVariable(%q): %v", name, err)
	} else if display := obj.Display(); display != new_string {
		t.Fatalf("Expected vm.FetchVariable(%q) to give %q, instead got: %s", name, new_string, display)
	} else if id != expected_id {
		t.Fatalf("Expected FetchVariable to give the id %d, but instead gave %d", expected_id, id)
	}
}
