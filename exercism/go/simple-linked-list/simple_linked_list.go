package linkedlist

import (
	"errors"
)

type List struct {
	Head, Tail *Node
	Length     int
}
type Node struct {
	V    int
	Next *Node
}

func New(elements []int) *List {
	list := &List{}
	for _, v := range elements {
		list.Push(v)
	}
	return list
}

func (l *List) Size() int {
	return l.Length
}

func (l *List) Push(element int) {
	node := &Node{element, nil}
	if l.Head == nil {
		l.Head = node
	}
	if l.Tail != nil {
		l.Tail.Next = node
	}
	l.Tail = node
	l.Length++
}

func (l *List) Pop() (v int, err error) {
	if l.Length == 0 {
		err = errors.New("pop: list is empty")
		return
	}

	v = l.Tail.V
	l.Length--
	if l.Length == 0 {
		l.Head, l.Tail = nil, nil
		return
	}

	for node := l.Head; ; {
		if node.Next == nil {
			break
		}
		l.Tail, node = node, node.Next
	}
	l.Tail.Next = nil
	return
}

func (l *List) Array() (a []int) {
	if l == nil {
		return
	}
	for node := l.Head; node != nil; node = node.Next {
		a = append(a, node.V)
	}
	return
}

func (l *List) Reverse() *List {
	if l == nil || l.Length == 0 {
		return l
	}
	var node, prev *Node
	for node = l.Head; node != nil; {
		node, prev, node.Next = node.Next, node, prev
	}
	l.Head, l.Tail = l.Tail, l.Head
	return l
}
