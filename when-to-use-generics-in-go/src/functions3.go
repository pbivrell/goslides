package main

import (
	"fmt"
)

// START_FUNC OMIT
func GetFirst[T any](s []T) T {
	var zero T
	if len(s) < 1 {
		return zero
	}
	return s[0]
}
// END_FUNC OMIT

func main() {
	// START_MAIN OMIT
	letter := GetFirst([]string{"a","b","c"})
	fmt.Println(letter)
	
	boolean := GetFirst([]bool{true, false, true, true})
	fmt.Println(boolean)
	
	integer := GetFirst([]int{})
	fmt.Println(integer)
	// END_MAIN OMIT
}
