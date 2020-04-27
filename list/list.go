package list

import "fmt"

type Element struct {
	Value interface{}
	prev  *Element
	next  *Element
}

type List struct {
	head *Element
	tail *Element
	size int
}

func New() *List {
	return &List{}
}

func (l *List) Append(value interface{}) {
	node := &Element{
		Value: value,
	}
	if l.size == 0 {
		l.head = node
		l.tail = node
		node.prev = nil
		node.next = nil
	} else {
		node.prev = l.tail
		node.next = nil
		l.tail.next = node
		l.tail = node
	}
	l.size++
}

func (l *List) Prepend(value interface{}) {
	node := &Element{
		Value: value,
	}
	if l.size == 0 {
		l.head = node
		l.tail = node
		node.prev = nil
		node.next = nil
	} else {
		node.prev = nil
		node.next = l.head
		l.head.prev = node
		l.head = node
	}
	l.size++
}

func (l *List) Last() *Element {
	return l.tail
}

func (l *List) First() *Element {
	return l.head
}

func (l *List) Len() int {
	return l.size
}

func (l *List) Print() {
	for e := l.First(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func (l *List) Remove(e *Element) interface{} {
	if e.isFirst() {
		l.head = e.next
		e.next.prev = nil
	} else if e.isLast() {
		l.tail = e.prev
		e.prev.next = nil
	} else {
		e.prev.next = e.next
		e.next.prev = e.prev
	}
	l.size--
	return e.Value
}

func (l *List) InsertAfter(value interface{}, e *Element) *Element {
	node := &Element{
		Value: value,
	}
	node.prev = e
	node.next = e.next
	e.next = node
	e.next.prev = node
	l.size++
	return node
}

func (l *List) InsertBefore(value interface{}, e *Element) *Element {
	node := &Element{
		Value: value,
	}
	node.prev = e.prev
	node.next = e
	e.prev.next = node
	e.prev = node
	l.size++
	return node
}

func (e *Element) Next() *Element {
	if e.isLast() {
		return nil
	}
	return e.next
}

func (e *Element) Prev() *Element {
	if e.isFirst() {
		return nil
	}
	return e.prev
}

func (e *Element) isLast() bool {
	return e.next == nil
}

func (e *Element) isFirst() bool {
	return e.prev == nil
}
