package utils

import (
	"errors"
	"sync"
)

type Stack[T any] struct {
	lock        sync.Mutex
	data        []T
	thread_safe bool
}

func CreateStack[T any]() *Stack[T] {
	return &Stack[T]{sync.Mutex{}, make([]T, 0), false}
}

func (stack *Stack[T]) Push(v T) {
	if stack.thread_safe {
		stack.lock.Lock()
		defer stack.lock.Unlock()
	}

	stack.PushUnsafe(v)
}

func (stack *Stack[T]) PushUnsafe(v T) {
	stack.data = append(stack.data, v)
}

func (stack *Stack[T]) Pop() (T, error) {
	if stack.thread_safe {
		stack.lock.Lock()
		defer stack.lock.Unlock()
	}

	return stack.PopUnsafe()
}

func (stack *Stack[T]) PopUnsafe() (T, error) {
	l := len(stack.data)
	if l == 0 {
		var nothing T
		return nothing, errors.New("attempted to pop from empty stack")
	}

	last := stack.data[l-1]
	stack.data = stack.data[:l-1]
	return last, nil
}

func (stack *Stack[T]) Peek() (*T, error) {
	if stack.thread_safe {
		stack.lock.Lock()
		defer stack.lock.Unlock()
	}

	return stack.PeekUnsafe()
}

func (stack *Stack[T]) PeekUnsafe() (*T, error) {
	if l := len(stack.data); l == 0 {
		return nil, errors.New("attempted to peek at empty stack")
	} else {
		return &stack.data[l-1], nil
	}
}

func (stack *Stack[T]) Size() int {
	return len(stack.data)
}

func (stack *Stack[T]) Empty() bool {
	return stack.Size() == 0
}
