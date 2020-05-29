package main

import (
	"list/list"
)

func main() {
	l := list.New()
	l.Append(1)
	l.Append(2)

	l.Remove(l.First())

	l.Prepend(3)

	l.Print() //expect: 3 2
}

