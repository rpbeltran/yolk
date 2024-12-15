package vm

import (
	"testing"
)

func TestLIFO(t *testing.T) {
	stack := CreateStack[int]()
	for i := 0; i < 100; i++ {
		stack.Push(i)
	}
	for i := 0; i < 100; i++ {
		expected := 99 - i
		if value, err := stack.Pop(); err != nil {
			t.Fatalf("stack.Pop() expected %d with no error, instead received the error %v", expected, err)
		} else if value != expected {
			t.Fatalf("stack.Pop() expected to return %d but returned %d", expected, value)
		}
	}
}
