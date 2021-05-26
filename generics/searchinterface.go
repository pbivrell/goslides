package main

import "fmt"

func main() {
	fmt.Println(Search("a", []interface{}{"a", "b", "c"}))
	fmt.Println(Search(2, []interface{}{1, 2, 3}))
}

func Search(find interface{}, list []interface{}) (item interface{}) {
	for _, v := range list {
		if v == find {
			return v
		}
	}
	return item
}
