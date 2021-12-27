package main

import (
	"fmt"
)

func main() {

	// START_CALLS OMIT
	value, ok := Search([]int{1,2,3,4,5}, 6)
	fmt.Println(value, ok)

	letter, ok := Search([]string{"a", "b", "c", "d"}, "b")
	fmt.Println(letter, ok)

	// Does this compile?
	// ints, ok := Search([][]int{ []int{1}, []int{2,3}, []int{4,5,6}}, []int{1})
	// fmt.Println(ints, ok)
	
	// Does this compile?
	// value, ok = Search([]string{"a", "b", "c", "d"}, "b")
	// fmt.Println(value, ok)

	// Does this compile?
	// anything, ok := Search[interface{}]([]interface{}{ "a", 1, 1.23, false}, "a")
	// fmt.Println(anything, ok)
	// END_CALLS OMIT
}


// START_FUNC OMIT
func Search[T comparable](slice []T, value T) (result T, ok bool) {

	for _, v := range slice {
		if value == v {
			return v, true
		}
	}

	return result, ok
}
// END_FUNC OMIT
