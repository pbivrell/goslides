package main

import "fmt"

func main() {
	fmt.Println(Search("a", []string{"a", "b", "c"}))
	fmt.Println(Search(2, []int{1, 2, 3}))
}

func Search (type Comparable) (find Comparable, list []Comparable) (item Comparable) {
	for _, v := range list {
		if v == find {
			return v
		}
	}
	return item
}
