package memory

import (
	"testing"
	"yolk/types"
)

const name = "foo"
const name2 = "foo2"
const type_name = "str"
const bad_name = "foobar"
const expected_str = "Hello World!!!"
const other_str = "duck duck goose."

func TestFetchVariable(t *testing.T) {
	mem := NewVMMemory()
	bound_id := memID(13)
	mem.globals[name] = &Variable{bound_id: bound_id}
	mem.objects[bound_id] = types.MakeString(expected_str)

	if obj, err := mem.FetchVariableByName(name); err != nil {
		t.Fatalf("Unexpected error calling mem.FetchVariableByName(%q): %v", name, err)
	} else if display := obj.Display(); display != expected_str {
		t.Fatalf("Expected mem.FetchVariableByName(%q) to give %q, instead got: %q", name, expected_str, display)
	}

	if obj, err := mem.FetchVariableByName(bad_name); err == nil {
		t.Fatalf("Expected error calling mem.FetchVariableByName(%q), instead returned: %q", bad_name, obj.Display())
	}
}

func TestStoreNewVariable(t *testing.T) {
	mem := NewVMMemory()
	id := mem.StorePrimitive(types.MakeString(expected_str))
	if err := mem.BindNewVariable(name, id); err != nil {
		t.Fatalf("Unexpected error calling mem.BindNewVariable(%q, %q): %v", name, expected_str, err)
	} else if obj, err := mem.FetchVariableByName(name); err != nil {
		t.Fatalf("Unexpected error calling mem.FetchVariableByName(%q): %v", name, err)
	} else if display := obj.Display(); display != expected_str {
		t.Fatalf("Expected mem.FetchVariableByName(%q) to give %q, instead got: %q", name, expected_str, display)
	}

	id2 := mem.StorePrimitive(types.MakeString(other_str))
	if err := mem.BindNewVariable(name, id2); err == nil {
		t.Fatalf("Expected an error calling mem.BindNewVariable(%q, %q) a second time, got none", name, expected_str)
	}

	bad_id := 999
	if err := mem.BindNewVariable(name, memID(bad_id)); err == nil {
		t.Fatalf("Expected an error calling mem.BindNewVariable(%q, %q), got none", name, memID(bad_id))
	}

	if err := mem.BindNewVariable(name2, id2); err != nil {
		t.Fatalf("Unexpected error calling mem.BindNewVariable(%q, %q): %v", name2, other_str, err)
	} else if obj, err := mem.FetchVariableByName(name2); err != nil {
		t.Fatalf("Unexpected error calling mem.FetchVariableByName(%q): %v", name2, err)
	} else if display := obj.Display(); display != other_str {
		t.Fatalf("Expected mem.FetchVariableByName(%q) to give %q, instead got: %q", name2, other_str, display)
	}
}

func TestUpdateVariable(t *testing.T) {
	mem := NewVMMemory()
	id := mem.StorePrimitive(types.MakeString(expected_str))
	if err := mem.BindNewVariable(name, id); err != nil {
		t.Fatalf("Unexpected error calling mem.BindNewVariable(%q, %q): %v", name, expected_str, err)
	} else if obj, err := mem.FetchVariableByName(name); err != nil {
		t.Fatalf("Unexpected error calling mem.FetchVariableByName(%q): %v", name, err)
	} else if display := obj.Display(); display != expected_str {
		t.Fatalf("Expected mem.FetchVariableByName(%q) to give %q, instead got: %q", name, expected_str, display)
	}

	id2 := mem.StorePrimitive(types.MakeString(other_str))
	if err := mem.RebindVariable(name, id2); err != nil {
		t.Fatalf("Unexpected error calling vm.UpdateVariable(%q, %q): %v", name, other_str, err)
	} else if obj, err := mem.FetchVariableByName(name); err != nil {
		t.Fatalf("Unexpected error calling mem.FetchVariableByName(%q): %v", name, err)
	} else if display := obj.Display(); display != other_str {
		t.Fatalf("Expected mem.FetchVariableByName(%q) to give %q, instead got: %s", name, other_str, display)
	}
}

func TestVariablesWithType(t *testing.T) {
	mem := NewVMMemory()

	id := mem.StorePrimitive(types.MakeString(expected_str))

	if err := mem.BindNewVariableWithType(name, type_name, id); err != nil {
		t.Fatalf("Unexpected error calling mem.BindNewVariableWithType(%q, %q, %q): %v", name, type_name, expected_str, err)
	}

	if obj, err := mem.FetchVariableByName(name); err != nil {
		t.Fatalf("Unexpected error calling mem.FetchVariableByName(%q): %v", name, err)
	} else if !obj.Equal(types.MakeString(expected_str)) {
		t.Fatalf("Expected mem.FetchVariableByName(%q) to give %q, instead got: %q", name, expected_str, obj.Display())
	}

	if err := mem.BindNewVariableWithType(name, type_name, id); err == nil {
		t.Fatalf("Expected an error calling mem.BindNewVariableWithType(%q, %q, %q) a second time, got none", name, type_name, expected_str)
	}

	id2 := mem.StorePrimitive(types.MakeString(other_str))
	if err := mem.RebindVariable(name, id2); err != nil {
		t.Fatalf("Unexpected error calling mem.RebindVariable(%q, %q): %v", name, other_str, err)
	}

	id_bad := mem.StorePrimitive(types.MakeBool(true))
	if err := mem.RebindVariable(name, id_bad); err == nil {
		t.Fatalf("Expected an error calling vm.UpdateVariable(%q, true) with the wrong type, got none", name)
	}

	if err := mem.BindNewVariableWithType(name2, type_name, id_bad); err == nil {
		t.Fatalf("Expected an error from vm.StoreNewVariableWithType(%q, %q, true), got none", bad_name, type_name)
	}

	bad_id := 999
	if err := mem.BindNewVariableWithType(name, type_name, memID(bad_id)); err == nil {
		t.Fatalf("Expected an error calling vm.StoreNewVariable(%q, %q), got none", name, memID(bad_id))
	}
}
