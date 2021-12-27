package main

import "fmt"

func main() {
	fmt.Println(Search("a", []string{"a", "b", "c"}))
}

func Search(find string, list []string) (item string) {
	for _, v := range list {
		if v == find {
			return v
		}
	}
	return item
}
