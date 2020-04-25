package list

import "fmt"

type Element struct {
	value interface{}
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
		value: value,
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
		value: value,
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
		fmt.Println(e.value)
	}
}

func (l *List) Remove(e *Element) interface{} {
	if e.IsFirst() {
		l.head = e.next
		e.next.prev = nil
	} else if e.IsLast() {
		l.tail = e.prev
		e.prev.next = nil
	} else {
		e.prev.next = e.next
		e.next.prev = e.prev
	}
	l.size--
	return e.value
}

func (l *List) Size() int {
	return l.size
}

func (l *List) InsertAfter(value interface{}, e *Element) *Element {
	node := &Element{
		value: value,
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
		value: value,
	}
	node.prev = e.prev
	node.next = e
	e.prev.next = node
	e.prev = node
	l.size++
	return node
}

func (e *Element) Next() *Element {
	if e.next != nil {
		return e.next
	}
	return nil
}

func (e *Element) Prev() *Element {
	if e.prev != nil {
		return e.Prev()
	}
	return e
}

func (e *Element) IsLast() bool {
	return e.next == nil
}

func (e *Element) IsFirst() bool {
	return e.prev == nil
}
