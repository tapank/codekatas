package linkedlist

import (
	"errors"
	"fmt"
)

type Node struct {
	Value int
	p     *Node // previous
	n     *Node // next
}

type List struct {
	Head *Node
	Tail *Node
}

func NewList(elements ...interface{}) *List {
	l := &List{}
	var last *Node
	for _, e := range elements {
		val, ok := e.(int)
		if !ok {
			fmt.Printf("Bad value %v cannot be converted to int", e)
			return nil
		}

		n := &Node{val, last, nil}
		if l.Head == nil {
			l.Head = n
		}
		l.Tail = n
		if last != nil {
			last.n = n
		}
		last = n
	}
	return l
}

func (n *Node) Next() *Node {
	if n == nil {
		return nil
	}
	return n.n
}

func (n *Node) Prev() *Node {
	if n == nil {
		return nil
	}
	return n.p
}

func (l *List) Unshift(v interface{}) {
	val, ok := v.(int)
	if !ok {
		fmt.Printf("Bad value %v cannot be converted to int", v)
		return
	}
	n := &Node{Value: val}
	n.n = l.Head
	if l.Head != nil {
		l.Head.p = n
	}
	l.Head = n
	if l.Tail == nil {
		l.Tail = n
	}
}

func (l *List) Push(v interface{}) {
	val, ok := v.(int)
	if !ok {
		fmt.Printf("Bad value %v cannot be converted to int", v)
		return
	}
	node := &Node{Value: val}
	if l.Head == nil {
		l.Head = node
	}
	if l.Tail != nil {
		l.Tail.n = node
	}
	node.p = l.Tail
	l.Tail = node
}

func (l *List) Shift() (interface{}, error) {
	if l.Head == nil {
		return nil, errors.New("nothing to shift in an empty list")
	}
	n := l.Head
	l.Head = l.Head.n
	if l.Head == nil {
		l.Tail = nil
	} else {
		l.Head.p = nil
	}
	return n.Value, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.Tail == nil {
		return nil, errors.New("nothing to pop in an empty list")
	}
	n := l.Tail
	l.Tail = n.p
	if l.Tail != nil {
		l.Tail.n = nil
	} else {
		l.Head = nil
	}
	return n.Value, nil
}

func (l *List) Reverse() {
	n := l.Head
	for n != nil {
		n.p, n.n = n.n, n.p
		n = n.p
	}
	l.Head, l.Tail = l.Tail, l.Head
}

func (l *List) First() *Node {
	return l.Head
}

func (l *List) Last() *Node {
	return l.Tail
}
