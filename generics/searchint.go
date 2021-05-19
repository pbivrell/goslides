package main

import "fmt"

func main() {
	fmt.Println(Search(2, []int{1, 2, 3}))
}

func Search(find int, list []int) (item int) {
	for _, v := range list {
		if v == find {
			return v
		}
	}
	return item
}
