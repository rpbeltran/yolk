package utils

import (
	"testing"
)

func TestFIFO(t *testing.T) {
	queue := CreateQueue[int]()
	for i := 0; i < 100; i++ {
		queue.Push(i)
	}

	if actual := queue.Size(); actual != 100 {
		t.Fatalf("Queue had size %d, expected 100", actual)
	}

	for i := 0; i < 100; i++ {
		expected := i
		if value, err := queue.Pop(); err != nil {
			t.Fatalf("queue.Pop() expected %d with no error, instead received the error %v", expected, err)
		} else if value != expected {
			t.Fatalf("queue.Pop() expected to return %d but returned %d", expected, value)
		}
	}

	expected_error := "attempted to pop from empty queue"
	if value, err := queue.Pop(); err == nil {
		t.Fatalf("queue.Pop() expected error %q, instead succeeded and returned %d", expected_error, value)
	} else if err.Error() != expected_error {
		t.Fatalf("queue.Pop() expected error %q, but actually gave the error %q", expected_error, err)
	}
}
