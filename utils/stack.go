package utils

import (
	"errors"
	"sync"
)

type Stack[T any] struct {
	lock sync.Mutex
	data []T
}

func CreateStack[T any]() *Stack[T] {
	return &Stack[T]{sync.Mutex{}, make([]T, 0)}
}

func (stack *Stack[T]) Push(v T) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	stack.PushUnsafe(v)
}

func (stack *Stack[T]) PushUnsafe(v T) {
	stack.data = append(stack.data, v)
}

func (stack *Stack[T]) Pop() (T, error) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	return stack.PopUnsafe()
}

func (stack *Stack[T]) PopUnsafe() (T, error) {
	l := len(stack.data)
	if l == 0 {
		var nothing T
		return nothing, errors.New("attempted to pop from empty stack")
	}

	res := stack.data[l-1]
	stack.data = stack.data[:l-1]
	return res, nil
}
