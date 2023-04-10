package main

import (
	"fmt"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) String() string {
	if l != nil {
		return fmt.Sprintf("val = %v, next = [%s] ", l.val, l.next.String())
	}
	return ""
}

func main() {

	l := List[int]{&List[int]{nil, 1}, 2}
	fmt.Println(l.String())
}
