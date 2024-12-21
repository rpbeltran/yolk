package utils

import (
	"errors"
	"sync"
)

type Queue[T any] struct {
	lock        sync.Mutex
	data        []T
	thread_safe bool
}

func CreateQueue[T any]() *Queue[T] {
	return &Queue[T]{sync.Mutex{}, make([]T, 0), false}
}

func (queue *Queue[T]) Push(v T) {
	if queue.thread_safe {
		queue.lock.Lock()
		defer queue.lock.Unlock()
	}

	queue.PushUnsafe(v)
}

func (queue *Queue[T]) PushUnsafe(v T) {
	queue.data = append(queue.data, v)
}

func (queue *Queue[T]) Pop() (T, error) {
	if queue.thread_safe {
		queue.lock.Lock()
		defer queue.lock.Unlock()
	}

	return queue.PopUnsafe()
}

func (queue *Queue[T]) PopUnsafe() (T, error) {
	l := len(queue.data)
	if l == 0 {
		var nothing T
		return nothing, errors.New("attempted to pop from empty queue")
	}

	first := queue.data[0]
	queue.data = queue.data[1:]
	return first, nil
}

func (queue *Queue[T]) Size() int {
	return len(queue.data)
}
