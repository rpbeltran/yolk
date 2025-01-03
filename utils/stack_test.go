package utils

import (
	"testing"
)

func TestLIFO(t *testing.T) {
	stack := CreateStack[int]()
	for i := 0; i < 100; i++ {
		stack.Push(i)
	}

	if actual := stack.Size(); actual != 100 {
		t.Fatalf("Stack had size %d, expected 100", actual)
	}

	if stack.Empty() {
		t.Fatalf("Stack.Empty() was true, expected false")
	}

	if value, err := stack.Peek(); err != nil {
		t.Fatalf("stack.Peek() got unexpected errror: %v", err)
	} else if *value != 99 {
		t.Fatalf("stack.Peek() expected 99 but got %d", *value)
	}

	for i := 0; i < 100; i++ {
		expected := 99 - i
		if value, err := stack.Pop(); err != nil {
			t.Fatalf("stack.Pop() expected %d with no error, instead received the error %v", expected, err)
		} else if value != expected {
			t.Fatalf("stack.Pop() expected to return %d but returned %d", expected, value)
		}
	}

	expected_error := "attempted to pop from empty stack"
	if value, err := stack.Pop(); err == nil {
		t.Fatalf("stack.Pop() expected error %q, instead succeeded and returned %d", expected_error, value)
	} else if err.Error() != expected_error {
		t.Fatalf("stack.Pop() expected error %q, but actually gave the error %q", expected_error, err)
	}

	if !stack.Empty() {
		t.Fatalf("Stack.Empty() was false, expected true")
	}
}

func BenchmarkLIFO(b *testing.B) {
	stack := CreateStack[int]()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
		if _, err := stack.Pop(); err != nil {
			b.Fatalf("stack.Pop() had unexpected error: %v", err)
		}
	}
}

func BenchmarkUnsafeLIFO(b *testing.B) {
	stack := CreateStack[int]()
	for i := 0; i < b.N; i++ {
		stack.PushUnsafe(i)
		if _, err := stack.PopUnsafe(); err != nil {
			b.Fatalf("stack.Pop() had unexpected error: %v", err)
		}
	}
}
