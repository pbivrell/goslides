package main

import "fmt"

func main() {
	fmt.Println(Search("a", []string{"a", "b", "c"}))
	fmt.Println(Search(2, []int{1, 2, 3}))
}

// START OMIT
func Search[T comparable](find T, list []T) (item T) {
	for _, v := range list {
		if v == find {
			return v
		}
	}
	return item
}
// END OMIT