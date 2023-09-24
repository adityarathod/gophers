package main

import (
	"errors"
	"fmt"
)

type LinkedList[T any] struct {
	head *Node[T]
}

type Node[T any] struct {
	value T
	next  *Node[T]
}

func makeLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{nil}
}

func (ll *LinkedList[T]) add(value T) {
	node := &Node[T]{value, nil}
	if ll.head == nil {
		ll.head = node
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
}

func (ll *LinkedList[T]) addMany(values ...T) {
	for _, value := range values {
		ll.add(value)
	}
}

func (ll *LinkedList[T]) get(index int) (T, error) {
	current := ll.head
	for i := 0; i < index; i++ {
		if current.next == nil {
			// wtf???
			return *new(T), errors.New("index out of bounds")
		}
		current = current.next
	}
	return current.value, nil
}

func (ll *LinkedList[T]) delete(index int) (T, error) {
	current := ll.head
	for i := 0; i < index-1; i++ {
		if current.next == nil {
			return *new(T), errors.New("index out of bounds")
		}
		current = current.next
	}
	deleted := current.next
	current.next = deleted.next
	return deleted.value, nil
}

func (ll *LinkedList[T]) toString() string {
	tmp := ""
	current := ll.head
	for current != nil {
		tmp += fmt.Sprintf("%v ", current.value)
		current = current.next
	}
	tmp += "\n"
	return tmp
}
