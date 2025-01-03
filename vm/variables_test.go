package vm

import (
	"testing"
	"yolk/types"
)

const name = "foo"
const bad_name = "foobar"
const expected_str = "Hello World!!!"

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
