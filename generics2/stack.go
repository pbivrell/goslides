package main

import (
	"fmt"
)

// START_STACK OMIT

// Stack represents a stack datastructure
type Stack[T any] []T

// Push appends to the end of the stack
func (s *Stack[T]) Push (elem T) {
	*s = append(*s, elem)
}

// Pop returns the last pushed item
func (s *Stack[T]) Pop() (elem T) {

	if len(*s) < 1 {
		return elem
	}

	return (*s)[len(*s)-1]
}

// END_STACK OMIT

func main() {
	
	// START_MAIN OMIT
	s := Stack[int]{1,2,3}
	s.Push(4)
	fmt.Println(s.Pop())

	var s1 Stack[string]
	fmt.Println(s1.Pop())
	s1.Push("a")
	s1.Push("b")
	s1.Push("c")
	fmt.Println(s1.Pop())

	type X struct {
		A string
	}

	s2 := Stack[X]{}
	s2.Push(X{ A: "A"})
	s2.Push(X{ A: "B"})
	// END_MAIN OMIT

}
